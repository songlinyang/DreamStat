/*
*回文数
 */
package main

////////
// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

// 例如，121 是回文，而 123 不是。

// 示例 1：

// 输入：x = 121
// 输出：true
// 示例 2：

// 输入：x = -121
// 输出：false
// 解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
// 示例 3：

// 输入：x = 10
// 输出：false
// 解释：从右向左读, 为 01 。因此它不是一个回文数。
// 提示：
// -2^31 <= x <= 2^31 - 1
// 进阶：你能不将整数转为字符串来解决这个问题吗？
// ////

func isPalindrome(x int) bool {
	if x > 10 {
		var intNumbers []int
		for {
			if x < 10 {
				intNumbers = append([]int{x}, intNumbers...)
				break
			} else {
				intNumbers = append([]int{x % 10}, intNumbers...)
				x /= 10
			}

		}
		//判断是否是回文数
		var aa string
		var bb string
		indexs := len(intNumbers) / 2
		if len(intNumbers)%2 != 0 {
			//是奇数，奇数判断前后
			for i := 0; i < indexs; i++ {
				aa += string(intNumbers[i])
			}
			for i := len(intNumbers) - 1; i > indexs; i-- {
				bb += string(intNumbers[i])
			}
			if aa == bb {
				return true
			} else {
				return false
			}
		} else {
			//是偶数
			for i := 0; i < indexs; i++ {
				aa += string(intNumbers[i])
			}
			for i := len(intNumbers) - 1; i >= indexs; i-- {
				bb += string(intNumbers[i])
			}
			if aa == bb {
				return true
			} else {
				return false
			}
		}
	} else if x >= 0 && x < 10 {
		return true //个位数判断 都是回文数
	} else {
		return false
	}
}
