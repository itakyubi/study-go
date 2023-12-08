package main

func main() {

}

func maxTaxiEarnings(n int, rides [][]int) int64 {
	rideDp := make([]int64, n+1)
	rideMap := make(map[int][][]int)

	for _, ride := range rides {
		rideMap[ride[1]] = append(rideMap[ride[1]], ride)
	}

	for i := 1; i <= n; i++ {
		rideDp[i] = rideDp[i-1]
		for _, ride := range rideMap[i] {
			rideDp[i] = maxTaxiEarningsMax(rideDp[i], rideDp[ride[0]]+int64(ride[1]-ride[0]+ride[2]))
		}
	}
	return rideDp[n]
}

func maxTaxiEarningsMax(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
