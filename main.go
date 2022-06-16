package main

import (
	"go-dsa/linked-list"
)

func main() {
	list := linked_list.SinglyList[string]{}
	list.Append("A")
	list.Append("B")
	list.Append("C")

	list.PrintList()

}
