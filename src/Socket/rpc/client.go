package main

import (
	"fmt"
	"log"
	"net/rpc"
	"rpc/utils"
)

func main() {
	var serverAddress = "localhost" //服务端地址

	client, err := rpc.DialHTTP("tcp", serverAddress+":8080")
	if err != nil {
		log.Fatal("建立与服务端连接失败:", err)
	}

	args := &utils.Args{10, 10} //接收参数 传递指针
	var reply int
	err = client.Call("MathService.Multiply", args, &reply) //同步调用
	if err != nil {
		log.Fatal("调用远程方法 MathService.Multiply 失败:", err)
	}
	fmt.Printf("%d*%d=%d\n", args.A, args.B, reply)

	divideCall := client.Go("MathService.Divide", args, &reply, nil) //异步方法调用
	for {
		select {
		case <-divideCall.Done:
			fmt.Printf("%d/%d=%d\n", args.A, args.B, reply)
			return
		}
	}
}
