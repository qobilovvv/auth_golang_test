package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/qobilovvv/test_tasks/auth/internal/services"
	"github.com/qobilovvv/test_tasks/auth/pkg/helpers"
)

type sysUserHandler struct {
	service services.SysUserService
}

func NewSysUserHandler(service *services.SysUserService) *sysUserHandler {
	return &sysUserHandler{service: *service}
}

func (h *sysUserHandler) CreateSysUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name     string   `json:"name"`
		Phone    string   `json:"phone"`
		Password string   `json:"password"`
		Roles    []string `json:"roles"` // uuid in string inside list
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	roleIDs := make([]uuid.UUID, len(req.Roles))
	for i, r := range req.Roles {
		id, err := uuid.Parse(r)
		if err != nil {
			helpers.ResponseError(w, http.StatusBadRequest, "invalid uuid of role")
			return
		}
		roleIDs[i] = id
	}

	userID, err := h.service.CreateSysUser(req.Name, req.Phone, req.Password, roleIDs)
	if err != nil {
		helpers.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.RespondJSON(w, http.StatusCreated, map[string]string{"id": userID.String()})
}
