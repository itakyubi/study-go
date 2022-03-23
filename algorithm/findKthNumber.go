package main

/**
https://leetcode-cn.com/problems/k-th-smallest-in-lexicographical-order/
*/

func main() {

}

func findKthNumber(n int, k int) int {
	cur := 1
	k--
	for k > 0 {
		count := getNodeCount(cur, n)
		if count <= k {
			k -= count
			cur++
		} else {
			cur *= 10
			k--
		}
	}
	return cur
}

func getNodeCount(prefix, n int) int {
	cur := prefix
	next := prefix + 1
	count := 0
	for cur <= n {
		count += min(next, n+1) - cur
		cur *= 10
		next *= 10
	}
	return count
}
