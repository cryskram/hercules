package seeder

import (
	"fmt"
	"strings"

	"github.com/cryskram/hercules/internal/models"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

func SeedBonds(db *gorm.DB, path string) error {

	file, err := excelize.OpenFile(path)
	if err != nil {
		return err
	}
	defer file.Close()

	sheet := file.GetSheetName(0)

	rows, err := file.GetRows(sheet)
	if err != nil {
		return err
	}

	if len(rows) < 2 {
		return fmt.Errorf("excel file contains no data")
	}

	headerMap := make(map[string]int)

	for i, header := range rows[0] {
		headerMap[strings.TrimSpace(header)] = i
	}

	get := func(row []string, column string) string {

		index, ok := headerMap[column]

		if !ok || index >= len(row) {
			return ""
		}

		return strings.TrimSpace(row[index])
	}

	var bonds []models.Bond

	for _, row := range rows[1:] {

		bond := models.Bond{

			ISIN: get(row, "ISIN"),

			BondName:  get(row, "Bond Name"),
			BrandName: get(row, "Brand Name"),

			LogoURL:     strPtr(get(row, "Logo URL")),
			Website:     strPtr(get(row, "Website")),
			Description: strPtr(get(row, "Description")),

			MinInvestment: parseFloat(get(row, "Min Investment")),
			MinUnits:      parseInt(get(row, "Min Units")),
			MaxUnits:      parseInt(get(row, "Max Units")),

			YieldPct:   parseFloat(get(row, "Yield")),
			CouponRate: parseFloat(get(row, "Coupon Rate")),

			CouponType:      get(row, "Coupon Type"),
			PayoutFrequency: strPtr(get(row, "Payout Frequency")),

			FaceValue: parseFloat(get(row, "Face Value")),

			Nature:      get(row, "Nature"),
			Seniority:   get(row, "Seniority"),
			ModeOfIssue: get(row, "Mode of Issue"),
			YieldType:   get(row, "Yield Type"),

			SecurityCover:    floatPtr(get(row, "Security Cover")),
			DebentureTrustee: strPtr(get(row, "Debenture Trustee")),

			DateOfIssue:  parseDate(get(row, "Date of Issue")),
			MaturityDate: parseDate(get(row, "Maturity Date")),

			Rating:       strPtr(get(row, "Rating")),
			RatingAgency: strPtr(get(row, "Rating Agency")),
			RatingDate:   datePtr(get(row, "Rating Date")),

			TotalInterest:  parseFloat(get(row, "Total Interest")),
			TotalPrincipal: parseFloat(get(row, "Total Principal")),
			TotalPayout:    parseFloat(get(row, "Total Payout")),

			InterestFrequency:  strPtr(get(row, "Interest Frequency")),
			PrincipalFrequency: strPtr(get(row, "Principal Frequency")),

			RemainingTenureMonths: parseInt(get(row, "Remaining Tenure (months)")),

			RiskBucket:  get(row, "Risk Bucket"),
			Sector:      get(row, "Sector"),
			YieldBucket: get(row, "Yield Bucket"),

			IsActive: true,
		}

		bonds = append(bonds, bond)
	}

	fmt.Printf("📄 Parsed %d bonds\n", len(bonds))

	return UpsertBonds(db, bonds)
}
