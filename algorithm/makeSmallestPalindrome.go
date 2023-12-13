package main

func main() {
	makeSmallestPalindrome("abcd")
}

func makeSmallestPalindrome(s string) string {
	l, r := 0, len(s)-1
	res := []byte(s)
	for l < r {
		if s[l] != s[r] {
			if s[l] <= s[r] {
				res[l] = s[l]
				res[r] = s[l]
			} else {
				res[l] = s[r]
				res[r] = s[r]
			}
		}
		l++
		r--
	}
	return string(res)
}
