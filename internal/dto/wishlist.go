package dto

type WishlistResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Color       *int    `json:"color"`

	BondCount int `json:"bondCount"`
}
