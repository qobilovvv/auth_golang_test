package tests

import (
	"testing"

	"github.com/google/uuid"
	"github.com/qobilovvv/test_tasks/auth/internal/models"
	"github.com/qobilovvv/test_tasks/auth/internal/repositories"
	"github.com/qobilovvv/test_tasks/auth/internal/services"
)

func newRoleService() services.RoleService {
	repo := repositories.NewRoleReposity(testDB)
	return services.NewRoleService(repo)
}

func TestGetAllRoles(t *testing.T) {
	service := newRoleService()

	rolesToInsert := []models.Role{
		{Id: uuid.New(), Name: "admin"},
		{Id: uuid.New(), Name: "user"},
		{Id: uuid.New(), Name: "waiter"},
	}

	if err := testDB.Create(&rolesToInsert).Error; err != nil {
		t.Fatal(err)
	}

	roles, err := service.GetAll()
	if err != nil {
		t.Fatalf("something went wrong, error: %v", err)
	}

	if roles[0].Name != "admin" && roles[1].Name != "admin" {
		t.Fatalf("expected admin role")
	}
}

func TestGetAllRoles_Fail(t *testing.T) {
	service := newRoleService()

	rolesToInsert := []models.Role{
		{Id: uuid.New(), Name: "admin"},
		{Id: uuid.New(), Name: "user"},
		{Id: uuid.New(), Name: "waiter"},
	}

	if err := testDB.Create(&rolesToInsert).Error; err != nil {
		t.Fatal(err)
	}

	roles, err := service.GetAll()
	if err != nil {
		t.Fatalf("something went wrong, error: %v", err)
	}

	if roles[0].Name != "admin" && roles[1].Name != "admin" {
		t.Fatalf("expected admin role")
	}
}

func TestCreateRole(t *testing.T) {
	service := newRoleService()

	role, err := service.CreateRole("seo")
	if err != nil {
		t.Fatal("something went wrong, error:", err)
	}

	if role.Name != "seo" {
		t.Fatal("expected it to be seo")
	}
}

func TestUpdateRole(t *testing.T) {
	service := newRoleService()

	createdRole, err := service.CreateRole("seo")
	if err != nil {
		t.Fatal("something went wrong, error: ", err)
	}

	updatedRole, err := service.UpdateRole(createdRole.Id, "usta")
	if err != nil {
		t.Fatal("something went wrong, error: ", err)
	}

	if updatedRole.Name != "usta" {
		t.Fatal("expected it to be usta")
	}
}
