package services

import (
	"github.com/abenee/bookstore_users-api/domain/users"
	"github.com/abenee/bookstore_users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.ErrorResp) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.ErrorResp) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
