package main

/**
https://leetcode-cn.com/problems/three-consecutive-odds/
*/

func main() {

}

func threeConsecutiveOdds(arr []int) bool {
	count := 0
	for _, a := range arr {
		if a&1 == 0 {
			count = 0
		} else {
			count++
		}
		if count == 3 {
			return true
		}
	}
	return false
}
