package main

import (
	"strconv"
	"study-go/utils"
)

/**
https://leetcode-cn.com/problems/base-7/
*/

func main() {

}

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	var res string
	sign := num < 0
	num = utils.Abs(num)

	for num > 0 {
		res = strconv.Itoa(num%7) + res
		num = num / 7
	}

	if sign {
		res = "-" + res
	}
	return res
}
