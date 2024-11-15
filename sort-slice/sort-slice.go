package main

import (
	"fmt"
	"sort"
)

func main() {
	arr_slice := make([]int, 0)
	arr_slice = append(arr_slice, 1, 3, -2, -4, -6, 5, -7, -9, -8)
	fmt.Println(arr_slice)
	neg_slice := make([]int, 0)
	pos_slice := make([]int, 0)
	for x := range arr_slice {
		if arr_slice[x] < 0 {
			neg_slice = append(neg_slice, arr_slice[x])
		} else {
			pos_slice = append(pos_slice, arr_slice[x])
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(neg_slice)))
	fmt.Println("neg_slice:", neg_slice)
	fmt.Println("pos_slice:", pos_slice)

	var result []int
	len1, len2 := len(pos_slice), len(neg_slice)
	maxLen := len1
	if len2 > len1 {
		maxLen = len2
	}

	for i := 0; i < maxLen; i++ {
		if i < len1 {
			result = append(result, pos_slice[i])
		}
		if i < len2 {
			result = append(result, neg_slice[i])
		}
	}

	fmt.Println("Resulting slice:", result)

}
