package piscine

func IsNumeric(str string) bool {
	for _, val := range str {
		if val < 48 || val > 57 {
			return false
		}
	}
	return true
}
