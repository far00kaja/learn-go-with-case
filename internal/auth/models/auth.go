package models

import (
	"time"

	"github.com/google/uuid"
)

type Auth struct {
	ID        uuid.UUID `gorm:"column:id;type:uuid;default:uuid_generate_v4()"`
	Msisdn    string    `gorm:"column:msisdn;unique"`
	Username  string    `gorm:"type:varchar(20);column:username;unique"`
	Password  string    `gorm:"column:password"`
	Name      string    `gorm:"column:name"`
	Salt      string    `gorm:"column:salt"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
