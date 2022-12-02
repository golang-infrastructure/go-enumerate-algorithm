# 排列、组合、笛卡尔积等等

# 一、排列
```go
slice := []int{1, 2, 3}
permutation := enumerate_algorithm.Permutation(slice)
fmt.Println(permutation)
// Output:
// [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 2 1] [3 1 2]]
```

# 二、组合

```go
slice := []int{1, 2, 3, 4, 5, 6}
combination := enumerate_algorithm.Combination(slice, 5)
fmt.Println(combination)
// Output:
// [[1 2 3 4 5] [1 2 3 4 6] [1 2 3 5 6] [1 2 4 5 6] [1 3 4 5 6] [2 3 4 5 6]]
```

# 三、笛卡尔积 
```go
sliceA := []string{"A", "B", "C"}
sliceB := []int{1, 2, 3}
r := enumerate_algorithm.Product(sliceA, sliceB)
t.Log(r)
// Output:
// [("A", 1) ("A", 2) ("A", 3) ("B", 1) ("B", 2) ("B", 3) ("C", 1) ("C", 2) ("C", 3)]
```
