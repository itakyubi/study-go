package main

/**
https://leetcode-cn.com/problems/add-digits/
*/

func main() {

}

func addDigits(num int) int {
	if num < 10 {
		return num
	}

	res := 0
	for num > 0 {
		res += num % 10
		num /= 10
	}
	return addDigits(res)
}

func addDigits2(num int) int {
	return (num-1)%9 + 1
}
