package main

import "study-go/utils"

func main() {

}

func shortestToChar(s string, c byte) []int {
	res := make([]int, len(s))

	idx := -len(s)
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			idx = i
		}
		res[i] = i - idx
	}

	idx = 2 * len(s)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			idx = i
		}
		res[i] = utils.Min(res[i], idx-i)
	}
	return res
}
