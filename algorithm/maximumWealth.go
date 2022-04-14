package main

import "study-go/utils"

func main() {

}

func maximumWealth(accounts [][]int) int {
	max := 0
	for _, account := range accounts {
		sum := 0
		for _, a := range account {
			sum += a
		}
		max = utils.Max(sum, max)
	}
	return max
}
