package models

import (
	"time"

	"gorm.io/gorm"
)

type StockHolding struct {
	gorm.Model
	Ticker string `json:"ticker" gorm:"not null"`
	Name   string `json:"name" gorm:"not null"`
}

type StockPrice struct {
	gorm.Model
	PriceDate      time.Time    `json:"priceDate" gorm:"not null"`
	UnitPrice      float64      `json:"unitPrice" gorm:"default:0;not null"`
	StockHoldingID uint         `json:"stockHoldingId" gorm:"not null"`
	StockHolding   StockHolding `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type StockDividend struct {
	gorm.Model
	ExDividendDate time.Time    `json:"exDividendDate" gorm:"not null"`
	UnitValue      float64      `json:"unitValue" gorm:"default:0;not null"`
	PayoutDate     time.Time    `json:"payoutDate" gorm:"not null"`
	StockHoldingID uint         `json:"stockHoldingId" gorm:"not null"`
	StockHolding   StockHolding `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type StockBrokerageAccount struct {
	gorm.Model
	ShortName   string `json:"shortName" gorm:"not null"`
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
}

type StockBrokerageAccountFee struct {
	gorm.Model
	CommissionPercentage    float64               `json:"commissionPercentage" gorm:"default:0;not null"`
	MinimumCommission       float64               `json:"minimumCommission" gorm:"default:0;not null"`
	EffectiveDate           time.Time             `json:"effectiveDate" gorm:"not null"`
	ExpiryDate              time.Time             `json:"expiryDate" gorm:"not null"`
	MinimumTradingSum       float64               `json:"minimumTradingSum" gorm:"default:0;not null"`
	MaximumTradingSum       float64               `json:"maximumTradingSum" gorm:"default:0;not null"`
	TradingCurrency         string                `json:"tradingCurrency" gorm:"not null"`
	StockBrokerageAccountID uint                  `json:"stockBrokerageAccountId" gorm:"not null"`
	StockBrokerageAccount   StockBrokerageAccount `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type StockPurchaseOrder struct {
	gorm.Model
	PurchaseDate            time.Time             `json:"purchaseDate" gorm:"not null"`
	PurchasePrice           float64               `json:"purchasePrice" gorm:"default:0;not null"`
	ShareCount              int32                 `json:"shareCount" gorm:"default:0;not null"`
	ContractSum             float64               `json:"contractSum" gorm:"default:0;not null"`
	Reinvest                bool                  `json:"reinvest" gorm:"default:false;not null"`
	StockHoldingID          uint                  `json:"stockHoldingId" gorm:"not null"`
	StockHolding            StockHolding          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StockBrokerageAccountID uint                  `json:"stockBrokerageAccountId" gorm:"not null"`
	StockBrokerageAccount   StockBrokerageAccount `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type StockSaleOrder struct {
	gorm.Model
	SaleDate                time.Time             `json:"saleDate" gorm:"not null"`
	SalePrice               float64               `json:"salePrice" gorm:"default:0;not null"`
	ShareCount              int32                 `json:"shareCount" gorm:"default:0;not null"`
	ContractSum             float64               `json:"contractSum" gorm:"default:0;not null"`
	StockHoldingID          uint                  `json:"stockHoldingId" gorm:"not null"`
	StockHolding            StockHolding          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StockBrokerageAccountID uint                  `json:"stockBrokerageAccountId" gorm:"not null"`
	StockBrokerageAccount   StockBrokerageAccount `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
