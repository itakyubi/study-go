package main

import "sort"

func main() {

}

func maxSubsequence(nums []int, k int) []int {
	var copyNums []int
	for _, num := range nums {
		copyNums = append(copyNums, num)
	}

	// 先排序
	sort.Ints(copyNums)

	// 找到k个最大的数
	cnt := make(map[int]int)
	for i := 0; i < k; i++ {
		cnt[copyNums[len(nums)-1-i]]++
	}

	var res []int
	for i := 0; i < len(nums); i++ {
		if cnt[nums[i]] > 0 {
			res = append(res, nums[i])
			cnt[nums[i]]--
		}
	}
	return res
}
