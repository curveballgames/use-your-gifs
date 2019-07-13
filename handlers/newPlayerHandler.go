package handlers

import "github.com/gin-gonic/gin"

// HandleNewPlayer handles a new player connecting to a game room
func HandleNewPlayer(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "player connected",
	})
}
