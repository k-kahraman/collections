package stack

import (
	"collections/linear"
	"collections/linear/vector"
)

type Stack[T comparable] struct {
	stackVector *vector.Vector[T]
}

func (stack Stack[T]) Push(element T) {
	stack.stackVector.AddBack(element)
}

func (stack Stack[T]) AddFront(element T) {
	stack.Push(element)
}

func (stack Stack[T]) Pop() (T, error) {
	if stack.stackVector.Size() == 0 {
		var result T
		return result, linear.Error("Stack is empty!")
	}
	lastIndex := stack.stackVector.Size() - 1
	elem, _ := stack.stackVector.Get(lastIndex)
	stack.stackVector.Remove(lastIndex)

	return elem, nil
}

func (stack Stack[T]) Size() int {
	return stack.stackVector.Size()
}

func (stack Stack[T]) IsEmpty() bool {
	return stack.stackVector.IsEmpty()
}

func (stack Stack[T]) Clear() {
	stack.stackVector.Clear()
}

func (stack Stack[T]) Values() []T {
	return stack.stackVector.Values()
}

func New[T comparable](elements ...T) *Stack[T] {
	vec := vector.New(elements...)
	return &Stack[T]{
		stackVector: vec,
	}
}
