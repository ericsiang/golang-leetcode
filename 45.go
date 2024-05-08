package main

//Input: nums = [2,3,1,1,4]
// Output: 2

// Input: nums = [2,3,0,1,4]
// Output: 2
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	var current_cover, next_max_cover, count int = 0, 0, 0
	for i, v := range nums {
		if i+v > next_max_cover {
			next_max_cover = i + v
		}

		if i == current_cover {
			count++
			current_cover = next_max_cover
			if current_cover >= len(nums)-1 {
				break
			}
		}
	}
	return count
}
