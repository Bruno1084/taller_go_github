package domain

import (
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	NickName  string    `json:"nickname"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Version   int       `json:"version"`
}

type UserUpdateFields struct {
	Name     *string `json:"name"`
	Address  *string `json:"address"`
	NickName *string `json:"nickname"`
}
