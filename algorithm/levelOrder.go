package main

func levelOrder(root *Node) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		size := len(queue)
		var list []int
		var tmp []*Node
		for i := 0; i < size; i++ {
			node := queue[i]
			list = append(list, node.Val)
			if len(node.Children) > 0 {
				tmp = append(tmp, node.Children...)
			}
		}
		res = append(res, list)
		queue = tmp
	}

	return res
}
