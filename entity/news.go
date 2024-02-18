package entity

import "time"

type NewsEntity struct {
	Id           int
	Category_Id  int
	News_Content string
	Created_At   *time.Time
	Updated_At   *time.Time
	Deleted_At   *time.Time
}
