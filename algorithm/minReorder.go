package main

func main() {

}

func minReorder(n int, connections [][]int) int {
	edges := make([][][]int, n)
	for _, c := range connections {
		edges[c[0]] = append(edges[c[0]], []int{c[1], 1})
		edges[c[1]] = append(edges[c[1]], []int{c[0], 0})
	}
	return minReorderDfs(0, -1, edges)
}

func minReorderDfs(cur, parent int, edges [][][]int) int {
	res := 0
	for _, edge := range edges[cur] {
		if edge[0] == parent {
			continue
		}
		res += edge[1] + minReorderDfs(edge[0], cur, edges)
	}
	return res
}
