// Package btree implements a btree that shouldn't be taken too seriously.
package btree

// TODO:
// - implement Delete
// - make comments less bad

import (
	"errors"
)

type btree struct {
	root   *node
	degree int
}

// node consists of one or many elements and many children and always maintains
// elements + 1 number of children.
type node struct {
	parent   *node
	elements []int
	children []*node
}

// New returns a tree with the given degree.
//
// When a node reaches the given degree of elements the node will split.
//
// This implementation allows degrees between 3-7 inclusive.
//
// Given `values` are provided the tree will be populated with the `values`. The
// complexity of this is O(n log n).
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

// Exists checks for the existense of the given `value`.
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
	// If the node is a leaf the value does not exist
	if len(n.children) == 0 {
		return false
	}
	// Recursively search next node
	var childNode *node
	for i, el := range n.elements {
		if el > value {
			childNode = n.children[i]
			break
		}
		if i == len(n.elements)-1 {
			childNode = n.children[i+1]
			break
		}
	}
	return bt.exists(childNode, value)
}

// Insert inserts an element into the tree.
//
// Given the value already exists in the tree, the value will still be inserted.
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
	// node is not a leaf, `insert` will recursively find a leaf node to insert
	// into.
	bt.insert(bt.root, value)
}

// insert recursively follows nodes until hitting a leaf node. Once a leaf is
// hit, an insert will be performed (whether or not the insert is allowed)
// followed by splitting to ensure the tree is in a valid state.
func (bt *btree) insert(n *node, value int) {
	// Insert if leaf.
	if len(n.children) == 0 {
		n.elements = append(n.elements, value)
		// Start splitting if needed.
		bt.split(n)
	} else {
		// Recursively try to insert on the next sub tree.
		for i, el := range n.elements {
			if el > value {
				bt.insert(n.children[i], value)
				break
			}
			if i == len(n.elements)-1 {
				bt.insert(n.children[i+1], value)
				break
			}
		}
	}
}

// split splits a node exceeding a tree's degree. split continues to recursively
// process parent nodes until all parent nodes have valid degrees.
func (bt *btree) split(n *node) {
	// Base case.
	// Done splitting, the current node and all parent nodes have valid degrees.
	if len(n.elements) < bt.degree {
		return
	}

	// Base case #2.
	// `n` is the root node and needs to be split.
	if n.parent == nil {
		bt.splitRoot(n)
		return
	}

	// `n` is an internal node that needs splitting.
	bt.splitInternal(n)

	// Recursively continue to split.
	bt.split(n.parent)
}

// splitRoot takes an invalid degree root node and creates a new root node with
// the previously invalid root node evenly split as the new root nodes children.
// This procedure is what grows the tree in height.
func (bt *btree) splitRoot(n *node) {
	middle, lefts, rights := getPartitionedElementsOf(n)
	leftChildren, rightChildren := getPartitionedChildrenOf(n)
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

// splitInternal takes an internal node `n` and inserts `n`'s middle node into
// the parent. The partitions to the left and right of `n`'s middle node then
// become children of `n`'s parent.
//
// An internal node is a node who has a parent and children.
func (bt *btree) splitInternal(n *node) {
	middleElement, leftElements, rightElements := getPartitionedElementsOf(n)
	leftChildren, rightChildren := getPartitionedChildrenOf(n)

	removeChildFromParent(n, leftElements)

	// Insert middle into parent and put lefts and rights as children.
	for i, element := range n.parent.elements {
		// End of list, append to the end.
		if element < middleElement && i+1 == len(n.parent.elements) {
			appendSplitInternal(
				n,
				middleElement,
				leftElements,
				leftChildren,
				rightElements,
				rightChildren,
			)
		}
		// Gone past by 1 insert in order.
		if element > middleElement {
			insertSplitInternal(
				n,
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

// removeChildFromParent takes a node `n` and removes the relation to `n` in
// `n`'s parent.
func removeChildFromParent(n *node, leftElements []int) {
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

// appendSplitInternal appends a split internal node `n` to the end of `n`'s
// parent `p` by appending the middle element to `p`'s elements and appending
// new children to the end of `p`'s children.
func appendSplitInternal(
	n *node,
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

// insertSplitInternal inserts a split internal node `n` made up of middle,
// lefts, and rights in order with `n`'s parent's elements and children.
func insertSplitInternal(
	n *node,
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

// getPartitionedElementsOf splits and returns the middle, left, and right
// elements of the given node.
func getPartitionedElementsOf(n *node) (int, []int, []int) {
	middleIndex := (len(n.elements) - 1) / 2
	middle := n.elements[middleIndex]
	lefts := n.elements[:middleIndex]
	rights := n.elements[middleIndex+1:]
	return middle, lefts, rights
}

// getPartitionedChildrenOf splits and returns the children of the given node
// into left and right partitions.
func getPartitionedChildrenOf(n *node) ([]*node, []*node) {
	middleIndex := (len(n.elements) - 1) / 2
	lefts := []*node{}
	rights := []*node{}
	if len(n.children) != 0 {
		lefts = n.children[:middleIndex+1]
		rights = n.children[middleIndex+1:]
	}
	return lefts, rights
}
