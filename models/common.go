package models

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel ...
type BaseModel struct {
	ID        int            `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
