package main

import (
	"fmt"
	"sync/atomic"
)

type Value struct {
	v interface{}
}

func main() {
	var i int32 = 1
	atomic.AddInt32(&i, 1)
	fmt.Println("i = i + 1 =", i)
	atomic.AddInt32(&i, -1)
	fmt.Println("i = i - 1 =", i)

	var a int32 = 1
	var b int32 = 2
	var c int32 = 2
	atomic.CompareAndSwapInt32(&a, a, b)
	atomic.CompareAndSwapInt32(&b, b, c)
	fmt.Println("a, b, c:", a, b, c)

	var x int32 = 100
	y := atomic.LoadInt32(&x)
	fmt.Println("x, y:", x, y)

	var m int32 = 100
	var n int32
	atomic.StoreInt32(&n, atomic.LoadInt32(&m))
	fmt.Println("m, n:", m, n)

	var j int32 = 1
	var k int32 = 2
	jOld := atomic.SwapInt32(&j, k)
	fmt.Println("old,new:", jOld, j)

	var v atomic.Value
	v.Store(100)
	fmt.Println("v:", v.Load())
}
