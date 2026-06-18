package repository

import (
	"bookmark-api/model"
	"database/sql"
	"fmt"
	"time"
)

// BookmarkRepository 定义数据访问接口
// Go 的 interface 是隐式实现的——只要 struct 实现了所有方法，就自动满足接口
// 对比 Java: 不需要写 "implements BookmarkRepository"
// 对比 Python: 类似 ABC，但更简洁
type BookmarkRepository interface {
	GetAll() ([]model.Bookmark, error)
	GetByID(id int64) (*model.Bookmark, error)
	Create(req model.CreateBookmarkRequest) (*model.Bookmark, error)
	Update(id int64, req model.UpdateBookmarkRequest) (*model.Bookmark, error)
	Delete(id int64) error
}

// sqliteBookmarkRepo 是 SQLite 实现
// 小写开头 = 包外不可见（类似 Java 的 private）
// 这是 Go 的访问控制：大写开头=public，小写开头=private
type sqliteBookmarkRepo struct {
	db *sql.DB
}

// NewBookmarkRepository 是构造函数
// Go 惯例：用 NewXxx 函数创建实例，而不是直接暴露 struct
// 对比 Python: def __init__(self, db)
// 对比 Java:   new BookmarkRepository(db)
func NewBookmarkRepository(db *sql.DB) BookmarkRepository {
	return &sqliteBookmarkRepo{db: db}
}

// GetAll 获取所有书签
// Go 的多返回值：(结果, 错误)
// 对比 Python: try/except 捕获异常
// 对比 Java:   throws Exception
func (r *sqliteBookmarkRepo) GetAll() ([]model.Bookmark, error) {
	query := `SELECT id, title, url, tag, created_at, updated_at FROM bookmarks ORDER BY created_at DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("查询书签失败: %w", err) // %w 包装原始错误
	}
	defer rows.Close() // defer: 函数结束时执行，类似 Python 的 with / Java 的 try-with-resources

	var bookmarks []model.Bookmark
	for rows.Next() {
		var b model.Bookmark
		err := rows.Scan(&b.ID, &b.Title, &b.URL, &b.Tag, &b.CreatedAt, &b.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("扫描书签行失败: %w", err)
		}
		bookmarks = append(bookmarks, b)
	}

	return bookmarks, nil
}

// GetByID 根据 ID 获取单个书签
func (r *sqliteBookmarkRepo) GetByID(id int64) (*model.Bookmark, error) {
	query := `SELECT id, title, url, tag, created_at, updated_at FROM bookmarks WHERE id = ?`

	var b model.Bookmark
	err := r.db.QueryRow(query, id).Scan(&b.ID, &b.Title, &b.URL, &b.Tag, &b.CreatedAt, &b.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // 没找到返回 nil, nil（不是错误）
		}
		return nil, fmt.Errorf("查询书签失败: %w", err)
	}

	return &b, nil // 返回指针，避免拷贝整个 struct
}

// Create 创建书签
func (r *sqliteBookmarkRepo) Create(req model.CreateBookmarkRequest) (*model.Bookmark, error) {
	query := `INSERT INTO bookmarks (title, url, tag, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`

	now := time.Now()
	result, err := r.db.Exec(query, req.Title, req.URL, req.Tag, now, now)
	if err != nil {
		return nil, fmt.Errorf("创建书签失败: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("获取新书签ID失败: %w", err)
	}

	return &model.Bookmark{
		ID:        id,
		Title:     req.Title,
		URL:       req.URL,
		Tag:       req.Tag,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

// Update 更新书签
func (r *sqliteBookmarkRepo) Update(id int64, req model.UpdateBookmarkRequest) (*model.Bookmark, error) {
	// 先获取现有书签
	existing, err := r.GetByID(id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, nil // 书签不存在
	}

	// 合并更新（只更新非空字段）
	if req.Title != "" {
		existing.Title = req.Title
	}
	if req.URL != "" {
		existing.URL = req.URL
	}
	if req.Tag != "" {
		existing.Tag = req.Tag
	}
	existing.UpdatedAt = time.Now()

	query := `UPDATE bookmarks SET title = ?, url = ?, tag = ?, updated_at = ? WHERE id = ?`
	_, err = r.db.Exec(query, existing.Title, existing.URL, existing.Tag, existing.UpdatedAt, id)
	if err != nil {
		return nil, fmt.Errorf("更新书签失败: %w", err)
	}

	return existing, nil
}

// Delete 删除书签
func (r *sqliteBookmarkRepo) Delete(id int64) error {
	query := `DELETE FROM bookmarks WHERE id = ?`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("删除书签失败: %w", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("书签不存在") // 返回 error 表示没找到
	}

	return nil
}
