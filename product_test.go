package enumerate_algorithm

import (
	"testing"
)

func TestProduct(t *testing.T) {
	sliceA := []string{"A", "B", "C"}
	sliceB := []int{1, 2, 3}
	r := Product(sliceA, sliceB)
	t.Log(r)
}
