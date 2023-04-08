package entity

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	ID        uuid.UUID  `json:"id"`
	OwnerName string     `json:"owner_name"`
	Infos     []string   `json:"infos"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
