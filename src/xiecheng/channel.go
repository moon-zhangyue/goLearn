package main

import (
	"fmt"
	"time"
)

func add(a, b int, ch chan int) {
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
	ch <- 1
}

func main() {
	start := time.Now()
	chs := make([]chan int, 10)

	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go add(1, i, chs[i])
	}

	for _, ch := range chs {
		<-ch
	}

	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}
