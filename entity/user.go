package entity

import "time"

type UserEntity struct {
	ID         int
	Name       string
	Username   string
	Password   string
	Created_At *time.Time
	Updated_At *time.Time
	Deleted_At *time.Time
}
