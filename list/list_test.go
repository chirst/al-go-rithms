package list

import (
	"testing"
)

func TestNew(t *testing.T) {
	l := New(1, 2, 3)
	checkLen(t, l, 3)
}

func TestLen(t *testing.T) {

	t.Run("prepend", func(t *testing.T) {
		l := New()

		checkLen(t, l, 0)
		l.Prepend(1)
		checkLen(t, l, 1)
		l.Prepend(2)
		checkLen(t, l, 2)
		l.Prepend(3)
		checkLen(t, l, 3)
	})

	t.Run("append", func(t *testing.T) {
		l := New()

		checkLen(t, l, 0)
		l.Append(1)
		checkLen(t, l, 1)
		l.Append(2)
		checkLen(t, l, 2)
		l.Append(3)
		checkLen(t, l, 3)
	})

	t.Run("shift", func(t *testing.T) {
		l := New(1, 2, 3)

		checkLen(t, l, 3)
		l.Shift()
		checkLen(t, l, 2)
		l.Shift()
		checkLen(t, l, 1)
		l.Shift()
		checkLen(t, l, 0)
		l.Shift()
		checkLen(t, l, 0)
	})

	t.Run("pop", func(t *testing.T) {
		l := New(1, 2, 3)

		checkLen(t, l, 3)
		l.Pop()
		checkLen(t, l, 2)
		l.Pop()
		checkLen(t, l, 1)
		l.Pop()
		checkLen(t, l, 0)
		l.Pop()
		checkLen(t, l, 0)
	})
}

func checkLen(t *testing.T, l *linkList, expectedLen int) {
	if len := l.Len(); len != expectedLen {
		t.Errorf("expected len to be %v got %v", expectedLen, len)
	}
}

func TestPrepend(t *testing.T) {
	l := New()

	l.Prepend(1)
	l.Prepend(2)
	l.Prepend(3)

	checkNodeValue(t, l, 0, 3)
	checkNodeValue(t, l, 1, 2)
	checkNodeValue(t, l, 2, 1)

	checkNodePrev(t, l, 0, nil)
	checkNodeNext(t, l, 0, l.getNode(1))

	checkNodePrev(t, l, 1, l.getNode(0))
	checkNodeNext(t, l, 1, l.getNode(2))

	checkNodePrev(t, l, 2, l.getNode(1))
	checkNodeNext(t, l, 2, nil)
}

func TestAppend(t *testing.T) {
	l := New()

	l.Append(1)
	l.Append(2)
	l.Append(3)

	checkNodeValue(t, l, 0, 1)
	checkNodeValue(t, l, 1, 2)
	checkNodeValue(t, l, 2, 3)

	checkNodePrev(t, l, 0, nil)
	checkNodeNext(t, l, 0, l.getNode(1))

	checkNodePrev(t, l, 1, l.getNode(0))
	checkNodeNext(t, l, 1, l.getNode(2))

	checkNodePrev(t, l, 2, l.getNode(1))
	checkNodeNext(t, l, 2, nil)
}

func TestShift(t *testing.T) {
	l := New()

	l.Append(1)
	l.Append(2)
	l.Append(3)

	l.Shift()
	checkNodeValue(t, l, 0, 2)
	checkNodeValue(t, l, 1, 3)
	checkNodeNil(t, l, 2)

	l.Shift()
	checkNodeValue(t, l, 0, 3)
	checkNodeNil(t, l, 1)
	checkNodeNil(t, l, 2)

	l.Shift()
	checkNodeNil(t, l, 0)
	checkNodeNil(t, l, 1)
	checkNodeNil(t, l, 2)
}

func TestPop(t *testing.T) {
	l := New()

	l.Append(1)
	l.Append(2)
	l.Append(3)

	l.Pop()
	checkNodeValue(t, l, 0, 1)
	checkNodeValue(t, l, 1, 2)
	checkNodeNil(t, l, 2)

	l.Pop()
	checkNodeValue(t, l, 0, 1)
	checkNodeNil(t, l, 1)
	checkNodeNil(t, l, 2)

	l.Pop()
	checkNodeNil(t, l, 0)
	checkNodeNil(t, l, 1)
	checkNodeNil(t, l, 2)
}

func checkNodeValue(t *testing.T, l *linkList, nodeIndex int, wantValue int) {
	n := l.getNode(nodeIndex)
	if n == nil {
		t.Errorf("expected node value at index: %v, not to be nil", nodeIndex)
	}
	if n.value != wantValue {
		t.Errorf(
			"expected node value at index: %v to be %v got %v",
			nodeIndex,
			wantValue,
			n.value,
		)
	}
}

func checkNodePrev(t *testing.T, l *linkList, nodeIndex int, wantNode *node) {
	n := l.getNode(nodeIndex)
	if n == nil {
		t.Errorf("expected node value at index: %v, not to be nil", nodeIndex)
	}
	if n.prev != wantNode {
		t.Errorf(
			"expected node prev at index: %v to be %v got %v",
			nodeIndex,
			wantNode,
			n.prev,
		)
	}
}

func checkNodeNext(t *testing.T, l *linkList, nodeIndex int, wantNode *node) {
	n := l.getNode(nodeIndex)
	if n == nil {
		t.Errorf("expected node value at index: %v, not to be nil", nodeIndex)
	}
	if n.next != wantNode {
		t.Errorf(
			"expected node next at index: %v to be %v got %v",
			nodeIndex,
			wantNode,
			n.next,
		)
	}
}

func checkNodeNil(t *testing.T, l *linkList, nodeIndex int) {
	n := l.getNode(nodeIndex)
	if n != nil {
		t.Errorf("expected node at index: %v to be nil", nodeIndex)
	}
}

func (ll *linkList) getNode(index int) *node {
	count := 0
	currentNode := ll.head
	for currentNode != nil {
		if count == index {
			return currentNode
		}
		count++
		currentNode = currentNode.next
	}
	return nil
}
