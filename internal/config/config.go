package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	// 数据库配置
	DBType     string // mysql 或 sqlite
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	
	// 服务器配置
	ServerHost string
	ServerPort string
	
	// 应用配置
	AppEnv   string
	AppDebug bool
}

func Load() *Config {
	// 加载 .env 文件
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}
	
	return &Config{
		DBType:     getEnv("DB_TYPE", "mysql"),
		DBHost:     getEnv("DB_HOST", "117.72.67.244"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "A123456"),
		DBName:     getEnv("DB_NAME", "topservice_db"),
		
		ServerHost: getEnv("SERVER_HOST", "0.0.0.0"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
		
		AppEnv:   getEnv("APP_ENV", "development"),
		AppDebug: getEnv("APP_DEBUG", "true") == "true",
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}