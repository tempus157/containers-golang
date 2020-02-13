package linkedlist

// Val is the ultimate base type of all Golang types.
type Val interface{}

// LinkedListNode represents a node in a LinkedList.
type LinkedListNode struct {
	val        Val
	list       *LinkedList
	prev, next *LinkedListNode
}

// List gets the LinkedList that the LinkedListNode belongs to.
func (node *LinkedListNode) List() *LinkedList {
	return node.list
}

// Next gets the next node in the LinkedList.
func (node *LinkedListNode) Next() *LinkedListNode {
	return node.next
}

// Prev gets the previous node in the LinkedList.
func (node *LinkedListNode) Prev() *LinkedListNode {
	return node.prev
}

// Val gets the value contained in the node.
func (node *LinkedListNode) Val() *LinkedListNode {
	return node.val
}

// NewNode initializes a new instance of the LinkedListNode struct, containing the specified value.
func NewNode(val Val) *LinkedListNode {
	return &LinkedListNode{val, nil, nil, nil}
}
