package avltree

import (
	"reflect"
	"testing"

	"golang.org/x/exp/constraints"
)

type validNode[V constraints.Ordered] struct {
	node     *node[V]
	key      V
	height   int
	nodePath string
}

func TestNew(t *testing.T) {
	type testCase[V constraints.Ordered] struct {
		name string
		want *Tree[V]
	}
	testInt := testCase[int]{
		name: "int empty tree",
		want: &Tree[int]{root: nil},
	}
	t.Run(testInt.name, func(t *testing.T) {
		if got := New[int](); !reflect.DeepEqual(got, testInt.want) {
			t.Errorf("CreateNode() = %v, want %v", got, testInt.want)
		}
	})

	testString := testCase[string]{
		name: "int empty tree",
		want: &Tree[string]{root: nil},
	}
	t.Run(testString.name, func(t *testing.T) {
		if got := New[int](); !reflect.DeepEqual(got, testInt.want) {
			t.Errorf("New() = %v, want %v", got, testInt.want)
		}
	})
}

func TestTree_Insert(t1 *testing.T) {
	type args[V constraints.Ordered] struct {
		key   V
		value any
	}
	type testCase[V constraints.Ordered] struct {
		name string
		t    *Tree[V]
		args args[V]
		want *Tree[V]
	}

	tree := New[int]()

	treeWithOneRightElement := Tree[int]{
		root: &node[int]{
			element: element[int]{
				key:   15,
				value: 15,
			},
			right: &node[int]{
				element: element[int]{
					key:   25,
					value: 25,
				},
			},
			height: 1,
		},
	}
	treeWithOneRightElement.root.right.parent = treeWithOneRightElement.root

	treeWithOneLeftElement := Tree[int]{
		root: &node[int]{
			element: element[int]{
				key:   15,
				value: 15,
			},
			left: &node[int]{
				element: element[int]{
					key:   10,
					value: 10,
				},
			},
			height: 1,
		},
	}
	treeWithOneLeftElement.root.left.parent = treeWithOneLeftElement.root

	tests := []testCase[int]{
		{
			name: "insert root",
			t:    tree,
			args: args[int]{key: 15, value: 15},
			want: &Tree[int]{
				root: &node[int]{
					element: element[int]{
						key:   15,
						value: 15,
					},
				},
			},
		},
		{
			name: "insert right node to root",
			t:    getTree([]int{15}),
			args: args[int]{key: 25, value: 25},
			want: &treeWithOneRightElement,
		},
		{
			name: "insert left node to root",
			t:    getTree([]int{15}),
			args: args[int]{key: 10, value: 10},
			want: &treeWithOneLeftElement,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			tt.t.Insert(tt.args.key, tt.args.value)
			if !reflect.DeepEqual(tt.t, tt.want) {
				t1.Errorf("Insert() = %#+v, want %#+v", tt.t, tt.want)
			}
		})
	}
}

func TestTree_Insert_right_rotate(t1 *testing.T) {
	t := getTree([]int{4, 2})

	validTree := []validNode[int]{
		{node: t.root, key: 4, height: 1, nodePath: "t.root"},
		{node: t.root.left, key: 2, height: 0, nodePath: "t.root.left"},
	}

	// check tree's structure and heights before insert
	for _, n := range validTree {
		checkNode(t1, &n)
	}

	// insert
	t.Insert(1, 1)

	// check tree's structure and heights after insert
	validTreeAfterInsert := []validNode[int]{
		{node: t.root, key: 2, height: 1, nodePath: "t.root"},
		{node: t.root.left, key: 1, height: 0, nodePath: "t.root.left"},
		{node: t.root.right, key: 4, height: 0, nodePath: "t.root.right"},
	}
	for _, n := range validTreeAfterInsert {
		checkNode(t1, &n)
	}
}

func TestTree_Insert_left_rotate(t1 *testing.T) {
	t := getTree([]int{4, 5})

	validTree := []validNode[int]{
		{node: t.root, key: 4, height: 1, nodePath: "t.root"},
		{node: t.root.right, key: 5, height: 0, nodePath: "t.root.left"},
	}

	// check tree's structure and heights before insert
	for _, n := range validTree {
		checkNode(t1, &n)
	}

	// insert
	t.Insert(6, 6)

	// check tree's structure and heights after insert
	validTreeAfterInsert := []validNode[int]{
		{node: t.root, key: 5, height: 1, nodePath: "t.root"},
		{node: t.root.left, key: 4, height: 0, nodePath: "t.root.left"},
		{node: t.root.right, key: 6, height: 0, nodePath: "t.root.right"},
	}
	for _, n := range validTreeAfterInsert {
		checkNode(t1, &n)
	}
}

func TestTree_Insert_right_left_rotate(t1 *testing.T) {
	t := getTree([]int{20, 10, 30, 29, 40})

	validTree := []validNode[int]{
		{node: t.root, key: 20, height: 2, nodePath: "t.root"},
		{node: t.root.left, key: 10, height: 0, nodePath: "t.root.left"},
		{node: t.root.right, key: 30, height: 1, nodePath: "t.root.right"},
		{node: t.root.right.left, key: 29, height: 0, nodePath: "t.root.right.left"},
		{node: t.root.right.right, key: 40, height: 0, nodePath: "t.root.right.left"},
	}

	// check tree's structure and heights before insert
	for _, n := range validTree {
		checkNode(t1, &n)
	}

	// insert
	t.Insert(28, 28)

	// check tree's structure and heights after insert
	validTreeAfterInsert := []validNode[int]{
		{node: t.root, key: 29, height: 2, nodePath: "t.root"},
		{node: t.root.left, key: 20, height: 1, nodePath: "t.root.left"},
		{node: t.root.right, key: 30, height: 1, nodePath: "t.root.right"},
		{node: t.root.left.left, key: 10, height: 0, nodePath: "t.root.left.left"},
		{node: t.root.left.right, key: 28, height: 0, nodePath: "t.root.left.right"},
		{node: t.root.right.right, key: 40, height: 0, nodePath: "t.root.right.right"},
	}
	for _, n := range validTreeAfterInsert {
		checkNode(t1, &n)
	}
}

func TestTree_Insert_left_right_rotate(t1 *testing.T) {
	t := getTree([]int{20, 10, 25, 7, 15})

	validTree := []validNode[int]{
		{node: t.root, key: 20, height: 2, nodePath: "t.root"},
		{node: t.root.left, key: 10, height: 1, nodePath: "t.root.left"},
		{node: t.root.right, key: 25, height: 0, nodePath: "t.root.right"},
		{node: t.root.left.left, key: 7, height: 0, nodePath: "t.root.left.left"},
		{node: t.root.left.right, key: 15, height: 0, nodePath: "t.root.left.right"},
	}

	// check tree's structure and heights before insert
	for _, n := range validTree {
		checkNode(t1, &n)
	}

	// insert
	t.Insert(16, 16)

	// check tree's structure and heights after insert
	validTreeAfterInsert := []validNode[int]{
		{node: t.root, key: 15, height: 2, nodePath: "t.root"},
		{node: t.root.left, key: 10, height: 1, nodePath: "t.root.left"},
		{node: t.root.right, key: 20, height: 1, nodePath: "t.root.right"},
		{node: t.root.left.left, key: 7, height: 0, nodePath: "t.root.left.left"},
		{node: t.root.right.left, key: 16, height: 0, nodePath: "t.root.right.left"},
		{node: t.root.right.right, key: 25, height: 0, nodePath: "t.root.right.right"},
	}
	for _, n := range validTreeAfterInsert {
		checkNode(t1, &n)
	}
}

func TestTree_Delete(t1 *testing.T) {
	type args[V constraints.Ordered] struct {
		key V
	}
	type testCase[V constraints.Ordered] struct {
		name string
		t    *Tree[V]
		args args[V]
		want *Tree[V]
	}

	tests := []testCase[int]{
		{
			name: "empty tree",
			t:    getTree([]int{}),
			args: args[int]{key: 1},
			want: getTree([]int{}),
		},
		{
			name: "tree only with root - without changes",
			t:    getTree([]int{15}),
			args: args[int]{key: 1},
			want: getTree([]int{15}),
		},
		{
			name: "tree only with root - delete root",
			t:    getTree([]int{15}),
			args: args[int]{key: 15},
			want: getTree([]int{}),
		},
		{
			name: "tree with elements - without changes",
			t:    getTree([]int{15, 25}),
			args: args[int]{key: 85},
			want: getTree([]int{15, 25}),
		},
		{
			name: "tree with elements - delete node without children",
			t:    getTree([]int{15, 25}),
			args: args[int]{key: 25},
			want: getTree([]int{15}),
		},
		{
			name: "delete root with left and right node",
			t:    getTree([]int{25, 15, 35}),
			args: args[int]{key: 25},
			want: getTree([]int{35, 15}),
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			tt.t.Delete(tt.args.key)
			if !treeEquals(tt.t, tt.want) {
				t1.Errorf("Delete() = %v, want %v", tt.t, tt.want)
			}
		})
	}
}

func getTree(elements []int) *Tree[int] {
	tree := New[int]()
	for _, el := range elements {
		tree.Insert(el, el)
	}

	return tree
}

func checkNode[V constraints.Ordered](t *testing.T, vn *validNode[V]) {
	if vn == nil {
		return
	}

	if vn.node.element.key != vn.key {
		t.Errorf("Error - Want key: %v, have key %v in %s\n",
			vn.key,
			vn.node.element.key,
			vn.nodePath,
		)
	}

	if vn.node.height != vn.height {
		t.Errorf("Error - Want height: %v, have height: %v in %s\n",
			vn.height,
			vn.node.height,
			vn.nodePath,
		)
	}
}

func treeEquals(tree1, tree2 *Tree[int]) bool {
	return nodesEquals(tree1.root, tree2.root)
}

func nodesEquals(node1, node2 *node[int]) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}

	return node1.element.key == node2.element.key &&
		node1.element.value == node2.element.value &&
		node1.height == node2.height &&
		nodesEquals(node1.left, node2.left) &&
		nodesEquals(node1.right, node2.right)
}
