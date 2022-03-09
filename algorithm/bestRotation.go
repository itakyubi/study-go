package main

func main() {

}

func bestRotation(nums []int) int {
	n := len(nums)

	diff := make([]int, n+1)
	for i := 0; i < n; i++ {
		if i < nums[i] {
			diff[i+1]++
			diff[n-(nums[i]-i)+1]--
		} else {
			diff[0]++
			diff[i-nums[i]+1]--
			diff[i+1]++
		}
	}

	max, sum, k := 0, 0, 0
	for i := 0; i < n; i++ {
		sum += diff[i]
		if sum > max {
			max = sum
			k = i
		}
	}

	return k
}
