package main

import (
	"sort"
)

/*
https://leetcode-cn.com/problems/longest-word-in-dictionary/
*/

func main() {

}

func longestWord(words []string) string {
	sort.Slice(words, func(i, j int) bool {
		s, t := words[i], words[j]
		return len(s) < len(t) || len(s) == len(t) && s > t
	})

	res := ""
	set := map[string]struct{}{"": {}}

	for _, word := range words {
		if _, ok := set[word[:len(word)-1]]; ok {
			res = word
			set[word] = struct{}{}
		}
	}

	return res
}
