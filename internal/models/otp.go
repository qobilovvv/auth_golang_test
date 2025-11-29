package models

import (
	"time"

	"github.com/google/uuid"
)

const (
	StatusConfirmed   = "confirmed"
	StatusUnconfirmed = "unconfirmed"
)

type OTP struct {
	Id        uuid.UUID `gorm:"type:uuid; primaryKey;not null" json:"id"`
	Email     string    `gorm:"type:varchar(255);not null"`
	Code      string    `gorm:"type:varchar(50);not null"`
	Status    string    `gorm:"type:varchar(50);not null;default:'unconfirmed'"`
	ExpiresAt time.Time
}
