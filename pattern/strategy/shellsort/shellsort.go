package shellsort

// shellsort 希尔排序（一种分组插入排序）
type ShellSort struct {
}

func (s ShellSort) Sort(values []int) {
	var (
		n    = len(values)
		gaps = []int{1}
		k    = 1
	)
	for {
		gap := s.element(2, k) + 1
		if gap > n-1 {
			break
		}
		gaps = append([]int{gap}, gaps...)
		k++
	}
	for _, gap := range gaps {
		for i := gap; i < n; i += gap {
			j := i
			for j > 0 {
				if values[j-gap] > values[j] {
					values[j-gap], values[j] = values[j], values[j-gap]
				}
				j = j - gap
			}
		}
	}
}

func (s ShellSort) element(a, b int) int {
	e := 1
	for b > 0 {
		if b&1 != 0 {
			e *= a
		}
		b >>= 1
		a *= a
	}
	return e
}
