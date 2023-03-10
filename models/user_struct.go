package models

import (
	"time"
)

type User struct {
	ID            uint      `json:"id" validate:"required"`
	First_Name    string    `json:"firstname" validate:"required"`
	Last_Name     string    `json:"lastname" validate:"required"`
	Password      string    `json:"Password" validate:"required,min=6"`
	Email         string    `json:"email" validate:"required"`
	PhoneNumber   string    `json:"phonenumber" validate:"required"`
	City          string    `json:"city" validate:"required"`
	Pin_Code      string    `json:"pincode" validate:"required"`
	State         string    `json:"state" validate:"required"`
	Location      string    `json:"location" validate:"required"`
	Token         string    `json:"token"`
	Refresh_Token string    `json:"refresh_token"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
	User_id       string    `json:"user_id"`
}
