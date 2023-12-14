package models

import (
	"time"

	"gorm.io/gorm"
)

type Merchant struct {
	gorm.Model
	ID           int
	Name         string
	Address      string
	MerchantCode string
	Bank         Bank
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
