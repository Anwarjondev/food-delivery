package controllers

import (
    "net/http"
    "food-delivery/models"
    "food-delivery/utils"
    "github.com/gin-gonic/gin"
)

func UpdateOrderStatus(c *gin.Context) {
    var input struct {
        Status string `json:"status" binding:"required"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var order models.Order
    if err := utils.DB.First(&order, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }

    order.Status = input.Status
    if err := utils.DB.Save(&order).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order status"})
        return
    }

    NotifyUser(order.UserID, "Your order status has been updated to: "+input.Status)

    c.JSON(http.StatusOK, gin.H{"message": "Order status updated successfully"})
}
