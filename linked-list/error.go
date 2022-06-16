package linked_list

import "fmt"

type LinkedListError string

func (err LinkedListError) Error() string {
	return fmt.Sprintf("LinkedListError! Cause: %v", err)
}
