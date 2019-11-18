package InsertSort

// InsertSort 插入排序
func InsertSort(values []int) {
	for i := 1; i < len(values); i++ {
		key := values[i]
		var j int
		for j = i - 1; j >= 0 && values[j] > key; j-- {
			values[j+1] = values[j]
		}
		values[j+1] = key
	}
}
