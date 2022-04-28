package main

func sortArrayByParity(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		for i < j && nums[i]&1 == 0 {
			i++
		}
		for i < j && nums[j]&1 == 1 {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}
