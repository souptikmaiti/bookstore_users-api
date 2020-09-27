package services

import (
	"github.com/souptikmaiti/bookstore_users-api/domain/users"
	"github.com/souptikmaiti/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.ValidateUser(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	u := users.User{Id: userId}
	err := u.Get()
	if err != nil {
		return nil, err
	}
	return &u, nil
}
