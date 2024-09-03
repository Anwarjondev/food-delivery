package models


type User struct {
    ID       uint   `json:"id" gorm:"primaryKey"`
    Name     string `json:"name"`
    Email    string `json:"email" gorm:"unique"`
    Password string `json:"pasword"`
    Role     string `json:"user"`
}


type Courier struct {
    ID           uint   `json:"id" gorm:"primaryKey"`
    Name         string `json:"name"`
    Email        string `json:"email" gorm:"unique"`
    Password     string `json:"password"` 
    Role         string `json:"curier"`
    Status       string `json:"status"`
    CurrentOrder uint   `json:"current_order"`
}
