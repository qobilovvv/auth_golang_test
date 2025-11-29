package repositories

import (
	"github.com/qobilovvv/test_tasks/auth/internal/models"
	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(role *models.Role) error
	GetAll() ([]models.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleReposity(db *gorm.DB) *roleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) Create(role *models.Role) error {
	return r.db.Create(role).Error
}

func (r *roleRepository) GetAll() ([]models.Role, error) {
	var users []models.Role
	err := r.db.Find(&users).Error
	return users, err
}
