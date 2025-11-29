package repositories

import (
	"github.com/google/uuid"
	"github.com/qobilovvv/test_tasks/auth/internal/models"
	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(role *models.Role) error
	GetAll() ([]models.Role, error)
	GetById(role_id uuid.UUID) (*models.Role, error)
	Update(role_id uuid.UUID, newName string) (*models.Role, error)
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

func (r *roleRepository) GetById(role_id uuid.UUID) (*models.Role, error) {
	var role models.Role
	err := r.db.Where("id = ?", role_id).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, err
}

func (r *roleRepository) GetAll() ([]models.Role, error) {
	var users []models.Role
	err := r.db.Select("id", "name", "created_at").Where("status = ?", models.StatusActive).Find(&users).Error
	return users, err
}

func (r *roleRepository) Update(roleID uuid.UUID, newName string) (*models.Role, error) {
	var role models.Role

	if err := r.db.First(&role, "id = ?", roleID).Error; err != nil {
		return nil, err
	}

	role.Name = newName
	if err := r.db.Model(&role).Where("id = ?", roleID).Update("name", newName).Error; err != nil {
		return nil, err
	}

	return &role, nil
}
