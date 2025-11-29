package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	RoleHandler *userHandler
}

func NewRouter(r Router) http.Handler {
	router := chi.NewRouter()

	router.Route("/roles", func(ro chi.Router) {
		ro.Post("/create", r.RoleHandler.CreateRole)
		ro.Get("/", r.RoleHandler.GetRoles)
		ro.Patch("/{id}", r.RoleHandler.UpdateRole)
	})

	return router
}
