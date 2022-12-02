package enumerate_algorithm

import (
	"fmt"
)

func ExampleCombination() {
	slice := []int{1, 2, 3, 4, 5, 6}
	combination := Combination(slice, 5)
	fmt.Println(combination)
	// Output:
	// [[1 2 3 4 5] [1 2 3 4 6] [1 2 3 5 6] [1 2 4 5 6] [1 3 4 5 6] [2 3 4 5 6]]
}
