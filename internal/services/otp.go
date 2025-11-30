package services

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/qobilovvv/test_tasks/auth/internal/models"
	"github.com/qobilovvv/test_tasks/auth/internal/repositories"
	"gopkg.in/gomail.v2"
)

type OTPService interface {
	SendOTP(email string) (*models.OTP, error)
}

type otpService struct {
	repo repositories.OTPRepository
}

func NewOTPService(repo repositories.OTPRepository) *otpService {
	return &otpService{repo: repo}
}

func (s *otpService) SendOTP(email string) (*models.OTP, error) {
	otp := &models.OTP{
		Id:        uuid.New(),
		Email:     email,
		Code:      fmt.Sprintf("%06d", rand.Intn(1000000)), // 6-digit code
		Status:    models.StatusUnconfirmed,
		ExpiresAt: time.Now().Add(3 * time.Minute),
	}

	// Save to DB
	if err := s.repo.Create(otp); err != nil {
		return nil, err
	}

	// Send email (simple example)
	if err := sendEmail(email, otp.Code); err != nil {
		return nil, err
	}

	return otp, nil
}

// Simple email sender using gomail
func sendEmail(toEmail, code string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "qobilovvodil@gmail.com")
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Your OTP Code")
	m.SetBody("text/plain", fmt.Sprintf("Your confirmation code: %s", code))

	d := gomail.NewDialer("smtp.example.com", 587, "username", "password")

	return d.DialAndSend(m)
}
