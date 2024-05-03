package main

// Input: nums = [1,1,1,2,2,3]
// Output: 5, nums = [1,1,2,2,3,_]

// Input: nums = [0,0,1,1,1,1,2,3,3]
// Output: 7, nums = [0,0,1,1,2,3,3,_,_]
func removeDuplicates2(nums []int) int {

	p := 0
	dupCounterMap := make(map[int]int)
    dupCounterMap[nums[p]] = 1
	for i := 1; i < len(nums); i++ {
        dupCounterMap[nums[i]] += 1
        if dupCounterMap[nums[i]] <= 2 {
            p++
            nums[p] = nums[i]
        }
    }

    return p + 1
    // for j := 0; j < len(nums); j++ {
    //     if i < 2 || nums[j] != nums[i-2] {
    //         nums[i] = nums[j]
    //         i++
    //     }
    // }
    // return i


	
	// p := 1
	// c := 1
	// for i := 1; i < len(nums); i++ {
	// 	if nums[i] == nums[i-1] {
	// 		c++
	// 	} else {
	// 		c = 1
	// 	}

	// 	if c <= 2 {
	// 		nums[p] = nums[i]
	// 		fmt.Println(nums)
	// 		p++
	// 	}
	// }

	// return p
}
