package utils

import (
	"net/http"

	"github.com/cryskram/hercules/internal/dto"

	"github.com/gin-gonic/gin"
)

func OK[T any](c *gin.Context, data T) {
	c.JSON(http.StatusOK, dto.APIResponse[T]{
		Success: true,
		Data:    data,
	})
}

func Created[T any](c *gin.Context, message string, data T) {
	c.JSON(http.StatusCreated, dto.APIResponse[T]{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Message(c *gin.Context, status int, message string) {
	c.JSON(status, dto.APIResponse[any]{
		Success: true,
		Message: message,
	})
}

func Error(c *gin.Context, status int, err string) {
	c.JSON(status, dto.APIResponse[any]{
		Success: false,
		Error:   err,
	})
}
