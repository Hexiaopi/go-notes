package mergesort

import (
	"math"
)

// MergeSort 合并排序
type MergeSort struct {
	Low  int
	Mid  int
	High int
}

// Merge 合并两个数组
func (m MergeSort) Merge(values []int, p, q, r int) {
	n1 := q - p + 1
	n2 := r - q
	// 申请额外的空间
	L := make([]int, n1+1)
	R := make([]int, n2+1)
	for i := 0; i < n1; i++ {
		L[i] = values[p+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = values[q+j+1]
	}
	L[n1] = math.MaxInt32 //哨兵
	R[n2] = math.MaxInt32 //哨兵
	i, j := 0, 0
	for k := p; k <= r; k++ {
		if L[i] <= R[j] {
			values[k] = L[i]
			i++
		} else {
			values[k] = R[j]
			j++
		}
	}
}

// Sort 合并排序，使用递归
func (m MergeSort) Sort(values []int) {
	if m.Low < m.High {
		mid := (m.Low + m.High) / 2
		m.Mid = mid
		m1 := MergeSort{Low: m.Low, High: m.Mid}
		m1.Sort(values)
		m2 := MergeSort{Low: m.Mid+1, High: m.High}
		m2.Sort(values)
		m.Merge(values, m.Low, m.Mid, m.High)
	}
}
