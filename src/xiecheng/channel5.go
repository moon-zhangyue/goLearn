package main

import (
	"fmt"
	"time"
)

func main() {
	// 初始化 ch 通道
	ch := make(chan int, 1)

	// 初始化 timeout 通道
	timeout := make(chan bool, 1)

	// 实现一个匿名超时等待函数
	go func() {
		time.Sleep(1e9) // 睡眠1秒钟
		timeout <- true
	}()

	// 借助 timeout 通道结合 select 语句实现 ch 通道读取超时效果
	select {
	case <-ch:
		fmt.Println("接收到 ch 通道数据")
	case <-timeout:
		fmt.Println("超时1秒，程序退出")
	}
}
