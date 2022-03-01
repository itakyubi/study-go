package main

/**
 *	https://leetcode-cn.com/problems/pyramid-transition-matrix/
 */

func main() {
	bottom := "BCD"
	allowed := []string{"BCG", "CDE", "GEA", "FFF"}
	bottom2 := "AABA"
	allowed2 := []string{"AAA", "AAB", "ABA", "ABB", "BAC"}
	bottom3 := "DBCDA"
	allowed3 := []string{"ACC", "ACB", "ACA", "AAC", "ACD", "BCD", "BCC", "BAB", "CCD",
		"CCA", "CCB", "DAD", "DAC", "AAB", "CAD", "ABB", "ABC", "ABD", "BDC", "BDB", "BBD", "BBC",
		"BBB", "ADD", "ADB", "DDA", "CDD", "CBC", "CDA", "CDB", "DBD", "BDA"}

	println(pyramidTransition(bottom, allowed))
	println(pyramidTransition(bottom2, allowed2))
	println(pyramidTransition(bottom3, allowed3))
}

var visited map[string]bool
var allowedMap map[string][]string

func pyramidTransition(bottom string, allowed []string) bool {
	visited = make(map[string]bool)
	allowedMap = make(map[string][]string)
	for _, allow := range allowed {
		prefix := allow[0:2]
		var endList []string
		if _, ok := allowedMap[prefix]; !ok {
			endList = []string{}
		} else {
			endList = allowedMap[prefix]
		}
		endList = append(endList, allow[2:])
		allowedMap[prefix] = endList
	}

	return pyramidTransitionHelper(bottom)
}

func pyramidTransitionHelper(bottom string) bool {
	if len(bottom) == 1 {
		return true
	}
	if _, ok := visited[bottom]; ok {
		return visited[bottom]
	}

	var nextBottoms []string
	for i := 0; i < len(bottom)-1; i++ {
		next := bottom[i : i+2]
		if _, ok2 := allowedMap[next]; !ok2 {
			return false
		} else {
			if len(nextBottoms) == 0 {
				for _, end := range allowedMap[next] {
					nextBottoms = append(nextBottoms, end)
				}
			} else {
				var tmp []string
				for _, nextBottom := range nextBottoms {
					for _, end := range allowedMap[next] {
						tmp = append(tmp, nextBottom+end)
					}
				}
				nextBottoms = tmp
			}
		}
	}

	for _, nextBottom := range nextBottoms {
		res := pyramidTransitionHelper(nextBottom)
		visited[nextBottom] = res
		if res {
			return true
		}
	}
	return false
}
