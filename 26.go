package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	pointer := 0
	for i := 0; i < len(nums); i++ {
		if nums[pointer] != nums[i] {
			pointer++
			nums[pointer] = nums[i]
			fmt.Println(nums)
		}
	}

	return pointer+1
}
