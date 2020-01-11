package piscine

func AppendRange(min, max int) []int {
	result := []int(nil)

	if min > max {
		return result
	}

	for i := min; i < max; i++ {
		result = append(result, i)
	}

	return result
}
