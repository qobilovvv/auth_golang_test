package repositories

import (
	"github.com/qobilovvv/test_tasks/auth/internal/models"
	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(user *models.Users) error
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleReposity(db *gorm.DB) *roleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) Create(user *models.Users) error {
	return r.db.Create(&user).Error
}
