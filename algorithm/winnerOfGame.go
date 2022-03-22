package main

/*
https://leetcode-cn.com/problems/remove-colored-pieces-if-both-neighbors-are-the-same-color/
*/

func main() {

}

func winnerOfGame(colors string) bool {
	countA, countB := 0, 0
	start := 0
	for start < len(colors) {
		c := colors[start]
		index := start + 1
		for index < len(colors) && colors[index] == c {
			index++
		}
		if index-start >= 3 {
			if c == 'A' {
				countA += index - start - 2
			} else {
				countB += index - start - 2
			}
		}
		start = index
	}
	return countA > countB
}
