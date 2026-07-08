package dto

type BondFilter struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`

	Search string `form:"search"`

	Rating string `form:"rating"`
	Sector string `form:"sector"`

	CouponType string `form:"coupon_type"`

	PayoutFrequency string `form:"payout_frequency"`

	PrincipalFrequency string `form:"principal_frequency"`

	Nature string `form:"nature"`

	Sort  string `form:"sort"`
	Order string `form:"order"`
}
