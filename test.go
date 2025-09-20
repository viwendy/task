package main

import (
	"fmt"
)

//反转字符串 (Reverse String)
//题目描述：反转一个字符串。输入 "abcde"，输出 "edcba"

func reverses(s string) string {
	runes := []rune(s)
	lens := len(runes)
	for i := 0; i < lens/2; i++ {
		runes[i], runes[lens-i-1] = runes[lens-i-1], runes[i]
	}
	return string(runes)
}

func main() {
	s := "123456"
	re := reverses(s)
	fmt.Println(re)
}
