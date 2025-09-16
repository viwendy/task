package main

import "fmt"

func main() {
	intArr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	count, intData := removeDuplicates(intArr)
	fmt.Println(count, ", nums =", intData)
}

func removeDuplicates(nums []int) (int, []int) {
	if len(nums) == 0 {
		return 0, nums
	}
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		// 当发现不同的元素时
		if nums[fast] != nums[slow] {
			// 将不重复元素放到slow+1位置
			slow++
			nums[slow] = nums[fast]
			//fmt.Println(nums)
			//break
		}
	}
	return slow + 1, nums[:slow+1]
}
