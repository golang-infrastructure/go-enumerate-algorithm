package enumerate_algorithm

import (
	"fmt"
)

func ExamplePermutation() {
	slice := []int{1, 2, 3}
	permutation := Permutation(slice)
	fmt.Println(permutation)
	// Output:
	// [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 2 1] [3 1 2]]
}
