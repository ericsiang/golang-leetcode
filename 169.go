package main

// Example 1:

// Input: nums = [3,2,3]
// Output: 3
// Example 2:

// Input: nums = [2,2,1,1,1,2,2]
// Output: 2
func majorityElement(nums []int) int {
	countMap := make(map[int]int)
	for _, v := range nums {
		countMap[v] += 1
	}

	for i, c := range countMap {
		if c > len(nums)/2 {
			return i
		}
	}
	return -1
}
