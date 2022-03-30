package main

/**
https://leetcode-cn.com/problems/find-servers-that-handled-most-number-of-requests/
*/

func main() {

}

func busiestServers(k int, arrival []int, load []int) []int {
	endTime := make([]int, k)
	tackCount := make([]int, k)
	for i := 0; i < len(arrival); i++ {
		startIndex := i % k
		index := startIndex

		for index < k && endTime[index] > arrival[i] {
			index++
		}
		if index < k {
			endTime[index] = arrival[i] + load[i]
			tackCount[index]++
			continue
		}

		index = 0
		for index < startIndex && endTime[index] > arrival[i] {
			index++
		}
		if index < startIndex {
			endTime[index] = arrival[i] + load[i]
			tackCount[index]++
		}
	}

	var res []int
	max := 0
	for i := 0; i < k; i++ {
		if tackCount[i] > max {
			res = []int{}
			res = append(res, i)
			max = tackCount[i]
		} else if tackCount[i] == max {
			res = append(res, i)
		}
	}
	return res
}
