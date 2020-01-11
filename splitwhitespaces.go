package piscine

func SplitWhiteSpaces(str string) []string {
	size := 1
	i_res := 0
	for i, char := range str {
		if (char == ' ' && str[i+1] != ' ' && i != 0) || char == '\n' || (char == '	' && str[i+1] != '	') {
			size++
		}
	}
	res := make([]string, size)
	word := ""
	for i, char := range str {
		if (char == ' ' && str[i+1] != ' ' && i != 0) || char == '\n' || (char == '	' && str[i+1] != '	') {
			res[i_res] = word
			i_res++
			word = ""
		} else if char != ' ' || char != '\n' || char != '	' {
			if !(str[i] == ' ') {
				word += string(char)
			}
		}
	}
	res[i_res] = word
	return res
}
