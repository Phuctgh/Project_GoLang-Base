package model

import "github.com/jinzhu/gorm"

// User model represents the user in the system
type User struct {
	gorm.Model
	Username   string `gorm:"unique;not null" json:"username"`
	Password   string `json:"password"`
	Email      string `gorm:"unique;not null" json:"email"`
	FullName   string `json:"full_name"`
	Status     string `gorm:"default:'active'" json:"status"`
	Role       string `gorm:"default:'user'" json:"role"`
	DeleteFlag bool   `gorm:"default:false" json:"delete_flag"`
}
