// TODO package level comment
package btree

import (
	"errors"
)

type btree struct {
	root   *node
	degree int
}

type node struct {
	parent   *node
	elements []int
	children []*node
}

// New returns a tree with the given degree. When a node reaches the given
// degree of elements the node will decide to split.
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

// Insert inserts an element into the tree
func (bt *btree) Insert(value int) {
	// No node at all so create a new one.
	if bt.root == nil {
		bt.root = &node{
			elements: []int{value},
		}
		return
	}
	bt.insert(bt.root, value)
}

func (bt *btree) insert(n *node, value int) {
	// Insert at leaf.
	if len(n.children) == 0 {
		n.elements = append(n.elements, value)
		// Start splitting if needed.
		bt.split(n)
	} else {
		// Try insert on sub tree.
		for i, el := range n.elements {
			if el > value {
				bt.insert(n.children[i], value)
				break
			}
			if i == len(n.elements)-1 {
				// Edge case because there is n+1 children for n elements.
				bt.insert(n.children[i+1], value)
				break
			}
		}
	}
}

func (bt *btree) split(n *node) {
	// Base case, done splitting.
	if len(n.elements) < bt.degree {
		return
	}

	// `n` is the root node.
	if n.parent == nil {
		middleIndex := len(n.elements) / 2
		value := n.elements[middleIndex]
		lefts := n.elements[:middleIndex]
		rights := n.elements[middleIndex+1:]

		bt.root = &node{
			elements: []int{value},
		}
		for _, l := range lefts {
			bt.root.children = append(bt.root.children, &node{
				parent:   bt.root,
				elements: []int{l},
			})
		}
		for _, r := range rights {
			bt.root.children = append(bt.root.children, &node{
				parent:   bt.root,
				elements: []int{r},
			})
		}
		n.parent = bt.root
		return
	} else {
		// Split child and gather pieces.
		middleIndex := (len(n.elements) - 1) / 2
		middleElement := n.elements[middleIndex]
		leftElements := n.elements[:middleIndex]
		leftChildren := n.children[:middleIndex-1]
		rightElements := n.elements[middleIndex+1:]
		rightChildren := []*node{}
		children := n.children
		if len(children) != 0 {
			rightChildren = n.children[middleIndex+1:]
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
	}
}
