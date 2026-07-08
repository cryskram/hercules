package repository

import (
	"errors"
	"fmt"

	"github.com/cryskram/hercules/internal/dto"
	"github.com/cryskram/hercules/internal/models"

	"gorm.io/gorm"
)

type BondRepository interface {
	GetAll(filter dto.BondFilter) ([]models.Bond, int64, error)
	GetByISIN(isin string) (*models.Bond, error)
}

type bondRepository struct {
	db *gorm.DB
}

func NewBondRepository(db *gorm.DB) BondRepository {
	return &bondRepository{
		db: db,
	}
}

func (r *bondRepository) GetByISIN(isin string) (*models.Bond, error) {

	var bond models.Bond

	err := r.db.
		Where("isin = ?", isin).
		First(&bond).
		Error

	if err != nil {
		return nil, err
	}

	return &bond, nil
}

func (r *bondRepository) GetAll(filter dto.BondFilter) ([]models.Bond, int64, error) {

	var (
		bonds []models.Bond
		total int64
	)

	query := r.db.Model(&models.Bond{})

	if filter.Rating != "" {
		query = query.Where("rating = ?", filter.Rating)
	}

	if filter.Sector != "" {
		query = query.Where("sector = ?", filter.Sector)
	}

	if filter.CouponType != "" {
		query = query.Where("coupon_type = ?", filter.CouponType)
	}

	if filter.Nature != "" {
		query = query.Where("nature = ?", filter.Nature)
	}

	if filter.PayoutFrequency != "" {
		query = query.Where("payout_frequency = ?", filter.PayoutFrequency)
	}

	if filter.PrincipalFrequency != "" {
		query = query.Where("principal_frequency = ?",
			filter.PrincipalFrequency,
		)
	}

	if filter.Search != "" {

		search := "%" + filter.Search + "%"

		query = query.Where(`
			bond_name ILIKE ?
			OR brand_name ILIKE ?
			OR isin ILIKE ?
			OR similarity(bond_name, ?) > 0.2
			OR similarity(brand_name, ?) > 0.2
		`,
			search,
			search,
			search,
			filter.Search,
			filter.Search,
		)

	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if filter.Page <= 0 {
		return nil, 0, errors.New("page must be greater than 0")
	}

	if filter.Limit <= 0 {
		return nil, 0, errors.New("limit cannot be negative or 0")
	}

	if filter.Limit > 100 {
		filter.Limit = 100
	}

	allowedSorts := map[string]string{
		"yield":       "yield_pct",
		"investment":  "min_investment",
		"maturity":    "maturity_date",
		"name":        "bond_name",
		"coupon_rate": "coupon_rate",
		"face_value":  "face_value",
		"brand":       "brand_name",
	}

	sortColumn := "yield_pct"

	if column, ok := allowedSorts[filter.Sort]; ok {
		sortColumn = column
	}

	order := "DESC"

	if filter.Order == "asc" || filter.Order == "ASC" {
		order = "ASC"
	}

	if filter.Sort == "rating" {

		ratingOrder := `
			CASE rating
				WHEN 'Sovereign' THEN 1
				WHEN 'AAA' THEN 2
				WHEN 'AA+' THEN 3
				WHEN 'AA' THEN 4
				WHEN 'AA-' THEN 5
				WHEN 'A+' THEN 6
				WHEN 'A' THEN 7
				WHEN 'A-' THEN 8
				WHEN 'BBB+' THEN 9
				WHEN 'BBB' THEN 10
				WHEN 'BBB-' THEN 11
				WHEN 'BB+' THEN 12
				WHEN 'BB' THEN 13
				WHEN 'BB-' THEN 14
				WHEN 'B+' THEN 15
				WHEN 'B' THEN 16
				WHEN 'B-' THEN 17
				ELSE 999
			END
			`

		query = query.Order(fmt.Sprintf("%s %s", ratingOrder, order))

	} else if filter.Sort == "principal_frequency" {

		principalOrder := `
			CASE principal_frequency
				WHEN 'Monthly' THEN 1
				WHEN 'Quarterly' THEN 2
				WHEN 'Semi Annually' THEN 3
				WHEN 'Annual' THEN 4
				ELSE 999
			END
			`
		query = query.Order(fmt.Sprintf("%s %s", principalOrder, order))

	} else {

		query = query.Order(fmt.Sprintf("%s %s", sortColumn, order))

	}
	err := query.
		Offset((filter.Page - 1) * filter.Limit).
		Limit(filter.Limit).
		Find(&bonds).
		Error

	if err != nil {
		return nil, 0, err
	}

	return bonds, total, nil
}
