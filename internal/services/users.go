package services

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/qobilovvv/test_tasks/auth/internal/config"
	"github.com/qobilovvv/test_tasks/auth/internal/errors"
	"github.com/qobilovvv/test_tasks/auth/internal/models"
	"github.com/qobilovvv/test_tasks/auth/internal/repositories"
)

type UserService interface {
	SignUpUser(token, email, name, password string) (string, error)
}

type userService struct {
	userRepo repositories.UserRepository
	otpRepo  repositories.OTPRepository
}

func NewUserService(
	userRepo repositories.UserRepository,
	otpRepo repositories.OTPRepository,
) *userService {
	return &userService{userRepo: userRepo, otpRepo: otpRepo}
}

func (s *userService) SignUpUser(token, email, name, password string) (string, error) {
	otpIDStr, exp, err := config.DecodeOtpToken(token)
    if !strings.Contains(email, "@") {
		return "", errors.ErrInvalidEmail
	}

	if err != nil {
		return "", errors.ErrOtpUnavailable
	}

	if time.Now().After(exp) {
		return "", errors.ErrOtpTokenExpired
	}

	otpID, err := uuid.Parse(otpIDStr)
	if err != nil {
		return "", errors.ErrOtpUnavailable
	}

	otp, err := s.otpRepo.GetOtpWithEmail(otpID, email)
	if err != nil || otp == nil {
		return "", errors.ErrOtpUnavailable
	}

	usr, err := s.userRepo.GetActiveUser(email)
	if err == nil && usr != nil {
		return "", errors.ErrUserAlreadyExists
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	newUser := &models.Users{
		Id:       uuid.New(),
		Email:    email,
		Name:     name,
		Password: string(hashed),
		Status:   "active",
	}

	created_user, err := s.userRepo.Create(newUser)
	if err != nil {
		return "", errors.ErrUserCreationFailed
	}
	return created_user.Id.String(), nil
}
