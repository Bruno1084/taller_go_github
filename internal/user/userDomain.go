package user

import (
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Nickname  string    `json:"nickname"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Version   int       `json:"version"`
}


type UserCreateParams struct {
	Name     *string `json:"name"`
	Address  *string `json:"address"`
	NickName *string `json:"nickname"`
}