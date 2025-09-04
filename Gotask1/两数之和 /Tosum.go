/*
* 两数之和 解题
 */
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	var result int
	var i1, j1 int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			result = nums[i] + nums[j]
			if result == target {
				i1 = i
				j1 = j
				break
			}
		}
		if j1 != 0 {
			break
		}
	}
	return []int{i1, j1}

}

func main() {
	aa := twoSum([]int{0, 4, 3, 0}, 0)
	fmt.Println(aa)
}
