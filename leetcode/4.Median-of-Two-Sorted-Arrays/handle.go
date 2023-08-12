package Median_of_Two_Sorted_Arrays

func handleOne(nums1 []int, nums2 []int) float64 {
	temp := make([]int, 0)
	n1, n2 := 0, 0
	for n1 < len(nums1) && n2 < len(nums2) {
		if nums1[n1] <= nums2[n2] {
			temp = append(temp, nums1[n1])
			n1++
		} else if nums1[n1] > nums2[n2] {
			temp = append(temp, nums2[n2])
			n2++
		}
	}
	for n1 < len(nums1) {
		temp = append(temp, nums1[n1])
		n1++
	}
	for n2 < len(nums2) {
		temp = append(temp, nums2[n2])
		n2++
	}
	half := len(temp) / 2
	if len(temp)%2 == 1 {
		return float64(temp[half])
	} else {
		return float64(temp[half-1]+temp[half]) / 2
	}
}
