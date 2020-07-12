package models

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	user := os.Getenv("mySQLUser")
	password := os.Getenv("mySQLPassword")
	host := os.Getenv("mySQLHost")
	port := os.Getenv("mySQLPort")
	dbName := os.Getenv("mySQLDatabase")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		user, password, host, port, dbName)
	// dsn := "root:Ethan0909@tcp(127.0.0.1:3306)/sql_test"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
