package linkedlist

// List is a data structure consisting of a group of nodes
// which together represent a sequence
type List struct {
	Size int
	Head *Node
	Tail *Node
}

// Node is an element of the list which stores a general data
// and information about the "neighbours" nodes in the list.
type Node struct {
	Data interface{}
	Prev *Node
	Next *Node
}

// NewList creates an empty list
func NewList() *List {
	list := &List{Size: 0}
	return list
}

// NewNode creates a node with the received data
func NewNode(data interface{}) *Node {
	return &Node{Data: data}
}

// Prepend add a node at the begining of the list
func (l *List) Prepend(data interface{}) {
	node := NewNode(data)
	if l.Size == 0 {
		l.Head = node
		l.Tail = node
	} else {
		oldHead := l.Head
		oldHead.Prev = node
		node.Next = oldHead
		l.Head = node
	}

	l.Size++
}

// Append add a node at the end of the list
func (l *List) Append(data interface{}) {
	node := NewNode(data)
	if l.Size == 0 {
		l.Head = node
		l.Tail = node
	} else {
		oldTail := l.Tail
		oldTail.Next = node
		node.Prev = oldTail
		l.Tail = node
	}

	l.Size++
}
