package main

// Input: ratings = [1,0,2]
// Output: 5

// Input: ratings = [1,2,2]
// Output: 4
func candy(ratings []int) int {
	child := make([]int, len(ratings))
	total := 0

	for i, _ := range child {
		child[i] = 1
	}

	// 右比左大
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			child[i] = child[i-1] + 1
		}
	}

	// 左比右大
	for i := len(child) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			if child[i] < (child[i+1] + 1) {
				child[i] = child[i+1] + 1
			}
		}
	}

	for _, v := range child {
		total += v
	}
	return total
}
