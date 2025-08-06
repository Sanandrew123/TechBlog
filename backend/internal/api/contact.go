package api

import (
	"net/http"
	"strconv"
	"math"
	
	"github.com/gin-gonic/gin"
	"techblog-api/backend/internal/database"
	"techblog-api/backend/internal/models"
)

type ContactHandler struct{}

func NewContactHandler() *ContactHandler {
	return &ContactHandler{}
}

// SubmitContact 提交联系消息
func (h *ContactHandler) SubmitContact(c *gin.Context) {
	var message models.ContactMessage
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Invalid request data",
			Error:   err.Error(),
		})
		return
	}
	
	db := database.GetDB()
	
	if err := db.Create(&message).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "Failed to submit message",
			Error:   err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusCreated, models.APIResponse{
		Success: true,
		Message: "Message submitted successfully",
		Data:    message,
	})
}

// GetMessages 获取联系消息列表（需要管理员权限）
func (h *ContactHandler) GetMessages(c *gin.Context) {
	// 查询参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	isRead := c.Query("is_read")
	isReplied := c.Query("is_replied")
	
	db := database.GetDB()
	var messages []models.ContactMessage
	var total int64
	
	// 构建查询
	query := db.Model(&models.ContactMessage{})
	
	// 过滤条件
	if isRead != "" {
		if isReadBool, err := strconv.ParseBool(isRead); err == nil {
			query = query.Where("is_read = ?", isReadBool)
		}
	}
	
	if isReplied != "" {
		if isRepliedBool, err := strconv.ParseBool(isReplied); err == nil {
			query = query.Where("is_replied = ?", isRepliedBool)
		}
	}
	
	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "Failed to count messages",
			Error:   err.Error(),
		})
		return
	}
	
	// 分页查询
	offset := (page - 1) * limit
	if err := query.Order("created_at DESC").Offset(offset).Limit(limit).Find(&messages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "Failed to fetch messages",
			Error:   err.Error(),
		})
		return
	}
	
	// 计算分页信息
	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	
	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Messages fetched successfully",
		Data:    messages,
		Meta: &models.PaginationMeta{
			Page:       page,
			Limit:      limit,
			Total:      total,
			TotalPages: totalPages,
		},
	})
}

// GetMessage 获取单个联系消息（需要管理员权限）
func (h *ContactHandler) GetMessage(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Invalid message ID",
			Error:   "invalid_id",
		})
		return
	}
	
	db := database.GetDB()
	var message models.ContactMessage
	
	if err := db.First(&message, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.APIResponse{
			Success: false,
			Message: "Message not found",
			Error:   "message_not_found",
		})
		return
	}
	
	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Message fetched successfully",
		Data:    message,
	})
}

// MarkAsRead 标记消息为已读（需要管理员权限）
func (h *ContactHandler) MarkAsRead(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Invalid message ID",
			Error:   "invalid_id",
		})
		return
	}
	
	db := database.GetDB()
	
	if err := db.Model(&models.ContactMessage{}).Where("id = ?", id).Update("is_read", true).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "Failed to mark message as read",
			Error:   err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Message marked as read",
	})
}

// MarkAsReplied 标记消息为已回复（需要管理员权限）
func (h *ContactHandler) MarkAsReplied(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Invalid message ID",
			Error:   "invalid_id",
		})
		return
	}
	
	db := database.GetDB()
	
	if err := db.Model(&models.ContactMessage{}).Where("id = ?", id).Updates(map[string]interface{}{
		"is_read":    true,
		"is_replied": true,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "Failed to mark message as replied",
			Error:   err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Message marked as replied",
	})
}

// DeleteMessage 删除联系消息（需要管理员权限）
func (h *ContactHandler) DeleteMessage(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Invalid message ID",
			Error:   "invalid_id",
		})
		return
	}
	
	db := database.GetDB()
	
	if err := db.Delete(&models.ContactMessage{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "Failed to delete message",
			Error:   err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Message deleted successfully",
	})
}