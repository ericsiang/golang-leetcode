package main

// Input: nums = [3,2,2,3], val = 3
// Output: 2, nums = [2,2,_,_]

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	pointer := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			// if i != pointer {
			// 	nums[i] ,nums[pointer] = nums[pointer] ,nums[i]
			// }
			// pointer++
			nums[pointer] = nums[i]
			pointer++
		}
	}
	return pointer
}
