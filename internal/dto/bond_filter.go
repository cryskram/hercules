package dto

type BondFilter struct {
	Search          string
	Rating          string
	Sector          string
	CouponType      string
	Nature          string
	PayoutFrequency string
	Sort            string
	Order           string
	Page            int
	Limit           int
}
