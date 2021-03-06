package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

type error interface {
	Error() string
}

func Foo(param int) (n int, err error) {
	// ...
	return n, err
}

func add(a, b int) (c int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("只支持非负整数相加")
		return
	}
	a *= 2
	b *= 3
	c = a + b
	return
}

func main() {
	//n, err := Foo(0)
	//if err != nil {
	//	// 错误处理
	//	fmt.Println(err)
	//} else {
	//	// 使用返回值 n
	//	fmt.Println(n)
	//}

	/*if len(os.Args) != 3 {
		fmt.Printf("Usage: %s num1 num2\n", filepath.Base(os.Args[0]))
		return
	}
	x, _ := strconv.Atoi(os.Args[1])
	y, _ := strconv.Atoi(os.Args[2])
	// 通过多返回值捕获函数调用过程中可能的错误信息
	z, err := add(x, y)
	// 通过「卫述语句」处理后续业务逻辑
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("add(%d, %d) = %d\n", x, y, z)
	}*/

	// 获取指定路径文件信息，对应类型是 FileInfo
	// 如果文件不存在，则返回 PathError 类型错误
	fi, err := os.Stat("test.txt")
	if err != nil {
		switch err.(type) {
		case *os.PathError:
			// do something
			fmt.Println("文件不存在")
		case *os.LinkError:
			// dome something
		case *os.SyscallError:
			// dome something
		case *exec.Error:
			// dome something
		}
	} else {
		// ...
		fmt.Println(fi)
	}
}

type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}
