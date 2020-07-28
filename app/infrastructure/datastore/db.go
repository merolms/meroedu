package datastore

import (
	"fmt"
	"log"
	"net/url"

	"github.com/jinzhu/gorm"
	"github.com/meroedu/course-api/app/config"
)

// NewDB ...
func NewDB() *gorm.DB {
	configs := config.C.Database
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.User, configs.Password, configs.Host, configs.Port, configs.DBName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Kathmandu")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		log.Fatalln("Database Connection error:", err)
	}

	return db
}
