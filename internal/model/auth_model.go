package model

type LoginAuthRequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterAuthRequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
