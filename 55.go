package main

// Input: nums = [2,3,1,1,4]
// Output: true
// Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
// Input: nums = [3,2,1,0,4]
// Output: false
// Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.

func canJump(nums []int) bool {
	// nums = [0]
    if nums[0]==0 && len(nums) == 1 {
        return true
    }
    max_len := 0
	for i, v := range nums {
		if max_len >= i && i+v > max_len{
			max_len =i+v	
		}	
	}	
	return max_len >= len(nums)-1
}
