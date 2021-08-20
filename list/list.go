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
//
// The complexity is O(n).
func New(values ...int) *linkList {
	l := &linkList{}
	for _, v := range values {
		l.Append(v)
	}
	return l
}

// Len returns the count of elements in the list.
//
// The complexity is O(1).
func (ll *linkList) Len() int {
	return ll.len
}

// Prepend creates a new element at the beginning of the list.
//
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
//
// Given the index is not in the set of indexes no item will be inserted.
//
// The complexity is O(n).
//
// Examples:
//	- given [1, 2, 3] Insert(0, 4) = [4, 1, 2, 3].
//	- given [1, 2, 3] Insert(1, 4) = [1, 4, 2, 3].
//	- given [1, 2, 3] Insert(2, 4) = [1, 2, 4, 3].
//	- given [1, 2, 3] Insert(3, 4) = [1, 2, 3, 4].
func (ll *linkList) Insert(index, value int) {
	if index == 0 {
		ll.Prepend(value)
		return
	}
	if index == ll.Len() {
		ll.Append(value)
		return
	}

	currentNode := ll.head
	currentIndex := 1
	for currentNode != nil {
		if currentIndex == index {
			next := currentNode.next
			nn := &node{
				prev:  currentNode,
				next:  next,
				value: value,
			}
			currentNode.next = nn
			next.prev = nn
			return
		}
		currentNode = currentNode.next
		currentIndex++
	}
}

// Append adds a new element to the end of the list.
//
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
//
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
	ll.head = ll.head.next
	ll.head.prev = nil
}

// Remove removes an element for a zero based index.
//
// Given the index is not in the set of indexes no item will be removed.
//
// The complexity is O(n).
func (ll *linkList) Remove(index int) {
	if index == 0 {
		ll.Shift()
		return
	}
	if index == ll.Len()-1 {
		ll.Pop()
		return
	}

	currentNode := ll.head
	currentIndex := 0
	for currentNode != nil {
		if currentIndex == index {
			prevNode := currentNode.prev
			nextNode := currentNode.next
			prevNode.next = nextNode
			nextNode.prev = prevNode
			return
		}
		currentNode = currentNode.next
		currentIndex++
	}
}

// Pop removes the last element in the list.
//
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
//
// Given indexA or indexB is not in the set of indexes no items will be swapped.
//
// Note this swaps values, but not references.
//
// The complexity is O(n).
func (ll *linkList) Swap(indexA, indexB int) {
	currentNode := ll.head
	currentIndex := 0
	var nodeA *node
	var nodeB *node
	for currentNode != nil {
		if currentIndex == indexA {
			nodeA = currentNode
		}
		if currentIndex == indexB {
			nodeB = currentNode
		}
		if nodeA != nil && nodeB != nil {
			break
		}
		currentNode = currentNode.next
		currentIndex++
	}
	if nodeA != nil && nodeB != nil {
		nodeA.value, nodeB.value = nodeB.value, nodeA.value
	}
}
