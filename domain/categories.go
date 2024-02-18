package domain

type CategoriesCreateResponse struct {
	Id int `json:"id"`
}
type CategoriesCreateRequest struct {
	Name string `validate:"required,min=3,max=30" json:"name"`
}

type CategoriesUpdateRequest struct {
	Id   int    `json:"id"`
	Name string `validate:"required,min=3,max=30" json:"name"`
}
type CategoriesUpdateResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CategoriesDeleteResponse struct {
	Id int `json:"id"`
}

type CategoriesGetResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
