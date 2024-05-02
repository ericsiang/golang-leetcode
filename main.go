package main

import "fmt"

func main() {
	//1.
	// nums := []int{3,2,4}
	// fmt.Println(twoSum(nums,6))

	//88.
	// nums1 := []int{1, 2, 3, 0, 0, 0}
	// m := 3
	// nums2 := []int{2, 5, 6}
	// n := 3
	// merge(nums1, m, nums2, n)
	// fmt.Println(nums1)

	//27.
	// nums27 := []int{3,2,2,3}
	// nums27_1 := []int{0,1,2,2,3,0,4,2}

	// removeElement(nums27,2)
	// fmt.Println(nums27)

	//26.
	nums26 := []int{1, 1, 2}
	fmt.Println("nums26 : ", nums26)
	k := removeDuplicates(nums26)
	fmt.Println("k : ", k)

	//217.
	// nums := []int{1, 2, 3, 1}
	// fmt.Println(containsDuplicate(nums))

}
