package services

import (
	"github.com/harlesbayu/bookstore-utils-go/rest_errors"
	"github.com/harlesbayu/bookstore_users-api/domain/users"
	"github.com/harlesbayu/bookstore_users-api/utils/crypto_utils"
	"github.com/harlesbayu/bookstore_users-api/utils/date_utils"
)

var (
	UserService userServiceInterface = &userService{}
)

type userServiceInterface interface {
	CreateUsers(users.User) (*users.User, rest_errors.RestErr)
	GetUsers(int64) (*users.User, rest_errors.RestErr)
	UpdateUsers(users.User) (*users.User, rest_errors.RestErr)
	DeleteUsers(int64) rest_errors.RestErr
	FindByStatus(string) (users.Users, rest_errors.RestErr)
	LoginUser(users.LoginRrequest) (*users.User, rest_errors.RestErr)
}

type userService struct{}

func (s *userService) CreateUsers(user users.User) (*users.User, rest_errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.DateCreated = date_utils.GetNowDBFormat()
	user.Status = users.StatusActive
	user.Password = crypto_utils.EncryptPassword(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *userService) GetUsers(userId int64) (*users.User, rest_errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *userService) UpdateUsers(user users.User) (*users.User, rest_errors.RestErr) {
	current, err := s.GetUsers(user.Id)

	if err != nil {
		return nil, err
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if user.FirstName != "" {
		current.FirstName = user.FirstName
	}
	if user.LastName != "" {
		current.LastName = user.LastName
	}
	if user.Email != "" {
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func (s *userService) DeleteUsers(userId int64) rest_errors.RestErr {
	user := &users.User{Id: userId}

	return user.Delete()
}

func (s *userService) FindByStatus(status string) (users.Users, rest_errors.RestErr) {
	dao := users.User{}
	return dao.FindByStatus(status)
}

func (s *userService) LoginUser(request users.LoginRrequest) (*users.User, rest_errors.RestErr) {
	dao := &users.User{
		Email: request.Email,
	}

	if err := dao.FindByEmail(); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid credential")
		return nil, restErr
	}

	password := crypto_utils.DecryptPassword(dao.Password)

	if request.Password != password {
		restErr := rest_errors.NewBadRequestError("invalid credential")
		return nil, restErr
	}
	return dao, nil
}
