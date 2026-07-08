package dto

type UpdateWishlistRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Color       *int    `json:"color"`
}
