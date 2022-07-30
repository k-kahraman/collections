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
	"collections/collection"
	"collections/linear"
)

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

type List[T comparable] struct {
	head   *Node[T]
	length int
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

func (list *List[T]) AddBack(element T) {
	list.Add(element)
}

// AddFront O(1) runtime
func (list *List[T]) AddFront(element T) {
	prevStartNode := list.head

	newStartNode := &Node[T]{
		value: element,
		next:  prevStartNode,
	}

	list.head = newStartNode
}

func (list *List[T]) InsertAt(index int, element T) error {
	if list.length >= index {
		return linear.Error("Index out of boundaries")
	}

	node := list.head
	for counter := 0; counter <= index; counter++ {
		node = node.next
	}

	node = &Node[T]{
		value: element,
		next:  node.next,
	}

	return nil
}

func (list *List[T]) Get(index int) (T, error) {
	if index > list.length {
		return nil, linear.Error("Index out of boundaries")
	}

	node := list.head
	for counter := 0; counter < index; counter++ {
		node = node.next
	}

	return node.value, nil
}

func (list *List[T]) Size() int {
	return list.length
}

func (list *List[T]) Contains(t T) bool {
	node := list.head

	for node.next != nil {
		if node.value == t {
			return true
		}
		node = node.next
	}

	return false
}

func (list *List[T]) ContainsAll(collection collection.Collection[T]) bool {
	collectionElements := collection.Values()

	for _, element := range collectionElements {
		if !list.Contains(element) {
			return false
		}
	}

	return true
}

func (list *List[T]) IsEmpty() bool {
	return list.length == 0
}

func (list *List[T]) Clear() {
	list.head = nil
}

func (list *List[T]) Values() []T {
	//TODO implement me
	panic("implement me")
}

func (list *List[T]) IndexOf(element T) int {
	node := list.head
	counter := 0
	found := false

	for node.next != nil {
		if element == node.value {
			found = true
			break
		}
		counter++
	}

	if found {
		return counter
	}

	return -1
}

func (list *List[T]) SubList(startIndex, endIndex int) *List[T] {
	if endIndex >= startIndex || startIndex < 0 || endIndex > list.length {
		return nil
	}

	newList := &List[T]{}
	node := list.head
	for startIndex < endIndex {
		newList.Add(node.value)

		node = node.next
		startIndex++
	}

	return newList
}

func (list *List[T]) Remove(index int) error {
	if index > list.length {
		return linear.Error("Index out of boundaries!")
	}

	node := list.head

	for counter := 0; counter < index; counter++ {
		node = node.next
	}

	nextNode := node.next.next
	prevNode := node

	prevNode.next = nextNode

	return nil
}

func New[T comparable](elements ...T) *List[T] {
	list := &List[T]{}

	for _, element := range elements {
		list.Add(element)
	}

	return list
}
