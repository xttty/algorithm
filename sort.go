package algorithm

// BubbleSort 冒泡排序
func BubbleSort(nums []int) {
	l := len(nums)
	if l <= 1 {
		return
	}
	flag := false
	for i := l; i > 0; i-- {
		if i < l && !flag {
			break
		}
		flag = false
		for j := 1; j < i; j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
				flag = true
			}
		}
	}
}

// InsertSort 插入排序
func InsertSort(nums []int) {
	l := len(nums)
	if l <= 1 {
		return
	}
	for i := 1; i < l; i++ {
		for j := i; j > 0 && nums[j] < nums[j-1]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
}

// SelectSort 选择排序
func SelectSort(nums []int) {
	l := len(nums)
	if l <= 1 {
		return
	}
	for i := 0; i < l; i++ {
		min := i
		for j := i + 1; j < l; j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
}

// QSort 快速排序
func QSort(nums []int) {
	qSort(nums, 0, len(nums)-1)
}

func qSort(nums []int, start, end int) {
	pos := partition(nums, start, end)
	if pos > start {
		qSort(nums, start, pos-1)
	}
	if pos < end {
		qSort(nums, pos+1, end)
	}
}

func partition(nums []int, start, end int) int {
	pos := nums[start]
	for start < end {
		for start < end && nums[end] > pos {
			end--
		}
		nums[start] = nums[end]
		for start < end && nums[start] < pos {
			start++
		}
		nums[end] = nums[start]
	}
	nums[start] = pos
	return start
}

// MergeSort 归并排序
func MergeSort(nums []int) {
	l := len(nums)
	if l <= 1 {
		return
	}
	for i := 1; i <= l/2; i *= 2 {
		start, mid, end := 0, i-1, 2*i-1
		for end < l {
			merge(nums, start, mid, end)
			start = end + 1
			mid, end = start+i-1, start+2*i-1
		}
	}
}

func merge(nums []int, start, mid, end int) {
	tempArr := make([]int, end-start+1)
	i, j, k := start, mid+1, 0
	for i <= mid && j <= end {
		if nums[i] < nums[j] {
			tempArr[k] = nums[i]
			i++
		} else {
			tempArr[k] = nums[j]
			j++
		}
		k++
	}
	for i <= mid {
		tempArr[k] = nums[i]
		k++
		i++
	}
	for j <= end {
		tempArr[k] = nums[j]
		k++
		j++
	}
	for i = 0; i < len(tempArr); i++ {
		nums[start+i] = tempArr[i]
	}
}

// HeapSort 堆排序
func HeapSort(nums []int) {
	len := len(nums)
	// 构建大顶堆
	for i := len/2 - 1; i >= 0; i-- {
		adjustHeap(nums, i, len-i)
	}
	for i := len - 1; i > 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		adjustHeap(nums, 0, i)
	}
}

func adjustHeap(nums []int, top, len int) {
	temp := nums[top]
	for i := top; i < top+len/2; {
		child := i*2 + 1
		if child+1 < len && nums[child+1] > nums[child] {
			child++
		}
		if nums[child] > nums[i] {
			nums[i] = nums[child]
			i = child
		} else {
			break
		}
		nums[i] = temp
	}
}
