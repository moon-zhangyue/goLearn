package main

import "fmt"

//func main() {
//var v1 int      // 整型
//var v2 string   // 字符串
//var v3 bool     // 布尔型
//var v4 [10]int  // 数组，数组元素类型为整型
//var v5 struct { // 结构体，成员变量 f 的类型为64位浮点型
//	f float64
//}
//var v6 *int            // 指针，指向整型
//var v7 map[string]int  // map（字典），key为字符串类型，value为整型
//var v8 func(a int) int // 函数，参数类型为整型，返回值类型为整型
//
//fmt.Println("v1:", v1)
//fmt.Println("v2:", v2)
//fmt.Println("v3:", v3)
//fmt.Println("v4:", v4)
//fmt.Println("v5:", v5)
//fmt.Println("v6:", v6)
//fmt.Println("v7:", v7)
//fmt.Println("v8:", v8)

/*变量初始化
如果声明变量时想要同时对变量值进行初始化，可以通过以下这些方式：*/

//var v1 int = 10   // 方式一，常规的初始化操作
//var v2 = 10       // 方式二，此时变量类型会被编译器自动推导出来
//v3 := 10          // 方式三，可以省略 var，编译器可以自动推导出v3的类型
//
//fmt.Println("v1:", v1)
//fmt.Println("v2:", v2)
//fmt.Println("v3:", v3)

//若只想获得 nickName，则函数调用语句可以用如下方式实现：
//_, nickName := GetName()
//fmt.Printf("v%", nickName)
//fmt.Printf("s%", nickName)
//fmt.Println(nickName)
//}

//匿名变量
//我们来看个例子，假设 GetName() 函数的定义如下，它返回两个值，分别为 userName 和 nickName：
//func GetName() (userName, nickName string) {
//	return "nonfu", "学院君"
//}

//func main() {
//	const Pi float64 = 3.14159265358979323846
//	const zero = 0.0 // 无类型浮点常量
//	const (          // 通过一个 const 关键字定义多个常量，和 var 类似
//		size int64 = 1024
//		eof        = -1 // 无类型整型常量
//	)
//	const u, v float32 = 0, 3   // u = 0.0, v = 3.0，常量的多重赋值
//	const a, b, c = 3, 4, "foo" // a = 3, b = 4, c = "foo", 无类型整型和字符串常量
//
//}

func main() {
	//const ( // iota 被重置为 0
	//	c0 = iota // c0 = 0
	//	c1 = iota // c1 = 1
	//	c2 = iota // c2 = 2
	//)
	//const (
	//	u = iota * 2 // u = 0
	//	v = iota * 2 // v = 2
	//	w = iota * 2 // w = 4
	//)
	//const x = iota // x = 0
	//const y = iota // y = 0

	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, numberOfDays)
}
