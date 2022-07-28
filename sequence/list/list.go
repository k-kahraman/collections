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

import (
	"collections/sequence"
)

type Node[T any] struct {
	value T
	next  *Node[T]
}

type List[T any] struct {
	head   *Node[T]
	length uint
}

func (list *List[T]) Add(element T) {
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

func (list *List[T]) InsertAt(index uint, element T) error {
	if list.length >= index {
		return sequence.Error("Index out of boundaries")
	}

	node := list.head
	for counter := uint(0); counter <= index; counter++ {
		node = node.next
	}

	node = &Node[T]{
		value: element,
		next:  node.next,
	}

	return nil
}

func New[T any](elements ...T) *List[T] {
	return &List[T]{}
}
