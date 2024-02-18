package entity

import "time"

type CommentEntity struct {
	Id         int
	News_Id    int
	Name       string
	Comment    string
	Created_At *time.Time
	Updated_At *time.Time
	Deleted_At *time.Time
}
