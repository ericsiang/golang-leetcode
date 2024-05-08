package main

// Input: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
// Output: 3

// Input: gas = [2,3,4], cost = [3,4,3]
// Output: -1
func canCompleteCircuit(gas []int, cost []int) int {
	curSum := 0
	totSum := 0
	start := 0
	for i, _ := range gas {
		curSum += gas[i] - cost[i]
		totSum += gas[i] - cost[i]
		if curSum < 0 {
			start = i + 1
			curSum = 0
		}
	}

	if totSum < 0{
		return -1
	}

	return start
}
