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

import "collections/list"

const (
	// GrowthCoefficient Determines how much the array list will grow,
	// giving a smaller coefficient is better for space complexity but the trade-off is time complexity,
	// because array will grow more frequently as more data added
	GrowthCoefficient = 2.0

	// ShrinkCoefficient Determines how much the array list will shrink
	ShrinkCoefficient = 0.25
)

type List[T any] struct {
	list.List[T]
	elements []T
	length   int
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

func New[T any](elements ...T) *List[T] {
	arrayList := &List[T]{}

	if len(elements) > 0 {
		arrayList.Add(elements...)
	}
	return arrayList
}
