package model

type LoginRequestDTO struct {
	Email    string `json:"email" validate:"required,min=5,max=150"`
	Password string `json:"password" validate:"required,min=8,max=120"`
}

type SignUpRequestDTO struct {
	Email    string `json:"email" validate:"required,min=5,max=150"`
	Username string `json:"username" validate:"required,min=3,max=120"`
	Password string `json:"password" validate:"required,min=8,max120"`
}
