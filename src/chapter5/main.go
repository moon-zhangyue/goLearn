package main

import "fmt"

type error interface {
	Error() string
}

func Foo(param int) (n int, err error) {
	// ...
	return n, err
}

func main() {
	n, err := Foo(0)
	if err != nil {
		// 错误处理
		fmt.Println(err)
	} else {
		// 使用返回值 n
		fmt.Println(n)
	}
}
