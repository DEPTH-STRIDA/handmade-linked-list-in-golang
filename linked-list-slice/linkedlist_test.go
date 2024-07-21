package main

import (
	"testing"
)

func TestPrepend(t *testing.T) {
	var list LinkedList
	list.Prepend(1)
	if list.GetElement(0) != 1 {
		t.Errorf("Expected 1, got %d", list.GetElement(0))
	}
}

func TestAppend(t *testing.T) {
	var list LinkedList
	list.Append(3)
	if list.GetElement(0) != 3 {
		t.Errorf("Expected 3, got %d", list.GetElement(0))
	}
}

func TestInsertAt(t *testing.T) {
	var list LinkedList
	list.Append(1)
	list.Append(3)
	list.InsertAt(2, 1)
	if list.GetElement(1) != 2 {
		t.Errorf("Expected 2, got %d", list.GetElement(1))
	}
}

func TestGetElement(t *testing.T) {
	var list LinkedList
	list.Append(1)
	if list.GetElement(0) != 1 {
		t.Errorf("Expected 1, got %d", list.GetElement(0))
	}
}

func TestDeleteAt(t *testing.T) {
	var list LinkedList
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.DeleteAt(1)
	if list.GetElement(1) != 3 {
		t.Errorf("Expected 3, got %d", list.GetElement(1))
	}
}

func TestDeleteFirst(t *testing.T) {
	var list LinkedList
	list.Append(1)
	list.Append(2)
	list.DeleteFirst()
	if list.GetElement(0) != 2 {
		t.Errorf("Expected 2, got %d", list.GetElement(0))
	}
}

func TestDeleteEnd(t *testing.T) {
	var list LinkedList
	list.Append(1)
	list.Append(2)
	list.DeleteEnd()
	if list.Length() != 1 || list.GetElement(0) != 1 {
		t.Errorf("Expected length 1 with element 1, got length %d with element %d", list.Length(), list.GetElement(0))
	}
}

func TestLength(t *testing.T) {
	var list LinkedList
	list.Append(1)
	list.Append(2)
	if list.Length() != 2 {
		t.Errorf("Expected length 2, got %d", list.Length())
	}
}

func TestReverse(t *testing.T) {
	var list LinkedList
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Reverse()
	if list.GetElement(0) != 3 || list.GetElement(1) != 2 || list.GetElement(2) != 1 {
		t.Errorf("Expected [3 2 1], got [%d %d %d]", list.GetElement(0), list.GetElement(1), list.GetElement(2))
	}
}
