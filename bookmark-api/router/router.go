package router

import (
	"bookmark-api/handler"
	"bookmark-api/middleware"
	"net/http"
	"strings"
)

// NewRouter 创建并配置路由
// Go 标准库的路由比较基础，真实项目会用 chi、gorilla/mux、gin 等
// 这里用标准库手动实现，帮助你理解路由原理
func NewRouter(bookmarkHandler *handler.BookmarkHandler) http.Handler {
	mux := http.NewServeMux()

	// 健康检查
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		handler.Success(w, map[string]string{"status": "ok"})
	})

	// 书签 API
	mux.HandleFunc("/api/bookmarks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			bookmarkHandler.GetAll(w, r)
		case http.MethodPost:
			bookmarkHandler.Create(w, r)
		default:
			handler.Error(w, http.StatusMethodNotAllowed, "不支持的请求方法")
		}
	})

	// 带 ID 的路由：/api/bookmarks/{id}
	mux.HandleFunc("/api/bookmarks/", func(w http.ResponseWriter, r *http.Request) {
		// 检查是否有 ID
		idStr := strings.TrimPrefix(r.URL.Path, "/api/bookmarks/")
		if idStr == "" {
			handler.Error(w, http.StatusBadRequest, "缺少书签ID")
			return
		}

		switch r.Method {
		case http.MethodGet:
			bookmarkHandler.GetByID(w, r)
		case http.MethodPut:
			bookmarkHandler.Update(w, r)
		case http.MethodDelete:
			bookmarkHandler.Delete(w, r)
		default:
			handler.Error(w, http.StatusMethodNotAllowed, "不支持的请求方法")
		}
	})

	// 应用中间件（从外到内执行）
	// Recovery → CORS → Logger → 实际处理器
	var app http.Handler = mux
	app = middleware.Logger(app)
	app = middleware.CORS(app)
	app = middleware.Recovery(app)

	return app
}
