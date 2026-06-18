package config

import (
	"os"
)

// Config 应用配置
// 真实项目会用环境变量或配置文件
// 对比 Python: os.environ.get() / Java: @Value
type Config struct {
	Port   string
	DBPath string
}

// Load 加载配置（从环境变量读取，有默认值）
func Load() *Config {
	return &Config{
		Port:   getEnv("PORT", "8080"),
		DBPath: getEnv("DB_PATH", "./bookmarks.db"),
	}
}

// getEnv 获取环境变量，如果不存在返回默认值
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
