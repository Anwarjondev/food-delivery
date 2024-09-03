package controllers

import (
    "net/http"
    "food-delivery/models"
    "food-delivery/utils"
    "github.com/gin-gonic/gin"
)
func GetAllOrders(c *gin.Context) {
    var orders []models.Order
    if err := utils.DB.Find(&orders).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
        return
    }
    c.JSON(http.StatusOK, orders)
}

func AdminUpdateOrderStatus(c *gin.Context) {
    var input struct {
        OrderID uint   `json:"order_id"`
        Status  string `json:"status"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var order models.Order
    if err := utils.DB.First(&order, input.OrderID).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Order not found"})
        return
    }

    order.Status = input.Status
    if err := utils.DB.Save(&order).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order status"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Order status updated successfully"})
}
