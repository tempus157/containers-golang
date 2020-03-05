package binarytree

// Value is the ultimate vase type of all Golang types.
type Value interface{}

// BinaryTree represents a binary tree.
type BinaryTree struct {
	val   Value
	left  *BinaryTree
	right *BinaryTree
}

// SetLeft adds the specified new tree at the left of the BinaryTree.
func (tree *BinaryTree) SetLeft(newTree *BinaryTree) {
	tree.left = newTree
}

// SetValLeft adds a new node containing the specified value at the left of the BInaryTree.
func (tree *BinaryTree) SetValLeft(val Value) *BinaryTree {
	newTree := New(val)
	tree.SetLeft(newTree)
	return newTree
}

// SetRight adds the specified new tree at the right of the BinaryTree.
func (tree *BinaryTree) SetRight(newTree *BinaryTree) {
	tree.right = newTree
}

// SetValRight adds a new node containing the specified value at the right of the BinaryTree.
func (tree *BinaryTree) SetValRight(val Value) *BinaryTree {
	newTree := New(val)
	tree.SetRight(newTree)
	return newTree
}

// Clear removes all binary trees from the BinaryTree.
func (tree *BinaryTree) Clear() {
	tree.val = nil
	tree.RemoveLeft()
	tree.RemoveRight()
}

// RemoveLeft removes the binary tree at the left of the BinaryTree.
func (tree *BinaryTree) RemoveLeft() {
	if tree.left == nil {
		return
	}

	tree.left.Clear()
	tree.left = nil
}

// RemoveRight removes the binary tree at the right of the BinaryTree.
func (tree *BinaryTree) RemoveRight() {
	if tree.right == nil {
		return
	}

	tree.right.Clear()
	tree.right = nil
}

// Left gets the left sub tree contained in the tree.
func (tree *BinaryTree) Left() *BinaryTree {
	return tree.left
}

// Right gets the right sub tree contained in the tree.
func (tree *BinaryTree) Right() *BinaryTree {
	return tree.right
}

// SetVal sets the value contained in the tree.
func (tree *BinaryTree) SetVal(newVal Value) {
	tree.val = newVal
}

// Val gets the value contained in the tree.
func (tree *BinaryTree) Val() Value {
	return tree.val
}

// New initializes a new instance of the BinaryTree struct, containing the specified value.
func New(val Value) *BinaryTree {
	return &BinaryTree{val, nil, nil}
}
