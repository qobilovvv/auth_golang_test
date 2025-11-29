package services

import (
	"github.com/google/uuid"
	"github.com/qobilovvv/test_tasks/auth/internal/models"
	"github.com/qobilovvv/test_tasks/auth/internal/repositories"
)

type RoleService interface {
	CreateRole(name string) (*models.Role, error)
	GetAll() ([]models.Role, error)
	UpdateRole(roleID uuid.UUID, name string) (*models.Role, error)
}

type roleService struct {
	repo repositories.RoleRepository
}

func NewRoleService(repo repositories.RoleRepository) *roleService {
	return &roleService{repo: repo}
}

func (s *roleService) CreateRole(name string) (*models.Role, error) {
	role := &models.Role{Id: uuid.New(), Name: name}
	if err := s.repo.Create(role); err != nil {
		return nil, err
	}
	return role, nil
}

func (s *roleService) GetAll() ([]models.Role, error) {
	return s.repo.GetAll()
}

func (s *roleService) UpdateRole(roleID uuid.UUID, name string) (*models.Role, error) {
	return s.repo.Update(roleID, name)
}