package controllers

import (
    "net/http"
    "food-delivery/models"
    "food-delivery/utils"
    "github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
    var input models.Product
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := utils.DB.Create(&input).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Product created successfully", "product": input})
}

func GetProducts(c *gin.Context) {
    var products []models.Product
    if err := utils.DB.Find(&products).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"products": products})
}
