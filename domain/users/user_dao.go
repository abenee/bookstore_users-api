package users

import (
	"fmt"
	"github.com/abenee/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.ErrorResp {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewRecordNotfoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.ErrorResp {
	current := usersDB[user.Id]
	if current != nil{
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s arlready exists", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d arlready exists", user.Id))
	}
	usersDB[user.Id] = user
	return nil
}
