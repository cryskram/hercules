package routes

import (
	"github.com/cryskram/hercules/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, bondHandler *handlers.BondHandler) {
	api := router.Group("/api")
	{
		bonds := api.Group("/bonds")
		{
			bonds.GET("", bondHandler.GetAll)
			bonds.GET("/:isin", bondHandler.GetByISIN)
		}
	}
}
