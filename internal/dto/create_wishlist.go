package dto

type CreateWishlistRequest struct {
	Name        string  `json:"name" binding:"required,max=100"`
	Description *string `json:"description"`
	Color       *string `json:"color"`
}
