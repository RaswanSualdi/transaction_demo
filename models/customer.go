package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	ID           int
	Name         string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
