package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const(
	mysql_db_username = "mysql_db_username"
	mysql_db_password = "mysql_db_password"
	mysql_db_host = "mysql_db_host"
	mysql_db_schema = "mysql_db_schema"
)

var (
	Client *sql.DB

	username	= os.Getenv("mysql_db_username")
	password 	= os.Getenv("mysql_db_password")
	host 		= os.Getenv("mysql_db_host")
	schema 		= os.Getenv("mysql_db_schema")
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		username,
		password,
		host,
		schema)
	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database Successfully configured")
}
