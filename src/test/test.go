package main

import "fmt"

//type Integer int
//
////比较
//func (a Integer) Equal(b Integer) bool {
//	return a == b
//}
//
//// 加法
//func (a Integer) Add(b Integer) Integer {
//	return a + b
//}
//
//// 乘法
//func (a Integer) Multiply(b Integer) Integer {
//	return a * b
//}
//
//func main() {
//	var x Integer
//	var y Integer
//	x, y = 10, 15
//	fmt.Println(x.Equal(y))
//	fmt.Println(x.Add(y))
//}
/*type Student struct {
	id    uint
	name  string
	male  bool
	score float64
}

//func NewStudent(id uint, name string, male bool, score float64) *Student {
//	return &Student{id, name, male, score}
//}

func NewStudent(id uint, name string, score float64) *Student {
	return &Student{id: id, name: name, score: score}
}

//获取名字
func (s Student) GetName() string {
	return s.name
}

func (s *Student) SetName(name string) {
	s.name = name
}

/*func (s Student) String() string {
	return fmt.Sprintf("{id: %d, name: %s, male: %t, score: %f}",
		s.id, s.name, s.male, s.score)
}*/

/*func (s *Student) String() string {
	return "My name is :" + s.name
}

func main() {
	//student := NewStudent(1, "嘟嘟胖胖", true, 10.0)
	//fmt.Println(student)

	//student := NewStudent(1, "aaaa", 100)
	//fmt.Println("Name:", student.GetName())
	//fmt.Println(student)
	//student.SetName("bbbb")
	//fmt.Println("Name:", student.GetName())

	var s Student
	s.name = "haha"
	student := NewStudent(1, "aaaa", 100)
	fmt.Println(s)
	fmt.Println(student)
}*/

//类的继承
type Animal struct {
	Name string
}

//Animal的属性
func (a Animal) Call() string {
	return "动物的叫声..."
}
func (a Animal) FavorFood() string {
	return "爱吃的食物..."
}
func (a Animal) GetName() string {
	return a.Name
}

//type Dog struct {
//	Animal
//}

type Pet struct {
	Name string
}

//起别名
type Dog struct {
	animal *Animal
	pet    Pet
}

func (d Dog) FavorFood() string {
	return "骨头"
}

func (d Dog) Call() string {
	return "汪汪汪"
}

func main() {
	//animal := Animal{"中华田园犬"}
	//dog := Dog{animal}
	//fmt.Println(dog.GetName())
	//fmt.Println(dog.Call())
	//fmt.Println(dog.FavorFood())
	//
	////调用父类
	//fmt.Print(dog.Animal.Call())
	//fmt.Println(dog.Call())
	//fmt.Print(dog.Animal.FavorFood())
	//fmt.Println(dog.FavorFood())

	animal := Animal{"中华田园犬"}
	pet := Pet{"宠物狗"}
	dog := Dog{&animal, pet}
	// 通过 animal 引用 Animal 类型实例
	fmt.Println(dog.animal.GetName())
	fmt.Print(dog.animal.Call())
	fmt.Println(dog.Call())
	fmt.Print(dog.animal.FavorFood())
	fmt.Println(dog.FavorFood())
}
