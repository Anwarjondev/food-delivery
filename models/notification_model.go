package models

import "time"

type Notification struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    UserID    uint      `json:"user_id"`
    Message   string    `json:"message"`
    CreatedAt time.Time `json:"created_at"`
}

