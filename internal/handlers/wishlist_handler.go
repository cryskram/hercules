package handlers

import (
	"net/http"

	"github.com/cryskram/hercules/internal/dto"
	"github.com/cryskram/hercules/internal/services"
	"github.com/cryskram/hercules/internal/utils"

	"github.com/gin-gonic/gin"
)

type WishlistHandler struct {
	service services.WishlistService
}

func NewWishlistHandler(service services.WishlistService) *WishlistHandler {
	return &WishlistHandler{
		service: service,
	}
}

// GET /api/wishlists
func (h *WishlistHandler) GetAll(c *gin.Context) {

	wishlists, err := h.service.GetAll()

	if err != nil {
		utils.Error(
			c,
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	utils.OK(c, wishlists)
}

// GET /api/wishlists/:id
func (h *WishlistHandler) GetByID(c *gin.Context) {

	id := c.Param("id")

	wishlist, err := h.service.GetByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Wishlist not found",
		})
		return
	}

	utils.OK(c, wishlist)
}

// POST /api/wishlists
func (h *WishlistHandler) Create(c *gin.Context) {
	var req dto.CreateWishlistRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.service.Create(req)

	if err != nil {
		switch err.Error() {
		case "maximum of 5 wishlists allowed",
			"wishlist can contain at most 10 bonds":
			utils.Error(c, http.StatusConflict, err.Error())
		default:
			utils.Error(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Wishlist created successfully",
	})
}

// PATCH /api/wishlists/:id
func (h *WishlistHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdateWishlistRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.service.Update(id, req)

	if err != nil {
		utils.Error(
			c,
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	utils.Message(
		c,
		http.StatusCreated,
		"Wishlist created successfully",
	)
}

// DELETE /api/wishlists/:id
func (h *WishlistHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(id)

	if err != nil {
		utils.Error(
			c,
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	utils.Message(
		c,
		http.StatusCreated,
		"Wishlist deleted successfully",
	)
}

// GET /api/wishlists/:id/bonds
func (h *WishlistHandler) GetWishlistBonds(c *gin.Context) {
	id := c.Param("id")
	bonds, err := h.service.GetWishlistBonds(id)

	if err != nil {
		utils.Error(
			c,
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	c.JSON(http.StatusOK, bonds)
}

// POST /api/wishlists/:id/bonds
func (h *WishlistHandler) AddBond(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		ISIN string `json:"isin" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.service.AddBond(id, body.ISIN)

	if err != nil {
		utils.Error(
			c,
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Bond added successfully",
	})
}

// DELETE /api/wishlists/:id/bonds/:isin
func (h *WishlistHandler) RemoveBond(c *gin.Context) {
	id := c.Param("id")
	isin := c.Param("isin")
	err := h.service.RemoveBond(id, isin)

	if err != nil {
		utils.Error(
			c,
			http.StatusInternalServerError,
			err.Error(),
		)

		return
	}

	utils.Message(
		c,
		http.StatusCreated,
		"Bond removed successfully",
	)
}
