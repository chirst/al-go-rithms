package list

import (
	"testing"
)

func TestNew(t *testing.T) {
	l := New(1, 2, 3)
	if l.Len() != 3 {
		t.Fatalf("expected new to generate 3 values")
	}
}

func TestLen(t *testing.T) {

	t.Run("prepend", func(t *testing.T) {
		l := New()

		if len := l.Len(); len != 0 {
			t.Fatalf("expected len to be 0 got %v", len)
		}

		l.Prepend(1)
		if len := l.Len(); len != 1 {
			t.Fatalf("expected len to be 1 got %v", len)
		}

		l.Prepend(2)
		if len := l.Len(); len != 2 {
			t.Fatalf("expected len to be 2 got %v", len)
		}

		l.Prepend(3)
		if len := l.Len(); len != 3 {
			t.Fatalf("expected len to be 3 got %v", len)
		}
	})

	t.Run("append", func(t *testing.T) {
		l := New()

		if len := l.Len(); len != 0 {
			t.Fatalf("expected len to be 0 got %v", len)
		}

		l.Append(1)
		if len := l.Len(); len != 1 {
			t.Fatalf("expected len to be 1 got %v", len)
		}

		l.Append(2)
		if len := l.Len(); len != 2 {
			t.Fatalf("expected len to be 2 got %v", len)
		}

		l.Append(3)
		if len := l.Len(); len != 3 {
			t.Fatalf("expected len to be 3 got %v", len)
		}
	})

	t.Run("shift", func(t *testing.T) {
		l := New()

		l.Append(1)
		l.Append(2)
		l.Append(3)

		l.Shift()
		if len := l.Len(); len != 2 {
			t.Fatalf("expected len to be 2 got %v", len)
		}
		l.Shift()
		if len := l.Len(); len != 1 {
			t.Fatalf("expected len to be 1 got %v", len)
		}
		l.Shift()
		if len := l.Len(); len != 0 {
			t.Fatalf("expected len to be 0 got %v", len)
		}
		l.Shift()
		if len := l.Len(); len != 0 {
			t.Fatalf("expected len to be 0 got %v", len)
		}
	})

	t.Run("pop", func(t *testing.T) {
		l := New()

		l.Append(1)
		l.Append(2)
		l.Append(3)

		l.Pop()
		if len := l.Len(); len != 2 {
			t.Fatalf("expected len to be 2 got %v", len)
		}
		l.Pop()
		if len := l.Len(); len != 1 {
			t.Fatalf("expected len to be 1 got %v", len)
		}
		l.Pop()
		if len := l.Len(); len != 0 {
			t.Fatalf("expected len to be 0 got %v", len)
		}
		l.Pop()
		if len := l.Len(); len != 0 {
			t.Fatalf("expected len to be 0 got %v", len)
		}
	})
}

func TestPrepend(t *testing.T) {
	l := New()

	l.Prepend(1)
	l.Prepend(2)
	l.Prepend(3)

	firstNode := l.getNode(0)
	secondNode := l.getNode(1)
	thirdNode := l.getNode(2)

	if firstNode == nil {
		t.Fatalf("expected first node not to be nil")
	}
	if secondNode == nil {
		t.Fatalf("expected second node not to be nil")
	}
	if thirdNode == nil {
		t.Fatalf("expected third node not to be nil")
	}

	if firstNode.value != 3 {
		t.Fatalf("expected first node value of 3")
	}
	if secondNode.value != 2 {
		t.Fatalf("expected second node value of 2")
	}
	if thirdNode.value != 1 {
		t.Fatalf("expected third node value of 1")
	}

	if firstNode.prev != nil {
		t.Fatalf("expected first node prev value to be nil")
	}
	if firstNode.next != secondNode {
		t.Fatalf("expected first node next to be second node")
	}

	if secondNode.prev != firstNode {
		t.Fatalf("expected second node prev to be first node")
	}
	if secondNode.next != thirdNode {
		t.Fatalf("expected second node next to be third node")
	}

	if thirdNode.prev != secondNode {
		t.Fatalf("expected third node prev to be second node")
	}
	if thirdNode.next != nil {
		t.Fatalf("expected third node next to be nil")
	}
}

func TestAppend(t *testing.T) {
	l := New()

	l.Append(1)
	l.Append(2)
	l.Append(3)

	firstNode := l.getNode(0)
	secondNode := l.getNode(1)
	thirdNode := l.getNode(2)

	if firstNode == nil {
		t.Fatalf("expected first node not to be nil")
	}
	if secondNode == nil {
		t.Fatalf("expected second node not to be nil")
	}
	if thirdNode == nil {
		t.Fatalf("expected third node not to be nil")
	}

	if firstNode.value != 1 {
		t.Fatalf("expected first node value of 1")
	}
	if secondNode.value != 2 {
		t.Fatalf("expected second node value of 2")
	}
	if thirdNode.value != 3 {
		t.Fatalf("expected third node value of 3")
	}

	if firstNode.prev != nil {
		t.Fatalf("expected first node prev value to be nil")
	}
	if firstNode.next != secondNode {
		t.Fatalf("expected first node next to be second node")
	}

	if secondNode.prev != firstNode {
		t.Fatalf("expected second node prev to be first node")
	}
	if secondNode.next != thirdNode {
		t.Fatalf("expected second node next to be third node")
	}

	if thirdNode.prev != secondNode {
		t.Fatalf("expected third node prev to be second node")
	}
	if thirdNode.next != nil {
		t.Fatalf("expected third node next to be nil")
	}
}

func TestShift(t *testing.T) {
	l := New()

	l.Append(1)
	l.Append(2)
	l.Append(3)

	l.Shift()
	if l.getNode(0) == nil {
		t.Fatalf("expected first node not to be nil")
	}
	if l.getNode(1) == nil {
		t.Fatalf("expected second node not to be nil")
	}
	if l.getNode(2) != nil {
		t.Fatalf("expected third node to be nil")
	}
	if l.getNode(0).value != 2 {
		t.Fatalf("expected first node value of 2")
	}
	if l.getNode(1).value != 3 {
		t.Fatalf("expected second node value of 3")
	}

	l.Shift()
	if l.getNode(0) == nil {
		t.Fatalf("expected first node not to be nil")
	}
	if l.getNode(1) != nil {
		t.Fatalf("expected second node to be nil")
	}
	if l.getNode(2) != nil {
		t.Fatalf("expected third node to be nil")
	}
	if l.getNode(0).value != 3 {
		t.Fatalf("expected first node value of 3")
	}

	l.Shift()
	if l.getNode(0) != nil {
		t.Fatalf("expected first node to be nil")
	}
	if l.getNode(1) != nil {
		t.Fatalf("expected second node to be nil")
	}
	if l.getNode(2) != nil {
		t.Fatalf("expected third node to be nil")
	}

	// assert can shift empty list
	l.Shift()
}

func TestPop(t *testing.T) {
	l := New()

	l.Append(1)
	l.Append(2)
	l.Append(3)

	l.Pop()
	if l.getNode(0) == nil {
		t.Fatalf("expected first node not to be nil")
	}
	if l.getNode(1) == nil {
		t.Fatalf("expected second node not to be nil")
	}
	if l.getNode(2) != nil {
		t.Fatalf("expected third node to be nil")
	}
	if l.getNode(0).value != 1 {
		t.Fatalf("expected first node value of 1")
	}
	if l.getNode(1).value != 2 {
		t.Fatalf("expected second node value of 2")
	}

	l.Pop()
	if l.getNode(0) == nil {
		t.Fatalf("expected first node not to be nil")
	}
	if l.getNode(1) != nil {
		t.Fatalf("expected second node to be nil")
	}
	if l.getNode(2) != nil {
		t.Fatalf("expected third node to be nil")
	}
	if l.getNode(0).value != 1 {
		t.Fatalf("expected first node value of 1")
	}

	l.Pop()
	if l.getNode(0) != nil {
		t.Fatalf("expected first node to be nil")
	}
	if l.getNode(1) != nil {
		t.Fatalf("expected second node to be nil")
	}
	if l.getNode(2) != nil {
		t.Fatalf("expected third node to be nil")
	}

	// assert can pop empty list
	l.Pop()
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
