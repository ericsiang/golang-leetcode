package main

import "fmt"

func main() {
	//1.
	// nums := []int{3,2,4}
	// fmt.Println(twoSum(nums,6))

	//2.
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
	// nums := []int{1, 2, 3, 1}
	// fmt.Println(containsDuplicate(nums))

}
