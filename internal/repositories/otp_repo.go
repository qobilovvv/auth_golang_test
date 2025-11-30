package repositories

import (
	"github.com/google/uuid"
	"github.com/qobilovvv/test_tasks/auth/internal/models"
	"gorm.io/gorm"
)

type OTPRepository interface {
	Create(otp *models.OTP) error
	GetOtp(uuid uuid.UUID, status string) (*models.OTP, error)
	UpdateOtp(otp *models.OTP) error
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

func (r *otpRepository) GetOtp(uuid uuid.UUID, status string) (*models.OTP, error) {
	var otp models.OTP
	err := r.db.Where("id = ? AND status = ?", uuid, status).First(&otp).Error
	if err != nil {
		return nil, err
	}
	return &otp, nil
}


func (r *otpRepository) UpdateOtp(otp *models.OTP) error {
	return r.db.Save(otp).Error
}