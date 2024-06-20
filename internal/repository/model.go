package repository

import (
	"firebase.google.com/go/v4/auth"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Task struct {
	ID             uuid.UUID `gorm:"primaryKey,type:uuid,default:gen_random_uuid(),not null,unique"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	User           User
	UserID         string
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
	SiteURL        string    `gorm:"type:text;not null"`
	ConditionQuery string    `gorm:"type:text;not null"`
	DurationDay    int       `gorm:"not null"`
	IsPublic       bool      `gorm:"not null"`
	LastCrawledAt  *time.Time
	LastContent    *string        `gorm:"type:text"`
	IsPaused       bool           `gorm:"not null"`
	DeletedAt      gorm.DeletedAt // TODO: add last triggered at
}

type User struct {
	ID         string    `gorm:"primaryKey,not null, unique"`
	CreateAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	Name       string    `gorm:"type:text;not null"`
	Email      string    `gorm:"type:text;not null, unique"`
	PictureURL string    `gorm:"type:text;not null"`
}

func TokenToUserModel(token *auth.Token) User {
	return User{
		ID:         token.UID,
		Name:       token.Claims["name"].(string),
		Email:      token.Claims["email"].(string),
		PictureURL: token.Claims["picture"].(string),
	}
}
