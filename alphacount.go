package piscine

func AlphaCount(str string) int {
	counter := 0

	for _, val := range str {
		Rval := rune(val)
		if Rval > 96 && Rval < 123 {
			counter++
		}
		if Rval > 64 && Rval < 91 {
			counter++
		}
	}
	return counter
}
