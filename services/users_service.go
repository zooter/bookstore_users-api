package services

import (
	"github.com/zooter/bookstore_users-api/domain/users"
	"github.com/zooter/bookstore_users-api/utils/errors"
)

func GetUser(userId int64) (*users.User_DTO, *errors.RestErr) {
	if userId < 0 {
		errors.NewBadRequestError("Invalid user id")
	}
	result := users.User_DTO{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return &result, nil
}

func CreateUser(user users.User_DTO) (*users.User_DTO, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
