package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/qobilovvv/test_tasks/auth/internal/config"
	"github.com/qobilovvv/test_tasks/auth/internal/handlers"
	"github.com/qobilovvv/test_tasks/auth/internal/repositories"
	"github.com/qobilovvv/test_tasks/auth/internal/services"
)

const (
	PORT = 3000
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("there is no .env file")
	}
	db := config.InitDB()

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	roleRepo := repositories.NewRoleReposity(db)
	otpRepo := repositories.NewOTPRepository(db)
	usersRepo := repositories.NewUserRepository(db)

	roleService := services.NewRoleService(roleRepo)
	otpService := services.NewOTPService(otpRepo)
	userService := services.NewUserService(usersRepo, otpRepo)

	roleHandler := handlers.NewRoleHandler(roleService)
	otpHandler := handlers.NewOTPHandler(otpService)
	userHandler := handlers.NewUserHandler(userService)

	router := handlers.NewRouter(handlers.Router{
		RoleHandler: roleHandler,
		OtpHandler:  otpHandler,
		UserHandler: userHandler,
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", PORT), router))
}
