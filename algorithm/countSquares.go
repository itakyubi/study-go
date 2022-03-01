package main

import "study-go/utils"

func main() {
	nums := [][]int{{0, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 1, 1}}
	nums2 := [][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}}

	println(countSquares(nums))
	println(countSquares(nums2))

}

func countSquares(matrix [][]int) int {
	count := 0
	rows := len(matrix)
	cols := len(matrix[0])
	dp := make([][]int, rows)

	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = matrix[i][j]
			} else if matrix[i][j] == 0 {
				dp[i][j] = 0
			} else {
				dp[i][j] = utils.Min(dp[i-1][j-1], utils.Min(dp[i-1][j], dp[i][j-1])) + 1
			}
			count += dp[i][j]
		}
	}
	return count
}
