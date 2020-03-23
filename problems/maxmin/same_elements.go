package maxmin


// Given two arrays, find the same elements in both of them
func SameElements(A []int, B []int ) []int {
	
	counterMap := make(map[int]int, 0)
	if len(A) > len(B) {
		counterMap = counter(A)
		return find(B, counterMap)
	}
	
	counterMap = counter(B)
	return find(A, counterMap)
}

func find(list []int, counterMap map[int]int) []int {
	result := make([]int, 0)

	for _, v := range list {
		num, ok := counterMap[v]
		if ok && num > 0 {
			result = append(result, v)
			counterMap[v]--
		}
	}

	return result
}

func counter(list []int) map[int]int {
	result := make(map[int]int, 0)

	for _, v := range list {
		result[v]++
	}

	return result
}