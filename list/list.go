// Package list is a doubly linked list that shouldn't be taken too seriously.
package list

// TODO:
// - Implement Sort

type linkList[T comparable] struct {
	head *node[T]
	tail *node[T]
	len  int
}

type node[T comparable] struct {
	prev  *node[T]
	next  *node[T]
	value T
}

// New returns an instance of a list with the given values.
//
// The complexity is O(n).
func New[T comparable](values ...T) *linkList[T] {
	l := &linkList[T]{}
	for _, v := range values {
		l.Append(v)
	}
	return l
}

// Len returns the count of elements in the list.
//
// The complexity is O(1).
func (ll *linkList[T]) Len() int {
	return ll.len
}

// Prepend creates a new element at the beginning of the list.
//
// The complexity is O(1).
func (ll *linkList[T]) Prepend(value T) {
	ll.len++
	if ll.head != nil {
		oldHead := ll.head
		ll.head = &node[T]{
			next:  oldHead,
			value: value,
		}
		ll.head.next.prev = ll.head
		return
	}
	ll.head = &node[T]{
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
func (ll *linkList[T]) Insert(index int, value T) {
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
			nn := &node[T]{
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
func (ll *linkList[T]) Append(value T) {
	ll.len++
	if ll.head == nil {
		ll.head = &node[T]{
			value: value,
		}
		ll.tail = ll.head
		return
	}
	ll.tail.next = &node[T]{
		prev:  ll.tail,
		value: value,
	}
	ll.tail = ll.tail.next
}

// Shift removes the first element in the list.
//
// Returns the value of the removed element or nil if the list is empty.
//
// The complexity is O(1).
func (ll *linkList[T]) Shift() *T {
	if ll.head == nil {
		return nil
	}
	ll.len--
	ret := ll.head.value
	if ll.head.next == nil {
		ll.head = nil
		ll.tail = nil
		return &ret
	}
	ll.head = ll.head.next
	ll.head.prev = nil
	return &ret
}

// Remove removes an element for a zero based index.
//
// Given the index is not in the set of indexes no item will be removed.
//
// Returns the value of the removed element or nil if nothing is removed.
//
// The complexity is O(n).
func (ll *linkList[T]) Remove(index int) *T {
	if index == 0 {
		return ll.Shift()
	}
	if index == ll.Len()-1 {
		return ll.Pop()
	}

	currentNode := ll.head
	currentIndex := 0
	for currentNode != nil {
		if currentIndex == index {
			prevNode := currentNode.prev
			nextNode := currentNode.next
			prevNode.next = nextNode
			nextNode.prev = prevNode
			return &currentNode.value
		}
		currentNode = currentNode.next
		currentIndex++
	}
	return nil
}

// Pop removes the last element in the list.
//
// Returns the value of the removed element or nil if the list is empty.
//
// The complexity is O(1).
func (ll *linkList[T]) Pop() *T {
	if ll.head == nil {
		return nil
	}
	ll.len--
	ret := ll.tail.value
	if ll.head.next == nil {
		ll.head = nil
		ll.tail = nil
		return &ret
	}
	ll.tail = ll.tail.prev
	ll.tail.next = nil
	return &ret
}

// Swap swaps two elements in the list for two zero based indexes.
//
// Given indexA or indexB is not in the set of indexes no items will be swapped.
//
// Note this swaps values, but not references.
//
// The complexity is O(n).
func (ll *linkList[T]) Swap(indexA, indexB int) {
	currentNode := ll.head
	currentIndex := 0
	var nodeA *node[T]
	var nodeB *node[T]
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

// Get returns the value of an element in the list for a zero based index. If no
// element matches the given index, nil is returned.
//
// The complexity is O(n)
func (ll *linkList[T]) Get(index int) *T {
	count := 0
	currentNode := ll.head
	for currentNode != nil {
		if count == index {
			return &currentNode.value
		}
		count++
		currentNode = currentNode.next
	}
	return nil
}
