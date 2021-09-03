package app

import (
	"github.com/gin-gonic/gin"
	"github.com/harlesbayu/bookstore-utils-go/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("starting the application")
	router.Run(":3000")
}
