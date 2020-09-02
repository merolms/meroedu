package datastore

import (
	"database/sql"
	"fmt"
	"net/url"
	"time"

	"github.com/meroedu/meroedu/internal/config"
	log "github.com/meroedu/meroedu/pkg/log"
)

// MaxDatabaseConnectionAttempts ...
const MaxDatabaseConnectionAttempts int = 10

var db *sql.DB

// NewDB ...
func NewDB() (*sql.DB, error) {
	configs := config.C.Database
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.User, configs.Password, configs.Host, configs.Port, configs.DBName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Kathmandu")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	// Open our database connection
	i := 0
	var err error
	for {
		db, err = sql.Open(`mysql`, dsn)
		if err == nil {
			break
		}
		if err != nil && i >= MaxDatabaseConnectionAttempts {
			log.Error(err)
			return nil, err
		}
		i++
		log.Warn("waiting for database to be up...")
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Info("Database connection success")
	return db, nil
}
