package domain

import (
	"database/sql"
	"reflect"
	"time"
)

// Status ...
type Status string

const (
	Success Status = "success"
	Error   Status = "error"
)

// APIResponseError example
type APIResponseError struct {
	ErrorCode    int
	ErrorMessage string
	CreatedAt    time.Time
}

// Response contains the attributes found in an API response
type Response struct {
	Message Status      `json:"message"`
	Data    interface{} `json:"data"`
}

// NullInt64 ...
type NullInt64 sql.NullInt64

// Scan implements the Scanner interface for NullInt64
func (ni *NullInt64) Scan(value interface{}) error {
	var i sql.NullInt64
	if err := i.Scan(value); err != nil {
		return err
	}
	// if nil the make Valid false
	if reflect.TypeOf(value) == nil {
		*ni = NullInt64{i.Int64, false}
	} else {
		*ni = NullInt64{i.Int64, true}
	}
	return nil
}
