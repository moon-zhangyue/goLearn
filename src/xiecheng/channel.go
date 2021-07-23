package main

import (
	"fmt"
	"time"
)

//func add(a, b int, ch chan int) {
//	c := a + b
//	fmt.Printf("%d + %d = %d\n", a, b, c)
//	ch <- 1
//}
//
//func main() {
//	start := time.Now()
//	chs := make([]chan int, 10)
//
//	for i := 0; i < 10; i++ {
//		chs[i] = make(chan int)
//		go add(1, i, chs[i])
//	}
//
//	for _, ch := range chs {
//		<-ch
//	}
//
//	end := time.Now()
//	consume := end.Sub(start).Seconds()
//	fmt.Println("程序执行耗时(s)：", consume)
//}

//缓冲通道
//func test(ch chan int) {
//	for i := 0; i < 100; i++ {
//		ch <- i
//	}
//	close(ch) //关闭通道
//}
//
//func main() {
//	start := time.Now()
//	ch := make(chan int, 20) //缓冲通道
//	//ch := make(chan int) //非缓冲通道
//	go test(ch)
//	for i := range ch {
//		fmt.Println("接收到的数据:", i)
//	}
//	end := time.Now()
//	consume := end.Sub(start).Seconds()
//	fmt.Println("程序执行耗时(s)：", consume)
//}

//单向通道
func test(ch chan<- int) {
	//func test(ch <-chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

//单项只读通道
func test1() <-chan int {
	ch := make(chan int, 20)
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
	return ch
}

func main() {
	start := time.Now()
	ch := make(chan int, 20)
	go test(ch)
	for i := range ch {
		fmt.Println("接收到的数据:", i)
	}
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}
