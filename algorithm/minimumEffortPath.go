package main

import "sort"

func minimumEffortPath(heights [][]int) int {
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	m, n := len(heights), len(heights[0])
	left, right, res := 0, 999999, 0
	for left <= right {
		mid := (left + right) / 2
		seen := make([]bool, m*n)
		seen[0] = true
		var queue [][]int
		queue = append(queue, []int{0, 0})
		for len(queue) > 0 {
			cell := queue[0]
			queue = queue[1:]
			x, y := cell[0], cell[1]
			for i := 0; i < len(dirs); i++ {
				nx, ny := x+dirs[i][0], y+dirs[i][1]
				if nx >= 0 && nx < m && ny >= 0 && ny < n && !seen[nx*n+ny] &&
					abs(heights[x][y]-heights[nx][ny]) <= mid {
					queue = append(queue, []int{nx, ny})
					seen[nx*n+ny] = true
				}
			}
		}
		if seen[m*n-1] {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}

func minimumEffortPath2(heights [][]int) int {
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	m, n := len(heights), len(heights[0])
	return sort.Search(1e6, func(maxHeightDiff int) bool {
		seen := make([]bool, m*n)
		seen[0] = true
		var queue [][]int
		queue = append(queue, []int{0, 0})
		for len(queue) > 0 {
			cell := queue[0]
			queue = queue[1:]
			x, y := cell[0], cell[1]
			if x == m-1 && y == n-1 {
				return true
			}

			for _, dir := range dirs {
				nx, ny := x+dir[0], y+dir[1]
				if nx >= 0 && nx < m && ny >= 0 && ny < n && !seen[nx*n+ny] &&
					abs(heights[x][y]-heights[nx][ny]) <= maxHeightDiff {
					queue = append(queue, []int{nx, ny})
					seen[nx*n+ny] = true
				}
			}

		}
		return false
	})
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
