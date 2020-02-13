package linkedlist

// Value is the ultimate base type of all Golang types.
type Value interface{}

// Node represents a node in a LinkedList.
type Node struct {
	isDummy    bool
	val        Value
	list       *LinkedList
	prev, next *Node
}

// List gets the LinkedList that the Node belongs to.
func (node *Node) List() *LinkedList {
	return node.list
}

// Next gets the next node in the LinkedList.
func (node *Node) Next() *Node {
	if node.next.isDummy {
		return nil
	}

	return node.next
}

// Prev gets the previous node in the LinkedList.
func (node *Node) Prev() *Node {
	if node.prev.isDummy {
		return nil
	}

	return node.prev
}

// Val gets the value contained in the node.
func (node *Node) Val() Value {
	return node.val
}

// NewNode initializes a new instance of the Node struct, containing the specified value.
func NewNode(val Value) *Node {
	return &Node{false, val, nil, nil, nil}
}
