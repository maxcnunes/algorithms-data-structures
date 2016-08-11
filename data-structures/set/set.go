package set

// Set data scrtucture.
// In computer science, a set is an abstract data type that can store certain values,
// without any particular order, and no repeated values.
type Set struct {
	Size  int
	items []interface{}
}

// NewSet creates an empty Set data scrtucture
func NewSet() *Set {
	return &Set{}
}

// IndexOf returns the index of the item inside the Set
func (s *Set) IndexOf(item interface{}) int {
	for i, v := range s.items {
		if v == item {
			return i
		}
	}
	return -1
}

// Contains check if the item is inside the Set
func (s *Set) Contains(item interface{}) bool {
	return s.IndexOf(item) != -1
}

// Add a new item into the Set
func (s *Set) Add(item interface{}) {
	if s.Contains(item) {
		return
	}

	s.items = append(s.items, item)
	s.Size++
}

// Remove an existing item from the Set
func (s *Set) Remove(item interface{}) {
	itemIndex := s.IndexOf(item)
	if itemIndex == -1 {
		return
	}

	// better performance moving the last item
	// to the removed item position
	lastIndex := s.Size - 1
	if itemIndex != lastIndex {
		s.items[itemIndex] = s.items[lastIndex]
	}

	s.items[lastIndex] = nil
	s.Size--
}
