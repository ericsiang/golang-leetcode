package main

func maxArea(height []int) int {
	maxWater := 0
	l, r := 0, len(height)-1
	for l < r {
		water := 0
		if height[l] < height[r] {
			water = height[l] * (r - l)
			l++
		} else {
			water = height[r] * (r - l)
			r--
		}
		if water > maxWater {
			maxWater = water
		}
	}
	return maxWater
}
