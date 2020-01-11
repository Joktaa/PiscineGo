package piscine

//import "fmt"

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSizeNodeI(l *NodeI) int {
	var result int
	head := l

	for head != nil {
		result++
		head = head.Next
	}

	return result
}

func ListSort(l *NodeI) *NodeI {
	res := &NodeI{}
	head := l
	headRes := res
	tmpHead := res
	var lowest int
	var index int

	for i := 0; i < ListSizeNodeI(l); i++ {
		head = l
		lowest = head.Data
		for head != nil {
			if lowest > head.Data {
				lowest = head.Data
				index = i
			}
			head = head.Next
		}
		headRes.Data = lowest
		headRes.Next = &NodeI{}
		headRes = headRes.Next

		tmpHead = l
		for j := 0; j < index; j++ {
			tmpHead = tmpHead.Next
		}
		tmpHead.Data = 2147483647
	}

	return res
}
