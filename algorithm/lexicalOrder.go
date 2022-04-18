package main

func main() {

}

func lexicalOrder(n int) []int {
	var res []int

	num := 1
	for i := 0; i < n; i++ {
		res = append(res, num)
		if num*10 <= n {
			num *= 10
		} else {
			for num%10 == 9 || num+1 > n {
				num /= 10
			}
			num++
		}
	}
	return res
}
