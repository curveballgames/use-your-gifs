package util

import (
	"math/rand"
	"strings"
)

const (
	randRoomCodeLetterCount = 4
)

var (
	letters = [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
)

// GenerateRoomCode creates a pseudorandom room code for sessions.
func GenerateRoomCode() string {
	stringBuilder := strings.Builder{}

	for i := 0; i < randRoomCodeLetterCount; i++ {
		stringBuilder.WriteString(letters[int32(rand.Float32()*26)])
	}

	return stringBuilder.String()
}
