package controllers

import (
    "net/http"
    "food-delivery/models"
    "food-delivery/utils"
    "github.com/gin-gonic/gin"
	"food-delivery/services"
)

func GetAvailableCouriers(c *gin.Context) {
    var couriers []models.Courier
    if err := utils.DB.Where("status = ?", "available").Find(&couriers).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve couriers"})
        return
    }
    c.JSON(http.StatusOK, couriers)
}


func AssignOrderToCourier(c *gin.Context) {
    var input struct {
        OrderID   uint `json:"order_id"`
        CourierID uint `json:"courier_id"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var courier models.Courier
    if err := utils.DB.First(&courier, input.CourierID).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Courier not found"})
        return
    }

    courier.Status = "busy"
    courier.CurrentOrder = input.OrderID
    if err := utils.DB.Save(&courier).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign order to courier"})
        return
    }

    var order models.Order
    if err := utils.DB.First(&order, input.OrderID).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Order not found"})
        return
    }
    order.Status = "assigned"
    if err := utils.DB.Save(&order).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order status"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Order assigned to courier successfully"})
	services.SendNotification("Order assigned to courier")

    c.JSON(http.StatusOK, gin.H{"message": "Order assigned to courier successfully"})
}
func ListAvailableOrders(c *gin.Context) {
    var orders []models.Order
    if err := utils.DB.Where("status = ?", "available").Find(&orders).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
        return
    }

    c.JSON(http.StatusOK, orders)
}
