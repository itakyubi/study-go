package main

import (
	"math"
	"strconv"
	"study-go/utils"
)

func main() {

}

func nearestPalindromic(n string) string {
	num, _ := strconv.Atoi(n)
	candidates := []int{int(math.Pow10(len(n)-1)) - 1, int(math.Pow10(len(n))) + 1}

	prefix, _ := strconv.Atoi(n[:(len(n)+1)/2])
	for i := prefix - 1; i <= prefix+1; i++ {
		pre := strconv.Itoa(i)
		suffix := reverseString(strconv.Itoa(i))

		candidate, _ := strconv.Atoi(pre + suffix[len(n)&1:])
		candidates = append(candidates, candidate)
	}

	res := -1
	for _, candidate := range candidates {
		if candidate == num {
			continue
		}

		if res == -1 {
			res = candidate
			continue
		}

		candidateDiff := utils.Abs(candidate - num)
		resDiff := utils.Abs(res - num)
		if candidateDiff < resDiff || (candidateDiff == resDiff && candidate < res) {
			res = candidate

		}
	}

	return strconv.Itoa(res)
}

func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
