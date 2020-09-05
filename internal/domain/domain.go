package domain

import (
	"database/sql"
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
type NullInt64 struct {
	sql.NullInt64
}
