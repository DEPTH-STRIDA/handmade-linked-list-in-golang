package main

// LinkedList represents a linked list structure that uses a slice to store data.
type LinkedList struct {
	data []int
}

// Append adds an element to the end of the linked list.
func (list *LinkedList) Append(num int) {
	list.data = append(list.data, num)
}

// DeleteEnd removes the last element from the linked list.
func (list *LinkedList) DeleteEnd() {
	if len(list.data) > 0 {
		list.data = list.data[:len(list.data)-1]
	}
}

// DeleteFirst removes the first element from the linked list.
func (list *LinkedList) DeleteFirst() {
	if len(list.data) > 0 {
		list.data = list.data[1:]
	}
}

// Prepend adds an element to the beginning of the linked list.
func (list *LinkedList) Prepend(num int) {
	list.data = append([]int{num}, list.data...)
}

// InsertAt inserts an element at a specific index in the linked list.
func (list *LinkedList) InsertAt(num int, index int) {
	if index == 0 {
		list.Prepend(num)
		return
	}
	if index == len(list.data) {
		list.Append(num)
		return
	}
	list.data = append(list.data[:index], append([]int{num}, list.data[index:]...)...)
}

// DeleteAt removes an element at a specific index in the linked list.
func (list *LinkedList) DeleteAt(index int) {
	if index == 0 {
		list.DeleteFirst()
		return
	}
	if index == len(list.data)-1 {
		list.DeleteEnd()
		return
	}
	list.data = append(list.data[:index], list.data[index+1:]...)
}

// Length returns the length of the linked list.
func (list *LinkedList) Length() int {
	return len(list.data)
}

// Reverse reverses the linked list.
func (list *LinkedList) Reverse() {
	for i, j := 0, len(list.data)-1; i < j; i, j = i+1, j-1 {
		list.data[i], list.data[j] = list.data[j], list.data[i]
	}
}

// GetElement returns the element at a specific index in the linked list.
func (list *LinkedList) GetElement(index int) int {
	return list.data[index]
}
