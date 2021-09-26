package btree

import "testing"

func TestInsertDegree3(t *testing.T) {
	bt, _ := New(3)

	t.Run("insert 1", func(t *testing.T) {
		bt.Insert(1)
		bt.root.checkElements(t, 1)
		bt.root.checkChildrenLength(t, 0)
	})

	t.Run("insert 2", func(t *testing.T) {
		bt.Insert(2)
		bt.root.checkElements(t, 1, 2)
		bt.root.checkChildrenLength(t, 0)
	})

	t.Run("insert 3", func(t *testing.T) {
		bt.Insert(3)
		bt.root.checkElements(t, 2)
		bt.root.checkChildrenLength(t, 2)
		bt.root.children[0].checkElements(t, 1)
		bt.root.children[0].checkChildrenLength(t, 0)
		bt.root.children[1].checkElements(t, 3)
		bt.root.children[1].checkChildrenLength(t, 0)
	})

	t.Run("insert 4", func(t *testing.T) {
		bt.Insert(4)
		bt.root.checkElements(t, 2)
		bt.root.checkChildrenLength(t, 2)
		bt.root.children[0].checkElements(t, 1)
		bt.root.children[0].checkChildrenLength(t, 0)
		bt.root.children[1].checkElements(t, 3, 4)
		bt.root.children[1].checkChildrenLength(t, 0)
	})

	t.Run("insert 5", func(t *testing.T) {
		bt.Insert(5)
		bt.root.checkElements(t, 2, 4)
		bt.root.checkChildrenLength(t, 3)
		bt.root.children[0].checkElements(t, 1)
		bt.root.children[0].checkChildrenLength(t, 0)
		bt.root.children[1].checkElements(t, 3)
		bt.root.children[1].checkChildrenLength(t, 0)
		bt.root.children[2].checkElements(t, 5)
		bt.root.children[2].checkChildrenLength(t, 0)
	})

	t.Run("insert 6", func(t *testing.T) {
		bt.Insert(6)
		bt.root.checkElements(t, 2, 4)
		bt.root.checkChildrenLength(t, 3)
		bt.root.children[0].checkElements(t, 1)
		bt.root.children[0].checkChildrenLength(t, 0)
		bt.root.children[1].checkElements(t, 3)
		bt.root.children[1].checkChildrenLength(t, 0)
		bt.root.children[2].checkElements(t, 5, 6)
		bt.root.children[2].checkChildrenLength(t, 0)
	})

	t.Run("insert 7", func(t *testing.T) {
		bt.Insert(7)

		// top level
		bt.root.checkElements(t, 4)
		bt.root.checkChildrenLength(t, 2)

		// second level left to right
		bt.root.children[0].checkElements(t, 2)
		bt.root.children[0].checkChildrenLength(t, 2)
		bt.root.children[0].children[0].checkElements(t, 1)
		bt.root.children[0].children[0].checkChildrenLength(t, 0)
		bt.root.children[0].children[1].checkElements(t, 3)
		bt.root.children[0].children[1].checkChildrenLength(t, 0)

		// third level left to right
		bt.root.children[1].checkElements(t, 6)
		bt.root.children[1].checkChildrenLength(t, 2)
		bt.root.children[1].children[0].checkElements(t, 5)
		bt.root.children[1].children[0].checkChildrenLength(t, 0)
		bt.root.children[1].children[1].checkElements(t, 7)
		bt.root.children[1].children[1].checkChildrenLength(t, 0)
	})
}

func TestInsertDegree4(t *testing.T) {
	bt, _ := New(4)

	t.Run("insert 1", func(t *testing.T) {
		bt.Insert(1)
		bt.root.checkElements(t, 1)
		bt.root.checkChildrenLength(t, 0)
	})

	t.Run("insert 2", func(t *testing.T) {
		bt.Insert(2)
		bt.root.checkElements(t, 1, 2)
		bt.root.checkChildrenLength(t, 0)
	})

	t.Run("insert 3", func(t *testing.T) {
		bt.Insert(3)
		bt.root.checkElements(t, 1, 2, 3)
		bt.root.checkChildrenLength(t, 0)
	})

	t.Run("insert 4", func(t *testing.T) {
		bt.Insert(4)
		bt.root.checkElements(t, 2)
		bt.root.checkChildrenLength(t, 2)
		bt.root.children[0].checkElements(t, 1)
		bt.root.children[0].checkChildrenLength(t, 0)
		bt.root.children[1].checkElements(t, 3, 4)
		bt.root.children[1].checkChildrenLength(t, 0)
	})

	t.Run("insert 5", func(t *testing.T) {
		bt.Insert(5)
		bt.root.checkElements(t, 2)
		bt.root.checkChildrenLength(t, 2)
		bt.root.children[0].checkElements(t, 1)
		bt.root.children[0].checkChildrenLength(t, 0)
		bt.root.children[1].checkElements(t, 3, 4, 5)
		bt.root.children[1].checkChildrenLength(t, 0)
	})

	t.Run("insert 6", func(t *testing.T) {
		bt.Insert(6)
		bt.root.checkElements(t, 2, 4)
		bt.root.checkChildrenLength(t, 3)
		bt.root.children[0].checkElements(t, 1)
		bt.root.children[0].checkChildrenLength(t, 0)
		bt.root.children[1].checkElements(t, 3)
		bt.root.children[1].checkChildrenLength(t, 0)
		bt.root.children[2].checkElements(t, 5, 6)
		bt.root.children[2].checkChildrenLength(t, 0)
	})

	t.Run("insert 7", func(t *testing.T) {
		bt.Insert(7)

		// top level
		bt.root.checkElements(t, 2, 4)
		bt.root.checkChildrenLength(t, 3)

		// second level left to right
		bt.root.children[0].checkElements(t, 1)
		bt.root.children[0].checkChildrenLength(t, 0)
		bt.root.children[1].checkElements(t, 3)
		bt.root.children[1].checkChildrenLength(t, 0)
		bt.root.children[2].checkElements(t, 5, 6, 7)
		bt.root.children[2].checkChildrenLength(t, 0)
	})

	t.Run("insert 8", func(t *testing.T) {
		bt.Insert(8)

		// top level
		bt.root.checkElements(t, 2, 4, 6)
		bt.root.checkChildrenLength(t, 4)

		// second level left to right
		bt.root.children[0].checkElements(t, 1)
		bt.root.children[0].checkChildrenLength(t, 0)
		bt.root.children[1].checkElements(t, 3)
		bt.root.children[1].checkChildrenLength(t, 0)
		bt.root.children[2].checkElements(t, 5)
		bt.root.children[2].checkChildrenLength(t, 0)
		bt.root.children[3].checkElements(t, 7, 8)
		bt.root.children[3].checkChildrenLength(t, 0)
	})

	t.Run("insert 9", func(t *testing.T) {
		bt.Insert(9)

		// top level
		bt.root.checkElements(t, 2, 4, 6)
		bt.root.checkChildrenLength(t, 4)

		// second level left to right
		bt.root.children[0].checkElements(t, 1)
		bt.root.children[0].checkChildrenLength(t, 0)
		bt.root.children[1].checkElements(t, 3)
		bt.root.children[1].checkChildrenLength(t, 0)
		bt.root.children[2].checkElements(t, 5)
		bt.root.children[2].checkChildrenLength(t, 0)
		bt.root.children[3].checkElements(t, 7, 8, 9)
		bt.root.children[3].checkChildrenLength(t, 0)
	})

	t.Run("insert 10", func(t *testing.T) {
		bt.Insert(10)

		// top level
		bt.root.checkElements(t, 4)
		bt.root.checkChildrenLength(t, 2)

		// second level left to right
		bt.root.children[0].checkElements(t, 2)
		bt.root.children[0].checkChildrenLength(t, 2)
		bt.root.children[1].checkElements(t, 6, 8)
		bt.root.children[1].checkChildrenLength(t, 3)

		// third level left to right
		bt.root.children[0].children[0].checkElements(t, 1)
		bt.root.children[0].children[0].checkChildrenLength(t, 0)
		bt.root.children[0].children[1].checkElements(t, 3)
		bt.root.children[0].children[1].checkChildrenLength(t, 0)

		bt.root.children[1].children[0].checkElements(t, 5)
		bt.root.children[1].children[0].checkChildrenLength(t, 0)
		bt.root.children[1].children[1].checkElements(t, 7)
		bt.root.children[1].children[1].checkChildrenLength(t, 0)
		bt.root.children[1].children[2].checkElements(t, 9, 10)
		bt.root.children[1].children[2].checkChildrenLength(t, 0)
	})
}

func TestInsertDegree5(t *testing.T) {
	bt, _ := New(5)

	t.Run("insert 1", func(t *testing.T) {
		bt.Insert(1)
		bt.root.checkElements(t, 1)
		bt.root.checkChildrenLength(t, 0)
	})

	t.Run("insert 2", func(t *testing.T) {
		bt.Insert(2)
		bt.root.checkElements(t, 1, 2)
		bt.root.checkChildrenLength(t, 0)
	})

	t.Run("insert 3", func(t *testing.T) {
		bt.Insert(3)
		bt.root.checkElements(t, 1, 2, 3)
		bt.root.checkChildrenLength(t, 0)
	})

	t.Run("insert 4", func(t *testing.T) {
		bt.Insert(4)
		bt.root.checkElements(t, 1, 2, 3, 4)
		bt.root.checkChildrenLength(t, 0)
	})

	t.Run("insert 5", func(t *testing.T) {
		bt.Insert(5)
		bt.root.checkElements(t, 3)
		bt.root.checkChildrenLength(t, 2)
		bt.root.children[0].checkElements(t, 1, 2)
		bt.root.children[0].checkChildrenLength(t, 0)
		bt.root.children[1].checkElements(t, 4, 5)
		bt.root.children[1].checkChildrenLength(t, 0)
	})

	t.Run("insert 6", func(t *testing.T) {
		bt.Insert(6)
		bt.root.checkElements(t, 3)
		bt.root.checkChildrenLength(t, 2)
		bt.root.children[0].checkElements(t, 1, 2)
		bt.root.children[0].checkChildrenLength(t, 0)
		bt.root.children[1].checkElements(t, 4, 5, 6)
		bt.root.children[1].checkChildrenLength(t, 0)
	})

	t.Run("insert 7", func(t *testing.T) {
		bt.Insert(7)
		bt.root.checkElements(t, 3)
		bt.root.checkChildrenLength(t, 2)
		bt.root.children[0].checkElements(t, 1, 2)
		bt.root.children[0].checkChildrenLength(t, 0)
		bt.root.children[1].checkElements(t, 4, 5, 6, 7)
		bt.root.children[1].checkChildrenLength(t, 0)
	})

	t.Run("insert 8", func(t *testing.T) {
		bt.Insert(8)
		bt.root.checkElements(t, 3, 6)
		bt.root.checkChildrenLength(t, 3)
		bt.root.children[0].checkElements(t, 1, 2)
		bt.root.children[0].checkChildrenLength(t, 0)
		bt.root.children[1].checkElements(t, 4, 5)
		bt.root.children[1].checkChildrenLength(t, 0)
		bt.root.children[2].checkElements(t, 7, 8)
		bt.root.children[2].checkChildrenLength(t, 0)
	})

	t.Run("insert 9", func(t *testing.T) {
		bt.Insert(9)
		bt.root.checkElements(t, 3, 6)
		bt.root.checkChildrenLength(t, 3)
		bt.root.children[0].checkElements(t, 1, 2)
		bt.root.children[0].checkChildrenLength(t, 0)
		bt.root.children[1].checkElements(t, 4, 5)
		bt.root.children[1].checkChildrenLength(t, 0)
		bt.root.children[2].checkElements(t, 7, 8, 9)
		bt.root.children[2].checkChildrenLength(t, 0)
	})

	t.Run("insert 10", func(t *testing.T) {
		bt.Insert(10)
		bt.root.checkElements(t, 3, 6)
		bt.root.checkChildrenLength(t, 3)
		bt.root.children[0].checkElements(t, 1, 2)
		bt.root.children[0].checkChildrenLength(t, 0)
		bt.root.children[1].checkElements(t, 4, 5)
		bt.root.children[1].checkChildrenLength(t, 0)
		bt.root.children[2].checkElements(t, 7, 8, 9, 10)
		bt.root.children[2].checkChildrenLength(t, 0)
	})

	t.Run("insert 11", func(t *testing.T) {
		bt.Insert(11)
		bt.root.checkElements(t, 3, 6, 9)
		bt.root.checkChildrenLength(t, 4)
		bt.root.children[0].checkElements(t, 1, 2)
		bt.root.children[0].checkChildrenLength(t, 0)
		bt.root.children[1].checkElements(t, 4, 5)
		bt.root.children[1].checkChildrenLength(t, 0)
		bt.root.children[2].checkElements(t, 7, 8)
		bt.root.children[2].checkChildrenLength(t, 0)
		bt.root.children[3].checkElements(t, 10, 11)
		bt.root.children[3].checkChildrenLength(t, 0)
	})

	t.Run("insert 12", func(t *testing.T) {
		bt.Insert(12)
		bt.root.checkElements(t, 3, 6, 9)
		bt.root.checkChildrenLength(t, 4)
		bt.root.children[0].checkElements(t, 1, 2)
		bt.root.children[0].checkChildrenLength(t, 0)
		bt.root.children[1].checkElements(t, 4, 5)
		bt.root.children[1].checkChildrenLength(t, 0)
		bt.root.children[2].checkElements(t, 7, 8)
		bt.root.children[2].checkChildrenLength(t, 0)
		bt.root.children[3].checkElements(t, 10, 11, 12)
		bt.root.children[3].checkChildrenLength(t, 0)
	})

	t.Run("insert 13", func(t *testing.T) {
		bt.Insert(13)
		bt.root.checkElements(t, 3, 6, 9)
		bt.root.checkChildrenLength(t, 4)
		bt.root.children[0].checkElements(t, 1, 2)
		bt.root.children[0].checkChildrenLength(t, 0)
		bt.root.children[1].checkElements(t, 4, 5)
		bt.root.children[1].checkChildrenLength(t, 0)
		bt.root.children[2].checkElements(t, 7, 8)
		bt.root.children[2].checkChildrenLength(t, 0)
		bt.root.children[3].checkElements(t, 10, 11, 12, 13)
		bt.root.children[3].checkChildrenLength(t, 0)
	})

	t.Run("insert 14", func(t *testing.T) {
		bt.Insert(14)
		bt.root.checkElements(t, 3, 6, 9, 12)
		bt.root.checkChildrenLength(t, 5)
		bt.root.children[0].checkElements(t, 1, 2)
		bt.root.children[0].checkChildrenLength(t, 0)
		bt.root.children[1].checkElements(t, 4, 5)
		bt.root.children[1].checkChildrenLength(t, 0)
		bt.root.children[2].checkElements(t, 7, 8)
		bt.root.children[2].checkChildrenLength(t, 0)
		bt.root.children[3].checkElements(t, 10, 11)
		bt.root.children[3].checkChildrenLength(t, 0)
		bt.root.children[4].checkElements(t, 13, 14)
		bt.root.children[4].checkChildrenLength(t, 0)
	})

	t.Run("insert 15", func(t *testing.T) {
		bt.Insert(15)
		bt.root.checkElements(t, 3, 6, 9, 12)
		bt.root.checkChildrenLength(t, 5)
		bt.root.children[0].checkElements(t, 1, 2)
		bt.root.children[0].checkChildrenLength(t, 0)
		bt.root.children[1].checkElements(t, 4, 5)
		bt.root.children[1].checkChildrenLength(t, 0)
		bt.root.children[2].checkElements(t, 7, 8)
		bt.root.children[2].checkChildrenLength(t, 0)
		bt.root.children[3].checkElements(t, 10, 11)
		bt.root.children[3].checkChildrenLength(t, 0)
		bt.root.children[4].checkElements(t, 13, 14, 15)
		bt.root.children[4].checkChildrenLength(t, 0)
	})

	t.Run("insert 16", func(t *testing.T) {
		bt.Insert(16)
		bt.root.checkElements(t, 3, 6, 9, 12)
		bt.root.checkChildrenLength(t, 5)
		bt.root.children[0].checkElements(t, 1, 2)
		bt.root.children[0].checkChildrenLength(t, 0)
		bt.root.children[1].checkElements(t, 4, 5)
		bt.root.children[1].checkChildrenLength(t, 0)
		bt.root.children[2].checkElements(t, 7, 8)
		bt.root.children[2].checkChildrenLength(t, 0)
		bt.root.children[3].checkElements(t, 10, 11)
		bt.root.children[3].checkChildrenLength(t, 0)
		bt.root.children[4].checkElements(t, 13, 14, 15, 16)
		bt.root.children[4].checkChildrenLength(t, 0)
	})

	t.Run("insert 17", func(t *testing.T) {
		bt.Insert(17)

		// top level
		bt.root.checkElements(t, 9)
		bt.root.checkChildrenLength(t, 2)

		// second level left to right
		bt.root.children[0].checkElements(t, 3, 6)
		bt.root.children[0].checkChildrenLength(t, 3)

		bt.root.children[1].checkElements(t, 12, 15)
		bt.root.children[1].checkChildrenLength(t, 3)

		// third level left to right
		bt.root.children[0].children[0].checkElements(t, 1, 2)
		bt.root.children[0].children[0].checkChildrenLength(t, 0)
		bt.root.children[0].children[1].checkElements(t, 4, 5)
		bt.root.children[0].children[1].checkChildrenLength(t, 0)
		bt.root.children[0].children[2].checkElements(t, 7, 8)
		bt.root.children[0].children[2].checkChildrenLength(t, 0)

		bt.root.children[1].children[0].checkElements(t, 10, 11)
		bt.root.children[1].children[0].checkChildrenLength(t, 0)
		bt.root.children[1].children[1].checkElements(t, 13, 14)
		bt.root.children[1].children[1].checkChildrenLength(t, 0)
		bt.root.children[1].children[2].checkElements(t, 16, 17)
		bt.root.children[1].children[2].checkChildrenLength(t, 0)
	})
}

func TestInsertDuplicates(t *testing.T) {
	bt, _ := New(3)
	bt.Insert(1)
	bt.Insert(1)
	bt.Insert(1)
	bt.Insert(1)
	bt.root.checkElements(t, 1)
	bt.root.children[0].checkElements(t, 1)
	bt.root.children[1].checkElements(t, 1, 1)
}

// checkElements asserts a node's elements match exactly the values for elements.
// order does matter.
func (n *node) checkElements(t *testing.T, elements ...int) {
	if len(n.elements) != len(elements) {
		t.Errorf(
			"Got node with %v elements, but want node with %v elements",
			len(n.elements),
			len(elements),
		)
	}
	for i, e := range elements {
		if n.elements[i] != e {
			t.Errorf("Invalid match %v with %v", n.elements[i], e)
		}
	}
}

func (n *node) checkChildrenLength(t *testing.T, expectedLength int) {
	if len(n.children) != expectedLength {
		t.Errorf(
			"Got %v children, but expected %v children",
			len(n.children),
			expectedLength,
		)
	}
}
