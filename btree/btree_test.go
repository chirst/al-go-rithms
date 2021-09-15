package btree

import "testing"

func TestInsert(t *testing.T) {
	bt, _ := New(3)

	bt.Insert(1)
	if bt.root == nil {
		t.Error("expected root not to be nil")
	}
	if len(bt.root.elements) != 1 {
		t.Error("expected root node to have 1 elements")
	}

	bt.Insert(2)
	if len(bt.root.elements) != 2 {
		t.Error("expected root node to have 2 elements")
	}

	bt.Insert(3)
	if len(bt.root.elements) == 3 {
		t.Error("expected root node not to have 3 elements")
	}
	if len(bt.root.elements) != 1 {
		t.Error("expected root node to have 1 elements")
	}
	if len(bt.root.children) != 2 {
		t.Error("expected root node to have 2 children")
	}

	bt.Insert(4)
	if len(bt.root.elements) == 3 {
		t.Error("expected root node not to have 3 elements")
	}
	if len(bt.root.elements) != 1 {
		t.Error("expected root node to have 1 elements")
	}
	if len(bt.root.children) != 2 {
		t.Error("expected root node to have 2 children")
	}
	if len(bt.root.children[1].elements) != 2 {
		// needs more informative about what the actual state of the tree is
		t.Error("expected right child to have 2 elements")
	}

	bt.Insert(5)
	if len(bt.root.elements) != 2 {
		t.Error("expected root to have 2 elements")
	}
	if len(bt.root.children) != 3 {
		t.Error("expected root to have 3 children")
	}
}
