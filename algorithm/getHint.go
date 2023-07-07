package main

import (
	"fmt"
)

func main() {

}

func getHint(secret string, guess string) string {
	var secretCnt [10]int
	var guessCnt [10]int
	bullCnt := 0
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bullCnt++
		} else {
			secretCnt[secret[i]-'0']++
			guessCnt[guess[i]-'0']++
		}
	}

	cowCnt := 0
	for i := 0; i < 10; i++ {
		cowCnt += min(secretCnt[i], guessCnt[i])
	}
	return fmt.Sprintf("%dA%dB", bullCnt, cowCnt)
}
