package piscine

func ConcatParams(args []string) string {
	res := ""

	for id, val := range args {
		if id != 0 {
			res += "\n"
		}
		res += val
	}

	return res
}
