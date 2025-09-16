package main

import "fmt"

func main() {
	intArr := []int{1, 1, 2}
	count, intData := removeDuplicates(intArr)
	fmt.Println(count, ", nums =", intData)
}

func removeDuplicates(nums []int) (int, []int) {
	newIntArr := []int{}
	intMap := make(map[int]int)
	for _, num := range nums {
		intMap[num] += 1
	}
	for v := range intMap {
		newIntArr = append(newIntArr, v)
	}
	return len(intMap), newIntArr
}
