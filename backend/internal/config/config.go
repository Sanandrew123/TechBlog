package config

import (
	"os"
	"strconv"
	"log"
	"github.com/joho/godotenv"
)

type Config struct {
	// 服务器配置
	ServerPort string
	ServerHost string
	
	// 数据库配置
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	
	// JWT配置
	JWTSecret string
	
	// 文件上传配置
	UploadPath string
	MaxFileSize int64
	
	// CORS配置
	AllowOrigins []string
	
	// 环境
	Environment string
}

func LoadConfig() *Config {
	// 加载.env文件（如果存在）
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}
	
	config := &Config{
		// 服务器配置
		ServerPort: getEnv("SERVER_PORT", "8080"),
		ServerHost: getEnv("SERVER_HOST", "0.0.0.0"),
		
		// 数据库配置
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "techblog"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),
		
		// JWT配置
		JWTSecret: getEnv("JWT_SECRET", "your-super-secret-jwt-key-change-this-in-production"),
		
		// 文件上传配置
		UploadPath:  getEnv("UPLOAD_PATH", "./uploads"),
		MaxFileSize: getEnvAsInt64("MAX_FILE_SIZE", 10*1024*1024), // 10MB默认
		
		// CORS配置
		AllowOrigins: []string{
			getEnv("FRONTEND_URL", "http://localhost:5173"),
			"http://localhost:3000",
			"http://18.178.203.0",
		},
		
		// 环境
		Environment: getEnv("ENVIRONMENT", "development"),
	}
	
	return config
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt64(key string, defaultValue int64) int64 {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseInt(valueStr, 10, 64); err == nil {
		return value
	}
	return defaultValue
}