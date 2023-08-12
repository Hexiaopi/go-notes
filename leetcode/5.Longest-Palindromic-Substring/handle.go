package Longest_Palindromic_Substring

import (
	"strings"
)

//暴力解法，时间复杂度0(n^3)，空间复杂度0(1)
func handleOne(s string) string {
	begin := 0
	max := 1
	slice := strings.Split(s, "")
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			if j-i+1 > max && validateSubstring(s, i, j) {
				max = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+max]
}

func validateSubstring(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

//中心扩散法，时间复杂度0(n^3)，空间复杂度0(1)
func handleTwo(s string) string {
	slice := strings.Split(s, "")
	if len(slice) < 2 {
		return s
	}
	max := 1
	begin := 0
	for i := 0; i < len(slice)-1; i++ {
		//奇数
		oddLen := expandAroundCenter(slice, i, i)
		//偶数
		evenLen := expandAroundCenter(slice, i, i+1)
		if oddLen > max {
			max = oddLen
			begin = i - (oddLen-1)/2
		}
		if evenLen > max {
			max = evenLen
			begin = i - (evenLen)/2 + 1
		}
	}
	return s[begin : begin+max]
}

func expandAroundCenter(slice []string, i, j int) int {
	for i >= 0 && j < len(slice) {
		if slice[i] == slice[j] {
			i--
			j++
		} else {
			break
		}
	}
	return j - i - 1
}

//动态规划法
func handleThree(s string) string {
	res, dp := "", make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if (j - i) < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && (res == "" || j-i+1 > len(res)) {
				res = s[i : j+1]
			}
		}
	}
	return res
}
