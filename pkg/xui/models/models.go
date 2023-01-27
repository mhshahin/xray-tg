package models

type User struct{}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
