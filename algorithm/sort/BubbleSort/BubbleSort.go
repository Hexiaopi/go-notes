package BubbleSort

//BubbleSort 冒泡排序
func BubbleSort(values []int) {
	for i := 0; i < len(values)-1; i++ {
		sorted := true
		//最大的泡泡冒至数组后面
		/* for j := 0; j < len(values)-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				sorted = false
			}
		} */
		//最小的泡泡冒至数组前面
		for j := len(values) - 1; j > i; j-- {
			if values[j-1] > values[j] {
				values[j-1], values[j] = values[j], values[j-1]
				sorted = false
			}
		}
		if sorted {
			break
		}
	}
}
