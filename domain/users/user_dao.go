package users

import (
	"fmt"

	"github.com/zooter/bookstore_users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User_DTO)
)

func (user *User_DTO) Get() *errors.RestErr {
	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User_DTO) Save() *errors.RestErr {
	if userDB[user.Id] != nil {
		return errors.NewBadRequestError(fmt.Sprintf("User %d already exists", user.Id))
	}
	userDB[user.Id] = user
	return nil
}
