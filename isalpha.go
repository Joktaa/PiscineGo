package piscine

func IsAlpha(str string) bool {
	for _, val := range str {
		if val < 48 || val > 122 {
			return false
		}
		if val > 90 && val < 97 {
			return false
		}
		if val > 57 && val < 65 {
			return false
		}
	}
	return true
}
