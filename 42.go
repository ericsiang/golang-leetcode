package main

func trap(height []int) int {
	res := 0
	len := len(height)
	left_max := make([]int, len)
	right_max := make([]int, len)
	left_max[0] = height[0]
	right_max[len-1] = height[len-1]
	// 左到右的高度最大值
	for i := 1; i < len; i++ {
		left_max[i] = trapMax(left_max[i-1], height[i])
	}

	// 右到左的高度最大值
	for i := len - 2; i >= 0; i-- {
		right_max[i] = trapMax(right_max[i+1], height[i])
	}

	for i := 0; i < len; i++ {
		res += trapMin(left_max[i], right_max[i]) - height[i]
	}

	return res
}

func trapMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func trapMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
