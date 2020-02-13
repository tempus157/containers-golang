package linkedlist

import "fmt"

// LinkedList represents a doubly linked list.
type LinkedList struct {
	first, last *Node
	cnt         int
}

// Cnt gets the number of nodes actually contained in the LinkedList.
func (list *LinkedList) Cnt() int {
	return list.cnt
}

// First gets the first node of the LinkedList.
func (list *LinkedList) First() *Node {
	if list.first.next.isDummy {
		return nil
	}

	return list.first.next
}

// Last gets the last node of the LinkedList.
func (list *LinkedList) Last() *Node {
	if list.last.prev.isDummy {
		return nil
	}

	return list.last.prev
}

// AddAfter adds the specified new node after the specified existing node in the LinkedList.
func (list *LinkedList) AddAfter(node, newNode *Node) error {
	if node == nil {
		return fmt.Errorf("node is nil")
	} else if newNode == nil {
		return fmt.Errorf("newNode is nil")
	} else if node.list != list {
		return fmt.Errorf("node is not in the current LinkedList")
	} else if newNode.list != nil {
		return fmt.Errorf("newNode belongs to another LinkedList")
	}

	newNode.list = list
	newNode.prev = node
	newNode.next = node.next
	node.next.prev = newNode
	node.next = newNode

	list.cnt++
	return nil
}

// AddValAfter adds a new node containing the specified value after the specified existing node in the LinkedList.
func (list *LinkedList) AddValAfter(node *Node, val Value) (*Node, error) {
	newNode := NewNode(val)
	err := list.AddAfter(node, newNode)
	return newNode, err
}

// AddBefore adds the specified new node before the specified existing node in the LinkedList.
func (list *LinkedList) AddBefore(node, newNode *Node) error {
	return list.AddAfter(node.prev, newNode)
}

// AddValBefore adds a new node containing the specified value before the specified existing node in the LinkedList.
func (list *LinkedList) AddValBefore(node *Node, val Value) (*Node, error) {
	return list.AddValAfter(node.prev, val)
}

// AddFirst adds the specified new node at the start of the LinkedList.
func (list *LinkedList) AddFirst(node *Node) error {
	return list.AddAfter(list.first, node)
}

// AddValFirst adds a new node containing the specified value at the start of the LinkedList.
func (list *LinkedList) AddValFirst(val Value) (*Node, error) {
	return list.AddValAfter(list.first, val)
}

// AddLast adds the specified new node at the end of the LinkedList.
func (list *LinkedList) AddLast(node *Node) error {
	return list.AddBefore(list.last, node)
}

// AddValLast adds a new node containing the specified value at the end of the LinkedList.
func (list *LinkedList) AddValLast(val Value) (*Node, error) {
	return list.AddValBefore(list.last, val)
}

// Clear removes all nodes from the LinkedList.
func (list *LinkedList) Clear() {
	for {
		if err := list.RemoveFirst(); err != nil {
			break
		}
	}
}

// Find finds the first node that contains the specified value.
func (list *LinkedList) Find(val Value) *Node {
	for node := range list.Iter() {
		if node.val == val {
			return node
		}
	}

	return nil
}

// Iter iterates through the LinkedList.
func (list *LinkedList) Iter() <-chan *Node {
	ch := make(chan *Node)

	go func() {
		defer close(ch)
		cur := list.first.next

		for !cur.isDummy {
			ch <- cur
			cur = cur.next
		}
	}()

	return ch
}

// Remove removes the specified node from the LinkedList.
func (list *LinkedList) Remove(node *Node) error {
	if node == nil {
		return fmt.Errorf("node is nil")
	} else if node.list != list {
		return fmt.Errorf("node is not in the current LinkedList")
	}

	node.prev.next = node.next
	node.next.prev = node.prev
	node.list = nil
	node.prev = nil
	node.next = nil

	list.cnt--
	return nil
}

// RemoveVal removes the first occurrence of the specified value from the LinkedList.
func (list *LinkedList) RemoveVal(val Value) bool {
	node := list.Find(val)
	list.Remove(node)
	return node != nil
}

// RemoveFirst removes the node at the start of the LinkedList.
func (list *LinkedList) RemoveFirst() error {
	node := list.First()

	if node == nil {
		return fmt.Errorf("the LinkedList is empty")
	}

	list.Remove(node)
	return nil
}

// RemoveLast removes the node at the end of the LinkedList.
func (list *LinkedList) RemoveLast() error {
	node := list.Last()

	if node == nil {
		return fmt.Errorf("the LinkedList is empty")
	}

	list.Remove(node)
	return nil
}

// String returns a string that represents the LinkedList.
func (list *LinkedList) String() string {
	vals := make([]Value, 0, list.cnt)

	for node := range list.Iter() {
		vals = append(vals, node.val)
	}

	return fmt.Sprint(vals)
}

// New initializes a new instance of the LinkedList struct that is empty.
func New() *LinkedList {
	list := &LinkedList{}

	first := &Node{true, nil, list, nil, nil}
	last := &Node{true, nil, list, nil, nil}
	first.next = last
	last.prev = first

	list.first = first
	list.last = last
	return list
}
