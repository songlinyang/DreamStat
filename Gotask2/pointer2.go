package main

import "fmt"

// 将切片中的每个元素乘以2
func slicePointSetValue(s *[]int) *[]int {
	for i := range *s {
		(*s)[i] *= 2
	}
	return s
}

func main12() {
	nums := []int{1, 2, 3, 4}
	answer := slicePointSetValue(&nums)
	fmt.Println(*answer) // 输出: [2 4 6 8]
}
