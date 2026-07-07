package repository

import (
	"github.com/cryskram/hercules/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WishlistRepository interface {
	Create(wishlist *models.Wishlist) error
	GetAll() ([]models.Wishlist, error)
	GetByID(id string) (*models.Wishlist, error)
	Update(wishlist *models.Wishlist) error
	Delete(id string) error
	AddBond(wishlistID string, bondISIN string) error
	RemoveBond(wishlistID string, bondISIN string) error
	GetWishlistBonds(wishlistID string) ([]models.Bond, error)
	GetBondCount(wishlistID string) (int64, error)
}

type wishlistRepository struct {
	db *gorm.DB
}

func NewWishlistRepository(db *gorm.DB) WishlistRepository {
	return &wishlistRepository{
		db: db,
	}
}

func (r *wishlistRepository) Create(
	wishlist *models.Wishlist,
) error {

	return r.db.Create(wishlist).Error
}

func (r *wishlistRepository) GetAll() ([]models.Wishlist, error) {

	var wishlists []models.Wishlist

	err := r.db.
		Order("created_at DESC").
		Find(&wishlists).
		Error

	return wishlists, err
}

func (r *wishlistRepository) GetByID(
	id string,
) (*models.Wishlist, error) {

	var wishlist models.Wishlist

	err := r.db.
		Where("id = ?", id).
		First(&wishlist).
		Error

	if err != nil {
		return nil, err
	}

	return &wishlist, nil
}

func (r *wishlistRepository) Update(
	wishlist *models.Wishlist,
) error {

	return r.db.Save(wishlist).Error
}

func (r *wishlistRepository) Delete(
	id string,
) error {

	return r.db.
		Delete(&models.Wishlist{}, "id = ?", id).
		Error
}

func (r *wishlistRepository) AddBond(
	wishlistID string,
	bondISIN string,
) error {
	item := models.WishlistBond{
		WishlistID: wishlistID,
		BondISIN:   bondISIN,
	}

	return r.db.
		Clauses(clause.OnConflict{
			DoNothing: true,
		}).
		Create(&item).
		Error
}

func (r *wishlistRepository) RemoveBond(
	wishlistID string,
	bondISIN string,
) error {

	return r.db.
		Delete(
			&models.WishlistBond{},
			"wishlist_id = ? AND bond_isin = ?",
			wishlistID,
			bondISIN,
		).
		Error
}

func (r *wishlistRepository) GetBondCount(
	wishlistID string,
) (int64, error) {

	var count int64

	err := r.db.
		Model(&models.WishlistBond{}).
		Where("wishlist_id = ?", wishlistID).
		Count(&count).
		Error

	return count, err
}

func (r *wishlistRepository) GetWishlistBonds(
	wishlistID string,
) ([]models.Bond, error) {

	var bonds []models.Bond

	err := r.db.
		Table("bonds").
		Joins(
			"JOIN wishlist_bonds wb ON wb.bond_isin = bonds.isin",
		).
		Where(
			"wb.wishlist_id = ?",
			wishlistID,
		).
		Order("yield_pct DESC").
		Find(&bonds).
		Error

	return bonds, err
}
