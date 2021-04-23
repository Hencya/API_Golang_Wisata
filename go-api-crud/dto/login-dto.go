package dto

//LoginDTO is a model that used by client when POST from /login url
type LoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required" validate:"email,alphanum"`
	Password string `json:"password" from:"password" binding:"required" validate:"min:8"`
}
