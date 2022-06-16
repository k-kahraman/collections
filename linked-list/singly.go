package linked_list

import "fmt"

type Node[T any] struct {
	value T
	next  *Node[T]
}

type SinglyList[T any] struct {
	head   *Node[T]
	length uint
}

func (list *SinglyList[T]) Append(element T) {
	if list.head == nil {
		list.head = &Node[T]{
			value: element,
			next:  nil,
		}
		list.length++
		return
	}

	head := list.head

	for head.next != nil {
		head = head.next
	}

	list.length++
	head.next = &Node[T]{
		value: element,
		next:  nil,
	}
}

// InsertAt How should this behave? insert empty nodes until the given index,
// or if given index is longer than the length of linked list return?
func (list *SinglyList[T]) InsertAt(index uint, element T) error {
	if list.length >= index {
		return LinkedListError("Index out of boundaries")
	}
	// TODO implement this lmao

	return nil
}

func (list *SinglyList[T]) PrintList() {
	head := list.head

	for head != nil {
		fmt.Println(head.value)
		head = head.next
	}
}

func InitSingly[T any]() SinglyList[T] {
	return SinglyList[T]{}
}
