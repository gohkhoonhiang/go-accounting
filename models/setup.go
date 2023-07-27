package models

import (
	"log"
	"os"

	"github.com/go-accounting/utils"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConfig map[string]string

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConfig := DbConfig{
		"host":     os.Getenv("DATABASE_HOST"),
		"port":     os.Getenv("DATABASE_PORT"),
		"user":     os.Getenv("DATABASE_USER"),
		"password": os.Getenv("DATABASE_PASS"),
		"dbname":   os.Getenv("DATABASE_NAME"),
		"Timezone": os.Getenv("TIMEZONE"),
		"sslmode":  "disable",
	}

	dsn := utils.MapToString(dbConfig, " ")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = db.AutoMigrate(
		&User{},
		&SavingsAccount{}, &ExpenseAccount{}, &AccountOwner{},
		&AccountBucket{}, &Asset{},
		&AccountTransaction{}, &ExpenseTransaction{},
		&Budget{}, &BudgetLine{},
		&StockBrokerageAccount{}, &StockBrokerageAccountFee{},
		&StockHolding{}, &StockPrice{}, &StockDividend{},
		&StockPurchaseOrder{}, &StockSaleOrder{},
	)

	if err != nil {
		return
	}

	DB = db
}
