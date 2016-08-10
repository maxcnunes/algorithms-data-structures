package linkedlist

import "errors"

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

// Add a new element in a specific position
func (l *List) Add(data interface{}, index int) error {
	if index > l.Size {
		return errors.New("Index out of range")
	}

	if index == 0 || l.Size == 0 {
		l.Prepend(data)
		return nil
	}

	if l.Size-1 == index {
		l.Append(data)
		return nil
	}

	newNode := NewNode(data)
	nextNode, _ := l.Get(index)
	prevNode := nextNode.Prev

	prevNode.Next = newNode
	newNode.Prev = prevNode

	nextNode.Prev = newNode
	newNode.Next = nextNode

	l.Size++

	return nil
}

// Get the element by index
func (l *List) Get(index int) (*Node, error) {
	if index > l.Size {
		return nil, errors.New("Index out of range")
	}

	node := l.Head
	for i := 0; i < index; i++ {
		node = node.Next
	}

	return node, nil
}
