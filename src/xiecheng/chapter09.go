package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var counter int = 0

func add(a, b int, lock *sync.Mutex) {
	c := a + b
	lock.Lock()
	counter++
	fmt.Printf("%d: %d + %d = %d\n", counter, a, b, c)
	lock.Unlock()
}

func main() {
	start := time.Now()
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go add(1, i, lock)
	}

	for {
		lock.Lock() //防止读取阶段被其他并发操作更新
		c := counter
		lock.Unlock()
		runtime.Gosched() //用于让出CPU时间片，让出当前goroutine的执行权限，调度器安排其它等待的任务运行，并在下次某个时候从该位置恢复执行
		if c >= 10 {
			break
		}
	}
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}
