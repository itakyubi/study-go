package main

func main() {

}

func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
	nexts := make([][]int, n)
	for _, edge := range edges {
		nexts[edge[0]] = append(nexts[edge[0]], edge[1])
		nexts[edge[1]] = append(nexts[edge[1]], edge[0])
	}

	cnt := make([]int, n)
	for _, trip := range trips {
		dfs(trip[0], -1, trip[1], nexts, cnt)
	}

	a, b := dp(0, -1, price, nexts, cnt)
	return min(a, b)
}

func dfs(node int, parent int, end int, nexts [][]int, cnt []int) bool {
	if node == end {
		cnt[node]++
		return true
	}

	for _, next := range nexts[node] {
		if next == parent {
			continue
		}

		if dfs(next, node, end, nexts, cnt) {
			cnt[node]++
			return true
		}
	}
	return false
}

func dp(node int, parent int, price []int, nexts [][]int, cnt []int) (a, b int) {
	a = price[node] * cnt[node]
	b = price[node] * cnt[node] / 2
	for _, next := range nexts[node] {
		if next == parent {
			continue
		}
		childA, childB := dp(next, node, price, nexts, cnt)
		a += min(childA, childB)
		b += childA
	}
	return a, b
}
