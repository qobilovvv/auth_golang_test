package tests

import (
	"os"
	"testing"

	"github.com/qobilovvv/test_tasks/auth/internal/db"
	"github.com/qobilovvv/test_tasks/auth/internal/models"
	"gorm.io/gorm"
)

var testDB *gorm.DB

func TestMain(m *testing.M) {
	var err error

	testDB, err = db.TestDB()
	if err != nil {
		panic(err)
	}
	if err := testDB.AutoMigrate(
		&models.Role{},
		&models.Users{},
		&models.SysUsers{},
		&models.SysUserRoles{},
	); err != nil {
		panic(err)
	}

	code := m.Run()
	os.Exit(code)
}
