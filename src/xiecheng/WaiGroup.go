package main

import (
	"fmt"
	"sync"
)

func add_num(a, b int, deferFunc func()) {
	defer func() {
		deferFunc()
	}()
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go add_num(i, 1, wg.Done)
	}
	wg.Wait()
}
