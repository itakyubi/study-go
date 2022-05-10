package main

import "study-go/utils"

func findLHS(nums []int) int {
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}

	max := 0
	for k, v := range numMap {
		if _, ok := numMap[k+1]; ok {
			max = utils.Max(max, numMap[k+1]+v)
		}
	}
	return max
}
