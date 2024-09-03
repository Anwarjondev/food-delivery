package main

import (
    "food-delivery/controllers"
    "food-delivery/middleware"
    "food-delivery/utils"
    "github.com/gin-gonic/gin"
)

func main() {
    
    utils.InitDB()

    r := gin.Default()

    
    r.POST("/register", controllers.RegisterUser)
    r.POST("/login", controllers.LoginUser)

    
    secured := r.Group("/api")
    secured.Use(middleware.AuthMiddleware())
    {
        secured.GET("/profile", controllers.GetProfile)
        secured.PUT("/orders/:id/status", controllers.UpdateOrderStatus)
    }

    r.Run(":8080")
}


