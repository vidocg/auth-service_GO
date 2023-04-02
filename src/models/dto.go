package models

import "github.com/google/uuid"

type UserDto struct {
	ID           uuid.UUID
	Email        string
	RefreshToken string
}
