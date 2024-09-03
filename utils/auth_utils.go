package utils

import (
    "time"
    "github.com/golang-jwt/jwt"
)

func GenerateToken(userID uint, email string, role string) (string, error) {
    claims := jwt.MapClaims{
        "id":    userID,
        "email": email,
        "role":  role,
        "exp":   time.Now().Add(time.Minute * 60).Unix(), 
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte("your-secret-key"))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
