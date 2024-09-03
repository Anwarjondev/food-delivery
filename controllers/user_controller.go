package controllers

import (
    "net/http"
    "food-delivery/models"
    "food-delivery/utils"
    "github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
    claims, _ := c.Get("claims")
    userID := claims.(map[string]interface{})["id"].(uint)

    var user models.User
    if err := utils.DB.Where("id = ?", userID).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, user)
}
