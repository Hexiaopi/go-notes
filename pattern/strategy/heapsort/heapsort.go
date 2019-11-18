package heapsort

type HeapSort struct {
}

// Parent 父节点
func (h HeapSort) Parent(i int) (j int) {
	return i / 2
}

// Left 左孩子节点
func (h HeapSort) Left(i int) (j int) {
	return 2 * i
}

// Right 右孩子节点
func (h HeapSort) Right(i int) (j int) {
	return 2*i + 1
}

// MaxHeapIfy 递归维护最大堆性质
func (h HeapSort) MaxHeapIfy(values []int, i int) {
	var largest int
	l := h.Left(i)
	r := h.Right(i)
	if l < len(values) && values[l] > values[i] {
		largest = l
	} else {
		largest = i
	}

	if r < len(values) && values[r] > values[largest] {
		largest = r
	}

	if largest != i {
		values[i], values[largest] = values[largest], values[i]
		h.MaxHeapIfy(values, largest)
	}
}

// BuildMaxHeap 构建最大堆
func (h HeapSort) BuildMaxHeap(values []int) {
	for i := len(values) / 2; i >= 0; i-- {
		h.MaxHeapIfy(values, i)
	}
}

// heapsort 堆排序算法
func (h HeapSort) Sort(values []int) {
	h.BuildMaxHeap(values)
	for i := len(values) - 1; i > 0; i-- {
		values[0], values[i] = values[i], values[0]
		values = values[:len(values)-1]
		h.MaxHeapIfy(values, 0)
	}
}
