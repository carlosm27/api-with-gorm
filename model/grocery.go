package model

import (
	"time"

	"gorm.io/gorm"
)

type Grocery struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
