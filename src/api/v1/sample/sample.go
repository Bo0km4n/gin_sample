package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Hello world!"})
}

func World(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "World is beautiful"})
}

func Issue01(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "practice pull request"})
}
