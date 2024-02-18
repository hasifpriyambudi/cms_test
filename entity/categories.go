package entity

import "time"

type CategoriesEntity struct {
	Id            int
	Name_Category string
	Created_At    *time.Time
	Updated_At    *time.Time
	Deleted_At    *time.Time
}
