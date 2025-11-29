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
)

const (
	PORT = 3000
)

func main() {
	db := config.InitDB()
	
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	roleRepo := repositories.NewRoleReposity(db)
	roleService := services.NewRoleService(roleRepo)
	roleHandler := handlers.NewUserHandler(roleService)

	router := handlers.NewRouter(handlers.Router{
		RoleHandler: roleHandler,
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", PORT), router))
}
