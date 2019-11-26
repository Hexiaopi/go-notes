package pancakesort

// pancakesort 煎饼排序
type PancakeSort struct {
}

func (p PancakeSort) Sort(values []int) {
	for i := len(values) - 1; i > 0; i-- {
		max, maxValue := 0, values[0]
		for j := 1; j <= i; j++ {
			if values[j] > maxValue {
				max, maxValue = j, values[j]
			}
		}
		p.flip(values, max)
		p.flip(values, i)
	}
}

func (p PancakeSort) flip(values []int, position int) {
	for i := 0; i < position; i, position = i+1, position-1 {
		values[i], values[position] = values[position], values[i]
	}
}
