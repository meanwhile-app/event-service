package models

import (
	"time"

	"github.com/nuntjw/go-gin-starter/models/schemas"
)

type UserModel struct{}

func (userModel *UserModel) GetUsers() ([]schemas.User, error) {
	users := []schemas.User{
		{ID: 1, Name: "Nunt", Email: "nuntjw@gmail.com", Password: "test", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 2, Name: "User2", Email: "user2@gmail.com", Password: "test-2", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	return users, nil
}
