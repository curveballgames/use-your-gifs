package handlers

import (
	"math/rand"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	randRoomCodeLetterCount = 4
)

// HandleNewRoom handles a request for a new room to be created
func HandleNewRoom(c *gin.Context) {
	letters := [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	stringBuilder := strings.Builder{}

	for i := 0; i < randRoomCodeLetterCount; i++ {
		stringBuilder.WriteString(letters[int32(rand.Float32()*26)])
	}

	roomCode := stringBuilder.String()

	c.JSON(200, gin.H{
		"message":  "Room created",
		"roomCode": roomCode,
	})
}
