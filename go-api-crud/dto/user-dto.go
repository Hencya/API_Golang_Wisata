package dto

//UserUpdateDTO is used by client when PUT update profile
type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email,alphanum"`
	Password string `json:"password,omiempty" from:" password,omiempty" validate:"min:8"`
}

type CreateUserDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email,alphanum"`
	Password string `json:"password,omiempty" from:" password,omiempty" validate:"min:8" binding:"required"`
}
