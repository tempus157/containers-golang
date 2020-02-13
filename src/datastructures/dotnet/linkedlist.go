package datastructures

// LinkedList represents a doubly linked list.
type LinkedList struct {
	first, last *LinkedListNode
	cnt         int
}

// Cnt gets the number of nodes actually contained in the LinkedList.
func (list *LinkedList) Cnt() int {
	return list.cnt
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

// AddValAfter adds a new node containing the specified value after the specified existing node in the LinkedList.
func (list *LinkedList) AddValAfter(node *LinkedListNode, val Val) *LinkedListNode {

}

// AddBefore adds the specified new node before the specified existing node in the LinkedList.
func (list *LinkedList) AddBefore(node, newNode *LinkedListNode) {

}

// AddValBefore adds a new node containing the specified value before the specified existing node in the LinkedList.
func (list *LinkedList) AddValBefore(node *LinkedListNode, val Val) *LinkedListNode {

}

// AddFirst adds the specified new node at the start of the LinkedList.
func (list *LinkedList) AddFirst(node *LinkedListNode) {

}

// AddValFirst adds a new node containing the specified value at the start of the LinkedList.
func (list *LinkedList) AddValFirst(val Val) *LinkedListNode {

}

// AddLast adds the specified new node at the end of the LinkedList.
func (list *LinkedList) AddLast(node *LinkedListNode) {

}

// AddValLast adds a new node containing the specified value at the end of the LinkedList.
func (list *LinkedList) AddValLast(val Val) *LinkedListNode {

}

// Clear removes all nodes from the LinkedList.
func (list *LinkedList) Clear() {

}

// Contains determines wheter a value is in the LinkedList.
func (list *LinkedList) Contains(val Val) bool {

}

// Find finds the first node that contains the specified value.
func (list *LinkedList) Find(val Val) *LinkedListNode {

}

// Iter iterates through the LinkedList.
func (list *LinkedList) Iter() <-chan interface{} {

}

// Remove removes the specified node from the LinkedList.
func (list *LinkedList) Remove(node *LinkedListNode) {

}

// RemoveVal removes the first occurrence of the specified value from the LinkedList.
func (list *LinkedList) RemoveVal(val Val) bool {

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
