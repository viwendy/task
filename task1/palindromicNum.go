package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 100001
	boolData := palindromicNum(x)
	fmt.Println(boolData)
}

func palindromicNum(x int) bool {
	stingX := strconv.Itoa(x)
	for i := 0; i < len(stingX)/2; i++ {
		if stingX[i] != stingX[len(stingX)-i-1] {
			return false
		}
	}
	return true
}
