package main


// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]

// Input: nums = [-1,1,0,-3,3]
// Output: [0,0,9,0,0]
func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	left := make([]int, len(nums))
	right := make([]int, len(nums))

	left[0], right[len(nums)-1] = 1, 1
	// fmt.Println(left,right)
	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	// fmt.Println(left)

	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}
	// fmt.Println(right)

	for i := 0; i < len(nums); i++ {
		answer[i] = left[i] * right[i]
	}
	// fmt.Println(answer)
	return answer
}
