package models

import (
	"time"

	"gorm.io/gorm"
)

type SavingsAccount struct {
	gorm.Model
	ID           uint           `json:"id" gorm:"primary_key"`
	ShortName    string         `json:"shortName" gorm:"not null"`
	Description  string         `json:"description" gorm:"not null"`
	TotalBalance float64        `json:"totalBalance" gorm:"default:0;not null"`
	CreatedAt    time.Time      `json:"createdAt" gorm:"not null"`
	UpdatedAt    time.Time      `json:"updatedAt" gorm:"not null"`
	Owners       []AccountOwner `gorm:"many2many:savings_account_ownerships"`
}

type ExpenseAccount struct {
	gorm.Model
	ID            uint           `json:"id" gorm:"primary_key"`
	ShortName     string         `json:"shortName" gorm:"not null"`
	Description   string         `json:"description" gorm:"not null"`
	StatementDate int            `json:"statementDate" gorm:"default:1;not null"`
	CreatedAt     time.Time      `json:"createdAt" gorm:"not null"`
	UpdatedAt     time.Time      `json:"updatedAt" gorm:"not null"`
	Owners        []AccountOwner `gorm:"many2many:expense_account_ownerships"`
}

type AccountOwner struct {
	gorm.Model
	ID              uint             `json:"id" gorm:"primary_key"`
	ShortName       string           `json:"shortName" gorm:"not null"`
	Description     string           `json:"description" gorm:"not null"`
	CreatedAt       time.Time        `json:"createdAt" gorm:"not null"`
	UpdatedAt       time.Time        `json:"updatedAt" gorm:"not null"`
	UserId          uint             `json:"userId" gorm:"not null"`
	User            User             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ExpenseAccounts []ExpenseAccount `gorm:"many2many:expense_account_ownerships"`
	SavingsAccounts []SavingsAccount `gorm:"many2many:savings_account_ownerships"`
}

type AccountBucket struct {
	gorm.Model
	ID               uint           `json:"id" gorm:"primary_key"`
	Name             string         `json:"name" gorm:"not null"`
	ShortName        string         `json:"shortName" gorm:"not null"`
	Description      string         `json:"description" gorm:"not null"`
	CreatedAt        time.Time      `json:"createdAt" gorm:"not null"`
	UpdatedAt        time.Time      `json:"updatedAt" gorm:"not null"`
	SavingsAccountId uint           `json:"savingsAccountId" gorm:"not null"`
	SavingsAccount   SavingsAccount `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
