package main



// Input: nums = [1,2,3,4,5,6,7], k = 3
// Output: [5,6,7,1,2,3,4]
func rotate(nums []int, k int) {
	orign_nums := make([]int, 0, len(nums))
	for _, v := range nums {
		orign_nums = append(orign_nums, v)
	}

	for i := 0; i < len(nums); i++ {
		newIndex := (i + k) % len(nums)
		nums[newIndex] = orign_nums[i]
	}
}
