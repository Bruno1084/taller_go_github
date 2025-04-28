package services

import (
	"github.com/google/uuid"
	"go_test/src/db"
	"go_test/src/domain"
	"time"
)

func GetAllUsers() ([]domain.User, error) {
	rows, err := db.DB.Query("SELECT id, name, address, nickname, version FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.ID, &user.Name, &user.NickName, &user.NickName, &user.Version)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUserById(id string) (domain.User, error) {
	var user domain.User
	row := db.DB.QueryRow("SELECT id, name, address, nickname, version FROM users WHERE id = ?", id)
	err := row.Scan(&user.ID, &user.Name, &user.Address, &user.NickName, &user.Version)

	return user, err
}

func CreateUser(user *domain.User) (*domain.User, error) {
	user.ID = uuid.NewString()
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	user.Version = 1

	query := `INSERT INTO users(id, name, address, nickname, createdAt, updatedAt, version) 
	VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := db.DB.Exec(query, user.ID, user.Name, user.Address, user.NickName, user.CreatedAt, user.UpdatedAt, user.Version)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func PatchUser(id string, user *domain.UserUpdateFields) (*domain.User, error) {
	_, err := db.DB.Exec("UPDATE users SET name = ?, address = ?, nickname = ? WHERE id = ?",
		user.Name, user.Address, user.NickName, id)
	if err != nil {
		return nil, err
	}

	updatedUser, err := GetUserById(id)
	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}
