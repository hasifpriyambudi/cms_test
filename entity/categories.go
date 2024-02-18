package entity

import "time"

type CategoriesEntity struct {
	Id            int        `json:"id"`
	Name_Category string     `json:"name_categories"`
	Created_At    *time.Time `json:"created_at"`
	Updated_At    *time.Time `json:"updated_at"`
	Deleted_At    *time.Time `json:"deleted_at"`
}
