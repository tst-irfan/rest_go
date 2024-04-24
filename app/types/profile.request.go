package types

import "time"

type ProfileRequest struct {
	FirstName string    `json:"first_name" binding:"required"`
	LastName  string    `json:"last_name" binding:"required"`
	DoB       time.Time `json:"dob" binding:"required"`
}
