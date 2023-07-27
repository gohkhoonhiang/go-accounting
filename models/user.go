package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `json:"username" gorm:"not null"`
	Password  string `json:"password" gorm:"not null"` // encrypted
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type UserWithoutPassword struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
