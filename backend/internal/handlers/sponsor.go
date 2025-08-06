/*
开发心理过程：
1. 实现微信支付集成，提供便捷的赞助功能
2. 使用模拟的二维码生成，实际生产环境需要调用微信支付API
3. 处理订单状态管理和支付回调
4. 提供赞助者列表展示功能
5. 考虑安全性和错误处理
*/

package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"techblog-api/backend/internal/models"
)

type SponsorHandler struct {
	db *gorm.DB
}

func NewSponsorHandler(db *gorm.DB) *SponsorHandler {
	return &SponsorHandler{db: db}
}

// CreateSponsorOrder 创建赞助订单
func (h *SponsorHandler) CreateSponsorOrder(c *gin.Context) {
	var req models.SponsorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "请求参数有误",
			Error:   err.Error(),
		})
		return
	}

	// 生成唯一订单号
	orderID := generateOrderID()

	// 创建订单记录
	order := models.SponsorOrder{
		OrderID:       orderID,
		Amount:        req.Amount,
		SponsorName:   req.SponsorName,
		Message:       req.Message,
		PaymentMethod: req.PaymentMethod,
		Status:        "pending",
	}

	if order.SponsorName == "" {
		order.SponsorName = "匿名赞助者"
	}

	if err := h.db.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "创建订单失败",
			Error:   err.Error(),
		})
		return
	}

	// 生成支付二维码（模拟）
	qrCodeURL := generateMockQRCode(orderID, req.Amount)

	// 更新订单二维码URL
	h.db.Model(&order).Update("qr_code_url", qrCodeURL)

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "订单创建成功",
		Data: models.SponsorResponse{
			OrderID: orderID,
			QRCode:  qrCodeURL,
			Amount:  req.Amount,
		},
	})
}

// GetOrderStatus 查询订单状态
func (h *SponsorHandler) GetOrderStatus(c *gin.Context) {
	orderID := c.Param("orderId")

	var order models.SponsorOrder
	if err := h.db.Where("order_id = ?", orderID).First(&order).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, models.APIResponse{
				Success: false,
				Message: "订单不存在",
				Error:   "order_not_found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "查询订单失败",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "查询成功",
		Data: gin.H{
			"orderId": order.OrderID,
			"status":  order.Status,
			"amount":  order.Amount,
			"paidAt":  order.PaidAt,
		},
	})
}

// GetSponsorList 获取赞助者列表
func (h *SponsorHandler) GetSponsorList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	var orders []models.SponsorOrder
	var total int64

	// 只查询已支付的订单
	query := h.db.Model(&models.SponsorOrder{}).Where("status = ?", "paid")
	
	// 统计总数
	query.Count(&total)

	// 分页查询
	offset := (page - 1) * limit
	if err := query.Order("paid_at DESC").Offset(offset).Limit(limit).Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "查询赞助列表失败",
			Error:   err.Error(),
		})
		return
	}

	// 计算分页信息
	totalPages := int((total + int64(limit) - 1) / int64(limit))

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "查询成功",
		Data:    orders,
		Meta: &models.PaginationMeta{
			Page:       page,
			Limit:      limit,
			Total:      total,
			TotalPages: totalPages,
		},
	})
}

// GetSponsorStats 获取赞助统计
func (h *SponsorHandler) GetSponsorStats(c *gin.Context) {
	var totalAmount float64
	var totalCount int64
	var monthlyCount int64

	// 查询总赞助金额和次数
	h.db.Model(&models.SponsorOrder{}).Where("status = ?", "paid").
		Select("COALESCE(SUM(amount), 0)").Scan(&totalAmount)
	
	h.db.Model(&models.SponsorOrder{}).Where("status = ?", "paid").Count(&totalCount)

	// 查询本月赞助次数
	now := time.Now()
	firstOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	h.db.Model(&models.SponsorOrder{}).
		Where("status = ? AND paid_at >= ?", "paid", firstOfMonth).
		Count(&monthlyCount)

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "查询成功",
		Data: gin.H{
			"totalAmount":   totalAmount,
			"totalCount":    totalCount,
			"monthlyCount":  monthlyCount,
		},
	})
}

// MockPaymentCallback 模拟支付回调（用于测试）
func (h *SponsorHandler) MockPaymentCallback(c *gin.Context) {
	orderID := c.Param("orderId")
	action := c.DefaultQuery("action", "pay") // pay or cancel

	var order models.SponsorOrder
	if err := h.db.Where("order_id = ?", orderID).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, models.APIResponse{
			Success: false,
			Message: "订单不存在",
		})
		return
	}

	now := time.Now()
	
	switch action {
	case "pay":
		// 模拟支付成功
		order.Status = "paid"
		order.PaidAt = &now
		order.TransactionID = generateTransactionID()
		
	case "cancel":
		// 模拟支付取消
		order.Status = "cancelled"
	default:
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "无效的操作",
		})
		return
	}

	if err := h.db.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "更新订单状态失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: fmt.Sprintf("订单已%s", map[string]string{"pay": "支付", "cancel": "取消"}[action]),
		Data: gin.H{
			"orderId": order.OrderID,
			"status":  order.Status,
		},
	})
}

// 生成订单号
func generateOrderID() string {
	timestamp := time.Now().Unix()
	bytes := make([]byte, 4)
	rand.Read(bytes)
	random := hex.EncodeToString(bytes)
	return fmt.Sprintf("SP%d%s", timestamp, random)
}

// 生成交易号
func generateTransactionID() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

// 生成模拟二维码（实际应调用微信支付API）
func generateMockQRCode(orderID string, amount float64) string {
	// 这里返回一个模拟的二维码URL
	// 实际项目中应该调用微信支付统一下单接口获取真实的支付二维码
	return fmt.Sprintf("https://api.qrserver.com/v1/create-qr-code/?size=200x200&data=mock_payment_%s_%.2f", orderID, amount)
}