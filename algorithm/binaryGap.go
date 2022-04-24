package main

import "study-go/utils"

func binaryGap(n int) int {
	last, max := -1, 0
	for i := 0; n > 0; i++ {
		if n&1 == 1 {
			if last != -1 {
				max = utils.Max(max, i-last)
			}
			last = i
		}
		n >>= 1
	}
	return max
}
