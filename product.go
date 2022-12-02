package enumerate_algorithm

import "github.com/golang-infrastructure/go-tuple"

// Product 计算两个切片的笛卡尔积
func Product[T1, T2 any](sliceA []T1, sliceB []T2) []*tuple.Tuple2[T1, T2] {
	result := make([]*tuple.Tuple2[T1, T2], len(sliceA)*len(sliceB))
	index := 0
	for _, valueA := range sliceA {
		for _, valueB := range sliceB {
			result[index] = tuple.New2(valueA, valueB)
			index++
		}
	}
	return result
}
