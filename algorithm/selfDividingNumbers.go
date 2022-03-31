package main

/**
https://leetcode-cn.com/problems/self-dividing-numbers/
*/

func main() {

}

func selfDividingNumbers(left int, right int) []int {
	var res []int
	for i := left; i <= right; i++ {
		if isSelfDividingNumber(i) {
			res = append(res, i)
		}
	}
	return res
}

func isSelfDividingNumber(num int) bool {
	tmp := num
	for tmp > 0 {
		digit := tmp % 10
		if digit == 0 || num%digit != 0 {
			return false
		}
		tmp /= 10
	}
	return true
}
