/*
开发心理过程：
1. 设计主入口文件，初始化所有必要的组件
2. 配置路由，区分公开API和需要认证的管理API
3. 设置优雅关闭，确保服务稳定性
4. 集成所有中间件：CORS、认证、日志等
5. 提供清晰的API文档和错误处理
*/

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	
	"github.com/gin-gonic/gin"
	"techblog-api/backend/internal/api"
	"techblog-api/backend/internal/config"
	"techblog-api/backend/internal/database"
	"techblog-api/backend/internal/handlers"
	"techblog-api/backend/internal/middleware"
	"techblog-api/backend/internal/models"
)

func main() {
	// 加载配置
	cfg := config.LoadConfig()
	
	// 设置Gin模式
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	
	// 初始化数据库
	database.InitDatabase(cfg)
	
	// 创建Gin引擎
	r := gin.New()
	
	// 全局中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORSMiddleware(cfg))
	
	// 创建API处理器
	blogHandler := api.NewBlogHandler()
	contactHandler := api.NewContactHandler()
	authHandler := api.NewAuthHandler(cfg)
	sponsorHandler := handlers.NewSponsorHandler(database.GetDB())
	
	// API路由组
	api := r.Group("/api/v1")
	{
		// 健康检查
		api.GET("/health", healthCheck)
		
		// 公开API - 博客相关
		api.GET("/posts", blogHandler.GetPosts)
		api.GET("/posts/:id", blogHandler.GetPost)
		
		// 公开API - 联系表单
		api.POST("/contact", contactHandler.SubmitContact)
		
		// 赞助相关API
		sponsor := api.Group("/sponsor")
		{
			sponsor.POST("/create", sponsorHandler.CreateSponsorOrder)
			sponsor.GET("/status/:orderId", sponsorHandler.GetOrderStatus)
			sponsor.GET("/list", sponsorHandler.GetSponsorList)
			sponsor.GET("/stats", sponsorHandler.GetSponsorStats)
			sponsor.POST("/mock-callback/:orderId", sponsorHandler.MockPaymentCallback)
		}
		
		// 认证相关
		auth := api.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
			
			// 需要认证的认证API
			authRequired := auth.Group("/")
			authRequired.Use(middleware.AuthMiddleware(cfg))
			{
				authRequired.GET("/profile", authHandler.GetProfile)
				authRequired.PUT("/profile", authHandler.UpdateProfile)
				authRequired.POST("/change-password", authHandler.ChangePassword)
				authRequired.POST("/refresh", authHandler.RefreshToken)
			}
		}
		
		// 管理员API
		admin := api.Group("/admin")
		admin.Use(middleware.AuthMiddleware(cfg))
		admin.Use(middleware.AdminMiddleware())
		{
			// 博客管理
			admin.POST("/posts", blogHandler.CreatePost)
			admin.PUT("/posts/:id", blogHandler.UpdatePost)
			admin.DELETE("/posts/:id", blogHandler.DeletePost)
			
			// 联系消息管理
			admin.GET("/messages", contactHandler.GetMessages)
			admin.GET("/messages/:id", contactHandler.GetMessage)
			admin.PUT("/messages/:id/read", contactHandler.MarkAsRead)
			admin.PUT("/messages/:id/replied", contactHandler.MarkAsReplied)
			admin.DELETE("/messages/:id", contactHandler.DeleteMessage)
		}
	}
	
	// 静态文件服务（用于上传的文件）
	r.Static("/uploads", cfg.UploadPath)
	
	// API文档路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "TechBlog API Server",
			"version": "v1.0.0",
			"status":  "running",
			"docs":    "/api/docs",
			"endpoints": gin.H{
				"public": gin.H{
					"GET /api/v1/health":        "健康检查",
					"GET /api/v1/posts":         "获取博客文章列表",
					"GET /api/v1/posts/:id":     "获取单个博客文章",
					"POST /api/v1/contact":      "提交联系消息",
					"POST /api/v1/sponsor/create":      "创建赞助订单",
					"GET /api/v1/sponsor/status/:orderId": "查询订单状态",
					"GET /api/v1/sponsor/list":         "获取赞助者列表",
					"GET /api/v1/sponsor/stats":        "获取赞助统计",
				},
				"auth": gin.H{
					"POST /api/v1/auth/login":           "用户登录",
					"GET /api/v1/auth/profile":          "获取用户信息（需要认证）",
					"PUT /api/v1/auth/profile":          "更新用户信息（需要认证）",
					"POST /api/v1/auth/change-password": "修改密码（需要认证）",
					"POST /api/v1/auth/refresh":         "刷新token（需要认证）",
				},
				"admin": gin.H{
					"POST /api/v1/admin/posts":                "创建博客文章（需要管理员权限）",
					"PUT /api/v1/admin/posts/:id":             "更新博客文章（需要管理员权限）",
					"DELETE /api/v1/admin/posts/:id":          "删除博客文章（需要管理员权限）",
					"GET /api/v1/admin/messages":              "获取联系消息列表（需要管理员权限）",
					"GET /api/v1/admin/messages/:id":          "获取单个联系消息（需要管理员权限）",
					"PUT /api/v1/admin/messages/:id/read":     "标记消息为已读（需要管理员权限）",
					"PUT /api/v1/admin/messages/:id/replied":  "标记消息为已回复（需要管理员权限）",
					"DELETE /api/v1/admin/messages/:id":       "删除联系消息（需要管理员权限）",
				},
			},
		})
	})
	
	// 创建HTTP服务器
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort),
		Handler: r,
	}
	
	// 启动服务器
	go func() {
		log.Printf("🚀 Server starting on %s:%s", cfg.ServerHost, cfg.ServerPort)
		log.Printf("🌍 Environment: %s", cfg.Environment)
		log.Printf("📚 API Documentation: http://%s:%s/", cfg.ServerHost, cfg.ServerPort)
		log.Printf("💚 Health Check: http://%s:%s/api/v1/health", cfg.ServerHost, cfg.ServerPort)
		
		if cfg.Environment == "development" {
			log.Println("🔧 Development mode - detailed logging enabled")
			log.Printf("🔑 Default admin: username=admin, password=admin123")
		}
		
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()
	
	// 等待中断信号以优雅地关闭服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("🛑 Shutting down server...")
	
	// 优雅关闭，等待现有连接完成
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("❌ Server forced to shutdown:", err)
	}
	
	log.Println("✅ Server exited")
}

func healthCheck(c *gin.Context) {
	// 检查数据库连接
	db := database.GetDB()
	sqlDB, err := db.DB()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "error",
			"error":  "Database connection failed",
		})
		return
	}
	
	if err := sqlDB.Ping(); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "error",
			"error":  "Database ping failed",
		})
		return
	}
	
	// 获取基础统计信息
	var postCount, messageCount int64
	db.Model(&models.BlogPost{}).Count(&postCount)
	db.Model(&models.ContactMessage{}).Count(&messageCount)
	
	c.JSON(http.StatusOK, gin.H{
		"status":    "healthy",
		"timestamp": time.Now().Unix(),
		"version":   "v1.0.0",
		"database":  "connected",
		"stats": gin.H{
			"posts":    postCount,
			"messages": messageCount,
		},
	})
}