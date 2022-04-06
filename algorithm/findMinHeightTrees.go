package main

/**
https://leetcode-cn.com/problems/minimum-height-trees/
*/

func main() {

}

func findMinHeightTrees(n int, edges [][]int) []int {
	var res []int
	if n == 1 {
		res = append(res, 0)
		return res
	}

	// 计算节点的度和邻接节点
	adj := make([][]int, n)
	degrees := make([]int, n)
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
		degrees[edge[0]]++
		degrees[edge[1]]++
	}

	// 找到所有度为1的节点
	var queue []int
	for i, degree := range degrees {
		if degree == 1 {
			queue = append(queue, i)
		}
	}

	// 删除所有度为1的节点
	for len(queue) > 0 {
		res = []int{}
		var tmp []int
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			res = append(res, node)
			for _, v := range adj[node] {
				degrees[v]--
				if degrees[v] == 1 {
					tmp = append(tmp, v)
				}
			}
		}
		queue = tmp
	}

	return res
}
