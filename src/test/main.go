package main

import (
	"fmt"
	"reflect"
)

//func add(a, b int) int {
//	a *= 2
//	b *= 3
//	return a + b
//}
//func main() {
//	x, y := 1, 2
//	z := add(x, y)
//	fmt.Printf("add(%d, %d) = %d\n", x, y, z)
//}

//func add(a, b *int) int {
//	*a *= 2
//	*b *= 3
//	return *a + *b
//}

//func myfunc(numbers ...int) {
//	for _, number := range numbers {
//		fmt.Println(number)
//	}
//}

func myPrintf(args ...interface{}) {
	for _, arg := range args {
		switch reflect.TypeOf(arg).Kind() {
		case reflect.Int:
			fmt.Println(arg, "is an int value.")
		case reflect.String:
			fmt.Printf("\"%s\" is a string value.\n", arg)
		case reflect.Array:
			fmt.Println(arg, "is an array type.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

func main() {
	//x, y := 1, 2
	//z := add(&x, &y)
	//fmt.Printf("add(%d, %d) = %d\n", x, y, z)

	//myfunc(1, 12, 3, 45)

	//slice := []int{1, 2, 3, 4, 5}
	//myfunc(slice...)
	//myfunc(slice[1:3]...)

	myPrintf(1, "1", [1]int{1}, true)
}
