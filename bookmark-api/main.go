package main

import (
	"bookmark-api/config"
	"bookmark-api/database"
	"bookmark-api/handler"
	"bookmark-api/repository"
	"bookmark-api/router"
	"bookmark-api/service"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// ========================================
	// 1. 加载配置
	// ========================================
	cfg := config.Load()
	log.Printf("📋 配置: Port=%s, DBPath=%s", cfg.Port, cfg.DBPath)

	// ========================================
	// 2. 初始化数据库
	// ========================================
	db, err := database.InitDB(cfg.DBPath)
	if err != nil {
		log.Fatalf("❌ 数据库初始化失败: %v", err)
	}
	defer db.Close() // 程序退出时关闭数据库

	// ========================================
	// 3. 依赖注入（手动，真实项目会用 wire 等 DI 框架）
	// ========================================
	// Repository → Service → Handler → Router
	// 这就是 Go 的依赖注入方式：手动组装，清晰明了
	bookmarkRepo := repository.NewBookmarkRepository(db)
	bookmarkSvc := service.NewBookmarkService(bookmarkRepo)
	bookmarkHandler := handler.NewBookmarkHandler(bookmarkSvc)
	app := router.NewRouter(bookmarkHandler)

	// ========================================
	// 4. 创建 HTTP 服务器
	// ========================================
	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      app,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// ========================================
	// 5. 优雅关闭（Graceful Shutdown）
	// ========================================
	// 监听系统信号：Ctrl+C 或 kill
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 在 goroutine 中启动服务器
	// goroutine 是 Go 的轻量级线程，用 go 关键字启动
	// 对比 Python: threading.Thread(target=...).start()
	// 对比 Java:   new Thread(() -> ...).start()
	go func() {
		fmt.Printf("🚀 服务器启动: http://localhost:%s\n", cfg.Port)
		fmt.Println("📖 API 文档:")
		fmt.Println("  GET    /health              - 健康检查")
		fmt.Println("  GET    /api/bookmarks       - 获取所有书签")
		fmt.Println("  GET    /api/bookmarks/{id}  - 获取单个书签")
		fmt.Println("  POST   /api/bookmarks       - 创建书签")
		fmt.Println("  PUT    /api/bookmarks/{id}  - 更新书签")
		fmt.Println("  DELETE /api/bookmarks/{id}  - 删除书签")
		fmt.Println()
		fmt.Println("按 Ctrl+C 停止服务器")

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ 服务器启动失败: %v", err)
		}
	}()

	// 等待退出信号
	<-quit
	log.Println("🛑 正在关闭服务器...")

	// 给正在处理的请求 5 秒时间完成
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("❌ 服务器关闭失败: %v", err)
	}

	log.Println("✅ 服务器已安全关闭")
}
