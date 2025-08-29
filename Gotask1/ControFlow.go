/*
*只出现一次的数字
 */
package main

import "fmt"

////////
///136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，
// 其余每个元素均出现两次。找出那个只出现了一次的元素。可以使用 for 循环遍历数组，
// 结合 if 条件判断和 map 数据结构来解决，
// 例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
///

// ////
var testList = []int{1, 2, 3, 4, 5, 6, 6, 8, 6, 8, 1, 3, 4, 5}

func controflow(testList []int) {
	var testMap = make(map[int]int) // 记录元素出现相同次数的累计数
	processed := make(map[int]bool) // 记录已经处理过的元素状态

	for i, value1 := range testList {
		// 判断已经比对过的元素，则进行跳过循环
		if processed[value1] {
			continue
		}
		count := 1 // 至少出现一次（当前元素）
		testMap[value1] = count
		for j := i + 1; j < len(testList); j++ {
			value2 := testList[j]
			if value1 == value2 {
				count++
				testMap[value1] = count
				processed[value1] = true
			}
		}
	}
	//找出次数为1的值,并存在新的数组中
	onlyOneList := []int{}
	for mapKey, mapValue := range testMap {
		fmt.Println(mapKey, mapValue)
		if mapValue == 1 {
			onlyOneList = append(onlyOneList, mapKey)
		}
	}
	fmt.Println("只出现一次的元素是以下元素：", onlyOneList)
}
