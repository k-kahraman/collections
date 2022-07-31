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

package vector

import (
	"collections/collection"
	"collections/linear"
)

const (
	// GrowthCoefficient Determines how much the array linear will extend,
	// giving a smaller coefficient is better for space complexity but the trade-off is time complexity,
	// because array will extend more frequently as more data added.
	// 50% increase is inspired from Java's ArrayList implementation.
	GrowthCoefficient = 1.5

	// ShrinkCoefficient Determines how much the array linear will shrink
	ShrinkCoefficient = 0.25
)

type Vector[T comparable] struct {
	elements []T
	length   int
}

func (vector *Vector[T]) extend(newElementCount int) {
	capacity := cap(vector.elements)
	if vector.length+newElementCount >= capacity {
		newCapacity := int(GrowthCoefficient * float32(capacity+newElementCount))
		vector.resize(newCapacity)
	}
}

func (vector *Vector[T]) shrink() {
	capacity := cap(vector.elements)

	if vector.length <= int(float32(capacity)*ShrinkCoefficient) {
		vector.resize(vector.length)
	}
}

func (vector *Vector[T]) resize(capacity int) {
	elements := make([]T, capacity, capacity)
	copy(elements, vector.elements)
	vector.elements = elements
}

// Contains Binary search. Can be optimized with sorting first and then looking with binary search
func (vector *Vector[T]) Contains(t T) bool {
	for _, elem := range vector.elements {
		if elem == t {
			return true
		}
	}
	return false
}

func (vector *Vector[T]) ContainsAll(collection collection.Collection[T]) bool {
	collectionElements := collection.Values()

	for _, element := range collectionElements {
		if !vector.Contains(element) {
			return false
		}
	}

	return true
}

func (vector *Vector[T]) IndexOf(element T) int {
	for index, value := range vector.elements {
		if value == element {
			return index
		}
	}

	return -1
}

func (vector *Vector[T]) SubList(startIndex, endIndex int) *Vector[T] {
	if endIndex >= startIndex || startIndex < 0 || endIndex > vector.length {
		return nil
	}

	newVec := &Vector[T]{}

	copy(newVec.elements, vector.elements[startIndex:endIndex])

	return newVec
}

func (vector *Vector[T]) Add(element T) {
	vector.extend(1)
	vector.elements[vector.length] = element
	vector.length++
}

func (vector *Vector[T]) AddBack(element T) {
	vector.Add(element)
}

func (vector *Vector[T]) AddFront(element T) {
	vector.extend(1)

	for index := vector.length - 1; index >= 0; index++ {
		vector.elements[index+1] = vector.elements[index]
	}

	vector.elements[0] = element

	vector.length++
}

func (vector *Vector[T]) Get(index int) (T, error) {

	if index > vector.length || index < 0 {
		var result T
		return result, linear.Error("Index out of boundaries!")
	}

	return vector.elements[index], nil
}

func (vector *Vector[T]) Size() int {
	return vector.length
}

func (vector *Vector[T]) Remove(index int) error {
	// Left shift according to index
	if index >= vector.length {
		return linear.Error("Index out of boundaries for removal!")
	}

	// Left shift starts from right of the index
	leftIndex := index
	rightIndex := index + 1
	for rightIndex < vector.length {
		vector.elements[leftIndex] = vector.elements[rightIndex]
		leftIndex++
		rightIndex++
	}

	vector.length--

	return nil
}

func (vector *Vector[T]) IsEmpty() bool {
	return vector.length == 0
}

func (vector *Vector[T]) Clear() {
	vector.length = 0
	vector.elements = make([]T, 0)
}

func (vector *Vector[T]) Values() []T {
	return vector.elements
}

func (vector *Vector[T]) InsertAt(element T, index int) error {
	if index > vector.length || index < 0 {
		return linear.Error("Index out of boundaries!")
	}

	vector.elements[index] = element

	return nil
}

func New[T comparable](elements ...T) *Vector[T] {
	vector := &Vector[T]{
		length: 0,
	}

	vector.extend(len(elements))

	if len(elements) < 0 {
		return vector
	}

	for _, element := range elements {
		vector.Add(element)
	}

	return vector
}
