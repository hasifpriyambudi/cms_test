package domain

type CustomPageCreateResponse struct {
	Id int `json:"id"`
}
type CustomPageCreateRequest struct {
	Custom_Url   string `validate:"required,min=3,max=30" json:"custom_url"`
	Page_Content string `validate:"required" json:"page_content"`
}

type CustomPageUpdateRequest struct {
	Id           int    `validate:"required" json:"id"`
	Custom_Url   string `validate:"required,min=3,max=30" json:"custom_url"`
	Page_Content string `validate:"required" json:"page_content"`
}
type CustomPageUpdateResponse struct {
	Id           int    `json:"id"`
	Custom_Url   string `json:"custom_url"`
	Page_Content string `json:"page_content"`
}

type CustomPageGetResponse struct {
	Id           int    `json:"id"`
	Custom_Url   string `json:"custom_url"`
	Page_Content string `json:"page_content"`
}
