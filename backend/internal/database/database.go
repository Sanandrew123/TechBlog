package database

import (
	"fmt"
	"log"
	"techblog-api/backend/internal/config"
	"techblog-api/backend/internal/models"
	
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDatabase(config *config.Config) {
	var err error
	
	// 构建数据库连接字符串
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBSSLMode,
	)
	
	// 设置GORM配置
	gormConfig := &gorm.Config{}
	
	// 在开发环境下启用详细日志
	if config.Environment == "development" {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}
	
	// 连接数据库
	DB, err = gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	
	log.Println("Database connected successfully")
	
	// 自动迁移数据库结构
	if err := AutoMigrate(); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}

func AutoMigrate() error {
	log.Println("Starting database migration...")
	
	err := DB.AutoMigrate(
		&models.User{},
		&models.BlogPost{},
		&models.ContactMessage{},
		&models.Category{},
		&models.Comment{},
		&models.SponsorOrder{},
	)
	
	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}
	
	log.Println("Database migration completed successfully")
	
	// 创建默认管理员用户（如果不存在）
	if err := CreateDefaultAdmin(); err != nil {
		log.Printf("Warning: Failed to create default admin: %v", err)
	}
	
	return nil
}

func CreateDefaultAdmin() error {
	var user models.User
	
	// 检查是否已存在管理员用户
	result := DB.Where("role = ?", "admin").First(&user)
	if result.Error == nil {
		log.Println("Admin user already exists")
		return nil
	}
	
	// 创建默认管理员
	// 注意：在生产环境中，密码应该通过安全的方式设置
	defaultAdmin := models.User{
		Username: "admin",
		Email:    "admin@techblog.com",
		Password: "$2a$12$LQv3c1yqBWVHxkd0LHAkCOYz6TtxMQJqhN8/LewdBPj6QSs3kRYc6", // "admin123" 的bcrypt哈希
		Role:     "admin",
		Bio:      "系统管理员",
		Active:   true,
	}
	
	if err := DB.Create(&defaultAdmin).Error; err != nil {
		return fmt.Errorf("failed to create default admin: %w", err)
	}
	
	log.Println("Default admin user created successfully")
	log.Println("Username: admin, Password: admin123")
	log.Println("请在生产环境中立即修改默认密码！")
	
	return nil
}

func GetDB() *gorm.DB {
	return DB
}