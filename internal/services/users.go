package services

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/qobilovvv/test_tasks/auth/internal/errors"
	"github.com/qobilovvv/test_tasks/auth/internal/models"
	"github.com/qobilovvv/test_tasks/auth/internal/repositories"
	"github.com/qobilovvv/test_tasks/auth/pkg/helpers"
)

type UserService interface {
	SignUpUser(token, email, name, password string) (string, error)
	Login(identifier, password, user_type string) (string, error)
}

type userService struct {
	userRepo    repositories.UserRepository
	otpRepo     repositories.OTPRepository
	sysUserRepo repositories.SysUserRepository
}

func NewUserService(
	userRepo repositories.UserRepository,
	otpRepo repositories.OTPRepository,
	sysUserRepo repositories.SysUserRepository,

) *userService {
	return &userService{userRepo: userRepo, otpRepo: otpRepo, sysUserRepo: sysUserRepo}
}

func (s *userService) SignUpUser(token, email, name, password string) (string, error) {
	otpIDStr, exp, err := helpers.DecodeJwtOtpToken(token)
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

func (s *userService) Login(identifier, password, user_type string) (string, error) {
	if user_type == "sysuser" {
		user, err := s.sysUserRepo.GetByPhone(identifier)
		if err != nil || user == nil {
			return "", errors.ErrInvalidCredentials
		}

		if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
			return "", errors.ErrInvalidCredentials
		}

		token, err := helpers.GenerateAccessToken(user.Id.String(), user_type, time.Minute*30)
		if err != nil {
			return "", err
		}
		return token, nil
	}

	// For normal users
	usr, err := s.userRepo.GetActiveUser(identifier)
	if err != nil || usr == nil {
		return "", errors.ErrInvalidCredentials
	}

	if bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(password)) != nil {
		return "", errors.ErrInvalidCredentials
	}

	token, err := helpers.GenerateAccessToken(usr.Id.String(), user_type, time.Minute*30)
	if err != nil {
		return "", err
	}
	return token, nil
}
