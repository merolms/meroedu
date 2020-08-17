package datastore

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	"github.com/meroedu/meroedu/app/config"
)

// NewDB ...
func NewDB() *sql.DB {
	configs := config.C.Database
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.User, configs.Password, configs.Host, configs.Port, configs.DBName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Kathmandu")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	db, err := sql.Open(`mysql`, dsn)

	if err != nil {
		log.Fatalln("Database Connection error:", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}
