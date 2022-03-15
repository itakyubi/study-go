package main

func main() {

}

var countMaxOrSubsetsMax, countMaxOrSubsetsCount int

func countMaxOrSubsets(nums []int) int {
	countMaxOrSubsetsMax = 0
	countMaxOrSubsetsCount = 0

	for _, num := range nums {
		countMaxOrSubsetsMax |= num
	}
	countMaxOrSubsetsHelper(0, 0, nums)
	return countMaxOrSubsetsCount
}

func countMaxOrSubsetsHelper(index int, or int, nums []int) {
	if or == countMaxOrSubsetsMax {
		countMaxOrSubsetsCount += 1 << (len(nums) - index)
		return
	}

	if index >= len(nums) {
		return
	}

	countMaxOrSubsetsHelper(index+1, or, nums)
	countMaxOrSubsetsHelper(index+1, or|nums[index], nums)
}
