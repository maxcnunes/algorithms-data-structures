package linkedlist

import "testing"

func TestLinkedList(t *testing.T) {
	list := NewList()
	if list == nil {
		t.Error("Linked list has not been created correctly", list)
		return
	}

	if list.Head != nil || list.Tail != nil || list.Size != 0 {
		t.Error("Linked list should have been created empty", list)
		return
	}

	list.Prepend("Jack")
	if list.Size != 1 {
		t.Error("Linked list has not incremented the size of the list for the new element", list)
		return
	}

	if list.Head.Data != "Jack" || list.Tail.Data != "Jack" {
		t.Error("Linked list has not prepend the element right", list)
		return
	}

	list.Prepend("Maria")
	if list.Size != 2 {
		t.Error("Linked list has not incremented the size of the list for the new element", list)
		return
	}

	if list.Head.Data != "Maria" || list.Tail.Data != "Jack" {
		t.Error("Linked list has not prepend the element right", list)
		return
	}

	list.Append("Tailor")
	if list.Size != 3 {
		t.Error("Linked list has not incremented the size of the list for the new element", list)
		return
	}

	if list.Head.Data != "Maria" || list.Tail.Data != "Tailor" {
		t.Error("Linked list has not append the element right", *list)
		return
	}
}
