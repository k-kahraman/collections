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
	"collections/list"
)

const (
	// GrowthCoefficient Determines how much the array list will grow,
	// giving a smaller coefficient is better for space complexity but the trade-off is time complexity,
	// because array will grow more frequently as more data added
	GrowthCoefficient = 2.0

	// ShrinkCoefficient Determines how much the array list will shrink
	ShrinkCoefficient = 0.25
)

type ArrayList[T any] struct {
	list.List[T]
	elements []T
	length   int
}

func (list *ArrayList[T]) grow(exceededSize int) {
	capacity := cap(list.elements)
	if list.length+exceededSize >= capacity {
		newCapacity := int(GrowthCoefficient * float32(capacity+exceededSize))
		list.resize(newCapacity)
	}
}

func (list *ArrayList[T]) shrink() {
	capacity := cap(list.elements)

	if list.length <= int(float32(capacity)*ShrinkCoefficient) {
		list.resize(list.length)
	}
}

func (list *ArrayList[T]) resize(capacity int) {
	elements := make([]T, capacity, capacity)
	copy(elements, list.elements)
	list.elements = elements
}

func (list *ArrayList[T]) Add(elements ...T) {
	list.grow(len(elements))
	for index, element := range elements {
		list.elements[index] = element
		list.length++
	}
}

func (list *ArrayList[T]) Get(index int) (T, error) {
	absIndex, err := absoluteIndex(index, list.length)
	if err != nil {
		var zeroValue T // To return type specific zero value
		return zeroValue, err
	}
	return list.elements[absIndex], nil
}

func (list *ArrayList[T]) Size() int {
	return list.length
}

// IndexOf searches the array using linear search runtime is O(N).
// This can be improved in the future using an auxiliary array.
func (list *ArrayList[T]) IndexOf(element T) int {
	// TODO implement this with along side of Comparator
	return -1
}

func (list *ArrayList[T]) SubList(startIndex, endIndex int) *ArrayList[T] {
	if startIndex < 0 || endIndex > 0 || startIndex >= endIndex {
		return nil
	} else {
		return &ArrayList[T]{
			elements: list.elements[startIndex:endIndex],
			length:   endIndex - startIndex,
		}
	}
}

func (list *ArrayList[T]) Contains(element T) bool {
	if list.IndexOf(element) > 0 {
		return true
	} else {
		return false
	}
}

func (list *ArrayList[T]) ContainsAll(compList list.List[T]) bool {
	return false
}

func (list *ArrayList[T]) Values() []T {
	return list.elements
}

func (list *ArrayList[T]) ForEachValue(action func(v T)) {
	for _, val := range list.elements {
		action(val)
	}
}

func (list *ArrayList[T]) ForEachRefer(action func(p *T)) {
	for _, val := range list.elements {
		action(&val)
	}
}

func New[T any](elements ...T) *ArrayList[T] {
	arrayList := &ArrayList[T]{}

	if len(elements) > 0 {
		arrayList.Add(elements...)
	}
	return arrayList
}

// absoluteIndex Gets the absolute index from the end, -1 means last element of the list,
// -2 means second last element from the index and so on, this goes on until abs(index - list.length) < list.length
func absoluteIndex(index, size int) (int, error) {
	if index < 0 {
		index = -index
	}

	if index > size {
		return -1, list.Error("Index out of boundaries")
	} else {
		return index, nil
	}
}
