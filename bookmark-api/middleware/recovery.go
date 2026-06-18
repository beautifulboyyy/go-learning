package middleware

import (
	"log"
	"net/http"
	"runtime/debug"
)

// Recovery panic 恢复中间件
// Go 没有 try/catch，用 recover() 捕获 panic
// 类似 Python 的 except Exception / Java 的 catch(Throwable)
// 真实项目必备：防止一个请求的 panic 导致整个服务器崩溃
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// 打印堆栈信息
				log.Printf("🚨 Panic recovered: %v\n%s", err, debug.Stack())

				// 返回 500 错误
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
