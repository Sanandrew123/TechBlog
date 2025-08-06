package api

import (
	"net/http"
	"strconv"
	"strings"
	"encoding/json"
	"math"
	"time"
	
	"github.com/gin-gonic/gin"
	"techblog-api/backend/internal/database"
	"techblog-api/backend/internal/models"
)

type BlogHandler struct{}

func NewBlogHandler() *BlogHandler {
	return &BlogHandler{}
}

// GetPosts 获取博客文章列表
func (h *BlogHandler) GetPosts(c *gin.Context) {
	var query models.BlogPostQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Invalid query parameters",
			Error:   err.Error(),
		})
		return
	}
	
	db := database.GetDB()
	var posts []models.BlogPost
	var total int64
	
	// 构建查询
	dbQuery := db.Model(&models.BlogPost{})
	
	// 搜索条件
	if query.Search != "" {
		searchTerm := "%" + query.Search + "%"
		dbQuery = dbQuery.Where("title ILIKE ? OR content ILIKE ? OR excerpt ILIKE ?", 
			searchTerm, searchTerm, searchTerm)
	}
	
	// 标签过滤
	if query.Tag != "" {
		dbQuery = dbQuery.Where("tags ILIKE ?", "%"+query.Tag+"%")
	}
	
	// 作者过滤
	if query.Author != "" {
		dbQuery = dbQuery.Where("author ILIKE ?", "%"+query.Author+"%")
	}
	
	// 发布状态过滤
	if query.Published != nil {
		dbQuery = dbQuery.Where("published = ?", *query.Published)
	} else {
		// 对于公开API，只显示已发布的文章
		dbQuery = dbQuery.Where("published = ?", true)
	}
	
	// 获取总数
	if err := dbQuery.Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "Failed to count posts",
			Error:   err.Error(),
		})
		return
	}
	
	// 排序
	orderBy := query.Sort + " " + strings.ToUpper(query.Order)
	if query.Sort == "created_at" || query.Sort == "updated_at" || query.Sort == "view_count" || query.Sort == "read_time" {
		dbQuery = dbQuery.Order(orderBy)
	} else {
		dbQuery = dbQuery.Order("created_at DESC")
	}
	
	// 分页
	offset := (query.Page - 1) * query.Limit
	if err := dbQuery.Offset(offset).Limit(query.Limit).Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "Failed to fetch posts",
			Error:   err.Error(),
		})
		return
	}
	
	// 注意：这里不需要修改Tags字段，它已经是JSON字符串格式
	
	// 计算分页信息
	totalPages := int(math.Ceil(float64(total) / float64(query.Limit)))
	
	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Posts fetched successfully",
		Data:    posts,
		Meta: &models.PaginationMeta{
			Page:       query.Page,
			Limit:      query.Limit,
			Total:      total,
			TotalPages: totalPages,
		},
	})
}

// GetPost 获取单个博客文章
func (h *BlogHandler) GetPost(c *gin.Context) {
	id := c.Param("id")
	
	db := database.GetDB()
	var post models.BlogPost
	
	// 尝试按ID查找
	if postID, err := strconv.ParseUint(id, 10, 32); err == nil {
		if err := db.Where("id = ? AND published = ?", postID, true).First(&post).Error; err != nil {
			c.JSON(http.StatusNotFound, models.APIResponse{
				Success: false,
				Message: "Post not found",
				Error:   "post_not_found",
			})
			return
		}
	} else {
		// 按slug查找
		if err := db.Where("slug = ? AND published = ?", id, true).First(&post).Error; err != nil {
			c.JSON(http.StatusNotFound, models.APIResponse{
				Success: false,
				Message: "Post not found",
				Error:   "post_not_found",
			})
			return
		}
	}
	
	// 增加浏览量
	db.Model(&post).Update("view_count", post.ViewCount+1)
	
	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Post fetched successfully",
		Data:    post,
	})
}

// CreatePost 创建博客文章（需要管理员权限）
func (h *BlogHandler) CreatePost(c *gin.Context) {
	var req models.BlogPostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Invalid request data",
			Error:   err.Error(),
		})
		return
	}
	
	// 生成slug
	slug := generateSlug(req.Title)
	
	// 检查slug是否已存在
	db := database.GetDB()
	var existingPost models.BlogPost
	if err := db.Where("slug = ?", slug).First(&existingPost).Error; err == nil {
		slug = slug + "-" + strconv.FormatInt(time.Now().Unix(), 10)
	}
	
	// 创建文章
	post := models.BlogPost{
		Title:      req.Title,
		Content:    req.Content,
		Excerpt:    req.Excerpt,
		Slug:       slug,
		Author:     req.Author,
		Tags:       serializeTags(req.Tags),
		CoverImage: req.CoverImage,
		Published:  req.Published,
		ReadTime:   calculateReadTime(req.Content),
	}
	
	// 如果excerpt为空，自动生成
	if post.Excerpt == "" {
		post.Excerpt = generateExcerpt(req.Content)
	}
	
	if err := db.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "Failed to create post",
			Error:   err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusCreated, models.APIResponse{
		Success: true,
		Message: "Post created successfully",
		Data:    post,
	})
}

// UpdatePost 更新博客文章（需要管理员权限）
func (h *BlogHandler) UpdatePost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Invalid post ID",
			Error:   "invalid_id",
		})
		return
	}
	
	var req models.BlogPostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Invalid request data",
			Error:   err.Error(),
		})
		return
	}
	
	db := database.GetDB()
	var post models.BlogPost
	
	if err := db.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.APIResponse{
			Success: false,
			Message: "Post not found",
			Error:   "post_not_found",
		})
		return
	}
	
	// 更新文章数据
	post.Title = req.Title
	post.Content = req.Content
	post.Excerpt = req.Excerpt
	post.Author = req.Author
	post.Tags = serializeTags(req.Tags)
	post.CoverImage = req.CoverImage
	post.Published = req.Published
	post.ReadTime = calculateReadTime(req.Content)
	
	// 如果excerpt为空，自动生成
	if post.Excerpt == "" {
		post.Excerpt = generateExcerpt(req.Content)
	}
	
	if err := db.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "Failed to update post",
			Error:   err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Post updated successfully",
		Data:    post,
	})
}

// DeletePost 删除博客文章（需要管理员权限）
func (h *BlogHandler) DeletePost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Invalid post ID",
			Error:   "invalid_id",
		})
		return
	}
	
	db := database.GetDB()
	
	if err := db.Delete(&models.BlogPost{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "Failed to delete post",
			Error:   err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Post deleted successfully",
	})
}

// 辅助函数
func parseTags(tagsJSON string) []string {
	var tags []string
	if tagsJSON != "" {
		json.Unmarshal([]byte(tagsJSON), &tags)
	}
	return tags
}

func serializeTags(tags []string) string {
	if tags == nil {
		return "[]"
	}
	tagsJSON, _ := json.Marshal(tags)
	return string(tagsJSON)
}

func generateSlug(title string) string {
	slug := strings.ToLower(title)
	slug = strings.ReplaceAll(slug, " ", "-")
	// 移除特殊字符，只保留字母、数字和连字符
	var result strings.Builder
	for _, r := range slug {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			result.WriteRune(r)
		}
	}
	return result.String()
}

func calculateReadTime(content string) int {
	words := len(strings.Fields(content))
	readTime := words / 200 // 假设每分钟阅读200字
	if readTime < 1 {
		readTime = 1
	}
	return readTime
}

func generateExcerpt(content string) string {
	words := strings.Fields(content)
	if len(words) <= 30 {
		return content
	}
	return strings.Join(words[:30], " ") + "..."
}