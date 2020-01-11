package piscine

func MakeRange(min, max int) []int {
	result := []int(nil)

	if min >= max {
		return result
	}

	len := max - min
	result = make([]int, len)

	i := 0
	for nombre := min; nombre < max; nombre++ {
		result[i] = nombre
		i++
	}

	return result
}
