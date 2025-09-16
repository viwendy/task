package main

import "fmt"

func main() {
	stringArr := []string{"flower", "flow", "flight"}
	stringData := longestCommonPrefix(stringArr)
	fmt.Println(stringData)
}

func longestCommonPrefix(stringArr []string) string {
	if len(stringArr) == 0 {
		return ""
	}
	prefix := stringArr[0]
	for i := 1; i < len(stringArr); i++ {
		for j := 0; j < len(prefix); j++ {
			if j == len(stringArr[i]) || stringArr[i][j] != prefix[j] {
				prefix = prefix[:j]
				break
			}
		}
	}
	return prefix
}
