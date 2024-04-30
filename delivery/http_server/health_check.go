package httpserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Every thing is Good!",
	})
}
