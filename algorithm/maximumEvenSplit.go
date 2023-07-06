package main

func main() {

}

func maximumEvenSplit(finalSum int64) []int64 {
	var res []int64
	if finalSum&1 == 1 {
		return res
	}

	i := int64(2)
	for i <= finalSum {
		res = append(res, i)
		finalSum -= i
		i += 2
	}

	if finalSum > 0 {
		res[len(res)-1] = res[len(res)-1] + finalSum
	}
	return res
}
