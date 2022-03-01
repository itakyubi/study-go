package main

func main() {
	nums := []int{4, 3, 2, 1, 0}
	nums2 := []int{1, 0, 2, 3, 4}
	println(maxChunksToSorted(nums))
	println(maxChunksToSorted(nums2))
}

func maxChunksToSorted(arr []int) int {
	res := 0
	endIndex := 0
	for i := 0; i < len(arr); i++ {
		if endIndex < arr[i] {
			endIndex = arr[i]
		}
		if i == endIndex {
			res++
		}
	}
	return res
}
