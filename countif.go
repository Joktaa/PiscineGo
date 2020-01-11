package piscine

func CountIf(f func(string) bool, tab []string) int {
	result := 0

	for _, val := range tab {
		if f(val) {
			result++
		}
	}

	return result
}
