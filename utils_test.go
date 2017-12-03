package go_tries

import (
	"testing"
)

func TestEnsureIndex(t *testing.T) {
	arr := make([]int, 10, 10)
	arr = EnsureIndex(arr, 25)

	if cap(arr) < 25 {
		t.Errorf("array index is not reachable at %v, capacity is %v", 25, cap(arr))
	}

	if len(arr) < 25 {
		t.Errorf("array index is not reachable at %v, length is %v", 25, len(arr))
	}
}