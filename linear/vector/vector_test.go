package vector

import "testing"

func TestVector_New(t *testing.T) {
	vec := New(10, 20, 30)

	for index := 0; index < 3; index++ {
		elem, _ := vec.Get(index)
		t.Logf("Vector element at index %d is %d", index, elem)
	}

}
