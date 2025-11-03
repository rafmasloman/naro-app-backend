package model

import "time"

type UserResponse struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserRequestDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
