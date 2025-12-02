package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/qobilovvv/test_tasks/auth/internal/db"
	"github.com/qobilovvv/test_tasks/auth/internal/handlers"
	"github.com/qobilovvv/test_tasks/auth/internal/repositories"
	"github.com/qobilovvv/test_tasks/auth/internal/services"
	"github.com/qobilovvv/test_tasks/auth/pkg/helpers"
)

const (
	PORT = 3000
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("there is no .env file")
	}
	db := db.InitDB()

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	roleRepo := repositories.NewRoleReposity(db)
	otpRepo := repositories.NewOTPRepository(db)
	usersRepo := repositories.NewUserRepository(db)
	sysUsersRepo := repositories.NewSysUserRepository(db)

	roleService := services.NewRoleService(roleRepo)
	otpService := services.NewOTPService(otpRepo)
	userService := services.NewUserService(usersRepo, otpRepo, sysUsersRepo)
	sysUsersService := services.NewSysUserService(sysUsersRepo)

	roleHandler := handlers.NewRoleHandler(roleService)
	otpHandler := handlers.NewOTPHandler(otpService)
	userHandler := handlers.NewUserHandler(userService)
	sysUserHandler := handlers.NewSysUserHandler(&sysUsersService)

	// admin create
	helpers.InitSuperAdmin(sysUsersRepo)

	router := handlers.NewRouter(handlers.Router{
		RoleHandler:    roleHandler,
		OtpHandler:     otpHandler,
		UserHandler:    userHandler,
		SysUserHandler: sysUserHandler,
	})

	// log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", PORT), router))

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", PORT),
		Handler: router,
	}

	go func() {
		log.Printf("Server running on port %d\n", PORT)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server gracefully...")

	// waits server 5 seconds to finish requests
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Forced shutdown: %v", err)
	}

	log.Println("Server stopped clearly")
}
