package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Task struct {
	ID             uuid.UUID `gorm:"primaryKey,type:uuid,default:uuid_generate_v4(),not null,unique"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	CreatedBy      uuid.UUID `gorm:"type:uuid not null"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
	SiteURL        string    `gorm:"type:text not null"`
	ConditionQuery string    `gorm:"type:text not null"`
	DurationDay    int32     `gorm:"not null"`
	IsPublic       bool      `gorm:"not null"`
	DeletedAt      gorm.DeletedAt
}
