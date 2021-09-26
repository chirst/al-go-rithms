// Package btree implements a btree that shouldn't be taken too seriously.
package btree

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

// New returns a tree with the given degree. When a node reaches the given
// degree of elements the node will split.
//
// This implementation allows degrees between 3-7 inclusive.
func New(degree int) (*btree, error) {
	if degree < 3 {
		return nil, errors.New("tree must not have degree less than 3")
	}
	if 7 < degree {
		return nil, errors.New("tree must not have degree greater than 7")
	}
	return &btree{
		degree: degree,
	}, nil
}

// TODO: implement Exists and document
func (bt *btree) Exists(value int) bool {
	return false
}

// TODO: insert duplicate value and document behavior. Likely will skip if the
// value already exists.
// TODO: consider making Insert functions smaller.
// TODO: complexity

// Insert inserts an element into the tree.
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
	// Base case, done splitting.
	if len(n.elements) < bt.degree {
		return
	}

	// `n` is the root node. This means the tree will grow in height.
	if n.parent == nil {
		// Gather left, right, and middle partitions of the current node.
		middleIndex := (len(n.elements) - 1) / 2
		value := n.elements[middleIndex]
		lefts := n.elements[:middleIndex]
		rights := n.elements[middleIndex+1:]
		leftChildren := []*node{}
		rightChildren := []*node{}
		if len(n.children) != 0 {
			leftChildren = n.children[:middleIndex+1]
			rightChildren = n.children[middleIndex+1:]
		}

		// Create a new root node from the middle partition (the middle
		// partition always being size 1) making the root node's two children
		// from the left and right partition.
		bt.root = &node{
			elements: []int{value},
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
		n.parent = bt.root // TODO: remove?
		return
	} else {
		// Gather left, right, and middle partitions of the current node.
		middleIndex := (len(n.elements) - 1) / 2
		middleElement := n.elements[middleIndex]
		leftElements := n.elements[:middleIndex]
		rightElements := n.elements[middleIndex+1:]
		leftChildren := []*node{}
		rightChildren := []*node{}
		if len(n.children) != 0 {
			rightChildren = n.children[middleIndex+1:]
			leftChildren = n.children[:middleIndex-1]
		}

		// Remove old parent/child relation.
		for i, element := range n.parent.elements {
			// End of list, remove last child.
			if element < leftElements[0] && i+1 == len(n.parent.elements) {
				n.parent.children = n.parent.children[:len(n.parent.children)-1]
				break
			}
			// Gone past by 1.
			if element > leftElements[0] {
				n.parent.children = append(n.parent.children[:i], n.parent.children[i+1:]...)
				break
			}
		}

		// Insert middle into parent and put lefts and rights as children.
		for i, element := range n.parent.elements {
			// End of list, append to the end.
			if element < middleElement && i+1 == len(n.parent.elements) {
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
			// Gone past by 1.
			if element > middleElement {
				// Insert middle into parent elements.
				leftParentElements := n.parent.elements[:i-1]
				leftParentElements = append(leftParentElements, middleElement)
				rightParentElements := n.parent.elements[i:]
				n.parent.elements = append(leftParentElements, rightParentElements...)

				// Insert lefts and rights into parent children.
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
		}
		// Recursively continue to split.
		bt.split(n.parent)
	}
}
