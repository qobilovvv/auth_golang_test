package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	handlers "github.com/qobilovvv/test_tasks/auth/internal/handlers"
	"github.com/qobilovvv/test_tasks/auth/internal/models"
	"github.com/qobilovvv/test_tasks/auth/internal/services"
)

// mockRoleService implements services.RoleService for testing.
type mockRoleService struct {
	roles []models.Role
	err   error
}

func (m *mockRoleService) CreateRole(name string) (*models.Role, error) {
	r := &models.Role{
		Id:        uuid.New(),
		Name:      name,
		CreatedAt: time.Now(),
	}
	return r, nil
}

func (m *mockRoleService) GetAll() ([]models.Role, error) {
	return m.roles, m.err
}

func (m *mockRoleService) UpdateRole(roleID uuid.UUID, name string) (*models.Role, error) {
	r := &models.Role{
		Id:        roleID,
		Name:      name,
		CreatedAt: time.Now(),
	}
	return r, nil
}

func TestGetRolesHandler_Success(t *testing.T) {
	now := time.Now().UTC().Truncate(time.Second)

	role1ID := uuid.New()
	role2ID := uuid.New()

	mockRoles := []models.Role{
		{
			Id:        role1ID,
			Name:      "admin",
			CreatedAt: now,
		},
		{
			Id:        role2ID,
			Name:      "user",
			CreatedAt: now.Add(time.Minute),
		},
	}

	svc := &mockRoleService{
		roles: mockRoles,
		err:   nil,
	}

	h := handlers.NewRoleHandler(svc)

	req := httptest.NewRequest(http.MethodGet, "/roles", nil)
	w := httptest.NewRecorder()

	h.GetRoles(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d, body: %s", w.Code, w.Body.String())
	}

	ctype := w.Header().Get("Content-Type")
	if ctype != "application/json" {
		t.Fatalf("expected Content-Type application/json, got %s", ctype)
	}

	var resp []map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if len(resp) != len(mockRoles) {
		t.Fatalf("expected %d roles, got %d", len(mockRoles), len(resp))
	}

	// Verify fields for each role
	if resp[0]["name"] != mockRoles[0].Name {
		t.Fatalf("expected first role name %s, got %v", mockRoles[0].Name, resp[0]["name"])
	}
	if resp[1]["name"] != mockRoles[1].Name {
		t.Fatalf("expected second role name %s, got %v", mockRoles[1].Name, resp[1]["name"])
	}

	// Ensure id fields are strings and match
	if _, ok := resp[0]["id"].(string); !ok {
		t.Fatalf("expected id of first role to be string, got %T", resp[0]["id"])
	}
	if resp[0]["id"] != mockRoles[0].Id.String() {
		t.Fatalf("expected first role id %s, got %v", mockRoles[0].Id.String(), resp[0]["id"])
	}
}

func TestGetRolesHandler_ServiceError(t *testing.T) {
	svc := &mockRoleService{
		roles: nil,
		err:   errTest("database failure"),
	}

	h := handlers.NewRoleHandler(svc)

	req := httptest.NewRequest(http.MethodGet, "/roles", nil)
	w := httptest.NewRecorder()

	h.GetRoles(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 500, got %d, body: %s", w.Code, w.Body.String())
	}

	var resp map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to unmarshal error response: %v", err)
	}

	if resp["error"] != "database failure" {
		t.Fatalf("expected error message %q, got %q", "database failure", resp["error"])
	}
}

// errTest is a tiny error implementation for tests.
type errTest string

func (e errTest) Error() string { return string(e) }

// Ensure mockRoleService implements the RoleService interface at compile time.
var _ services.RoleService = (*mockRoleService)(nil)
