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

func main() {
	// var wg sync.WaitGroup
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			time.Sleep(100 * time.Millisecond)
			if yushu(i) == 0 {
				fmt.Println("goroutine协程1->偶数：", i)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			time.Sleep(100 * time.Millisecond)
			if yushu(i) != 0 {
				fmt.Println("goroutine协程2->奇数: ", i)
			}
		}
	}()

	wg.Wait()
	fmt.Println("main goroutine done")
	// say("hello")
}
