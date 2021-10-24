package car

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverageAge(t *testing.T) {
	cars := []Car{{2000}, {2001}, {2002}}
	avg, err := AverageAge(cars)

	assert.NoError(t, err)
	assert.Equal(
		t,
		float32(20),
		avg,
		fmt.Sprintf("Expecting 20, got: %v", avg))
}
