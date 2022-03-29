package main

import (
	"study-go/utils"
)

/**
https://leetcode-cn.com/problems/maximize-the-confusion-of-an-exam/
*/

func main() {

}

func maxConsecutiveAnswers(answerKey string, k int) int {
	return utils.Max(maxConsecutiveAnswersHelper(answerKey, k, 'T'), maxConsecutiveAnswersHelper(answerKey, k, 'F'))
}

func maxConsecutiveAnswersHelper(answerKey string, k int, c byte) int {
	left, right, count, max := 0, 0, 0, 0
	for right < len(answerKey) {
		if answerKey[right] != c {
			count++
		}
		for count > k {
			if answerKey[left] != c {
				count--
			}
			left++
		}
		max = utils.Max(max, right-left+1)
		right++
	}
	return max
}
