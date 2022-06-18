/*
 * MIT License
 *
 * Copyright (c) 2022 Kaan Kahraman
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package list

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
		return Error("Index out of boundaries")
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
