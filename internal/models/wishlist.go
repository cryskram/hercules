package models

import "time"

type Wishlist struct {
	ID          string    `gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey"`
	Name        string    `gorm:"column:name"`
	Description *string   `gorm:"column:description"`
	Color       *int      `gorm:"column:color"`
	IsDefault   bool      `gorm:"column:is_default;default:false"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (Wishlist) TableName() string {
	return "wishlists"
}
