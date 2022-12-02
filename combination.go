package enumerate_algorithm

// Combination 组合
func Combination[T any](slice []T, n int) [][]T {
	return _combination(slice, n, 0, make([]T, 0))
}

func _combination[T any](slice []T, n int, stepIndex int, selected []T) [][]T {
	// 选够元素了
	if len(selected) == n {
		result := make([]T, 0)
		result = append(result, selected...)
		return [][]T{result}
	}
	// 到达边界了
	if stepIndex >= len(slice) {
		return [][]T{}
	}
	result := make([][]T, 0)
	// 选中当前元素
	selected = append(selected, slice[stepIndex])
	result = append(result, _combination(slice, n, stepIndex+1, selected)...)
	// 不选中当前元素
	selected = selected[0 : len(selected)-1]
	result = append(result, _combination(slice, n, stepIndex+1, selected)...)
	return result
}
