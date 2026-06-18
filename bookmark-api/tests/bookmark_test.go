package tests

import (
	"bookmark-api/model"
	"bookmark-api/service"
	"testing"
)

// mockBookmarkRepo 是一个 mock 实现
// Go 的 interface 天然支持 mock：只要实现接口方法就行
// 对比 Java: Mockito / Python: unittest.mock
type mockBookmarkRepo struct {
	bookmarks map[int64]*model.Bookmark
	nextID    int64
}

func newMockRepo() *mockBookmarkRepo {
	return &mockBookmarkRepo{
		bookmarks: make(map[int64]*model.Bookmark),
		nextID:    1,
	}
}

func (m *mockBookmarkRepo) GetAll() ([]model.Bookmark, error) {
	var result []model.Bookmark
	for _, b := range m.bookmarks {
		result = append(result, *b)
	}
	return result, nil
}

func (m *mockBookmarkRepo) GetByID(id int64) (*model.Bookmark, error) {
	if b, ok := m.bookmarks[id]; ok {
		return b, nil
	}
	return nil, nil
}

func (m *mockBookmarkRepo) Create(req model.CreateBookmarkRequest) (*model.Bookmark, error) {
	b := &model.Bookmark{
		ID:    m.nextID,
		Title: req.Title,
		URL:   req.URL,
		Tag:   req.Tag,
	}
	m.bookmarks[m.nextID] = b
	m.nextID++
	return b, nil
}

func (m *mockBookmarkRepo) Update(id int64, req model.UpdateBookmarkRequest) (*model.Bookmark, error) {
	b, ok := m.bookmarks[id]
	if !ok {
		return nil, nil
	}
	if req.Title != "" {
		b.Title = req.Title
	}
	if req.URL != "" {
		b.URL = req.URL
	}
	if req.Tag != "" {
		b.Tag = req.Tag
	}
	return b, nil
}

func (m *mockBookmarkRepo) Delete(id int64) error {
	if _, ok := m.bookmarks[id]; !ok {
		return nil
	}
	delete(m.bookmarks, id)
	return nil
}

// TestCreateBookmark 表驱动测试
// 这是 Go 的经典测试模式：用切片定义多个测试用例
// 对比 Python: @pytest.mark.parametrize
func TestCreateBookmark(t *testing.T) {
	// 表驱动：定义测试用例
	tests := []struct {
		name    string
		req     model.CreateBookmarkRequest
		wantErr bool
	}{
		{
			name: "正常创建",
			req: model.CreateBookmarkRequest{
				Title: "Go 官方文档",
				URL:   "https://go.dev",
				Tag:   "学习",
			},
			wantErr: false,
		},
		{
			name: "标题为空",
			req: model.CreateBookmarkRequest{
				Title: "",
				URL:   "https://go.dev",
			},
			wantErr: true,
		},
		{
			name: "URL为空",
			req: model.CreateBookmarkRequest{
				Title: "Test",
				URL:   "",
			},
			wantErr: true,
		},
		{
			name: "URL格式错误",
			req: model.CreateBookmarkRequest{
				Title: "Test",
				URL:   "not-a-url",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		// t.Run 创建子测试，可以在 IDE 中单独运行
		t.Run(tt.name, func(t *testing.T) {
			repo := newMockRepo()
			svc := service.NewBookmarkService(repo)

			bm, err := svc.Create(tt.req)

			if tt.wantErr {
				if err == nil {
					t.Errorf("期望错误，但得到了 nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("意外错误: %v", err)
			}
			if bm == nil {
				t.Fatal("返回了 nil 书签")
			}
			if bm.Title != tt.req.Title {
				t.Errorf("标题不匹配: got %q, want %q", bm.Title, tt.req.Title)
			}
		})
	}
}

// TestGetByID 测试获取单个书签
func TestGetByID(t *testing.T) {
	repo := newMockRepo()
	svc := service.NewBookmarkService(repo)

	// 先创建一个
	bm, _ := svc.Create(model.CreateBookmarkRequest{
		Title: "Test",
		URL:   "https://test.com",
	})

	// 测试获取
	tests := []struct {
		name    string
		id      int64
		wantNil bool
	}{
		{"存在的ID", bm.ID, false},
		{"不存在的ID", 999, true},
		{"无效的ID", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := svc.GetByID(tt.id)
			if tt.id <= 0 {
				if err == nil {
					t.Error("期望错误，但得到了 nil")
				}
				return
			}
			if tt.wantNil && result != nil {
				t.Error("期望 nil，但得到了结果")
			}
			if !tt.wantNil && result == nil {
				t.Error("期望结果，但得到了 nil")
			}
		})
	}
}

// TestDeleteBookmark 测试删除
func TestDeleteBookmark(t *testing.T) {
	repo := newMockRepo()
	svc := service.NewBookmarkService(repo)

	bm, _ := svc.Create(model.CreateBookmarkRequest{
		Title: "To Delete",
		URL:   "https://delete.me",
	})

	// 删除
	err := svc.Delete(bm.ID)
	if err != nil {
		t.Fatalf("删除失败: %v", err)
	}

	// 验证已删除
	result, _ := svc.GetByID(bm.ID)
	if result != nil {
		t.Error("书签应该已被删除")
	}
}
