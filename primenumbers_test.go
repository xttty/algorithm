package algorithm_test

import (
	"algorithm"
	"testing"
)

func TestPrimeNumbersChanBuilder(t *testing.T) {
	numbers := algorithm.PrimeNumbersChanBuilder(100)
	// numbers := algorithm.PrimeNumbersBuilder(100)
	t.Log(numbers)
}
