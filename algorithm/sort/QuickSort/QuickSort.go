package QuickSort

import "math/rand"

// QuickSort 快速排序
func QuickSort(values []int) []int {
	if len(values) < 2 {
		return values
	}
	left, right := 0, len(values)-1
	pivot := rand.Int() % len(values)
	values[pivot], values[right] = values[right], values[pivot]
	for i := range values {
		if values[i] < values[right] {
			values[left], values[i] = values[i], values[left]
			left++
		}
	}
	values[left], values[right] = values[right], values[left]
	QuickSort(values[:left])
	QuickSort(values[left+1:])
	return values
}
