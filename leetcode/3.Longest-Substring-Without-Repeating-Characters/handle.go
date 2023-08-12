package Longest_Substring_Without_Repeating_Characters

import "strings"

func handleOne(s string) int {
	slice := strings.Split(s, "")
	result := 0
	for i := 0; i < len(slice); i++ {
		temp := make(map[string]int)
		temp[slice[i]] = 1
		count := 1
		for j := i + 1; j < len(slice); j++ {
			if _, ok := temp[slice[j]]; ok {
				break
			}
			temp[slice[j]] = 1
			count++
		}
		result = max(result, count)
	}
	return result
}

func handleTwo(s string) int {
	m := map[byte]int{}
	n := len(s)
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}
		ans = max(ans, rk-i+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
