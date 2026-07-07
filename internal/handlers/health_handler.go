package handlers

import (
	"github.com/cryskram/hercules/internal/utils"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	utils.OK(c, gin.H{
		"status": "healthy",
	})
}
