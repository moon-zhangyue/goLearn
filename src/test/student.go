package main

import "fmt"

type Student struct {
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

//func (s Student) String() string {
//	return fmt.Sprintf("{id: %d, name: %s, male: %t, score: %f}",
//		s.id, s.name, s.male, s.score)
//}

func (s Student) String() string {
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
	s.name = "hello"
	fmt.Println(s)
	s.SetName("nihao")
	fmt.Println(s)
	//student := NewStudent(1, "aaaa", 100) //初始化
	//fmt.Println(student)
}
