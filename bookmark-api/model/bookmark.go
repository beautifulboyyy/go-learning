package model

import "time"

// Bookmark 是书签的数据模型
// 对比 Python: @dataclass class Bookmark
// 对比 Java:   public class Bookmark
//
// `json:"xxx"` 是 struct tag，控制 JSON 序列化行为
// 类似 Java 的 @JsonProperty / Python 的 field(alias=...)
type Bookmark struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	Tag       string    `json:"tag"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateBookmarkRequest 是创建书签的请求体
// Go 惯例：把请求/响应的 DTO 定义在 model 包或 handler 包
type CreateBookmarkRequest struct {
	Title string `json:"title"`
	URL   string `json:"url"`
	Tag   string `json:"tag"`
}

// UpdateBookmarkRequest 是更新书签的请求体
type UpdateBookmarkRequest struct {
	Title string `json:"title,omitempty"` // omitempty: 空值时不出现在 JSON 中
	URL   string `json:"url,omitempty"`
	Tag   string `json:"tag,omitempty"`
}
