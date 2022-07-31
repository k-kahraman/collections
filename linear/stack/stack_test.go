package stack

import "testing"

func TestStack_Push(t *testing.T) {
	testStack := New(10, 20, 30)

	t.Logf("Stack size is %d", testStack.Size())

	for index := 0; index < 3; index++ {
		elem, _ := testStack.Pop()
		t.Logf("Stack first item %d", elem)

	}
}
