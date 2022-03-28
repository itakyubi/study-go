package main

/*
https://leetcode-cn.com/problems/binary-number-with-alternating-bits/
*/

func main() {

}

func hasAlternatingBits(n int) bool {
	prev := n & 1
	n = n >> 1
	for n > 0 {
		cur := n & 1
		if cur == prev {
			return false
		}
		prev = cur
		n = n >> 1
	}
	return true
}
