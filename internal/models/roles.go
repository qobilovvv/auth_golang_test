package models

import (
	"time"

	"github.com/google/uuid"
)

const (
	StatusActive  = "active"
	StatusDeleted = "deleted"
)

type Role struct {
	Id        uuid.UUID  `gorm:"type:uuid;primaryKey;not null" json:"id"`
	Name      string     `gorm:"type:varchar(50);not null" json:"name"`
	Status    string     `gorm:"type:varchar(10);not null;default:'active'" json:"status"` // active and deleted
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy *uuid.UUID `gorm:"type:uuid" json:"created_by"` // nullable field
}

type SysUserRoles struct {
	Id        uuid.UUID `gorm:"type:uuid;primaryKey;not null" json:"id"`
	SysUserId uuid.UUID `gorm:"not null" json:"sysuser_id"`
	RoleId    uuid.UUID `gorm:"type:uuid;not null" json:"role_id"`
}

type RoleResponse struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
