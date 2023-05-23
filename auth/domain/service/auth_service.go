package service

import (
	"heptaber/auth/domain/model"

	"github.com/google/uuid"
)

type IAuthService interface {
	Signup(signUpRequestDTO model.SignUpRequestDTO) (model.User, error)
	Login(loginRequestDTO model.LoginRequestDTO) (string, string, error)
	Logout(userId string) error
	GetNewAccessToken(userId string) (string, error)
	VerifyUser(code uuid.UUID) error
}
