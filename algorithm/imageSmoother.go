package main

/**
https://leetcode-cn.com/problems/image-smoother/
*/

func main() {

}

func imageSmoother(img [][]int) [][]int {
	res := make([][]int, len(img))

	for i := 0; i < len(img); i++ {
		res[i] = make([]int, len(img[0]))
		for j := 0; j < len(img[0]); j++ {
			sum, count := 0, 0

			for x := i - 1; x <= i+1; x++ {
				for y := j - 1; y <= j+1; y++ {
					if x >= 0 && x < len(img) && y >= 0 && y < len(img[0]) {
						sum += img[x][y]
						count++
					}
				}
			}

			res[i][j] = sum / count
		}
	}
	return res
}
