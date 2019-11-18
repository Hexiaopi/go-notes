package quicksort

import "math/rand"

// quicksort 快速排序
type QuickSort struct {
}

func (q QuickSort) Sort(values []int) {
	if len(values) < 2 {
		return
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
	q.Sort(values[:left])
	q.Sort(values[left+1:])
	return
}
