package models

import "time"

type User struct {
	ID        int       `gorm:"primaryKey"`
	Username  string    `gorm:"size:50;not null"`
	Email     string    `gorm:"size:100;not null;unique"`
	Password  string    `gorm:"size:255;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}