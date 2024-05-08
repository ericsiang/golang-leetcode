package main

import (
	"sort"
)

// Input: citations = [3,0,6,1,5]
// Output: 3

func hIndex(citations []int) int {
	//   4 3 2 1  h
	// 0,1,3,5,6
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] < citations[j]
	})

	h := 1
	for i := len(citations) - 1; i >= 0; i-- {
		if citations[i] >= h {
			h++
		} else {
			break
		}
	}
	return h - 1
}
