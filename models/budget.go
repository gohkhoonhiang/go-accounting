package models

import "time"

type Budget struct {
	gorm.Model
	ID                uint      `json:"id" gorm:"primary_key"`
	StartDate         time.Time `json:"startDate" gorm:"not null"`
	EndDate           time.Time `json:"endDate" gorm:"not null"`
	TotalBudgeted     float64   `json:"totalBudgeted" gorm:"default:0;not null"`
	TotalAllocated    float64   `json:"totalAllocated" gorm:"default:0;not null"`
	BalanceToAllocate float64   `json:"balanceToAllocate" gorm:"default:0;not null"`
	CreatedAt         time.Time `json:"createdAt" gorm:"not null"`
	UpdatedAt         time.Time `json:"updatedAt" gorm:"not null"`
}
