// Goroutine题目1
package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex

func yushu(i int) int {
	lock.Lock()
	defer lock.Unlock()
	return i % 2
}

func main1111() {
	// var wg sync.WaitGroup
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			number := i + 1
			time.Sleep(100 * time.Millisecond)
			if yushu(number) == 0 {
				fmt.Println("goroutine协程1->偶数：", number)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			number := i + 1
			time.Sleep(100 * time.Millisecond)
			if yushu(number) != 0 {
				fmt.Println("goroutine协程2->奇数: ", number)
			}
		}
	}()

	wg.Wait()
	fmt.Println("main goroutine done")
	// say("hello")
}
