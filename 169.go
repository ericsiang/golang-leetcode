package main

// Example 1:

// Input: nums = [3,2,3]
// Output: 3
// Example 2:

// Input: nums = [2,2,1,1,1,2,2]
// Output: 2
func majorityElement(nums []int) int {
	maxIndex := 0
	countMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		countMap[nums[i]] += 1
	}
	tmp := 0
	for i, v := range countMap {
		if v > tmp {
			tmp = v
			maxIndex = i
		}
	}

	return maxIndex
}
