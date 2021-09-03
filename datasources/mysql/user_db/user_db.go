package user_db

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/harlesbayu/bookstore-utils-go/logger"
	"log"

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

	var err error


	Client, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	mysql.SetLogger(logger.GetLogger())
	log.Printf("database successfully configured")
}
