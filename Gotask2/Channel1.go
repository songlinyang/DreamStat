package main

import (
	"fmt"
	"time"
)

// 接收channel数据 ch <- chan type
func receiveChannel(chOutput <-chan int, chInput chan<- int) {
	for {
		v := <-chOutput
		chInput <- v
	}
}

// 发送channel数据 ch chan <- type
func sendChannel(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("send data to channel:", i)
	}
	close(ch)
}

func main() {
	// 创建一个channel ,无缓冲
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 启动一个协程，发送数据到channel
	go sendChannel(ch1)
	// 启动一个协程，从channel接收数据
	go receiveChannel(ch1, ch2)
	for {
		select {
		case v, ok := <-ch2:
			if !ok {
				fmt.Println("Channel已关闭，退出程序")
				return
			}
			fmt.Println("主goroutine接收到数据:", v)
			time.Sleep(500 * time.Millisecond)
		case <-time.After(10 * time.Second):
			fmt.Println("10秒钟没有接收到数据，退出程序")
			return
		default:
			fmt.Println("没有数据，继续等待...")
			time.Sleep(200 * time.Millisecond)
		}

	}
}
