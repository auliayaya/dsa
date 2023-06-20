package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	number1 := []float64{3.5, 2.4, 12.8, 9.1}
	number2 := []float64{3.5, 2.4, 12.8, 9.1}

	name1 := []string{"Ali", "Ade", "Rocky", "Cerce"}
	name2 := []string{"Ali", "Ade", "Rocky", "Cerce"}
	bubbleSort[float64](number1)
	assert.NotEqual(t, number1, number2)
	bubbleSort[string](name1)
	assert.NotEqual(t, name1, name2)
}
