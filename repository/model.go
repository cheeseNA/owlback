package repository

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	ID           uuid.UUID `gorm:"primaryKey,type:uuid,default:uuid_generate_v4(),not null,unique"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	SiteURL      string    `gorm:"type:text not null"`
	Condition    string    `gorm:"type:text not null"`
	DurationDays int32     `gorm:"not null"`
}
