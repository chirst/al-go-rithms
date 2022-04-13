// Package btree implements a btree that shouldn't be taken too seriously.
package btree

// TODO:
// - Implement Delete
// - Test coverage

import (
	"errors"
)

// btree represents a single btree data structure made up of nodes.
type btree struct {
	// root is the entry node of the tree.
	root *node
	// degree is the maximum amount of elements a node in the btree can contain.
	// When the maximum is exceeded the node will perform a split operation.
	degree int
}

// node makes up a btree. There are three different kinds of nodes in this tree:
// Root, a node with no parent; Internal, a node with a parent and children;
// Leaf, a node with a parent and no children.
type node struct {
	// parent is nil when the node is the root of the tree.
	parent *node
	// elements are ordered from least to greatest. elements do not exceed the
	// degree of their associated btree.
	elements []int
	// A node maintains elements + 1 children at all times.
	children []*node
}

// New returns a tree with the given degree.
//
// When a node reaches the given degree of elements the node will split.
//
// This implementation allows degrees between 3-7 inclusive.
//
// Given values are provided, the tree will be populated with the values. The
// complexity of inserting these values is O(n log n).
func New(degree int, values ...int) (*btree, error) {
	if degree < 3 {
		return nil, errors.New("tree must not have degree less than 3")
	}
	if 7 < degree {
		return nil, errors.New("tree must not have degree greater than 7")
	}
	nt := &btree{
		degree: degree,
	}
	for _, v := range values {
		nt.Insert(v)
	}
	return nt, nil
}

// Exists checks for the existence of the given value.
//
// The complexity is O(log n).
func (bt *btree) Exists(value int) bool {
	if bt.root == nil {
		return false
	}
	return bt.exists(bt.root, value)
}

func (bt *btree) exists(n *node, value int) bool {
	// Check if value is in the current node.
	for _, e := range n.elements {
		if e == value {
			return true
		}
	}
	// If the node is a leaf the value does not exist.
	if len(n.children) == 0 {
		return false
	}
	// Existence is still unknown, determine what child node to search next,
	// then recursively search the next child node.
	childNode := n.getChildContaining(value)
	return bt.exists(childNode, value)
}

// Insert inserts an element into the tree.
//
// Given the value already exists in the tree, the value will still be inserted
// as a duplicate.
//
// The complexity is O(log n).
func (bt *btree) Insert(value int) {
	// No nodes at all so create a root node.
	if bt.root == nil {
		bt.root = &node{
			elements: []int{value},
		}
		return
	}
	// Root node exists, attempt to insert into the root node. In case the root
	// node is not a leaf, insert will recursively find a leaf node to insert
	// into.
	bt.insert(bt.root, value)
}

// insert recursively follows nodes until hitting a leaf node. Once a leaf is
// hit, an insert will be performed (whether or not the insert is allowed based
// off of the btree's degree). Given the insert leaves the node in an invalid
// state, splitting is performed to make the tree valid again.
func (bt *btree) insert(n *node, value int) {
	// Insert if leaf.
	if len(n.children) == 0 {
		n.addElement(value)
		// Start splitting if needed.
		bt.split(n)
	} else {
		// Recursively try to insert on the next sub tree.
		childNode := n.getChildContaining(value)
		bt.insert(childNode, value)
	}
}

// split splits a node exceeding a tree's degree. split continues to recursively
// process parent nodes until all parent nodes have valid degrees.
func (bt *btree) split(n *node) {
	// Done splitting, the current node is a valid degree.
	if len(n.elements) < bt.degree {
		return
	}

	// The current node is the root node and needs to be split.
	if n.parent == nil {
		bt.splitRoot()
		return
	}

	// The current node is an internal node that needs splitting.
	bt.splitInternal(n)

	// Recursively continue to split.
	bt.split(n.parent)
}

// splitRoot creates a new root node with the middle element of the root node as
// the new root node. The left and right partitions become the new root nodes
// children.
// This procedure is what grows the tree in height.
func (bt *btree) splitRoot() {
	middle, lefts, rights := bt.root.getPartitionedElements()
	leftChildren, rightChildren := bt.root.getPartitionedChildren()
	bt.root = &node{
		elements: []int{middle},
	}
	bt.root.children = append(bt.root.children, &node{
		parent:   bt.root,
		elements: lefts,
		children: leftChildren,
	})
	bt.root.children = append(bt.root.children, &node{
		parent:   bt.root,
		elements: rights,
		children: rightChildren,
	})
	return
}

// splitInternal takes an internal node and inserts it's middle element into the
// parent. The partitions to the left and right of the middle element then
// become children of the parent.
func (bt *btree) splitInternal(n *node) {
	middleElement, leftElements, rightElements := n.getPartitionedElements()
	leftChildren, rightChildren := n.getPartitionedChildren()

	n.removeChildFromParent(leftElements)

	// Insert middle into parent and put lefts and rights as children.
	for i, element := range n.parent.elements {
		// End of list, append to the end.
		if element < middleElement && i+1 == len(n.parent.elements) {
			n.appendSplitInternal(
				middleElement,
				leftElements,
				leftChildren,
				rightElements,
				rightChildren,
			)
		}
		// Gone past by 1 insert in order.
		if element > middleElement {
			n.insertSplitInternal(
				i,
				middleElement,
				leftElements,
				leftChildren,
				rightElements,
				rightChildren,
			)
		}
	}
}

// addElement adds an element to a leaf node while maintaining ordering of the
// elements.
func (n *node) addElement(value int) {
	for i, e := range n.elements {
		if value < e {
			n.elements = append(n.elements[:i+1], n.elements[i:]...)
			n.elements[i] = value
			return
		}
	}
	n.elements = append(n.elements, value)
}

// getChildContaining returns the child node potentially containing the given
// value.
func (n *node) getChildContaining(value int) *node {
	for i, el := range n.elements {
		if el > value {
			return n.children[i]
		}
	}
	return n.children[len(n.children)-1]
}

// getPartitionedElements splits and returns the middle, left, and right
// elements of the given node.
func (n *node) getPartitionedElements() (int, []int, []int) {
	middleIndex := (len(n.elements) - 1) / 2
	middle := n.elements[middleIndex]
	lefts := n.elements[:middleIndex]
	rights := n.elements[middleIndex+1:]
	return middle, lefts, rights
}

// getPartitionedChildren splits and returns the children of the given node
// into left and right partitions.
func (n *node) getPartitionedChildren() ([]*node, []*node) {
	middleIndex := (len(n.elements) - 1) / 2
	lefts := []*node{}
	rights := []*node{}
	if len(n.children) != 0 {
		lefts = n.children[:middleIndex+1]
		rights = n.children[middleIndex+1:]
	}
	return lefts, rights
}

// removeChildFromParent removes the relation between the node and it's parent.
func (n *node) removeChildFromParent(leftElements []int) {
	for i, element := range n.parent.elements {
		// End of list, remove last child.
		if i+1 == len(n.parent.elements) {
			n.parent.children = n.parent.children[:len(n.parent.children)-1]
			break
		}
		// Gone past by 1.
		if element > leftElements[0] {
			n.parent.children = append(n.parent.children[:i], n.parent.children[i+1:]...)
			break
		}
	}
}

// appendSplitInternal appends a split internal node to the end of the nodes
// parent by appending the split partitions to the nodes parent.
func (n *node) appendSplitInternal(
	middleElement int,
	leftElements []int,
	leftChildren []*node,
	rightElements []int,
	rightChildren []*node,
) {
	n.parent.elements = append(n.parent.elements, middleElement)
	newLeft := &node{
		parent:   n.parent,
		elements: leftElements,
		children: leftChildren,
	}
	n.parent.children = append(n.parent.children, newLeft)
	newRight := &node{
		parent:   n.parent,
		elements: rightElements,
		children: rightChildren,
	}
	n.parent.children = append(n.parent.children, newRight)
}

// insertSplitInternal inserts a split internal node made up of middle,
// lefts, and rights in order with the parent node elements and children.
func (n *node) insertSplitInternal(
	i int,
	middleElement int,
	leftElements []int,
	leftChildren []*node,
	rightElements []int,
	rightChildren []*node,
) {
	leftParentElements := n.parent.elements[:i-1]
	leftParentElements = append(leftParentElements, middleElement)
	rightParentElements := n.parent.elements[i:]
	n.parent.elements = append(leftParentElements, rightParentElements...)

	leftParentChildren := n.parent.children[:i-1]
	newLeft := &node{
		parent:   n.parent,
		elements: leftElements,
		children: leftChildren,
	}
	leftParentChildren = append(leftParentChildren, newLeft)
	newRight := &node{
		parent:   n.parent,
		elements: rightElements,
		children: rightChildren,
	}
	leftParentChildren = append(leftParentChildren, newRight)
	rightParentChildren := n.parent.children[i:]
	n.parent.children = append(leftParentChildren, rightParentChildren...)
}
