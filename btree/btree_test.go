package btree

import "testing"

func TestInsert(t *testing.T) {
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

		// check top level (root)
		bt.root.checkElements(t, 4)
		bt.root.checkChildrenLength(t, 2)

		// check second level
		bt.root.children[0].checkElements(t, 2)
		bt.root.children[0].checkChildrenLength(t, 2)
		bt.root.children[0].children[0].checkElements(t, 1)
		bt.root.children[0].children[0].checkChildrenLength(t, 0)
		bt.root.children[0].children[1].checkElements(t, 3)
		bt.root.children[0].children[1].checkChildrenLength(t, 0)

		// check third level (leafs)
		bt.root.children[1].checkElements(t, 6)
		bt.root.children[1].checkChildrenLength(t, 2)
		bt.root.children[1].children[0].checkElements(t, 5)
		bt.root.children[1].children[0].checkChildrenLength(t, 0)
		bt.root.children[1].children[1].checkElements(t, 7)
		bt.root.children[1].children[1].checkChildrenLength(t, 0)
	})
}

// checkElements asserts nodes elements matches exactly the values for elements.
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
