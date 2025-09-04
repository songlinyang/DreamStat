// 指针题目1
package main

import "fmt"

var value int = 10
var golablPointer *int = &value

// Topic1
func pointerOperation(a *int) {
	fmt.Printf("in func pointerOperation a=%p \n", a)
	*a += 10
	fmt.Println("in func pointerOperation a=", *a)
}

func main() {
	fmt.Printf("before golbalPointer pointer p=%p\n", &golablPointer)
	pointerOperation(golablPointer)
	fmt.Printf("after golbalPointer pointer p=%p\n", &golablPointer)
	fmt.Println("glbalPointer value = ", *golablPointer)
}
