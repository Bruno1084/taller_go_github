package domain

import (
	"time"
)

type Sale struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Amount    float32   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Version   int       `json:"version"`
}

type SaleUpdateFields struct {
	Status *string `json:"status"`
}
