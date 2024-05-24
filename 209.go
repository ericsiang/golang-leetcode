package main

func minSubArrayLen(target int, nums []int) int {
	l, r, sublen, sum, res := 0, 0, 0, 0, len(nums)+1
	for r < len(nums) {
		sum += nums[r]
		for sum >= target {
			sublen = r - l + 1
			if sublen < res {
				res = sublen
			}
			sum -= nums[l]
			l++
		}
		r++
	}

	if res == len(nums)+1 {
		return 0
	}
	return res
}
