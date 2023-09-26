package avltree

import (
	"golang.org/x/exp/constraints"
)

// node is the structure of tree's node.
// node's key is any ordered type for type of
// node's value has type any
type node[V constraints.Ordered] struct {
	element element[V]
	parent  *node[V]
	left    *node[V]
	right   *node[V]
	height  int
}
type element[V constraints.Ordered] struct {
	key   V
	value any
}

// new - internal function for creating empty node
func new[V constraints.Ordered](key V, value any) *node[V] {
	return &node[V]{element: element[V]{
		key:   key,
		value: value,
	},
		left:   nil,
		right:  nil,
		parent: nil,
		height: 0,
	}
}

// swap - internal function for swapping key and value between two nodes
func swap[V constraints.Ordered](a, b *node[V]) {
	ak := a.element.key
	a.element.key = b.element.key
	b.element.key = ak
	av := a.element.value
	a.element.value = b.element.value
	b.element.value = av
}

// height - internal function for getting height of node
func height[V constraints.Ordered](n *node[V]) int {
	if n == nil {
		return -1
	}

	return n.height
}

// leftRotate - internal function for left rotating in tree
func leftRotate[V constraints.Ordered](n *node[V]) {
	if n == nil || n.right == nil {
		return
	}

	swap(n, n.right)
	buf := n.left
	n.left = n.right
	n.right = n.left.right
	n.left.right = n.left.left
	n.left.left = buf

	balance(n.left)
	balance(n)
}

// rightRotate - internal function for right rotating in tree
func rightRotate[V constraints.Ordered](n *node[V]) {
	if n == nil || n.left == nil {
		return
	}

	swap(n, n.left)
	buf := n.right
	n.right = n.left
	n.left = n.right.left
	n.right.left = n.right.right
	n.right.right = buf

	balance(n.right)
	balance(n)
}

// balanceFactor - internal function for getting balance factor
// this function can return such values: -2, -1, 1, 0, 1, 2
func balanceFactor[V constraints.Ordered](n *node[V]) int {
	if n == nil {
		return 0
	}

	return height(n.right) - height(n.left)
}

// balance - internal function for balancing avl tree
// - param n - starting node for balancing (from n to tree's root)
func balance[V constraints.Ordered](n *node[V]) {
	if n == nil {
		return
	}

	for n != nil {
		n.height = maxValue(height(n.left), height(n.right)) + 1
		fixBalance(n)
		n = n.parent
	}
}

// fixBalance - additional internal function for function balance: in this function we check node's balance
//
//	and do some rotations if it needs
//
// - param n - node which we want to check
func fixBalance[V constraints.Ordered](n *node[V]) {
	b := balanceFactor(n)

	if b == -2 {
		if balanceFactor(n.left) == 1 {
			leftRotate(n.left)
		}
		rightRotate(n)
	}

	if b == 2 {
		if balanceFactor(n.right) == -1 {
			rightRotate(n.right)
		}
		leftRotate(n)
	}
}

// search - internal function for searching node by key in tree
func search[V constraints.Ordered](n *node[V], key V) *node[V] {
	for n != nil && key != n.element.key {
		if key < n.element.key {
			n = n.left
			continue
		}
		n = n.right
	}

	return n
}

// min - internal function for searching min node in tree
func min[V constraints.Ordered](n *node[V]) *node[V] {
	if n == nil {
		return nil
	}

	for n.left != nil {
		n = n.left
	}

	return n
}

// isRoot - internal function for checking node is root or not
func isRoot[V constraints.Ordered](n *node[V]) bool {
	return n.parent == nil
}

// isLeftChild - internal function for checking node is left child of her parent
func isLeftChild[V constraints.Ordered](n *node[V]) bool {
	return n == n.parent.left
}

// maxValue - little internal function: get max value between two int items
func maxValue(i, j int) int {
	if i > j {
		return i
	}

	return j
}
