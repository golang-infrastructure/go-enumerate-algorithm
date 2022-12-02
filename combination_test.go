package enumerate_algorithm

import (
	"fmt"
	"testing"
)

func TestCombination(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}
	combination := Combination(slice, 5)
	fmt.Print(combination)
}
