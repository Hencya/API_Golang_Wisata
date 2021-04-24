package dto

//DestinationUpdateDTO is a model that client use when updating a destination
type DestinationUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Title       string `json:"title" from:"title" binding:"required"`
	Description string `json:"description" from:"description" binding:"required"`
	Keyword     string `json:"keyword" from:"keyword" binding:"required"`
	Images      string `json:"images" from:"images" binding:"required"`
	Address     string `json:"address" from:"images" binding:"required"`
	Views       uint64 `json:"views" from:"views" binding:"required"`
	UserID      uint64 `json:"user_id,omiempty" from:"user_id,omiempty"`
}

//DestinationCreateDTO is is a model that clinet use when create a new destination
type DestinationCreateDTO struct {
	Title       string `json:"title" from:"title" binding:"required"`
	Description string `json:"description" from:"description" binding:"required"`
	Keyword     string `json:"keyword" from:"keyword" binding:"required"`
	Images      string `json:"images" from:"images" binding:"required"`
	Address     string `json:"address" from:"images" binding:"required"`
	Views       uint64 `json:"views" from:"views" binding:"required"`
	UserID      uint64 `json:"user_id,omiempty" from:"user_id,omiempty"`
}
