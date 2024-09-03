package controllers

import (
    "net/http"
    "food-delivery/models"
    "food-delivery/utils"
    "time"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt"
    "golang.org/x/crypto/bcrypt"
)

func LoginUser(c *gin.Context) {
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    if err := utils.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id":    user.ID,
        "email": user.Email,
        "role":  user.Role,
        "exp":   time.Now().Add(time.Hour * 72).Unix(),
    })

    tokenString, err := token.SignedString([]byte("your-secret-key"))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
func GenerateToken(userID uint, email string, role string) (string, error) {
    claims := jwt.MapClaims{
        "id":    userID,
        "email": email,
        "role":  role,
        "exp":   time.Now().Add(time.Hour * 72).Unix(), // Token expires in 72 hours
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte("your-secret-key"))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
