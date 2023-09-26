package avltree

import "golang.org/x/exp/constraints"

type Tree[V constraints.Ordered] struct {
	root *node[V]
}

// New is a function for creation empty tree
// - param should be `ordered type` (`int`, `string`, `float`, etc.)
func New[V constraints.Ordered]() *Tree[V] {
	return &Tree[V]{}
}

// Insert is a function for inserting element into Tree
// - param key should be `ordered type` (`int`, `string`, `float` etc.)
// - param value can be any type
func (t *Tree[V]) Insert(key V, value any) {
	newNode := new[V](key, value)

	if t.root == nil {
		t.root = newNode
		return
	}

	current := t.root
	parentNode := t.root

	for current != nil {
		parentNode = current
		if newNode.element.key < current.element.key {
			current = current.left
			continue
		}
		current = current.right
	}

	newNode.parent = parentNode
	if newNode.element.key < parentNode.element.key {
		parentNode.left = newNode
	} else {
		parentNode.right = newNode
	}

	balance(newNode)
}

// Delete is a function for deleting node in rbtree
// - param key should be `ordered type` (`int`, `string`, `float` etc)
func (t *Tree[V]) Delete(key V) {
	n := search(t.root, key)
	if n == nil {
		return
	}

	if n.left == nil || n.right == nil {
		child := n.left
		if n.right != nil {
			child = n.right
		}
		t.transplant(n, child)
		balance(n)

		return
	}

	m := min(n.right)
	if m.parent != n {
		t.transplant(m, m.right)
		m.right = n.right
		if m.right != nil {
			m.right.parent = m
		}
	}

	t.transplant(n, m)
	m.left = n.left
	if m.left != nil {
		m.left.parent = m
	}

	balance(m)
}

// transplant - internal function for substitution u node to v node
func (t *Tree[V]) transplant(u, v *node[V]) {
	if isRoot(u) {
		t.root = v
	} else if isLeftChild(u) {
		u.parent.left = v
	} else {
		u.parent.right = v
	}

	if v != nil {
		v.parent = u.parent
	}
}
