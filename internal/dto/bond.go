package dto

type BondResponse struct {
	ISIN string `json:"isin"`

	BondName  string `json:"bondName"`
	BrandName string `json:"brandName"`

	LogoURL *string `json:"logoUrl"`

	YieldPct float64 `json:"yield"`

	Rating *string `json:"rating"`

	Sector string `json:"sector"`

	MinInvestment float64 `json:"minInvestment"`

	MaturityDate string `json:"maturityDate"`
}
