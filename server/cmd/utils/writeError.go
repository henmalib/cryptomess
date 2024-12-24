package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WriteError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"status": "error",
		"error":  err.Error(),
	})
}
