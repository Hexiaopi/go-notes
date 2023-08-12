package Valid_Parentheses

import "strings"

func handleOne(s string) bool {
	sl := strings.Split(s, "")
	if len(sl) == 0 {
		return true
	}
	stack := make([]string, 0)
	for _, v := range sl {
		if v == "[" || v == "{" || v == "(" {
			stack = append(stack, v)
		} else {
			length := len(stack)
			if (v == "]" && len(stack) > 0 && stack[length-1] == "[") ||
				(v == "}" && len(stack) > 0 && stack[length-1] == "{") ||
				(v == ")" && len(stack) > 0 && stack[length-1] == "(") {
				stack = stack[:length-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
