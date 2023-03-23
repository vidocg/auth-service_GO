package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Password     string    `gorm:"type:varchar(255);not null"`
	Email        string    `gorm:"uniqueIndex;not null"`
	RefreshToken string
}
