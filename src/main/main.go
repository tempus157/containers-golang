package main

import (
	"containers/binarytree"
	"fmt"
)

func main() {
	tree := binarytree.New(0)
	left := tree.SetValLeft(1)
	left.SetValLeft(2)
	left.SetValRight(3)
	right := tree.SetValRight(4)
	right.SetValLeft(5)

	for n := range tree.Iter() {
		fmt.Println(n.Val())
	}
}
