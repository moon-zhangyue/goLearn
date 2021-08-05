package main

import "fmt"

func main() {
	var varI interface{} = 1

	v1 := varI.(int)
	fmt.Printf("v1: %T: %d\n", v1, v1)

	v2, ok := varI.(int)
	if ok {
		fmt.Printf("v2: %T: %d\n", v2, v2)
	}
}
