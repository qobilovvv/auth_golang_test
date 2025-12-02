package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/qobilovvv/test_tasks/auth/internal/models"
	"github.com/qobilovvv/test_tasks/auth/internal/services"
	"github.com/qobilovvv/test_tasks/auth/pkg/helpers"
)

type roleHandler struct {
	service services.RoleService
}

func NewRoleHandler(service services.RoleService) *roleHandler {
	return &roleHandler{service: service}
}

func (h *roleHandler) CreateRole(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.ResponseError(w, http.StatusBadRequest, "Invalid credentials") // 400
		return
	}
	user, err := h.service.CreateRole(req.Name)
	if err != nil {
		helpers.RespondJSON(w, http.StatusInternalServerError, "something went wrong") // 500
		log.Println("Error in handler:", err)
		return
	}
	helpers.RespondJSON(w, http.StatusOK, user)
}

func (h *roleHandler) GetRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := h.service.GetAll()
	if err != nil {
		helpers.ResponseError(w, http.StatusInternalServerError, "something went wrong") // status = 500
		log.Println("Error in handler:", err)
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

	helpers.RespondJSON(w, http.StatusOK, res)
}


func (h *roleHandler) UpdateRole(w http.ResponseWriter, r *http.Request) {
	roleIDStr := chi.URLParam(r, "id")
	roleID, err := uuid.Parse(roleIDStr)
	if err != nil {
		helpers.ResponseError(w, http.StatusBadRequest, "Invalid role ID")
		return
	}


	var req struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Name == "" {
		helpers.ResponseError(w, http.StatusBadRequest, "Invalid or missing name")
		return
	}

	updatedRole, err := h.service.UpdateRole(roleID, req.Name)
	if err != nil {
		helpers.ResponseError(w, http.StatusInternalServerError, "something went wrong")
		log.Println("Error in handler:", err)
		return
	}

	helpers.RespondJSON(w, http.StatusOK,updatedRole)
}