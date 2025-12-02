package tests

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/qobilovvv/test_tasks/auth/internal/models"
	"github.com/qobilovvv/test_tasks/auth/internal/repositories"
	"github.com/qobilovvv/test_tasks/auth/internal/services"
)

func newSysUserService() services.SysUserService {
	repo := repositories.NewSysUserRepository(testDB)
	return services.NewSysUserService(repo)
}

func TestCreateSysUser(t *testing.T) {
	service := newSysUserService()

	rolesToInsert := []models.Role{
		{Id: uuid.New(), Name: "admin"},
		{Id: uuid.New(), Name: "user"},
		{Id: uuid.New(), Name: "waiter"},
	}

	if err := testDB.Create(&rolesToInsert).Error; err != nil {
		t.Fatal(err)
	}

	roleIDs := make([]uuid.UUID, len(rolesToInsert))
	for i, r := range rolesToInsert {
		roleIDs[i] = r.Id
	}

	uid, err := service.CreateSysUser("odil", "998977479811", "1", roleIDs)
	if err != nil {
		t.Fatal("something went wrong", err)
	}

	fmt.Println("PASSED TEST CREATE_SYS_USER, user_id: ", uid)
}
