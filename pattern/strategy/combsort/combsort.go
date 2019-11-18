package combsort

// CombSort 梳排序,像梳子一样
type CombSort struct {
}

// 对冒泡排序的改进，每次比较的不是相邻的两个数
func (c CombSort) Sort(values []int) {
	gap := len(values)
	shirnk := 1.3
	sorted := true
	for sorted {
		if gap > 1 {
			gap = int(float64(gap) / shirnk)
		} else {
			gap = 1
			sorted = false
		}
		for i := 0; i+gap < len(values); i++ {
			if values[i] > values[i+gap] {
				values[i+gap], values[i] = values[i], values[i+gap]
			}
		}
	}
}
