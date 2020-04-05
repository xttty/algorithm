package algorithm

import (
	"testing"
)

func TestPrimeNumbersChanBuilder(t *testing.T) {
	numbers := PrimeNumbersChanBuilder(100)
	// numbers := algorithm.PrimeNumbersBuilder(100)
	t.Log(numbers)
}
