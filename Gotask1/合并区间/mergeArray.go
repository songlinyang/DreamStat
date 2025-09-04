package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	// 1. 按照区间的起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	//合并数组，放到切片
	merge := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		lastMerged := merge[len(merge)-1]

		// 检查当前区间是否与最后一个合并区间重叠
		if current[0] <= lastMerged[1] {
			// 有重叠，合并区间（取结束位置的较大值）
			if current[1] > lastMerged[1] {
				lastMerged[1] = current[1]
			}
		} else {
			// 没有重叠，直接添加到结果中
			merge = append(merge, current)
		}
	}

	return merge

}
func main() {
	a := [][]int{{4, 5}, {1, 4}, {0, 1}}
	b := merge(a)
	fmt.Println(b)
}
