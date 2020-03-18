package algorithm_test

import (
	"algorithm"
	"testing"
)

func TestSort(t *testing.T) {
	nums := []int{1, 23, 423, 5, 345, 22, 90, 122}
	// algorithm.QSort(nums)
	// algorithm.BubbleSort(nums)
	// algorithm.InsertSort(nums)
	// algorithm.SelectSort(nums)
	// algorithm.MergeSort(nums)
	algorithm.HeapSort(nums)
	t.Log(nums)
}
