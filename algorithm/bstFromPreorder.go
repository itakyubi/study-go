package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	preorder := []int{8, 5, 1, 7, 10, 12}
	root := bstFromPreorder(preorder)
	println(root)
}

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := TreeNode{Val: preorder[0]}
	index := 1
	for index < len(preorder) {
		if preorder[0] < preorder[index] {
			break
		}
		index++
	}

	root.Left = bstFromPreorder(preorder[1:index])
	root.Right = bstFromPreorder(preorder[index:])
	return &root
}
