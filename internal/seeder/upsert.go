package seeder

import (
	"fmt"

	"github.com/cryskram/hercules/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func UpsertBonds(db *gorm.DB, bonds []models.Bond) error {

	if len(bonds) == 0 {
		return nil
	}

	tx := db.Begin()

	if tx.Error != nil {
		return tx.Error
	}

	err := tx.
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "isin"},
			},
			DoUpdates: clause.AssignmentColumns([]string{
				"bond_name",
				"brand_name",
				"logo_url",
				"website",
				"description",
				"min_investment",
				"min_units",
				"max_units",
				"yield_pct",
				"coupon_rate",
				"coupon_type",
				"payout_frequency",
				"face_value",
				"nature",
				"seniority",
				"mode_of_issue",
				"yield_type",
				"security_cover",
				"debenture_trustee",
				"date_of_issue",
				"maturity_date",
				"rating",
				"rating_agency",
				"rating_date",
				"total_interest",
				"total_principal",
				"total_payout",
				"interest_frequency",
				"principal_frequency",
				"remaining_tenure_months",
				"risk_bucket",
				"sector",
				"yield_bucket",
				"is_active",
				"updated_at",
			}),
		}).
		CreateInBatches(&bonds, 100).
		Error

	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	fmt.Printf("✅ Upserted %d bonds\n", len(bonds))

	return nil
}
