package main

func main() {

}

func findRestaurant(list1 []string, list2 []string) []string {
	listMap := make(map[string]int)
	for i, s := range list1 {
		listMap[s] = i
	}

	var res []string
	indexSum := -1
	for i := 0; i < len(list2); i++ {
		if _, ok := listMap[list2[i]]; ok {
			tmpSum := listMap[list2[i]] + i
			if indexSum == -1 || tmpSum == indexSum {
				res = append(res, list2[i])
				indexSum = tmpSum
			} else if tmpSum < indexSum {
				indexSum = tmpSum
				res = []string{}
				res = append(res, list2[i])
			}
		}
	}

	return res
}
