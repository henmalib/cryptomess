package utils

import (
	"github.com/gin-gonic/gin"
)

func WriteResponse(ctx *gin.Context, status int, response gin.H) {
	ctx.JSON(status, gin.H{
		"status":   "ok",
		"response": response,
	})
}
