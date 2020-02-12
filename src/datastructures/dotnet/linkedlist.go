package datastructures

// Item is the ultimate base type of all Golang types.
type Item interface{}

// LinkedListNode represents a node in a LinkedList.
type LinkedListNode struct {
	item       Item
	prev, next *LinkedListNode
}

// LinkedList represents a doubly linked list.
type LinkedList struct {
	first, last *LinkedListNode
	count       int
}

// Count gets the number of nodes actually contained in the LinkedList.
func (list *LinkedList) Count() int {
	return list.count
}

// First gets the first node of the LinkedList.
func (list *LinkedList) First() *LinkedListNode {
	return list.first
}

// Last gets the last node of the LinkedList.
func (list *LinkedList) Last() *LinkedListNode {
	return list.last
}

// AddAfter adds the specified new node after the specified existing node in the LinkedList.
func (list *LinkedList) AddAfter(node, newNode *LinkedListNode) {

}

// AddItemAfter adds a new node containing the specified value after the specified existing node in the LinkedList.
func (list *LinkedList) AddItemAfter(node *LinkedListNode, item Item) *LinkedListNode {

}

// AddBefore adds the specified new node before the specified existing node in the LinkedList.
func (list *LinkedList) AddBefore(node, newNode *LinkedListNode) {

}

// AddItemBefore adds a new node containing the specified value before the specified existing node in the LinkedList.
func (list *LinkedList) AddItemBefore(node *LinkedListNode, item Item) *LinkedListNode {

}

// AddFirst adds the specified new node at the start of the LinkedList.
func (list *LinkedList) AddFirst(node *LinkedListNode) {

}

// AddItemFirst adds a new node containing the specified value at the start of the LinkedList.
func (list *LinkedList) AddItemFirst(item Item) *LinkedListNode {

}

// AddLast adds the specified new node at the end of the LinkedList.
func (list *LinkedList) AddLast(node *LinkedListNode) {

}

// AddItemLast adds a new node containing the specified value at the end of the LinkedList.
func (list *LinkedList) AddItemLast(item Item) *LinkedListNode {

}

// Clear removes all nodes from the LinkedList.
func (list *LinkedList) Clear() {

}

// Contains determines wheter a value is in the LinkedList.
func (list *LinkedList) Contains(item Item) bool {

}

// Find finds the first node that contains the specified value.
func (list *LinkedList) Find(item Item) *LinkedListNode {

}

// Iterate iterates through the LinkedList.
func (list *LinkedList) Iterate() <-chan interface{} {

}

// Remove removes the specified node from the LinkedList.
func (list *LinkedList) Remove(node *LinkedListNode) {

}

// RemoveItem removes the first occurrence of the specified value from the LinkedList.
func (list *LinkedList) RemoveItem(item Item) bool {

}

// RemoveFirst removes the node at the start of the LinkedList.
func (list *LinkedList) RemoveFirst() {

}

// RemoveLast removes the node at the end of the LinkedList.
func (list *LinkedList) RemoveLast() {

}

// String returns a string that represents the LinkedList.
func (list *LinkedList) String() string {

}
