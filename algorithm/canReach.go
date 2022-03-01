package main

/**
https://leetcode-cn.com/problems/jump-game-iii/
*/
func main() {

}

func canReach(arr []int, start int) bool {
	visited := make([]bool, len(arr))
	return canReachHelper(arr, start, visited)
}

func canReachHelper(arr []int, start int, visited []bool) bool {
	if start < 0 || start >= len(arr) || visited[start] {
		return false
	}

	visited[start] = true
	if arr[start] == 0 {
		return true
	}

	return canReachHelper(arr, start+arr[start], visited) || canReachHelper(arr, start-arr[start], visited)
}
