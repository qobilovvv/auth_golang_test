package models

import "github.com/google/uuid"

type SysUsers struct {
	Id       uuid.UUID `gorm:"type:uuid" json:"id"`
	Name     string    `gorm:"type:varchar(255);not null" json:"name"`
	Phone    string    `gomr:"type:varchar(20); not null;unique"`
	Password string    `gorm:"type:varchar(255);not null" json:"password"`
	Status   string    `gorm:"type:varchar(15);not null;default:'active'"` // active, blocked
}