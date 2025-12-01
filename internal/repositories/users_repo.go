package repositories

import (
	"github.com/qobilovvv/test_tasks/auth/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.Users) (*models.Users, error)
	GetActiveUser(email string) (*models.Users, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.Users) (*models.Users, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetActiveUser(email string) (*models.Users, error) {
	var user models.Users
	err := r.db.Where("email = ? AND status = ?", email, "active").First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}