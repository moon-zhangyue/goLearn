package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)
	// 发送方
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("发送方: 发送数据 %v...\n", i)
			ch <- i
		}
		fmt.Println("发送方: 关闭通道...")
		close(ch)
	}()
	// 接收方
	for {
		time.Sleep(5 * time.Second) //加上超时 可以看出缓冲容量的作用
		num, ok := <-ch
		if !ok {
			fmt.Println("接收方: 通道已关闭")
			break
		}
		fmt.Printf("接收方: 接收数据: %v\n", num)
	}
	fmt.Println("程序退出")
}
