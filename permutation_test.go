package enumerate_algorithm

import (
	"fmt"
	"testing"
)

func TestPermutation(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	permutation := Permutation(slice)
	fmt.Println(permutation)
}
