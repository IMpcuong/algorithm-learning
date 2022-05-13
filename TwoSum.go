package main

func twoSumBruteForce(nums []int, target int) []int {
	arr := make([]int, 2)
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				arr := []int{i, j}
				return arr
			}
		}
	}
	return arr
}

func twoSumRecursive(nums []int, target int) []int {
	arr := make([]int, 2)
	return arr
}
