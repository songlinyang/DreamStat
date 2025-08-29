/*
* 有效的括号字符串 解题
 */
package main

///
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

// 有效字符串需满足：

// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。

type Stack struct {
	items []byte
}

// 压入栈
func (s *Stack) Push(item byte) {
	s.items = append(s.items, item)
}

// 出栈
func (s *Stack) Pop() byte {
	if len(s.items) == 0 {
		return 0
	}
	// 后进先出
	item := s.items[len(s.items)-1]

	s.items = s.items[:len(s.items)-1]
	return item
}

// 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func isMatch(left, right byte) bool {
	switch right {
	case ')':
		return left == '('
	case ']':
		return left == '['
	case '}':
		return left == '{'
	default:
		return false
	}
}

func isValid(s string) bool {
	stack := Stack{}

	for i := 0; i < len(s); i++ {
		char := s[i]
		if char == '(' || char == '[' || char == '{' {
			stack.Push(char)
		} else {
			if stack.IsEmpty() {
				return false // 栈为空，说明右括号没有对应的左括号
			}

			top := stack.Pop()
			if !isMatch(top, char) {
				return false
			}
		}
	}
	// 最后检查栈是否为空，不为空说明有左括号没有被匹配
	return stack.IsEmpty()

}

// /
// func main() {
// 	//'('，')'，'{'，'}'，'['，']'
// 	var inputString string = "{[()]}"
// 	result := isValid(inputString)
// 	fmt.Println(result)
// }
