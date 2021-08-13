package app

import (
	"github.com/harlesbayu/bookstore_users-api/controllers/ping_controllers"
	"github.com/harlesbayu/bookstore_users-api/controllers/user_controllers"
)

func mapUrls() {
	router.GET("/ping", ping_controllers.Ping)
	router.GET("users", user_controllers.FindByStatus)
	router.GET("users/:user_id", user_controllers.GetUsers)
	router.POST("users", user_controllers.CreateUsers)
	router.PUT("users/:user_id", user_controllers.UpdateUsers)
	router.DELETE("users/:user_id", user_controllers.DeleteUsers)
	router.POST("users/login", user_controllers.Login)
}
