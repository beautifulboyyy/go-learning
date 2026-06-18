package service

import (
	"bookmark-api/model"
	"bookmark-api/repository"
	"fmt"
	"strings"
)

// BookmarkService 定义业务逻辑接口
// Service 层负责：参数校验、业务规则、调用 Repository
// 对比 Java: @Service 注解的类
type BookmarkService interface {
	GetAll() ([]model.Bookmark, error)
	GetByID(id int64) (*model.Bookmark, error)
	Create(req model.CreateBookmarkRequest) (*model.Bookmark, error)
	Update(id int64, req model.UpdateBookmarkRequest) (*model.Bookmark, error)
	Delete(id int64) error
}

type bookmarkService struct {
	repo repository.BookmarkRepository // 依赖注入：通过接口，不依赖具体实现
}

// NewBookmarkService 构造函数
// Go 的依赖注入：把接口作为参数传入，而不是直接创建实现
// 对比 Java: @Autowired / constructor injection
func NewBookmarkService(repo repository.BookmarkRepository) BookmarkService {
	return &bookmarkService{repo: repo}
}

// GetAll 获取所有书签
func (s *bookmarkService) GetAll() ([]model.Bookmark, error) {
	return s.repo.GetAll()
}

// GetByID 根据 ID 获取书签
func (s *bookmarkService) GetByID(id int64) (*model.Bookmark, error) {
	if id <= 0 {
		return nil, fmt.Errorf("无效的书签ID: %d", id)
	}
	return s.repo.GetByID(id)
}

// Create 创建书签（带业务校验）
func (s *bookmarkService) Create(req model.CreateBookmarkRequest) (*model.Bookmark, error) {
	// 参数校验
	if strings.TrimSpace(req.Title) == "" {
		return nil, fmt.Errorf("标题不能为空")
	}
	if strings.TrimSpace(req.URL) == "" {
		return nil, fmt.Errorf("URL不能为空")
	}
	if !strings.HasPrefix(req.URL, "http://") && !strings.HasPrefix(req.URL, "https://") {
		return nil, fmt.Errorf("URL必须以 http:// 或 https:// 开头")
	}

	return s.repo.Create(req)
}

// Update 更新书签
func (s *bookmarkService) Update(id int64, req model.UpdateBookmarkRequest) (*model.Bookmark, error) {
	if id <= 0 {
		return nil, fmt.Errorf("无效的书签ID: %d", id)
	}

	// 校验 URL 格式（如果提供了的话）
	if req.URL != "" && !strings.HasPrefix(req.URL, "http://") && !strings.HasPrefix(req.URL, "https://") {
		return nil, fmt.Errorf("URL必须以 http:// 或 https:// 开头")
	}

	bm, err := s.repo.Update(id, req)
	if err != nil {
		return nil, err
	}
	if bm == nil {
		return nil, fmt.Errorf("书签不存在: %d", id)
	}
	return bm, nil
}

// Delete 删除书签
func (s *bookmarkService) Delete(id int64) error {
	if id <= 0 {
		return fmt.Errorf("无效的书签ID: %d", id)
	}
	return s.repo.Delete(id)
}
