package database

import (
	"fmt"
	"log"
	"topService/internal/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitializeSQLite(cfg *config.Config) (*gorm.DB, error) {
	// 使用SQLite数据库文件
	dbPath := "topservice.db"
	
	// 配置GORM日志级别
	var logLevel logger.LogLevel
	if cfg.AppDebug {
		logLevel = logger.Info
	} else {
		logLevel = logger.Silent
	}
	
	// 连接SQLite数据库
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	
	if err != nil {
		return nil, fmt.Errorf("failed to connect to SQLite database: %w", err)
	}
	
	log.Println("SQLite database connected successfully")
	return db, nil
}