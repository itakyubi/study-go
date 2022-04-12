package main

func main() {

}

func numberOfLines(widths []int, s string) []int {
	sum, level := 0, 0
	for i := 0; i < len(s); i++ {
		width := widths[s[i]-'a']
		if sum+width <= 100 {
			sum += width
		} else {
			level++
			sum = width
		}
	}
	return []int{level + 1, sum}
}
