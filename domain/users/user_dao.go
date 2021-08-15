package users

import (
	"fmt"

	"github.com/harlesbayu/bookstore_users-api/datasources/mysql/user_db"
	"github.com/harlesbayu/bookstore_users-api/logger"
	"github.com/harlesbayu/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser   = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser      = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id = ?;"
	queryUpdateUser   = "Update users SET first_name=?, last_name=?, email=? WHERE id = ?;"
	queryDeleteUser   = "DELETE FROM users where id=?;"
	queryFIndByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status = ?;"
	queryFindByEmail  = "SELECT id, first_name, last_name, email, date_created, password, status FROM users WHERE email=? and status=?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := user_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		logger.Error("error when trying to get user by user id", getErr)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := user_db.Client.Prepare(queryInsertUser)

	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)

	if saveErr != nil {
		logger.Error("error when trying to save user", saveErr)
		return errors.NewInternalServerError("database error")
	}

	userId, err := insertResult.LastInsertId()

	if err != nil {
		logger.Error("error when trying to get last user id after creating a new user", err)
		return errors.NewInternalServerError("database error")
	}

	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := user_db.Client.Prepare(queryUpdateUser)

	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)

	if err != nil {
		logger.Error("error when trying to update user", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := user_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.Id); err != nil {
		logger.Error("error when trying to delete user", err)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := user_db.Client.Prepare(queryFIndByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find users by status statement", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)

	if err != nil {
		logger.Error("error when trying to find users by status", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer rows.Close()

	result := make([]User, 0)

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when scan user row into user struct", err)
			return nil, errors.NewInternalServerError("database error")
		}
		result = append(result, user)
	}

	if len(result) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}

	return result, nil
}

func (user *User) FindByEmail() *errors.RestErr {
	stmt, err := user_db.Client.Prepare(queryFindByEmail)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, StatusActive)

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Password, &user.Status); getErr != nil {
		logger.Error("error when trying to get user by email", getErr)
		return errors.NewInternalServerError("database error")
	}

	return nil
}
