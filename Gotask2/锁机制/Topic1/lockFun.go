// 锁机制 Topic1
package main

import (
	"fmt"
	"sync"
)

// 线程安全锁
var lockCounter sync.Mutex

func counterFun(count *int64) {
	lockCounter.Lock()
	defer lockCounter.Unlock()
	*count++
}
func main444() {
	var countValue int64 = 0
	// 启动10个协程，每个协程对count加1000次
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < 1000; j++ {
				fmt.Printf("goroutine-%d:%d \n", i, j)
				counterFun(&countValue)
			}
		}()
	}
	wg.Wait()
	fmt.Println("count:", countValue)
}
