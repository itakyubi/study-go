package main

import "study-go/utils"

func main() {
	nums := []int{1, 5, 2}
	nums2 := []int{1, 5, 233, 7}
	println(PredictTheWinner2(nums))
	println(PredictTheWinner2(nums2))
}

func PredictTheWinner(nums []int) bool {
	return helper(nums, 0, len(nums)-1, 1) >= 0
}

func helper(nums []int, start, end, turn int) int {
	if start == end {
		return nums[start] * turn
	}

	scoreStart := nums[start]*turn + helper(nums, start+1, end, -turn)
	scoreEnd := nums[end]*turn + helper(nums, start, end-1, -turn)

	if turn == 1 {
		return utils.Max(scoreStart, scoreEnd)
	} else {
		return utils.Min(scoreStart, scoreEnd)
	}
}

// dp[i][j] 表示数组下标为i到j时，当前玩家与另一位玩家的分数最大差值
// dp[i][i] = nums[i]
// dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
func PredictTheWinner2(nums []int) bool {
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, len(nums))
		dp[i][i] = nums[i]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			dp[i][j] = utils.Max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}

	return dp[0][len(nums)-1] >= 0
}
