package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	RoleHandler *roleHandler
	OtpHandler  *otpHandler
	UserHandler *userHandler
	SysUserHandler *sysUserHandler
}

func NewRouter(r Router) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Route("/roles", func(ro chi.Router) {
		ro.Post("/create", r.RoleHandler.CreateRole)
		ro.Get("/", r.RoleHandler.GetRoles)
		ro.Patch("/{id}", r.RoleHandler.UpdateRole)
	})

	router.Route("/otp", func(ro chi.Router) {
		ro.Post("/send", r.OtpHandler.SendOTP)
		ro.Post("/confirm", r.OtpHandler.ConfirmOTP)
	})
	router.Route("/users", func(ro chi.Router) {
		ro.Post("/signup", r.UserHandler.SignUp)
		ro.Post("/login", r.UserHandler.Login)
	})

	router.Route("/sysusers", func(ro chi.Router) {
		ro.Post("/create", r.SysUserHandler.CreateSysUser)
	})

return router
}
