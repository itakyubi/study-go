package main

/**
https://leetcode-cn.com/problems/factorial-trailing-zeroes/
*/
func main() {

}

func trailingZeroes(n int) int {
	count := 0
	for i := 5; i <= n; i += 5 {
		tmp := i
		for tmp%5 == 0 {
			count++
			tmp /= 5
		}
	}
	return count
}
