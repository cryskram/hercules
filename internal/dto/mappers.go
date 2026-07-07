package dto

import (
	"time"

	"github.com/cryskram/hercules/internal/models"
)

func ToBondResponse(b models.Bond) BondResponse {
	return BondResponse{
		ISIN:          b.ISIN,
		BondName:      b.BondName,
		BrandName:     b.BrandName,
		LogoURL:       b.LogoURL,
		YieldPct:      b.YieldPct,
		Rating:        b.Rating,
		Sector:        b.Sector,
		MinInvestment: b.MinInvestment,
		MaturityDate:  b.MaturityDate.Format(time.DateOnly),
	}
}
