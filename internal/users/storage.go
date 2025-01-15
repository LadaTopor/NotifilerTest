package users

import "time"

type User struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"hashed_password"`
	Token          string    `json:"jwt_token"`
	CreatedAt      time.Time `json:"created_at"`
}
