package users

import (
	"net/mail"
	"strings"

	"github.com/harlesbayu/bookstore_users-api/utils/errors"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	DateCreated string `json:"dateCreated"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

const (
	StatusActive    = "active"
	StatusNonActive = "non_active"
)

type Users []User

func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid email address")
	} else {
		_, err := mail.ParseAddress(user.Email)
		if err != nil {
			return errors.NewBadRequestError("Invalid email address")
		}
	}

	return nil
}
