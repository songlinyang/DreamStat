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
func (ts *TaskScheduler) addTask(name string, task func()) {
	ts.mutex.Lock()
	defer ts.mutex.Unlock()
	if ts.taskMap == nil {
		ts.taskMap = make(map[string]func())
	}else{
		go func ()  {
			// TODO
		}
	}
}

func runFuncTotalTime(f func()) {
	start := time.Now()
	f()
	end := time.Now()
	fmt.Println("func runFuncTotalTime use time:", end.Sub(start))
}
