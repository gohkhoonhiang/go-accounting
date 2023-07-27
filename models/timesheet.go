package models

import (
	"time"

	"gorm.io/gorm"
)

type Timesheet struct {
	gorm.Model
	StartDate           time.Time `json:"startDate" gorm:"not null"`
	EndDate             time.Time `json:"endDate" gorm:"not null"`
	HoursWorked         float64   `json:"hoursWorked" gorm:"default:0;not null"`
	HoursToClock        float64   `json:"hoursToClock" gorm:"default:0;not null"`
	AverageHoursToClock float64   `json:"averageHoursToClock" gorm:"default:0;not null"`
	ActualEarnings      float64   `json:"actualEarnings" gorm:"-:migration"`
	ActualTakeHome      float64   `json:"actualTakeHome" gorm:"-:migration"`
}

type TimesheetLine struct {
	gorm.Model
	EntryDate    time.Time `json:"entryDate" gorm:"not null"`
	ClockedHours float64   `json:"clockedHours" gorm:"default:0;not null"`
	TimesheetID  uint      `json:"budgetId" gorm:"not null"`
	Timesheet    Timesheet `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
