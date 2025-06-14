package models

import "time"

type User struct {
    ID   uint   `json:"id"`
    Name string `json:"name"`
}

type status string

const (
	PENDING     status = "PENDING"
	IN_PROGRESS status = "IN_PROGRESS"
	COMPLETED   status = "COMPLETED"
	CANCELLED   status = "CANCELLED"
)

type Issue struct {
    ID          uint      `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Status      status    `json:"status"`
    User        *User     `json:"user,omitempty"`
    CreatedAt   time.Time `json:"createdAt"`
    UpdatedAt   time.Time `json:"updatedAt"`
}