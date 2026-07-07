package models

import "time"

type WishlistBond struct {
	WishlistID string    `gorm:"column:wishlist_id;primaryKey"`
	BondISIN   string    `gorm:"column:bond_isin;primaryKey"`
	AddedAt    time.Time `gorm:"column:added_at"`
}

func (WishlistBond) TableName() string {
	return "wishlist_bonds"
}
