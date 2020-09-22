package domain

import (
	"database/sql"
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
	CreatedAt    int64
}

// Response contains the attributes found in an API response
type Response struct {
	Message Status      `json:"message"`
	Data    interface{} `json:"data"`
}

// Summaries contains the attributes for GetAll API pagination response
type Summaries struct {
	Response
	Total int64 `json:"total"`
}

// NullInt64 ...
type NullInt64 struct {
	sql.NullInt64
}
