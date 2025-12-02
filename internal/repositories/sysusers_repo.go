package repositories

import (
	"errors"

	"github.com/google/uuid"
	"github.com/qobilovvv/test_tasks/auth/internal/models"
	"gorm.io/gorm"
)

type SysUserRepository interface {
	Create(user *models.SysUsers) (*models.SysUsers, error)
	AddRoles(userId uuid.UUID, roleIds []uuid.UUID) error
	CheckRoleExists(roleID uuid.UUID) (bool, error)
	GetByPhone(phone string) (*models.SysUsers, error)
	GetByEmail(email string) (*models.SysUsers, error)
	Count() (int64, error)
}

type sysuserRepository struct {
	db *gorm.DB
}

func NewSysUserRepository(db *gorm.DB) *sysuserRepository {
	return &sysuserRepository{db: db}
}

func (r *sysuserRepository) GetByPhone(phone string) (*models.SysUsers, error) {
	var user models.SysUsers
	err := r.db.Where("phone = ? AND status = ?", phone, "active").First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

func (r *sysuserRepository) GetByEmail(email string) (*models.SysUsers, error) {
	var user models.SysUsers
	err := r.db.Where("email = ? AND status = ?", email, "active").First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

func (r *sysuserRepository) Create(user *models.SysUsers) (*models.SysUsers, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *sysuserRepository) AddRoles(userId uuid.UUID, roleIds []uuid.UUID) error {
	for _, roleId := range roleIds {
		suser_role := models.SysUserRoles{
			Id:        uuid.New(),
			SysUserId: userId,
			RoleId:    roleId,
		}
		err := r.db.Create(suser_role).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *sysuserRepository) CheckRoleExists(roleID uuid.UUID) (bool, error) {
	var role models.Role
	err := r.db.Where("id = ? AND status = ?", roleID, "active").First(&role).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, err
}

func (r *sysuserRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&models.SysUsers{}).Where("status = ?", "active").Count(&count).Error
	return count, err
}
