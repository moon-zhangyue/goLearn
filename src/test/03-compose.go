package main

import (
	"fmt"
)

func main() {
	animal := Animal{Name: "中华田园犬"}
	pet := Pet{Name: "宠物狗"}
	dog := Dog{Animal: &animal, Pet: pet}
	fmt.Println(dog.Animal.GetName())
	fmt.Print(dog.Animal.Call())
	fmt.Println(dog.Call())
	fmt.Print(dog.Animal.FavorFood())
	fmt.Println(dog.FavorFood())
}
