package models

import (
	"time"

	"gorm.io/gorm"
)

type Bank struct {
	gorm.Model
	ID            int
	BankName      string
	AccountNumber string
	Saldo         int64
	MerchantID    uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
