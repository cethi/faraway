package proof_of_work

import (
	"math/rand"
	"time"
)

// POWFunction interface to all proof of work functions
type POWFunction interface {
	Compute(payload string) (string, error)
	Validate(task string, solution string) (bool, error)
}

func GenerateRandomTask(length int) string {
	rand.Seed(time.Now().UnixNano())
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
