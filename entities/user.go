package entities

import "time"

type User struct {
	ID            int
	First_name    string    `json:"first_name" validate:"required,min=2,max=100"`
	Last_name     string    `json:"last_name" validate:"required,min=2,max=100"`
	Email         string    `json:"email" validate:"email,required"`
	Password      string    `json:"password" validate:"required,min=6"`
	Avatar        string    `json:"avatar"`
	Phone         string    `json:"phone" validate:"required"`
	Token         string    `json:"token"`
	Refresh_Token string    `json:"refresh_token"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
	User_id       string    `json:"user_id" validate:"required"`
}
