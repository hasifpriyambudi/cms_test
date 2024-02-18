package domain

type NewsCreateResponse struct {
	Id int `json:"id"`
}
type NewsCreateRequest struct {
	Category_Id  int    `validate:"required" json:"category_id"`
	News_Content string `validate:"required" json:"news_content"`
}

type NewsUpdateRequest struct {
	Id           int    `validate:"required" json:"id"`
	Category_Id  int    `validate:"required" json:"category_id"`
	News_Content string `validate:"required" json:"news_content"`
}
type NewsUpdateResponse struct {
	Id           int    `json:"id"`
	Category_Id  int    `json:"category_id"`
	News_Content string `json:"news_content"`
}

type NewsGetResponse struct {
	Id           int    `json:"id"`
	Category_Id  int    `json:"category_id"`
	News_Content string `json:"news_content"`
}
