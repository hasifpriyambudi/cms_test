package domain

type AuthRequest struct {
	Username string `validate:"required,min=3,max=30" json:"username"`
	Password string `validate:"required,min=6" json:"password"`
}

type AuthDatabase struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Access_Token string `json:"access_token"`
}
