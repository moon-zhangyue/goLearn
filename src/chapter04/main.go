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

type Integer int

// 加法运算
func (a Integer) Add(b Integer) Integer {
	return a + b
}

// 乘法运算
func (a Integer) Multiply(b Integer) Integer {
	return a * b
}

type Math interface {
	Add(i Integer) Integer
	Multiply(i Integer) Integer
}

func main() {
	var a Integer = 1
	var m Math = a
	fmt.Println(m.Add(1))

	var b Integer = 1
	var n Math = &b //传递指针
	fmt.Println(n.Add(1))
}