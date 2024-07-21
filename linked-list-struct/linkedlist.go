package main

// Node represents a node in the linked list.
type Node struct {
	Value int
	Next  *Node
}

// LinkedList represents a linked list structure.
type LinkedList struct {
	Head *Node
}

// Append adds an element to the end of the linked list.
func (list *LinkedList) Append(value int) {
	newNode := &Node{Value: value}
	if list.Head == nil {
		list.Head = newNode
		return
	}
	current := list.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// DeleteEnd removes the last element from the linked list.
func (list *LinkedList) DeleteEnd() {
	if list.Head == nil {
		return
	}
	if list.Head.Next == nil {
		list.Head = nil
		return
	}
	current := list.Head
	for current.Next.Next != nil {
		current = current.Next
	}
	current.Next = nil
}

// DeleteFirst removes the first element from the linked list.
func (list *LinkedList) DeleteFirst() {
	if list.Head != nil {
		list.Head = list.Head.Next
	}
}

// Prepend adds an element to the beginning of the linked list.
func (list *LinkedList) Prepend(value int) {
	newNode := &Node{Value: value, Next: list.Head}
	list.Head = newNode
}

// InsertAt inserts an element at a specific index in the linked list.
func (list *LinkedList) InsertAt(value int, index int) {
	if index == 0 {
		list.Prepend(value)
		return
	}
	newNode := &Node{Value: value}
	current := list.Head
	for i := 0; current != nil && i < index-1; i++ {
		current = current.Next
	}
	if current == nil {
		return
	}
	newNode.Next = current.Next
	current.Next = newNode
}

// DeleteAt removes an element at a specific index in the linked list.
func (list *LinkedList) DeleteAt(index int) {
	if list.Head == nil {
		return
	}
	if index == 0 {
		list.Head = list.Head.Next
		return
	}
	current := list.Head
	for i := 0; current.Next != nil && i < index-1; i++ {
		current = current.Next
	}
	if current.Next == nil {
		return
	}
	current.Next = current.Next.Next
}

// Length returns the length of the linked list.
func (list *LinkedList) Length() int {
	length := 0
	current := list.Head
	for current != nil {
		length++
		current = current.Next
	}
	return length
}

// Reverse reverses the linked list.
func (list *LinkedList) Reverse() {
	var prev, next *Node
	current := list.Head
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	list.Head = prev
}

// GetElement returns the element at a specific index in the linked list.
func (list *LinkedList) GetElement(index int) int {
	current := list.Head
	for i := 0; current != nil && i < index; i++ {
		current = current.Next
	}
	if current == nil {
		return -1 // or some error value, as list is not long enough
	}
	return current.Value
}
