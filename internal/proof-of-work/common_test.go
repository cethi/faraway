package proof_of_work

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateTaskRandomFixedLength(t *testing.T) {
	task := GenerateRandomTask(10)
	assert.Len(t, task, 10)
}

func TestGenerateRandomTask(t *testing.T) {
	task1 := GenerateRandomTask(100)
	task2 := GenerateRandomTask(100)
	assert.NotEqual(t, task1, task2)
}

func TestGenerateRandomTaskContainsLetters(t *testing.T) {
	task := GenerateRandomTask(10)
	for _, r := range task {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			t.Fatalf("random task contains incorrect characters: %s", task)
		}
	}

}
