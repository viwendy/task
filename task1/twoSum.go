package main

import "fmt"

func main() {
	intArr := []int{2, 7, 11, 15}
	target := 9
	arr := twoSum(intArr, target)
	fmt.Println(arr)
}

func twoSum(intArr []int, target int) []int {
	for k, v := range intArr {
		for i := k + 1; i < len(intArr); i++ {
			if v+intArr[i] == target {
				return []int{k, i}
			}
		}
	}
	return nil
}
