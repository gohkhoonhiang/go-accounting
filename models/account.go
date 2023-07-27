package models

import (
	"gorm.io/gorm"
)

type SavingsAccount struct {
	gorm.Model
	ID           uint           `json:"id" gorm:"primary_key"`
	ShortName    string         `json:"shortName" gorm:"not null"`
	Description  string         `json:"description" gorm:"not null"`
	TotalBalance float64        `json:"totalBalance" gorm:"default:0;not null"`
	Owners       []AccountOwner `gorm:"many2many:savings_account_ownerships"`
}

type ExpenseAccount struct {
	gorm.Model
	ID            uint           `json:"id" gorm:"primary_key"`
	ShortName     string         `json:"shortName" gorm:"not null"`
	Description   string         `json:"description" gorm:"not null"`
	StatementDate int            `json:"statementDate" gorm:"default:1;not null"`
	Owners        []AccountOwner `gorm:"many2many:expense_account_ownerships"`
}

type AccountOwner struct {
	gorm.Model
	ID              uint             `json:"id" gorm:"primary_key"`
	ShortName       string           `json:"shortName" gorm:"not null"`
	Description     string           `json:"description" gorm:"not null"`
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
	SavingsAccountId uint           `json:"savingsAccountId" gorm:"not null"`
	SavingsAccount   SavingsAccount `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Asset struct {
	gorm.Model
	ID              uint          `json:"id" gorm:"primary_key"`
	Name            string        `json:"name" gorm:"not null"`
	Category        string        `json:"category" gorm:"not null"`
	RiskLevel       string        `json:"riskLevel" gorm:"not null"`
	YearlyYield     float64       `json:"yearlyYield" gorm:"default:0;not null"`
	AccountBucketId uint          `json:"accountBucketId" gorm:"not null"`
	AccountBucket   AccountBucket `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
