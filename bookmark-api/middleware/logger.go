package middleware

import (
	"log"
	"net/http"
	"time"
)

// responseWriter 包装 http.ResponseWriter，捕获状态码
// Go 的 struct 嵌入（embedding）：类似组合，但更简洁
type responseWriter struct {
	http.ResponseWriter // 嵌入，自动获得所有方法
	statusCode          int
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// Logger 日志中间件
// 记录每个请求的方法、路径、状态码、耗时
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// 包装 ResponseWriter 以捕获状态码
		wrapped := newResponseWriter(w)

		// 调用下一个处理器
		next.ServeHTTP(wrapped, r)

		// 记录日志
		log.Printf(
			"[%s] %s %s → %d (%v)",
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
			wrapped.statusCode,
			time.Since(start),
		)
	})
}
