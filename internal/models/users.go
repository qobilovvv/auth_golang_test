package models

import "github.com/google/uuid"

type Users struct {
	Id       uuid.UUID `gorm:"type:varchar(50); primaryKey; not null" json:"id"`
	Status   string    `gorm:"type:varchar(50)" json:"status"`
	Name     string    `gorm:"type:varchar(255)" json:"name"`
	Email    string    `gorm:"type:varchar(255)" json:"email"`
	Password string    `gorm:"type:varchar(255)" json:"password"`
}
