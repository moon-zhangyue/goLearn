package main

import "fmt"

//func main() {
//	//animal := Animal{Name: "中华田园犬"}
//	//pet := Pet{Name: "宠物狗"}
//	//dog := Dog{Animal: &animal, Pet: pet}
//	//
//	//fmt.Println(dog.Animal.GetName())
//	//fmt.Print(dog.Animal.Call())
//	//fmt.Println(dog.Call())
//	//fmt.Print(dog.Animal.FavorFood())
//	//fmt.Println(dog.FavorFood())
//
//	animal := NewAnimal("中华田园犬")
//	pet := NewPet("宠物狗")
//	dog := NewDog(&animal, pet)
//
//	fmt.Println(dog.GetName())
//	fmt.Println(dog.Call())
//	fmt.Println(dog.FavorFood())
//}

//type A interface {
//	Foo()
//}
//type B interface {
//	A
//	Bar()
//}
//type T struct{}
//
//func (t T) Foo() {
//	fmt.Println("call Foo function from interface A.")
//}
//func (t T) Bar() {
//	fmt.Println("call Bar function from interface B.")
//}

//值传递
//type Integer int
//
//// 加法运算
//func (a Integer) Add(b Integer) Integer {
//	return a + b
//}
//
//// 乘法运算
//func (a Integer) Multiply(b Integer) Integer {
//	return a * b
//}
//
//type Math interface {
//	Add(i Integer) Integer
//	Multiply(i Integer) Integer
//}
//
//func main() {
//	var a Integer = 1
//	var m Math = a
//	fmt.Println(m.Add(1))
//
//	var b Integer = 1
//	var n Math = &b //传递指针
//	fmt.Println(n.Add(1))
//}

//指针传递
//type Integer int
//
//func (a *Integer) Add(b Integer) {
//	*a = (*a) + b
//}
//func (a Integer) Multiply(b Integer) Integer {
//	return a * b
//}
//
//type Math interface {
//	Add(i Integer)
//	Multiply(i Integer) Integer
//}
//
//func main() {
//	var a Integer = 1
//	var m Math = &a
//	m.Add(2)
//	fmt.Printf("1 + 2 = %d\n", a)
//}

type Number1 interface {
	Equal(i int) bool
	LessThan(i int) bool
	MoreThan(i int) bool
}

type Number2 interface {
	Equal(i int) bool
	MoreThan(i int) bool
	LessThan(i int) bool
	Add(i int)
}

type Number int

func (n Number) Equal(i int) bool {
	return int(n) == i
}

func (n Number) LessThan(i int) bool {
	return int(n) < i
}

func (n Number) MoreThan(i int) bool {
	return int(n) > i
}

func (n *Number) Add(i int) {
	*n = *n + Number(i)
}

var num1 Number = 1
var num2 Number2 = &num1
var num3 Number1 = num2

func main() {
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
}
