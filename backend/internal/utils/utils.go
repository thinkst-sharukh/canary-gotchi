package utils

import (
	"math/rand"
	"os"
	"strings"
	"time"
)

var directions = []string{"up", "right", "left", "down"}

func IsDev() bool {
	return os.Getenv("ENVIRONMENT") == "development"
}

func GenerateRandomSequence(length int) string {
	// Create a new random generator with a unique seed
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	sequence := make([]string, length)
	for i := 0; i < length; i++ {
		sequence[i] = directions[rng.Intn(len(directions))]
	}
	return strings.Join(sequence, ",")
}
