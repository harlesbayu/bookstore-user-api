package user_db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	config "github.com/harlesbayu/bookstore_users-api/shared/config"
)

var (
	Client *sql.DB
)

func MysqlConnection(config *config.Config) {
	dataSource := fmt.Sprintf("%s:%s@(%s:%d)/%s",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)

	fmt.Println("database connection success")

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	Client = db
}
