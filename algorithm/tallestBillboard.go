package main

import "study-go/utils"

func main() {

}

func tallestBillboard(rods []int) int {
	sum := 0
	for _, rod := range rods {
		sum += rod
	}

	dp := make([][]int, len(rods)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, sum+1)
	}

	for i := 1; i <= len(rods); i++ {
		for j := 0; j <= sum; j++ {
			if dp[i-1][j] < j {
				continue
			}

			dp[i][j] = utils.Max(dp[i][j], dp[i-1][j])

			k := j + rods[i-1]
			dp[i][k] = utils.Max(dp[i][k], dp[i-1][j]+rods[i-1])

			k = utils.Abs(j - rods[i-1])
			dp[i][k] = utils.Max(dp[i][k], dp[i-1][j]+rods[i-1])
		}
	}

	return dp[len(rods)][0] / 2
}
