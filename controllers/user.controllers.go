package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/harlesbayu/bookstore_users-api/domain/users"
	"github.com/harlesbayu/bookstore_users-api/services"
	"github.com/harlesbayu/bookstore_users-api/utils/errors"
)

func CreateUsers(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}
	result, saveErr := services.CreateUsers(user)

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
	return
}

func GetUsers(c *gin.Context) {
	userId, idErr := GetUserId(c.Param("user_id"))

	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user, getErr := services.GetUsers(userId)

	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
	return
}

func UpdateUsers(c *gin.Context) {
	userId, idErr := GetUserId(c.Param("user_id"))

	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	_, getErr := services.GetUsers(userId)

	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	user.Id = userId

	result, updateErr := services.UpdateUsers(user)

	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}

	c.JSON(http.StatusOK, result)
	return
}

func DeleteUsers(c *gin.Context) {
	userId, idErr := GetUserId(c.Param("user_id"))

	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	_, getErr := services.GetUsers(userId)

	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	if err := services.DeleteUsers(userId); err != nil {
		c.JSON(err.Status, err)
	}

	c.JSON(http.StatusOK, map[string]string{"message": "deleted"})

}

func GetUserId(userIdParams string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParams, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("user id should be a number")
	}
	return userId, nil
}

func FindByStatus(c *gin.Context) {
	status := c.Query("status")

	users, err := services.FindByStatus(status)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users)
	return
}
