package main

func secondGreaterElement(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1
	}
	var list1, list2 []int
	for i := 0; i < n; i++ {
		num := nums[i]
		for len(list1) > 0 && nums[list1[len(list1)-1]] < num {
			res[list1[len(list1)-1]] = num
			list1 = list1[:len(list1)-1]
		}
		pos := len(list2) - 1
		for pos >= 0 && nums[list2[pos]] < num {
			pos--
		}
		list1 = append(list1, list2[pos+1:]...)
		list2 = append(list2[:pos+1], i)
	}
	return res
}
