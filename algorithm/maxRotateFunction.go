package main

import "study-go/utils"

func maxRotateFunction(nums []int) int {
	n, sum := len(nums), 0
	f := make([]int, n)
	for i := 0; i < n; i++ {
		sum += nums[i]
		f[0] += i * nums[i]
	}

	max := f[0]
	for i := 1; i < n; i++ {
		f[i] = f[i-1] + sum - n*nums[n-i]
		max = utils.Max(max, f[i])
	}
	return max
}
