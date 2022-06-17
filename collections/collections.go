package collections

// Collection is the base interface where all collections will implement
type Collection[T any] interface {
	IsEmpty() bool
	Size() uint
	Clear()
	Values()
}
