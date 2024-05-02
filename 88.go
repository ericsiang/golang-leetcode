package main

// Example 1:

// Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// Output: [1,2,2,3,5,6]
// Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
// The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.

// Example 2:
// Input: nums1 = [1], m = 1, nums2 = [], n = 0
// Output: [1]
// Explanation: The arrays we are merging are [1] and [].
// The result of the merge is [1].

// 說明： https://anj910.medium.com/leetcode-88-merge-sorted-array-%E4%B8%AD%E6%96%87-c0fac1a22343

// 合并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1

	for k := len(nums1) - 1; k >= 0; k-- {
		if j < 0 || i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
}
