package common

import (
	"time"

)

// Base struct for Requests (business model)
type RequestBase struct {
	References []Reference
	SendAt     time.Time
}

// Response business model
type Response struct {
	Id        string
	CreatedAt time.Time
}

// Reference for storing additional information
type Reference struct {
	Id   string
	Type string
}
