package enumerate_algorithm

import (
	"fmt"
)

func ExampleProduct() {
	sliceA := []string{"A", "B", "C"}
	sliceB := []int{1, 2, 3}
	r := Product(sliceA, sliceB)
	fmt.Println(r)
	// Output:
	// [("A", 1) ("A", 2) ("A", 3) ("B", 1) ("B", 2) ("B", 3) ("C", 1) ("C", 2) ("C", 3)]
}
