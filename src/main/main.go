package main

import (
	"containers/linkedlist"
	"fmt"
)

func main() {
	list := linkedlist.New()
	node0, _ := list.AddValLast(1)

	list.AddValFirst(2)
	list.AddValAfter(node0, 3)
	list.AddValBefore(node0, 4)
	list.AddLast(linkedlist.NewNode(5))
	list.AddFirst(linkedlist.NewNode(6))
	list.AddAfter(node0, linkedlist.NewNode(7))
	list.AddBefore(node0, linkedlist.NewNode(8))
	fmt.Println(list, list.Cnt(), list.First().Val(), list.Last().Val(), list.Find(5).Val())

	list.Remove(node0)
	list.RemoveFirst()
	list.RemoveLast()
	list.RemoveVal(2)

	for n := range list.Iter() {
		fmt.Print(n.Val())
	}

	list.Clear()
	fmt.Println("\nClear: ", list, list.Cnt())
}
