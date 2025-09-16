package main

import (
	"fmt"
)

func main() {
	intArr := []int{2, 3, 1}
	num := appearOnlyOnceNum(intArr)
	fmt.Println(num)
}

func appearOnlyOnceNum(nums []int) int {
	var intMap = make(map[int]int)
	for _, num := range nums {
		intMap[num] += 1
	}
	fmt.Println(intMap)
	for k, v := range intMap {
		if v == 1 {
			return k
		}
	}
	return -1
}
