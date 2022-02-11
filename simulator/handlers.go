package simulator

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type setArenaRequest struct {
	Arena [][]string `json:"arena"`
}

func HandleSetArena(c *gin.Context) {
	var body setArenaRequest
	if err := c.BindJSON(&body); err != nil {
		return
	}

	setArena(body.Arena)
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func HandleGetHamiltonPath(c *gin.Context) {
	plan := getHamiltonPath()

	c.JSON(http.StatusOK, gin.H{
		"cost":         plan.Cost,
		"path":         plan.Path,
		"transition":   plan.Path.ToTransition().ToStringArray(),
		"order":        plan.Order,
		"detect_image": plan.DetectImageIndices,
	})
}
