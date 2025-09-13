package main

import (
	"log"
	"topService/internal/config"
	"topService/internal/database"
	"topService/internal/handler"
	"topService/internal/middleware"
	"topService/internal/router"
	"topService/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	// 加载配置
	cfg := config.Load()
	
	// 初始化数据库
	var db *gorm.DB
	var err error
	
	if cfg.DBType == "sqlite" {
		db, err = database.InitializeSQLite(cfg)
	} else {
		db, err = database.Initialize(cfg)
	}
	
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	
	// 自动迁移数据库表
	if err := database.AutoMigrate(db); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	
	// 初始化服务层
	userService := service.NewUserService(db)
	productService := service.NewProductService(db)
	movieService := service.NewMovieService(db)
	
	// 初始化处理器层
	userHandler := handler.NewUserHandler(userService)
	productHandler := handler.NewProductHandler(productService)
	movieHandler := handler.NewMovieHandler(movieService)
	
	// 设置运行模式
	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	
	// 创建路由器
	r := gin.New()
	
	// 添加中间件
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())
	r.Use(middleware.CORS())
	
	// 设置路由
	router.SetupRoutes(r, userHandler, productHandler, movieHandler)
	
	// 启动服务器
	addr := cfg.ServerHost + ":" + cfg.ServerPort
	log.Printf("Server starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}