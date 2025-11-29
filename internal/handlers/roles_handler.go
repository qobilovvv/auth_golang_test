package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/qobilovvv/test_tasks/auth/internal/services"
)

type userHandler struct {
	service services.RoleService
}

func NewUserHandler(service services.RoleService) *userHandler {
	return &userHandler{service: service}
}

func (h *userHandler) CreateRole(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponseError(w, http.StatusBadRequest, "Invalid credentials") // 400
		return
	}
	user, err := h.service.CreateRole(req.Name)
	if err != nil {
		RespondJSON(w, http.StatusInternalServerError, err.Error()) // 500
		return
	}
	RespondJSON(w, http.StatusOK, user)
}

func (h *userHandler) GetRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := h.service.GetAll()
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error()) // status = 500
		return
	}
	RespondJSON(w, http.StatusOK, roles)
}
