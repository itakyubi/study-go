package main

import "sort"

func main() {

}

func matrixSum(nums [][]int) int {
	for i := 0; i < len(nums); i++ {
		sort.Ints(nums[i])
	}

	sum := 0
	for j := 0; j < len(nums[0]); j++ {
		max := 0
		for i := 0; i < len(nums); i++ {
			if nums[i][j] > max {
				max = nums[i][j]
			}
		}
		sum += max
	}
	return sum
}
