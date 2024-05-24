package main

func twoSum2(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if target == sum {
			return []int{l + 1, r + 1}
		}

		if sum > target {
			r--
		} else {
			l++
		}

	}

	return []int{}
}
