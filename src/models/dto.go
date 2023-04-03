package models

import "github.com/google/uuid"

type UserDto struct {
	ID           uuid.UUID
	Email        string
	RefreshToken string
}
type UserCreateDto struct {
	Password string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"uniqueIndex;not null"`
}
