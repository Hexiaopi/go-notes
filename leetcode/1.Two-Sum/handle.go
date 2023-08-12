package Two_Sum

func handleOne(nums []int, target int) []int {
	var numIndex = make(map[int]int)
	for k, v := range nums {
		numIndex[v] = k
	}
	for i := 0; i < len(nums); i++ {
		want := target - nums[i]
		if j, ok := numIndex[want]; ok {
			if j == i {
				continue
			}
			return []int{i, j}
		}
	}
	return nil
}

func handleTwo(nums []int, target int) []int {
	var numIndex = make(map[int]int)
	for j, v := range nums {
		want := target - v
		if i, ok := numIndex[want]; ok {
			return []int{i, j}
		}
		numIndex[v] = j
	}
	return nil
}
