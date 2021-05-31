package users

import "github.com/zooter/bookstore_users-api/utils/errors"

type User_DTO struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User_DTO) Validate() *errors.RestErr {
	if user.Email == "" {
		return errors.NewBadRequestError("bad_request")
	}
	return nil
}
