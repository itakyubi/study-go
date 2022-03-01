package main

func main() {
	nums := []int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19}
	res := subSort(nums)
	println(res[0])
	println(res[1])

}

func subSort(array []int) []int {
	res := []int{-1, -1}
	length := len(array)

	if length == 0 {
		return res
	}

	min := array[length-1]
	max := array[0]
	for i := 0; i < length; i++ {
		if array[i] >= max {
			max = array[i]
		} else {
			res[1] = i
		}

		if array[length-1-i] <= min {
			min = array[length-1-i]
		} else {
			res[0] = length - 1 - i
		}
	}

	return res
}

func subSort2(array []int) []int {
	left,right := -1,-1

	//  从前向后找第一个不满足升序的index

	for i := 1; i < len(array); i++ {
		if array[i] > array[i-1] {
			continue;
		} else {
			break;
		}
	}

	// 向前遍历找>=当前值的位置，不满足就b
	left = i
	for j := i-1; j >=0; j-- {
		if arry[j] <=array[j-1] {
			left--
		} else {
			break
		}

	}

	//  从后向前找第一个不满足降序的index
	k := len(array)-1
	for k > 0 {
		if array[k-1] <= array[k] {
			k++
		} else {
			break
		}
	}
}
