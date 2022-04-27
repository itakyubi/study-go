package main

func main() {

}

var h [][]int
var rows, cols int

func pacificAtlantic(heights [][]int) [][]int {
	rows = len(heights)
	cols = len(heights[0])
	h = heights
	visitedA := make([][]bool, rows)
	visitedP := make([][]bool, rows)
	for i := range visitedA {
		visitedA[i] = make([]bool, cols)
		visitedP[i] = make([]bool, cols)
	}

	for i := 0; i < rows; i++ {
		pacificAtlanticHelper(i, 0, visitedA)
	}
	for j := 0; j < cols; j++ {
		pacificAtlanticHelper(0, j, visitedA)
	}
	for i := 0; i < rows; i++ {
		pacificAtlanticHelper(i, cols-1, visitedP)
	}
	for j := 0; j < cols; j++ {
		pacificAtlanticHelper(rows-1, j, visitedP)
	}

	var res [][]int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if visitedA[i][j] && visitedP[i][j] {
				var tmp []int
				tmp = append(tmp, i)
				tmp = append(tmp, j)
				res = append(res, tmp)
			}
		}
	}
	return res
}

func pacificAtlanticHelper(i, j int, visited [][]bool) {
	if visited[i][j] {
		return
	}
	visited[i][j] = true

	if checkIndex(i-1, j) && h[i-1][j] >= h[i][j] {
		pacificAtlanticHelper(i-1, j, visited)
	}
	if checkIndex(i+1, j) && h[i+1][j] >= h[i][j] {
		pacificAtlanticHelper(i+1, j, visited)
	}
	if checkIndex(i, j-1) && h[i][j-1] >= h[i][j] {
		pacificAtlanticHelper(i, j-1, visited)
	}
	if checkIndex(i, j+1) && h[i][j+1] >= h[i][j] {
		pacificAtlanticHelper(i, j+1, visited)
	}
}

func checkIndex(i, j int) bool {
	return i >= 0 && i < rows && j >= 0 && j < cols
}
