package main

import (
	"containers/linkedlist"
	"fmt"
)

func main() {
	list := linkedlist.New()
	list.AddValLast(1)
	list.AddValLast(2)
	list.AddValLast(3)
	list.AddValLast(4)
	list.AddValLast(5)

	for val := range list.Iter() {
		fmt.Println(val.Val())
	}
}
