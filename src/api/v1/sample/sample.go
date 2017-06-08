package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"errors"
	"models"
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

func Post(ctx *gin.Context) {
	var team models.Team
	if err := ctx.BindJSON(&team); err != nil {
		ctx.JSON(400, gin.H{"message": "Invalid reuqest"})
		ctx.AbortWithError(400, errors.New("Failed to bind json"))
	} else {
		fmt.Println("team_name:", team.Team_name)
		ctx.JSON(200, team)
	}
}
