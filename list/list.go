// Package list is a doubly linked list
package list

type linkList struct {
	head *node
	tail *node
	len  int
}

type node struct {
	prev  *node
	next  *node
	value int
}

// New returns an instance of a list
func New(vs ...int) *linkList {
	l := &linkList{}
	for v := range vs {
		l.Append(v)
	}
	return l
}

// Len returns the count of elements in the list
func (ll *linkList) Len() int {
	return ll.len
}

// Prepend creates a new element at the beginning of the list
func (ll *linkList) Prepend(value int) {
	ll.len++
	if ll.head != nil {
		oldHead := ll.head
		ll.head = &node{
			next:  oldHead,
			value: value,
		}
		ll.head.next.prev = ll.head
		return
	}
	ll.head = &node{
		value: value,
	}
	ll.tail = ll.head
}

// Append adds a new element to the end of the list
func (ll *linkList) Append(value int) {
	ll.len++
	if ll.head == nil {
		ll.head = &node{
			value: value,
		}
		ll.tail = ll.head
		return
	}
	ll.tail.next = &node{
		prev:  ll.tail,
		value: value,
	}
	ll.tail = ll.tail.next
}

// Shift removes the first element in the list
func (ll *linkList) Shift() {
	if ll.head == nil {
		return
	}
	ll.len--
	if ll.head.next == nil {
		ll.head = nil
		ll.tail = nil
		return
	}
	ll.head.next.prev = nil
	ll.head = ll.head.next
}

// Pop removes the last element in the list
func (ll *linkList) Pop() {
	if ll.head == nil {
		return
	}
	ll.len--
	if ll.head.next == nil {
		ll.head = nil
		ll.tail = nil
		return
	}
	ll.tail = ll.tail.prev
	ll.tail.next = nil
}
