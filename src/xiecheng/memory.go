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
	lock := &sync.Mutex{} //锁机制
	for i := 0; i < 10; i++ {
		go add(1, i, lock)
	}

	for {
		lock.Lock()
		c := counter
		lock.Unlock()//因为add()里的counter和main的counter是同一个变量，add()和main()是并行运行的，所以在main中对counter进行处理的话也要加锁。这里并没有做什么处理，所以这里的锁没必要加。
		runtime.Gosched() //电脑如果是单核 CPU，那么到 for 循环这里就会堵塞卡死，Gosched() 的作用是遇到堵塞情况交给同线程下的其他协程继续处理当前任务。
		if c >= 10 {
			break
		}
	}
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}
