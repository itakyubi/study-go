package main

/**
https://leetcode-cn.com/problems/diagonal-traverse/
*/

func main() {

}

func findDiagonalOrder(mat [][]int) []int {
	rows, cols := len(mat), len(mat[0])
	total := rows * cols
	res := make([]int, total)

	index, i, j := 0, 0, 0
	for res[total-1] == 0 {
		for i >= 0 && j < cols {
			res[index] = mat[i][j]
			index++
			i--
			j++
		}
		if j >= cols {
			i = i + 2
			j = j - 1
		} else {
			i = i + 1
		}

		for i < rows && j >= 0 {
			res[index] = mat[i][j]
			index++
			i++
			j--
		}
		if i >= rows {
			i = i - 1
			j = j + 2
		} else {
			j = j + 1
		}
	}
	return res
}
