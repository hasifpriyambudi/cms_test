package entity

import "time"

type CustomPageEntity struct {
	Id           int
	Custom_Url   string
	Page_Content string
	Created_At   *time.Time
	Updated_At   *time.Time
	Deleted_At   *time.Time
}
