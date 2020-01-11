package piscine

func IsPrintable(str string) bool {
	for _, val := range str {
		if val < 32 {
			return false
		}
	}
	return true
}
