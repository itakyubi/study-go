package main

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	res = append(res, root.Val)

	for _, node := range root.Children {
		res = append(res, preorder(node)...)
	}

	return res
}

func preorder2(root *Node) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	stack := []*Node{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}
	return res
}
