package app

import (
	"github.com/harlesbayu/bookstore_users-api/controllers"
)

func mapUrls() {
	router.GET("/ping", controllers.Ping)
	router.GET("users", controllers.FindByStatus)
	router.GET("users/:user_id", controllers.GetUsers)
	router.POST("users", controllers.CreateUsers)
	router.PUT("users/:user_id", controllers.UpdateUsers)
	router.DELETE("users/:user_id", controllers.DeleteUsers)
}
