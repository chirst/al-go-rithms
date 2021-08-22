// Package list is a doubly linked list that shouldn't be taken too seriously.
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

// New returns an instance of a list with the given values.
// The complexity is O(n).
func New(values ...int) *linkList {
	l := &linkList{}
	for _, v := range values {
		l.Append(v)
	}
	return l
}

// Len returns the count of elements in the list.
// The complexity is O(1).
func (ll *linkList) Len() int {
	return ll.len
}

// Prepend creates a new element at the beginning of the list.
// The complexity is O(1).
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

// Insert inserts an element for a zero based index.
// Given the index is not in the set of indexes no item will be inserted.
// The complexity is O(n).
func (ll *linkList) Insert(index, value int) {

}

// Append adds a new element to the end of the list.
// The complexity is O(1).
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

// Shift removes the first element in the list.
// The complexity is O(1).
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

// Remove removes an element for a zero based index.
// Given the index is not in the set of indexes no item will be removed.
// The complexity is O(n).
func (ll *linkList) Remove(index int) {

}

// Pop removes the last element in the list.
// The complexity is O(1).
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

// Swap swaps two elements in the list for two zero based indexes.
// Given indexA or indexB is not in the set of indexes no items will be swapped.
// The complexity is O(n).
func (ll *linkList) Swap(indexA, indexB int) {

}
