package main

import "study-go/utils"

func main() {
	s1 := "cdbcbbaaabab"
	x1, y1 := 4, 5

	s2 := "aabbaaxybbaabb"
	x2, y2 := 5, 4

	s3 := "aabbabkbbbfvybssbtaobaaaabataaadabbbmakgabbaoapbbbbobaabvqhbbzbbkapabaavbbeghacabamdpaaqbqabbjbababmbakbaabajabasaabbwabrbbaabbafubayaazbbbaababbaaha"
	x3, y3 := 1926, 4320

	s4 := "aabbab"
	x4, y4 := 4, 5

	println(maximumGain2(s1, x1, y1))
	println(maximumGain2(s2, x2, y2))
	println(maximumGain2(s3, x3, y3))
	println(maximumGain2(s4, x4, y4))
}

func maximumGain(s string, x int, y int) int {
	res := 0

	var stack []uint8

	max := utils.Max(x, y)
	min := utils.Min(x, y)

	var priority uint8
	if x > y {
		priority = 'b'
	} else {
		priority = 'a'
	}

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == 'a' || c == 'b' {
			if len(stack) != 0 && c == priority && stack[len(stack)-1] != priority {
				stack = stack[:len(stack)-1]
				res += max
			} else {
				stack = append(stack, c)
			}
		} else {
			res += calRes(stack, min)
			stack = stack[:0]
		}
	}

	res += calRes(stack, min)
	return res
}

func calRes(stack []uint8, score int) int {
	res := 0
	var tmpStack []uint8
	for len(stack) > 1 {
		tmpStack = append(tmpStack, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
		for len(stack) > 0 && len(tmpStack) > 0 && stack[len(stack)-1] != tmpStack[len(tmpStack)-1] {
			stack = stack[:len(stack)-1]
			tmpStack = tmpStack[:len(tmpStack)-1]
			res += score
		}
	}
	return res
}

func maximumGain2(s string, x int, y int) int {
	res := 0
	cFirst, cSecond := 0, 0

	var max, min int
	var first, second uint8
	if x < y {
		max = y
		min = x
		first = 'a'
		second = 'b'
	} else {
		max = x
		min = y
		first = 'b'
		second = 'a'
	}

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == second {
			cSecond++
		} else if c == first {
			if cSecond > 0 {
				cSecond--
				res += max
			} else {
				cFirst++
			}
		} else {
			res += min * utils.Min(cFirst, cSecond)
			cFirst = 0
			cSecond = 0
		}
	}

	return res + min*utils.Min(cFirst, cSecond)
}
