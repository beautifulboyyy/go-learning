package handler

import (
	"bookmark-api/model"
	"bookmark-api/service"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// BookmarkHandler HTTP 处理器
// Handler 层负责：解析请求、调用 Service、返回响应
// 对比 Java: @RestController
// 对比 Python: Flask route handler
type BookmarkHandler struct {
	service service.BookmarkService
}

// NewBookmarkHandler 构造函数
func NewBookmarkHandler(svc service.BookmarkService) *BookmarkHandler {
	return &BookmarkHandler{service: svc}
}

// GetAll GET /api/bookmarks
func (h *BookmarkHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	bookmarks, err := h.service.GetAll()
	if err != nil {
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	// 空切片返回空数组而不是 null
	if bookmarks == nil {
		bookmarks = []model.Bookmark{}
	}

	Success(w, bookmarks)
}

// GetByID GET /api/bookmarks/{id}
func (h *BookmarkHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	// 从 URL 路径提取 ID
	// Go 标准库没有路由参数，需要手动解析
	// 真实项目会用 chi、gin 等路由库
	idStr := strings.TrimPrefix(r.URL.Path, "/api/bookmarks/")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		Error(w, http.StatusBadRequest, "无效的ID")
		return
	}

	bm, err := h.service.GetByID(id)
	if err != nil {
		Error(w, http.StatusBadRequest, err.Error())
		return
	}
	if bm == nil {
		Error(w, http.StatusNotFound, "书签不存在")
		return
	}

	Success(w, bm)
}

// Create POST /api/bookmarks
func (h *BookmarkHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateBookmarkRequest

	// json.NewDecoder 解析请求体
	// 对比 Python: request.get_json()
	// 对比 Java:   @RequestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		Error(w, http.StatusBadRequest, "无效的JSON格式")
		return
	}

	bm, err := h.service.Create(req)
	if err != nil {
		Error(w, http.StatusBadRequest, err.Error())
		return
	}

	Created(w, bm)
}

// Update PUT /api/bookmarks/{id}
func (h *BookmarkHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/bookmarks/")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		Error(w, http.StatusBadRequest, "无效的ID")
		return
	}

	var req model.UpdateBookmarkRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		Error(w, http.StatusBadRequest, "无效的JSON格式")
		return
	}

	bm, err := h.service.Update(id, req)
	if err != nil {
		Error(w, http.StatusBadRequest, err.Error())
		return
	}

	Success(w, bm)
}

// Delete DELETE /api/bookmarks/{id}
func (h *BookmarkHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/bookmarks/")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		Error(w, http.StatusBadRequest, "无效的ID")
		return
	}

	if err := h.service.Delete(id); err != nil {
		Error(w, http.StatusBadRequest, err.Error())
		return
	}

	Success(w, map[string]string{"message": "删除成功"})
}
