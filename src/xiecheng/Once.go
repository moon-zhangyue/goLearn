package main

import (
	"fmt"
	"sync"
	"time"
)

func dosomething(o *sync.Once) {
	fmt.Println("Start:")
	o.Do(func() {
		fmt.Println("Do Something...")
	})
	fmt.Println("Finished.")
}

func main() {
	o := &sync.Once{}
	go dosomething(o)
	go dosomething(o)
	time.Sleep(time.Second * 1)
}
