package selectsort

// SelectSort 选择排序
type SelectSort struct {
}

func (s SelectSort) Sort(values []int) {
	for i := 0; i < len(values)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(values); j++ {
			if values[j] < values[minIndex] {
				minIndex = j
			}
		}
		values[i], values[minIndex] = values[minIndex], values[i]
	}
}
