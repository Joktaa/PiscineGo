package piscine

func ListSize(l *List) int {
	var result int
	head := l.Head

	for head != nil {
		result++
		head = head.Next
	}

	return result
}
