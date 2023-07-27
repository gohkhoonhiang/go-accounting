package models

import (
	"time"

	"gorm.io/gorm"
)

type Budget struct {
	gorm.Model
	StartDate         time.Time `json:"startDate" gorm:"not null"`
	EndDate           time.Time `json:"endDate" gorm:"not null"`
	TotalBudgeted     float64   `json:"totalBudgeted" gorm:"default:0;not null"`
	TotalAllocated    float64   `json:"totalAllocated" gorm:"default:0;not null"`
	BalanceToAllocate float64   `json:"balanceToAllocate" gorm:"default:0;not null"`
}

type BudgetLine struct {
	gorm.Model
	Category       string  `json:"category" gorm:"not null"`
	SubCategory    string  `json:"subCategory" gorm:"not null"`
	Remarks        string  `json:"remarks"`
	BudgetedAmount float64 `json:"budgetedAmount" gorm:"default:0;not null"`
	ActualAmount   float64 `json:"actualAmount" gorm:"default:0;not null"`
	BudgetID       uint    `json:"budgetId" gorm:"not null"`
	Budget         Budget  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
