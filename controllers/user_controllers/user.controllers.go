package user_controllers

import (
	"github.com/harlesbayu/bookstore-utils-go/rest_errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/harlesbayu/bookstore-oauth-package-go/oauth"
	"github.com/harlesbayu/bookstore_users-api/domain/users"
	"github.com/harlesbayu/bookstore_users-api/services"
)

func CreateUsers(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := rest_errors.NewBadRequestError("Invalid json body")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	result, saveErr := services.UserService.CreateUsers(user)

	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("x-public") == "true"))
}

func GetUsers(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	if callerId := oauth.GetCallerId(c.Request); callerId == 0 {
		err := rest_errors.NewUnauthorizedError("resource not available")

		c.JSON(err.Status(), err)
		return
	}

	userId, idErr := GetUserId(c.Param("user_id"))

	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	user, getErr := services.UserService.GetUsers(userId)

	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	if oauth.GetCallerId(c.Request) == user.Id {
		c.JSON(http.StatusOK, user.Marshall(false))
		return
	}

	c.JSON(http.StatusOK, user.Marshall(oauth.IsPublic(c.Request)))
}

func UpdateUsers(c *gin.Context) {
	userId, idErr := GetUserId(c.Param("user_id"))

	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	_, getErr := services.UserService.GetUsers(userId)

	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := rest_errors.NewBadRequestError("Invalid json body")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	user.Id = userId

	result, updateErr := services.UserService.UpdateUsers(user)

	if updateErr != nil {
		c.JSON(updateErr.Status(), updateErr)
		return
	}

	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("x-public") == "true"))
}

func DeleteUsers(c *gin.Context) {
	userId, idErr := GetUserId(c.Param("user_id"))

	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	_, getErr := services.UserService.GetUsers(userId)

	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	if err := services.UserService.DeleteUsers(userId); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"message": "deleted"})
}

func GetUserId(userIdParams string) (int64, rest_errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParams, 10, 64)
	if userErr != nil {
		return 0, rest_errors.NewBadRequestError("user id should be a number")
	}
	return userId, nil
}

func FindByStatus(c *gin.Context) {
	status := c.Query("status")

	users, err := services.UserService.FindByStatus(status)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("x-public") == "true"))
}

func Login(c *gin.Context) {
	var request users.LoginRrequest

	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	user, err := services.UserService.LoginUser(request)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("x-public") == "true"))
}
