package main

func main() {
	nums := []int{4, 5, 6, 7, 1, 2, 3, 8}
	println(findKthLargest(nums, 3))
}

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, l, r, target int) int {
	index := partition(nums, l, r)
	if index == target {
		return nums[index]
	} else if index < target {
		return quickSelect(nums, index+1, r, target)
	} else {
		return quickSelect(nums, l, index-1, target)
	}
}

func partition(nums []int, l, r int) int {
	pivot := nums[l]
	small := l - 1
	for i := l; i <= r; i++ {
		if nums[i] <= pivot {
			small++
			nums[small], nums[i] = nums[i], nums[small]
		}
	}
	nums[small], nums[l] = nums[l], nums[small]
	return small
}
