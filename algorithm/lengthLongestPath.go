package main

import "study-go/utils"

func lengthLongestPath(input string) int {
	var stack []int
	i, max := 0, 0
	for i < len(input) {
		level := 1
		for i < len(input) && input[i] == '\t' {
			i++
			level++
		}

		isFile := false
		length := 0
		for i < len(input) && input[i] != '\n' {
			if input[i] == '.' {
				isFile = true
			}
			length++
			i++
		}

		i++

		for len(stack) >= level {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			length += stack[len(stack)-1] + 1
		}
		if isFile {
			max = utils.Max(max, length)
		} else {
			stack = append(stack, length)
		}
	}
	return max
}
