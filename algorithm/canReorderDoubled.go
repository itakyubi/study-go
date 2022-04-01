package main

import (
	"sort"
	"study-go/utils"
)

/**
https://leetcode-cn.com/problems/array-of-doubled-pairs/
*/

func main() {
	println(canReorderDoubled([]int{4, -2, 2, -4}))
}

func canReorderDoubled(arr []int) bool {
	countMap := make(map[int]int)
	for _, a := range arr {
		countMap[a]++
	}

	list := make([]int, 0, len(countMap))
	for k := range countMap {
		list = append(list, k)
	}
	sort.Slice(list, func(i, j int) bool {
		return utils.Abs(list[i]) < utils.Abs(list[j])
	})

	for _, num := range list {
		double := 2 * num
		if countMap[double] < countMap[num] {
			return false
		}
		countMap[double] -= countMap[num]
	}

	return true
}
