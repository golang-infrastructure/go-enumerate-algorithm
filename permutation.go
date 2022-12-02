package enumerate_algorithm

// Permutation 全排列
func Permutation[T any](slice []T) [][]T {
	return _permutation(slice, 0)
}

func _permutation[T any](slice []T, i int) [][]T {
	// 如果已经没有腾挪的余地了，则认为是OK了
	if i+1 == len(slice) {
		result := make([]T, len(slice))
		for index, value := range slice {
			result[index] = value
		}
		return [][]T{result}
	}
	result := make([][]T, 0)
	for j := i; j < len(slice); j++ {
		slice[j], slice[i] = slice[i], slice[j]
		result = append(result, _permutation(slice, i+1)...)
		slice[j], slice[i] = slice[i], slice[j]
	}
	return result
}
