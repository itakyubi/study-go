package main

import "strings"

func toGoatLatin(sentence string) string {
	vowels := map[byte]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
		'A': {},
		'E': {},
		'I': {},
		'O': {},
		'U': {},
	}

	res := &strings.Builder{}

	for i, cnt, n := 0, 1, len(sentence); i < n; i++ {
		if cnt > 1 {
			res.WriteByte(' ')
		}

		start := i
		for i++; i < n && sentence[i] != ' '; i++ {

		}
		cnt++

		if _, ok := vowels[sentence[start]]; ok {
			res.WriteString(sentence[start:i])
		} else {
			res.WriteString(sentence[start+1 : i])
			res.WriteByte(sentence[start])
		}
		res.WriteByte('m')
		res.WriteString(strings.Repeat("a", cnt))
	}
	return res.String()
}
