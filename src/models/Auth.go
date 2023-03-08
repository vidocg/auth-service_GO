package models

type AuthRequest struct {
	Password string
	Email string
}

type AuthResponse struct {
	Jwt string
	Refresh string
}
