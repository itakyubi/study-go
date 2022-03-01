package main

/**
https://leetcode-cn.com/problems/2bCMpM/
*/
func main() {
	mat := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	mat2 := [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}

	println(updateMatrix(mat))
	println(updateMatrix(mat2))
}

func updateMatrix(mat [][]int) [][]int {
	dp := make([][]int, len(mat))
	for i := 0; i < len(mat); i++ {
		dp[i] = make([]int, len(mat[0]))
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				dp[i][j] = 0
			} else {
				dp[i][j] = len(mat) + len(mat[0])
			}
		}
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] != 0 {
				if i > 0 {
					dp[i][j] = min(dp[i][j], dp[i-1][j]+1)
				}
				if j > 0 {
					dp[i][j] = min(dp[i][j], dp[i][j-1]+1)
				}
			}
		}
	}

	for i := len(mat) - 1; i >= 0; i-- {
		for j := len(mat[0]) - 1; j >= 0; j-- {
			if mat[i][j] != 0 {
				if i < len(mat)-1 {
					dp[i][j] = min(dp[i][j], dp[i+1][j]+1)
				}
				if j < len(mat[0])-1 {
					dp[i][j] = min(dp[i][j], dp[i][j+1]+1)
				}
			}
		}
	}

	return dp
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
