// Package list is a doubly linked list
package list

type linkList struct {
	head *node
	len  int
}

type node struct {
	prev  *node
	next  *node
	value int
}

// New returns an instance of a list
func New() *linkList {
	return &linkList{}
}

// Len returns the count of elements in the list
func (ll *linkList) Len() int {
	return ll.len
}

// Prepend creates a new element at the beginning of the list
func (ll *linkList) Prepend(value int) {
	if ll.head != nil {
		oldHead := ll.head
		ll.head = &node{
			next:  oldHead,
			value: value,
		}
		ll.head.next.prev = ll.head
		ll.len++
		return
	}
	ll.head = &node{
		value: value,
	}
	ll.len++
}

// Append adds a new element to the end of the list
func (ll *linkList) Append(value int) {
	currentNode := ll.head
	for currentNode != nil {
		if currentNode.next == nil {
			currentNode.next = &node{
				prev:  currentNode,
				value: value,
			}
			ll.len++
			return
		}
		currentNode = currentNode.next
	}
	ll.head = &node{
		value: value,
	}
	ll.len++
}

// Shift removes the first element in the list
func (ll *linkList) Shift() {
	if ll.head == nil {
		return
	}
	if ll.head.next == nil {
		ll.head = nil
		ll.len = 0
		return
	}
	ll.head.next.prev = nil
	ll.head = ll.head.next
	ll.len--
}

// Pop removes the last element in the list
func (ll *linkList) Pop() {
	if ll.head == nil {
		return
	}
	if ll.head.next == nil {
		ll.head = nil
		ll.len = 0
		return
	}
	currentNode := ll.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.prev.next = nil
	ll.len--
}
