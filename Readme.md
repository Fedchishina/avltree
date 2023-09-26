avltree
=======================

Library for work with AVL trees.

You can create a AVL tree and use a list of functions to work with it.
## Tree functions
- [Empty tree's creation example](#empty-trees-creation-example)
- [Insert element to tree](#insert-element-to-tree)
- [Delete element by key from tree](#delete-element-by-key-from-tree)

### Empty tree's creation example

```
t := tree.New[int]() // empty int tree
t := tree.New[string]() // empty string tree
```

### Insert element to tree
```
t := tree.New[int]() // empty int tree
t.Insert(22, 22) // insert to tree element: key=22, value=22
t.Insert(8, 8) // insert to tree element: key=8, value=8
t.Insert(4, 4) // insert to tree element: key=4, value=4
```

### Delete element by key from tree
```
t := tree.New[int]()
t.Insert(22, 22)
t.Insert(8, 8)
t.Insert(4, 4)

err := t.Delete(22) // without err
```