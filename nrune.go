package piscine

func NRune(s string, n int) rune {
	for id, val := range s {
		if id == n-1 {
			return val
		}
	}
	return 0
}
