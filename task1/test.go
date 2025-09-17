package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 3}
	b := []int{6, 5, 3}
	c := mergeIntArr(a, b)
	fmt.Println(c)
}

func mergeIntArr(a []int, b []int) []int {
	sort.Ints(a)
	sort.Ints(b)
	clen := len(a) + len(b)
	c := make([]int, clen)
	for i := 0; i < len(a); i++ {
		c[i] = a[i]
	}
	for i := 0; i < len(b); i++ {
		c[i+len(a)] = b[i]
	}
	return c[:]
}
