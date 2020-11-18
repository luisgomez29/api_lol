package models

import (
	"time"
)

type Character struct {
	ID          uint32     `json:"id"`
	Name        string     `gorm:"not null;size:40" json:"name"`
	Position    string     `gorm:"size:40;" json:"position"`
	Description string     `gorm:"size:255;" json:"description,omitempty"`
	CreatedAt   *time.Time `gorm:"not null;default:now()" json:"created_at,omitempty"`
	UpdatedAt   *time.Time `gorm:"not null;default:now()" json:"updated_at,omitempty"`
}
