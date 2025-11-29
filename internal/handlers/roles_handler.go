package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/qobilovvv/test_tasks/auth/internal/models"
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

	var res []models.RoleResponse
	for _, role := range roles {
		res = append(res, models.RoleResponse{
			Id:        role.Id,
			Name:      role.Name,
			CreatedAt: role.CreatedAt,
		})
	}

	RespondJSON(w, http.StatusOK, res)
}


func (h *userHandler) UpdateRole(w http.ResponseWriter, r *http.Request) {
	roleIDStr := chi.URLParam(r, "id")
	roleID, err := uuid.Parse(roleIDStr)
	if err != nil {
		http.Error(w, "Invalid role ID", http.StatusBadRequest)
		return
	}


	var req struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Name == "" {
		http.Error(w, "Invalid or missing name", http.StatusBadRequest)
		return
	}

	updatedRole, err := h.service.UpdateRole(roleID, req.Name)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondJSON(w, http.StatusOK,updatedRole)
}