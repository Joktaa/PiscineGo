package piscine

func Map(f func(int) bool, arr []int) []bool {
	var len int
	for range arr {
		len++
	}

	result := make([]bool, len)

	for id, val := range arr {
		result[id] = f(val)
	}

	return result
}
