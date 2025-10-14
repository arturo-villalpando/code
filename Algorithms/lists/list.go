package lists

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Append(val int) {
	newNode := &Node{Val: val, Next: nil}
	if l.Head == nil {
		l.Head = newNode
		return
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}
