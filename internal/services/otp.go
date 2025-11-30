package services

import (
	"fmt"
	"log"
	"math/rand"
	"os"
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
		Code:      fmt.Sprintf("%06d", rand.Intn(1000000)),
		Status:    models.StatusUnconfirmed,
		ExpiresAt: time.Now().Add(3 * time.Minute),
	}

	if err := s.repo.Create(otp); err != nil {
		return nil, err
	}

	go func(email, code string) {
        if err := sendEmail(email, code); err != nil {
            log.Println("did not send email:", err)
        }
    }(email, otp.Code)
	return otp, nil
}

func sendEmail(toEmail, code string) error {
	from := os.Getenv("GOOGLE_EMAIL")
	pass := os.Getenv("GOOGLE_PSW")
	fmt.Println("email:", from)
	fmt.Println("pass:", pass)


	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Your OTP Code")
	m.SetBody("text/plain", fmt.Sprintf("Your confirmation code: %s", code))

	d := gomail.NewDialer("smtp.gmail.com", 587, from, pass)

	return d.DialAndSend(m)
}