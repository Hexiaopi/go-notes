package TwoSum

func handleOne(nums []int, target int) []int {
	var numIndex = make(map[int]int)
	for k, v := range nums {
		numIndex[v] = k
	}
	for k := 0; k < len(nums); k++ {
		i := nums[k]
		j := target - i

		if index, ok := numIndex[j]; ok {
			if index == k {
				continue
			}
			return []int{k, index}
		}
	}
	return nil
}

func handleTwo(nums []int, target int) []int {
	var numIndex = make(map[int]int)
	for k, v := range nums {
		j := target - v
		if index, ok := numIndex[j]; ok {
			return []int{index,k}
		}
		numIndex[v] = k
	}
	return nil
}
