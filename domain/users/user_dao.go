package users

import (
	"fmt"

	"github.com/zooter/bookstore_users-api/datasources/mysql/users_db"
	"github.com/zooter/bookstore_users-api/utils/date_utils"
	"github.com/zooter/bookstore_users-api/utils/errors"
)

const(
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?, ? ,? ,?)"
)

var (
	userDB = make(map[int64]*User_DTO)
)

func (user *User_DTO) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
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
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return  errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if err != nil {
		return  errors.NewInternalServerError(fmt.Sprintf("Error when trying to save user: %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId();
	if err != nil {
		return  errors.NewInternalServerError(fmt.Sprintf("Error when trying to save user: %s", err.Error()))
	}

	user.Id = userId

	if userDB[user.Id] != nil {
		return errors.NewBadRequestError(fmt.Sprintf("User %d already exists", user.Id))
	}
	user.DateCreated = date_utils.GetNowString()
	userDB[user.Id] = user
	return nil
}
