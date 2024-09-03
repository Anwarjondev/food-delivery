package controllers

import (
    "fmt"
    "net/http"
    "food-delivery/models"
    "food-delivery/utils"
    "github.com/gin-gonic/gin"
)

func NotifyUser(userID uint, message string) {
    notification := models.Notification{
        UserID:  userID,
        Message: message,
    }

    if err := utils.DB.Create(&notification).Error; err != nil {
        fmt.Println("Failed to create notification:", err)
    }
}

func GetNotifications(c *gin.Context) {
    var notifications []models.Notification
    userID := c.MustGet("userID").(uint)

    if err := utils.DB.Where("user_id = ?", userID).Find(&notifications).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notifications"})
        return
    }

    c.JSON(http.StatusOK, notifications)
}
