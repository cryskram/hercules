package models

import "time"

type Bond struct {
	ISIN string `gorm:"column:isin;type:char(12);primaryKey"`

	BondName  string `gorm:"column:bond_name"`
	BrandName string `gorm:"column:brand_name"`

	LogoURL     *string `gorm:"column:logo_url"`
	Website     *string `gorm:"column:website"`
	Description *string `gorm:"column:description"`

	MinInvestment float64 `gorm:"column:min_investment"`
	MinUnits      int     `gorm:"column:min_units"`
	MaxUnits      int     `gorm:"column:max_units"`

	YieldPct   float64 `gorm:"column:yield_pct"`
	CouponRate float64 `gorm:"column:coupon_rate"`

	CouponType      string  `gorm:"column:coupon_type"`
	PayoutFrequency *string `gorm:"column:payout_frequency"`

	FaceValue float64 `gorm:"column:face_value"`

	Nature      string `gorm:"column:nature"`
	Seniority   string `gorm:"column:seniority"`
	ModeOfIssue string `gorm:"column:mode_of_issue"`
	YieldType   string `gorm:"column:yield_type"`

	SecurityCover    *float64 `gorm:"column:security_cover"`
	DebentureTrustee *string  `gorm:"column:debenture_trustee"`

	DateOfIssue  time.Time `gorm:"column:date_of_issue"`
	MaturityDate time.Time `gorm:"column:maturity_date"`

	Rating       *string    `gorm:"column:rating"`
	RatingAgency *string    `gorm:"column:rating_agency"`
	RatingDate   *time.Time `gorm:"column:rating_date"`

	TotalInterest  float64 `gorm:"column:total_interest"`
	TotalPrincipal float64 `gorm:"column:total_principal"`
	TotalPayout    float64 `gorm:"column:total_payout"`

	InterestFrequency  *string `gorm:"column:interest_frequency"`
	PrincipalFrequency *string `gorm:"column:principal_frequency"`

	RemainingTenureMonths int `gorm:"column:remaining_tenure_months"`

	RiskBucket  string `gorm:"column:risk_bucket"`
	Sector      string `gorm:"column:sector"`
	YieldBucket string `gorm:"column:yield_bucket"`

	IsActive bool `gorm:"column:is_active"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Bond) TableName() string {
	return "bonds"
}
