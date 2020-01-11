package piscine

func ListMerge(l1 *List, l2 *List) {
	head := l1.Head
	var prev *NodeL

	for head != nil {
		prev = head
		head = head.Next
	}
	if prev != nil {
		prev.Next = l2.Head
	} else {
		l1.Head = l2.Head
	}
}
