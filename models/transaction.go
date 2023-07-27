package models

import (
	"time"

	"gorm.io/gorm"
)

type AccountTransaction struct {
	gorm.Model
	Description     string        `json:"description" gorm:"not null"`
	TransactionDate time.Time     `json:"transactionDate" gorm:"not null"`
	Amount          float64       `json:"amount" gorm:"not null"`
	AccountBucketID uint          `json:"accountBucketId" gorm:"not null"`
	AccountBucket   AccountBucket `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ExpenseTransaction struct {
	gorm.Model
	Description      string         `json:"description" gorm:"not null"`
	TransactionDate  time.Time      `json:"transactionDate" gorm:"not null"`
	Amount           float64        `json:"amount" gorm:"not null"`
	PaidBy           string         `json:"paidBy" gorm:"not null"`
	Category         string         `json:"category" gorm:"not null"`
	Paid             bool           `json:"paid" gorm:"default:false;not null"`
	PaidDate         time.Time      `json:"paidDate"`
	Allocated        bool           `json:"allocated" gorm:"default:false;not null"`
	AllocatedDate    time.Time      `json:"allocatedDate"`
	ExpenseAccountID uint           `json:"expenseAccountId" gorm:"not null"`
	ExpenseAccount   ExpenseAccount `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
