package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/qobilovvv/test_tasks/auth/internal/config"
	"github.com/qobilovvv/test_tasks/auth/internal/handlers"
	"github.com/qobilovvv/test_tasks/auth/internal/repositories"
	"github.com/qobilovvv/test_tasks/auth/internal/services"
	"github.com/joho/godotenv"
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
	r.Use(middleware.Recoverer)

	roleRepo := repositories.NewRoleReposity(db)
	otpRepo := repositories.NewOTPRepository(db)

	roleService := services.NewRoleService(roleRepo)
	otpService := services.NewOTPService(otpRepo)

	roleHandler := handlers.NewUserHandler(roleService)
	otpHandler := handlers.NewOTPHandler(otpService)

	router := handlers.NewRouter(handlers.Router{
		RoleHandler: roleHandler,
		OtpHandler:  otpHandler,
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", PORT), router))
}
