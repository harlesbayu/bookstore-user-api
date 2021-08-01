package main

import (
	"github.com/harlesbayu/bookstore_users-api/app"
	"github.com/harlesbayu/bookstore_users-api/datasources/mysql/user_db"
	"github.com/harlesbayu/bookstore_users-api/shared/config"
)

func main() {
	user_db.MysqlConnection(config.NewConfig("./"))
	app.StartApplication()
}
