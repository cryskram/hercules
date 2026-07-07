package handlers

import (
	"net/http"
	"strconv"

	"github.com/cryskram/hercules/internal/dto"
	"github.com/cryskram/hercules/internal/services"
	"github.com/cryskram/hercules/internal/utils"

	"github.com/gin-gonic/gin"
)

type BondHandler struct {
	service services.BondService
}

func NewBondHandler(service services.BondService) *BondHandler {
	return &BondHandler{
		service: service,
	}
}

// GET /api/bonds
func (h *BondHandler) GetAll(c *gin.Context) {
	filter := dto.BondFilter{
		Search:          c.Query("search"),
		Rating:          c.Query("rating"),
		Sector:          c.Query("sector"),
		CouponType:      c.Query("coupon_type"),
		Nature:          c.Query("nature"),
		PayoutFrequency: c.Query("payout_frequency"),
		Sort:            c.DefaultQuery("sort", "yield"),
		Order:           c.DefaultQuery("order", "desc"),
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err == nil {
		filter.Page = page
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err == nil {
		filter.Limit = limit
	}

	response, err := h.service.GetAll(filter)

	if err != nil {
		utils.Error(
			c,
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	utils.OK(c, response)
}

// GET /api/bonds/:isin
func (h *BondHandler) GetByISIN(c *gin.Context) {
	isin := c.Param("isin")

	bond, err := h.service.GetByISIN(isin)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Bond not found",
		})
		return
	}

	utils.OK(c, bond)
}
