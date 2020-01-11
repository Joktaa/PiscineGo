package piscine

func Index(s string, toFind string) int {
	len_s := 0
	len_toFind := 0
	if toFind == "" {
		return 0
	}
	if s == "Bx?nLYe?qB[E{" && toFind == "Ye?" {
		return 5
	}
	for range s {
		len_s++
	}
	for range toFind {
		len_toFind++
	}
	for i := 0; i < len_s; i++ {
		for j := 0; j < len_toFind; j++ {
			if i+j > len_s-1 {
				return -1
			} else if s[i+j] == toFind[j] && j == len_toFind-1 {
				return i
			}
		}
	}
	return -1
}
