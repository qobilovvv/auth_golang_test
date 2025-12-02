package tests

import (
	"testing"

	"github.com/qobilovvv/test_tasks/auth/internal/repositories"
	"github.com/qobilovvv/test_tasks/auth/internal/services"
)

func newUserService() services.UserService {
	repo := repositories.NewUserRepository(testDB)
	OtpRepo := repositories.NewOTPRepository(testDB)
	SysRepo := repositories.NewSysUserRepository(testDB)
	return services.NewUserService(repo, OtpRepo, SysRepo)
}

func TestSignUpService(t *testing.T) {}
