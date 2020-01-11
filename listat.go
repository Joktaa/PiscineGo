package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	i := 1
	head := l

	for i <= pos {
		if head == nil {
			return nil
		}
		head = head.Next
		i++
	}

	return head
}
