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

package arraylist

import (
	"collections/collection"
	"collections/sequence"
)

const (
	// GrowthCoefficient Determines how much the array sequence will grow,
	// giving a smaller coefficient is better for space complexity but the trade-off is time complexity,
	// because array will grow more frequently as more data added
	GrowthCoefficient = 2.0

	// ShrinkCoefficient Determines how much the array sequence will shrink
	ShrinkCoefficient = 0.25
)

type List[T any] struct {
	elements []T
	length   int
}

// Contains Binary search. Can be optimized with sorting first and then looking with binary search
func (list *List[T]) Contains(t T) bool {
	// TODO implement custom comparable
	panic("Implementation awaits!")
}

func (list *List[T]) ContainsAll(collection collection.Collection[T]) bool {
	// TODO implement custom comparable
	panic("implement me")
}

func (list *List[T]) IndexOf(element T) (int, error) {
	// TODO implement custom comparable
	panic("implement me")
}

func (list *List[T]) SubList(startIndex, endIndex int) (sequence.Sequence[T], error) {
	// TODO implement custom comparable
	panic("implement me")
}

func (list *List[T]) grow(exceededSize int) {
	capacity := cap(list.elements)
	if list.length+exceededSize >= capacity {
		newCapacity := int(GrowthCoefficient * float32(capacity+exceededSize))
		list.resize(newCapacity)
	}
}

func (list *List[T]) shrink() {
	capacity := cap(list.elements)

	if list.length <= int(float32(capacity)*ShrinkCoefficient) {
		list.resize(list.length)
	}
}

func (list *List[T]) resize(capacity int) {
	elements := make([]T, capacity, capacity)
	copy(elements, list.elements)
	list.elements = elements
}

func (list *List[T]) Add(elements ...T) {
	list.grow(len(elements))
	for index, element := range elements {
		list.elements[index] = element
		list.length++
	}
}

func (list *List[T]) Get(index int) (T, error) {
	return list.elements[index], nil
}

func (list *List[T]) Size() int {
	return list.length
}

func (list *List[T]) Remove(index int) error {
	// Left shift according to index
	if index >= list.length {
		return sequence.Error("Index out of boundaries for removal!")
	}

	// Left shift the whole array
	leftIndex := index
	rightIndex := index + 1
	for rightIndex < list.length {
		list.elements[leftIndex] = list.elements[rightIndex]
		leftIndex++
		rightIndex++
	}

	list.length--

	return nil
}

func (list *List[T]) IsEmpty() bool {
	return list.length == 0
}

func (list *List[T]) Clear() {
	list.length = 0
	list.elements = make([]T, 0)
}

func (list *List[T]) Values() []T {
	return list.elements
}

func (list *List[T]) InsertAt(element T, index int) error {
	if index > list.length {
		return sequence.Error("Index out of boundaries!")
	}

	list.elements[index] = element

	return nil
}

func New[T any](elements ...T) *List[T] {
	arrayList := &List[T]{
		length: 0,
	}

	if len(elements) > 0 {
		arrayList.Add(elements...)
	}
	return arrayList
}
