package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	multipRes := [][]int{}
	//排序是為了方便操作跳過重複值
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	// -4 -1 -1 0 1 2
	// i   j         k
	// i+j+k = 0
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := 0 - nums[i] // l + r 值
		l, r := i+1, len(nums)-1
		for l < r {
			res := []int{}
			towSum := nums[l] + nums[r]
			if target == towSum {
				res = append(res, nums[i])
				res = append(res, nums[l])
				res = append(res, nums[r])
				multipRes = append(multipRes, res)
				// l 的值已重複跳過
				for (l+1) < len(nums) && nums[l] == nums[l+1] {
					l++
				}
				// r 的值已重複跳過
				for (r-1) >= 0 && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if target < towSum {
				r--
			} else {
				l++
			}
		}
	}
	return multipRes
}
