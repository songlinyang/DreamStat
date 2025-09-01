package main

import (
	"fmt"
	"sync"
	"time"
)

// 定义接口类型
type schedulerInterf interface {
	//添加任务
	addTask(string, func())
	//执行任务
	executeTask()
	//清空任务
	clearTask()
}

// 定义结构体类型
type TaskScheduler struct {
	taskMap   map[string]func()
	mutex     sync.RWMutex
	waitGroup sync.WaitGroup
}

// 实现接口方法
// 添加任务列表
func (ts *TaskScheduler) addTask(name string, task func()) {
	ts.mutex.Lock()
	defer ts.mutex.Unlock()
	fmt.Print("add task:", name, "\n")
	if ts.taskMap == nil {
		ts.taskMap = make(map[string]func())
	}

	//添加任务
	ts.taskMap[name] = task
}

// 执行任务
func (ts *TaskScheduler) executeTask() {
	ts.mutex.RLock()
	defer ts.mutex.RUnlock()
	for name, task := range ts.taskMap {
		ts.waitGroup.Add(1)
		fmt.Print("execute task:", name, "\n")
		//执行任务
		go func() {
			defer ts.waitGroup.Done()
			runFuncTotalTime(name, task)
		}()
	}
	ts.waitGroup.Wait()
}

// 清空任务
func (ts *TaskScheduler) clearTask() {
	ts.mutex.Lock()
	defer ts.mutex.Unlock()
	fmt.Print("clear task\n")
	//清空任务
	ts.taskMap = make(map[string]func())
}

// 计算函数执行时间
func runFuncTotalTime(name string, f func()) {
	start := time.Now()
	f()
	end := time.Now()
	fmt.Println("Task "+name+" use time:", end.Sub(start))
}

// 测试
func Test(scheduler schedulerInterf) {
	for i := 0; i < 10; i++ {
		//添加任务
		scheduler.addTask(fmt.Sprint("task", i), func() {
			time.Sleep(1 * time.Second)
			fmt.Printf("task%d done!\n", i)
		})
	}
	//执行任务
	scheduler.executeTask()
}

func main() {
	var ts TaskScheduler = TaskScheduler{}
	var scheduler schedulerInterf = &ts
	Test(scheduler)
}

// 学习点：
// 1.接口的定义和实现
// 2.结构体的定义和方法实现
// 3.读写锁的使用
// 4.等待组的使用
// 5.函数作为参数传递和匿名函数的使用
// 6.计算函数执行时间
// 7.ts.waitGroup.Done()的使用，以及如何防止死锁操作
