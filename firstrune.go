package piscine

func FirstRune(s string) rune {
	for id, val := range s {
		if id == 0 {
			return rune(val)
		}
	}
	return 0
}
