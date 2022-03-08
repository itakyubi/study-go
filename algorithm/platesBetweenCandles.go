package main

/**
https://leetcode-cn.com/problems/plates-between-candles/
*/

func main() {

}

func platesBetweenCandles(s string, queries [][]int) []int {
	res := make([]int, len(queries))

	platesCount := make([]int, len(s)) // platesCount[i]代表下标为i左边的盘子数量（包含i位置）
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			count++
		}
		platesCount[i] = count
	}

	left := make([]int, len(s))  // left[i]代表位置i左边第一根蜡烛的下标，如果没有则为-1
	right := make([]int, len(s)) // right[i]代表位置i右边第一根蜡烛的下标，如果没有则为-1
	l := -1
	for i := 0; i < len(s); i++ {
		if s[i] == '|' {
			l = i
		}
		left[i] = l
	}
	r := -1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '|' {
			r = i
		}
		right[i] = r
	}

	for i := 0; i < len(queries); i++ {
		ll := queries[i][0]
		rr := queries[i][1]
		if right[ll] != -1 && left[rr] != -1 && right[ll] < left[rr] {
			res[i] = platesCount[left[rr]] - platesCount[right[ll]]
		}
	}

	return res
}
