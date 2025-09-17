package main

import (
	"fmt"
)

func main() {
	interArr := [][]int{{1, 3}, {3, 6}, {10, 8}, {9, 11}}
	intervals := mergeArr(interArr)
	fmt.Println(intervals)
}

func mergeArr(interArr [][]int) [][]int {
	if len(interArr) <= 1 {
		return interArr
	}
	// 先按起始位置排序
	for i := 0; i < len(interArr)-1; i++ {
		for j := i + 1; j < len(interArr); j++ {
			if interArr[i][0] > interArr[j][0] {
				interArr[i], interArr[j] = interArr[j], interArr[i]
			}
		}
	}
	//fmt.Println(interArr)
	//os.Exit(0)
	// 合并重叠区间
	result := [][]int{interArr[0]}
	for i := 1; i < len(interArr); i++ {
		last := result[len(result)-1]
		current := interArr[i]
		// 如果当前区间与结果中最后一个区间重叠，则合并
		if last[1] >= current[0] {
			last[1] = max(last[1], current[1])
		} else {
			// 否则将当前区间添加到结果中
			result = append(result, current)
		}
	}
	return result
}
