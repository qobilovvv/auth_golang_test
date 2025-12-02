package helpers

import (
	"encoding/json"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/qobilovvv/test_tasks/auth/internal/models"
	"github.com/qobilovvv/test_tasks/auth/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

func InitSuperAdmin(repo repositories.SysUserRepository) {
	count, err := repo.Count()
	if err != nil {
		log.Println("failed to count sysusers:", err)
	}

	if count == 0 {
		hashed, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("ADMIN_PASSWORD")), bcrypt.DefaultCost)
		if err != nil {
			log.Println("error", err)
		}

		superAdmin := &models.SysUsers{
			Id:       uuid.New(),
			Phone:    os.Getenv("ADMIN_PHONE"),
			Name:     "Super Admin",
			Password: string(hashed),
			Status:   "active",
		}

		user, err := repo.Create(superAdmin)

		if err != nil {
			log.Fatalf("failed to create superadmin: %v", err)
		}

		user_info := map[string]string{
			"id":     user.Id.String(),
			"name":   user.Name,
			"phone":  user.Phone,
			"status": user.Status,
		}

		json_user, _ := json.Marshal(user_info)

		log.Println("created super user: ", string(json_user))
	}
}
