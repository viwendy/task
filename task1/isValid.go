package main

import "fmt"

func main() {
	stringX := "({)}[]"
	boolData := isValid(stringX)
	fmt.Println(boolData)
}

func isValid(stringX string) bool {
	var stack []rune
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, v := range stringX {
		left, exists := pairs[v]
		if exists {
			// 右括号
			if len(stack) == 0 || stack[len(stack)-1] != left {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			// 左括号
			stack = append(stack, v)
		}
	}
	// 当栈为空时，才说明所有括号都正确匹配
	return len(stack) == 0
}
