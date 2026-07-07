package routes

import (
	"net/http"

	"github.com/cryskram/hercules/internal/handlers"
	"github.com/cryskram/hercules/internal/utils"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, bondHandler *handlers.BondHandler, wishlistHandler *handlers.WishlistHandler) {
	router.GET("/health", handlers.Health)
	router.GET("/", func(c *gin.Context) {
		utils.Message(c, http.StatusOK, "Welcome to hercules")
	})

	api := router.Group("/api")
	{

		bonds := api.Group("/bonds")
		{
			bonds.GET("", bondHandler.GetAll)
			bonds.GET("/:isin", bondHandler.GetByISIN)
		}

		wishlists := api.Group("/wishlists")
		{
			wishlists.GET("", wishlistHandler.GetAll)
			wishlists.POST("", wishlistHandler.Create)
			wishlists.GET("/:id", wishlistHandler.GetByID)
			wishlists.PATCH("/:id", wishlistHandler.Update)
			wishlists.DELETE("/:id", wishlistHandler.Delete)
			wishlists.GET("/:id/bonds", wishlistHandler.GetWishlistBonds)
			wishlists.POST("/:id/bonds", wishlistHandler.AddBond)
			wishlists.DELETE("/:id/bonds/:isin", wishlistHandler.RemoveBond)
		}
	}
}
