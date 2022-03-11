package main

/*
https://leetcode-cn.com/problems/count-nodes-with-the-highest-score/
*/

func main() {

	parents := []int{-1, 2, 0, 2, 0}
	parents2 := []int{-1, 2, 0}
	parents3 := []int{-1, 3, 10, 12, 0, 9, 4, 6, 9, 4, 12, 10, 0, 3}
	parents4 := []int{-1, 8, 9, 7, 6, 2, 9, 8, 0, 0}

	println(countHighestScoreNodes(parents))
	println(countHighestScoreNodes(parents2))
	println(countHighestScoreNodes(parents3))
	println(countHighestScoreNodes(parents4))

}

type node struct {
	children []*node
}

var total, count int
var max int64

func countHighestScoreNodes(parents []int) int {
	max, total, count = 0, len(parents), 0

	nodes := make([]*node, total)
	for i := 0; i < total; i++ {
		nodes[i] = &node{children: []*node{}}
	}
	for i := 1; i < total; i++ {
		nodes[parents[i]].children = append(nodes[parents[i]].children, nodes[i])
	}

	countHighestScoreNodesHelper(nodes[0])

	return count
}

func countHighestScoreNodesHelper(root *node) int64 {
	if root == nil {
		return 0
	}

	var score, childrenCount, parentCount int64
	score = 1
	childrenCount = 0
	parentCount = int64(total - 1)
	for _, n := range root.children {
		nodeCount := countHighestScoreNodesHelper(n)
		if nodeCount > 0 {
			score *= nodeCount
			childrenCount += nodeCount
			parentCount -= nodeCount
		}
	}

	if parentCount > 0 {
		score *= parentCount
	}

	if score > max {
		max = score
		count = 1
	} else if score == max {
		count++
	}

	return childrenCount + 1
}
