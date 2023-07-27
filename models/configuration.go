package models

import (
	"gorm.io/gorm"
)

type Configuration struct {
	gorm.Model
	FieldName  string `json:"fieldName" gorm:"not null"`
	FieldKey   string `json:"fieldKey" gorm:"not null"`
	FieldValue string `json:"fieldValue" gorm:"not null"`
}
