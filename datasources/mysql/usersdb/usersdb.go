package usersdb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/renishb10/grossbuch_users_api/utils/config"
)

var (
	Client          *sql.DB
	userdb_username = config.GetEnv("MYSQL_USERSDB_USERNAME")
	userdb_password = config.GetEnv("MYSQL_USERSDB_PASSWORD")
	userdb_host     = config.GetEnv("MYSQL_USERSDB_HOST")
	userdb_dbname   = config.GetEnv("MYSQL_USERSDB_NAME")
)

func init() {
	fmt.Println(userdb_username)
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		userdb_username, userdb_password, userdb_host, userdb_dbname,
	)

	var err error
	Client, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database successfully configured/connected")
}
