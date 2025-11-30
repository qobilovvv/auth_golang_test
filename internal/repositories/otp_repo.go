package repositories

import (
	"github.com/qobilovvv/test_tasks/auth/internal/models"
	"gorm.io/gorm"
)

type OTPRepository interface {
	Create(otp *models.OTP) error
}

type otpRepository struct {
	db *gorm.DB
}

func NewOTPRepository(db *gorm.DB) *otpRepository {
	return &otpRepository{db: db}
}

func (r *otpRepository) Create(otp *models.OTP) error {
	return r.db.Create(otp).Error
}
