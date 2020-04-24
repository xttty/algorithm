package algorithm

import (
	"testing"
)

func TestSort(t *testing.T) {
	nums := []int{1, 23, 423, 5, 345, 22, 90, 122}
	// QSort(nums)
	// BubbleSort(nums)
	// InsertSort(nums)
	// SelectSort(nums)
	// MergeSort(nums)
	HeapSort(nums)
	t.Log(nums)
}
