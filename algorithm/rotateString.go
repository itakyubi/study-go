package main

/**
https://leetcode-cn.com/problems/rotate-string/
*/

func main() {

}

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	for i := 0; i < len(s); i++ {
		if s[i] == goal[0] && (s[i:]+s[:i]) == goal {
			return true
		}
	}
	return false
}
