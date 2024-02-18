package domain

type CommentCreateRequest struct {
	News_Id int    `validate:"required,number" json:"news_id"`
	Name    string `validate:"required,min=3,max=50" json:"name"`
	Comment string `validate:"required,min=3" json:"comment"`
}
