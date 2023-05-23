package service

import (
	"encoding/json"
	"errors"
	"heptaber/auth/app/helper"
	"heptaber/auth/domain/model"
	"heptaber/auth/domain/repository"
	"heptaber/auth/infrastructure/rabbitmq"

	"log"
	"time"

	"github.com/google/uuid"
)

type authService struct {
	repository.IUserRepository
	repository.IVerificationCodeRepository
	repository.ITokenRepository
}

func NewAuthService(
	userRepository repository.IUserRepository,
	vCodeRepository repository.IVerificationCodeRepository,
	tokenRepository repository.ITokenRepository) *authService {
	return &authService{
		IUserRepository:             userRepository,
		IVerificationCodeRepository: vCodeRepository,
		ITokenRepository:            tokenRepository,
	}
}

func (as *authService) Signup(signUpRequestDTO model.SignUpRequestDTO) (savedUser model.User, err error) {
	currentTime := helper.GetUTCCurrentTimeRFC3339()
	savedUser = model.User{
		Email:      signUpRequestDTO.Email,
		Password:   helper.HashPassword(signUpRequestDTO.Password),
		Username:   signUpRequestDTO.Username,
		Role:       model.REGULAR,
		CreatedAt:  currentTime,
		UpdatedAt:  currentTime,
		IsVerified: false,
	}

	if err := as.IUserRepository.Save(&savedUser); err != nil {
		return model.User{}, err
	}

	vCode := model.VerificationCode{
		CreatedAt: currentTime,
		ExpiresAt: currentTime.Add(time.Hour * time.Duration(24)),
		UserID:    savedUser.ID,
	}
	err = as.IVerificationCodeRepository.Save(&vCode)
	if err != nil {
		log.Fatal("error while saving verification code")
	}

	vEmail := model.VerificationEmailNotification{
		Recipient:        savedUser.Email,
		VerificationCode: vCode.ID.String(),
	}
	jsonData, err := json.Marshal(vEmail)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %s", err.Error())
		return
	}
	// should it be used in goroutine?
	sendMessageRabbit(jsonData)

	return
}

func (as *authService) Login(loginRequestDTO model.LoginRequestDTO) (string, string, error) {
	var foundUser *model.User
	foundUser, err := as.IUserRepository.FindByEmail(loginRequestDTO.Email)
	if err != nil {
		err = errors.New("invalid email or password provided")
		return "", "", err
	}

	if !foundUser.IsVerified {
		err = errors.New("account not yet verified")
		return "", "", err
	}

	if err = helper.VerifyPassword(foundUser.Password, loginRequestDTO.Password); err != nil {
		err = errors.New("invalid email or password provided")
		return "", "", err
	}

	currentTime := time.Now().UTC()
	accessToken, err := helper.GenerateAccessToken(*foundUser, currentTime)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := helper.GenerateRefreshToken(foundUser.ID, currentTime)
	if err != nil {
		log.Fatal("error on generating refresh token")
	}

	token := &model.Token{
		UserID:       foundUser.ID,
		RefreshToken: refreshToken,
		ExpiresAt:    currentTime.Add(time.Hour * time.Duration(4320)),
		CreatedAt:    currentTime,
		UpdatedAt:    currentTime,
	}

	as.ITokenRepository.UpdateToken(token)

	return accessToken, refreshToken, nil
}

func (as *authService) Logout(userId string) error {
	uid := uuid.MustParse(userId)
	return as.ITokenRepository.DeleteByUserID(uid)
}

func (as *authService) GetNewAccessToken(userId string) (string, error) {
	user, err := as.IUserRepository.FindByID(uuid.MustParse(userId))
	if err != nil {
		return "", err
	}
	accessToken, err := helper.GenerateAccessToken(*user, time.Now().UTC())
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (as *authService) VerifyUser(codeId uuid.UUID) error {
	foundCode, err := as.IVerificationCodeRepository.FindById(codeId)
	if err != nil {
		err = errors.New("no such code was found")
		return err
	}
	if foundCode.ExpiresAt.Unix() < time.Now().UTC().Unix() {
		err = errors.New("verification code expired")
		return err
	}

	if err = as.IUserRepository.SetUserVerifiedTrueByID(foundCode.UserID); err != nil {
		log.Fatal("error while updating user on verification email")
		return err
	}
	if err = as.IVerificationCodeRepository.DeleteById(codeId); err != nil {
		log.Fatal("error while deleting verification code")
	}

	return nil
}

func sendMessageRabbit(jsonData []byte) {
	rmq, err := rabbitmq.NewRabbitMQ()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err.Error())
		return
	}
	defer rmq.CloseConnection()

	rmq.SetUpForNotification()

	err = rmq.PublishMessage(
		rabbitmq.NotificationExchangeName,
		rabbitmq.VerificaitonEmailRoutingKey,
		jsonData,
	)
	if err != nil {
		log.Fatalf("Failed to publish message to RabbitMQ: %s", err.Error())
	}
}
