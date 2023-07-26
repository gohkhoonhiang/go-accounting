package models

import "time"

type Budget struct {
	ID                uint      `json:"id" gorm:"primary_key"`
	StartDate         time.Time `json:"startDate"`
	EndDate           time.Time `json:"endDate"`
	TotalBudgeted     float64   `json:"totalBudgeted"`
	TotalAllocated    float64   `json:"totalAllocated"`
	BalanceToAllocate float64   `json:"balanceToAllocate"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}
