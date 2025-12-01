package services

import (
	"github.com/google/uuid"
	"github.com/qobilovvv/test_tasks/auth/internal/errors"
	"github.com/qobilovvv/test_tasks/auth/internal/models"
	"github.com/qobilovvv/test_tasks/auth/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type SysUserService interface {
	CreateSysUser(name, phone, password string, roleIDs []uuid.UUID) (uuid.UUID, error)
}

type sysUserService struct {
	repo repositories.SysUserRepository
}

func NewSysUserService(repo repositories.SysUserRepository) SysUserService {
	return &sysUserService{repo: repo}
}

func (s *sysUserService) CreateSysUser(name, phone, password string, roleIDs []uuid.UUID) (uuid.UUID, error) {
	existingUser, err := s.repo.GetByPhone(phone)
	if err != nil {
		return uuid.Nil, err
	}
	if existingUser != nil {
		return uuid.Nil, errors.ErrSysUserAlreadyExists
	}

	for _, roleID := range roleIDs {
		exists, err := s.repo.CheckRoleExists(roleID)
		if err != nil {
			return uuid.Nil, err
		}
		if !exists {
			return uuid.Nil, errors.ErrRoleNotFound
		}
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return uuid.Nil, err
	}

	user := &models.SysUsers{
		Id:       uuid.New(),
		Name:     name,
		Phone:    phone,
		Password: string(hashed),
		Status:   "active",
	}

	createdUser, err := s.repo.Create(user)
	if err != nil {
		return uuid.Nil, err
	}

	err = s.repo.AddRoles(createdUser.Id, roleIDs)
	if err != nil {
		return uuid.Nil, err
	}

	return createdUser.Id, nil
}
