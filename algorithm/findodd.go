package algorithm

func FindOdd(in []int) int {
	mapIn := make(map[int]int, len(in))
	for _, v := range in {
		mapIn[v] = mapIn[v] + 1
	}

	for m, v := range mapIn {
		if v%2 != 0 {
			return m
		}
	}
	return -1
}
