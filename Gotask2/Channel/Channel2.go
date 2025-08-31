// Channel题目1
package main

import (
	"fmt"
	"time"
)

// 接收channel数据 ch <- chan type
func receiveChannel2(ch <-chan int) {
	for v := range ch {
		fmt.Println("receive data from channel:", v)
	}
}

// 发送channel数据 ch chan <- type
func sendChannel2(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("send data to channel:", i)
	}
	close(ch)
}

func main() {
	// 创建一个channel ,
	ch := make(chan int, 100)
	// 启动一个协程，发送数据到channel
	go sendChannel2(ch)
	// 启动一个协程，从channel接收数据
	go receiveChannel2(ch)
	for {
		time.Sleep(500 * time.Millisecond)
		v, ok := <-ch
		if !ok {
			fmt.Println("Channel已关闭，退出程序")
			return
		}
		fmt.Println("主goroutine接收到数据:", v)

	}
}
