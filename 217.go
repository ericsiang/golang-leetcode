package main


func containsDuplicate(nums []int) bool {
	var check = make(map[int]bool, len(nums))
	for _, num := range nums {
		if _, ok := check[num]; ok {
			return true
		} else {
			check[num] = true
		}
	}
	return false
}
