/*
开发心理过程：
1. 设计数据模型，包括博客文章、联系消息、用户等核心实体
2. 使用GORM标签进行数据库映射和验证
3. 考虑SEO友好的URL slug设计
4. 包含创建时间、更新时间等审计字段
5. 为后续扩展预留空间（如评论、分类等）
*/

package models

import (
	"time"
	"gorm.io/gorm"
)

// BlogPost 博客文章模型
type BlogPost struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"size:200;not null" json:"title" binding:"required"`
	Content     string         `gorm:"type:text;not null" json:"content" binding:"required"`
	Excerpt     string         `gorm:"size:500" json:"excerpt"`
	Slug        string         `gorm:"size:200;uniqueIndex;not null" json:"slug"`
	Author      string         `gorm:"size:100;not null" json:"author"`
	Tags        string         `gorm:"size:500" json:"tags"` // JSON字符串存储标签数组
	CoverImage  string         `gorm:"size:500" json:"coverImage"`
	Published   bool           `gorm:"default:false" json:"published"`
	ReadTime    int            `gorm:"default:5" json:"readTime"` // 预估阅读时间（分钟）
	ViewCount   int            `gorm:"default:0" json:"viewCount"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// User 用户模型（管理员）
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"size:50;uniqueIndex;not null" json:"username"`
	Email     string         `gorm:"size:100;uniqueIndex;not null" json:"email"`
	Password  string         `gorm:"size:255;not null" json:"-"` // 密码不返回给前端
	Role      string         `gorm:"size:20;default:admin" json:"role"`
	Avatar    string         `gorm:"size:500" json:"avatar"`
	Bio       string         `gorm:"size:500" json:"bio"`
	Active    bool           `gorm:"default:true" json:"active"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// ContactMessage 联系消息模型
type ContactMessage struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name" binding:"required"`
	Email     string    `gorm:"size:100;not null" json:"email" binding:"required,email"`
	Subject   string    `gorm:"size:200;not null" json:"subject" binding:"required"`
	Message   string    `gorm:"type:text;not null" json:"message" binding:"required"`
	IsRead    bool      `gorm:"default:false" json:"isRead"`
	IsReplied bool      `gorm:"default:false" json:"isReplied"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Category 分类模型（为后续扩展准备）
type Category struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"size:100;uniqueIndex;not null" json:"name"`
	Slug        string         `gorm:"size:100;uniqueIndex;not null" json:"slug"`
	Description string         `gorm:"size:500" json:"description"`
	Color       string         `gorm:"size:7;default:#6366f1" json:"color"` // 十六进制颜色
	PostCount   int            `gorm:"default:0" json:"postCount"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// Comment 评论模型（为后续扩展准备）
type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	PostID    uint      `gorm:"not null;index" json:"postId"`
	Post      BlogPost  `gorm:"foreignKey:PostID" json:"post,omitempty"`
	Author    string    `gorm:"size:100;not null" json:"author"`
	Email     string    `gorm:"size:100;not null" json:"email"`
	Website   string    `gorm:"size:200" json:"website"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	Approved  bool      `gorm:"default:false" json:"approved"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// APIResponse 通用API响应结构
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Meta    *PaginationMeta `json:"meta,omitempty"`
}

// PaginationMeta 分页元数据
type PaginationMeta struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int `json:"totalPages"`
}

// LoginRequest 登录请求结构
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse 登录响应结构
type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

// BlogPostRequest 博客文章请求结构
type BlogPostRequest struct {
	Title      string   `json:"title" binding:"required"`
	Content    string   `json:"content" binding:"required"`
	Excerpt    string   `json:"excerpt"`
	Author     string   `json:"author" binding:"required"`
	Tags       []string `json:"tags"`
	CoverImage string   `json:"coverImage"`
	Published  bool     `json:"published"`
}

// BlogPostQuery 博客文章查询参数
type BlogPostQuery struct {
	Page      int    `form:"page,default=1"`
	Limit     int    `form:"limit,default=10"`
	Search    string `form:"search"`
	Tag       string `form:"tag"`
	Author    string `form:"author"`
	Published *bool  `form:"published"`
	Sort      string `form:"sort,default=created_at"`
	Order     string `form:"order,default=desc"`
}

// SponsorOrder 赞助订单模型
type SponsorOrder struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	OrderID      string    `gorm:"uniqueIndex;size:100" json:"orderId"`
	Amount       float64   `gorm:"not null" json:"amount"`
	SponsorName  string    `gorm:"size:100;default:'匿名赞助者'" json:"sponsorName"`
	Message      string    `gorm:"type:text" json:"message"`
	PaymentMethod string   `gorm:"size:20;default:'wechat'" json:"paymentMethod"`
	Status       string    `gorm:"size:20;default:'pending'" json:"status"` // pending, paid, failed, cancelled
	QRCodeURL    string    `gorm:"size:500" json:"qrCodeUrl,omitempty"`
	TransactionID string   `gorm:"size:100" json:"transactionId,omitempty"`
	PaidAt       *time.Time `json:"paidAt,omitempty"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

// SponsorRequest 赞助请求结构
type SponsorRequest struct {
	Amount        float64 `json:"amount" binding:"required,min=1"`
	SponsorName   string  `json:"sponsorName"`
	Message       string  `json:"message"`
	PaymentMethod string  `json:"paymentMethod,default=wechat"`
}

// SponsorResponse 赞助响应结构
type SponsorResponse struct {
	OrderID string `json:"orderId"`
	QRCode  string `json:"qrCode"`
	Amount  float64 `json:"amount"`
}