package entity

import "time"

type UserEntity struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Username   string     `json:"username"`
	Password   string     `json:"password"`
	Created_At *time.Time `json:"created_at"`
	Updated_At *time.Time `json:"updated_at"`
	Deleted_At *time.Time `json:"deleted_at"`
}
