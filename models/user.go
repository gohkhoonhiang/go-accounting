package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"not null"`
	Password string `json:"password" gorm:"not null"` // encrypted
}
