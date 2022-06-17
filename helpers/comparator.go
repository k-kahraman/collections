package helpers

import "fmt"

type Comparator[T Comparable] func(first, second T) bool

type IntComparable interface {
	int | int8 | int16 | int32 | int64
}

type UIntComparable interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type FloatComparable interface {
	float32 | float64
}

type StringComparable interface {
	string
}

type Comparable interface {
	IntComparable
	UIntComparable
	FloatComparable
	StringComparable
	Compare()
}

func Compare[T Comparable](first, second interface{}) bool {
	fmt.Println(first)
	return false
}
