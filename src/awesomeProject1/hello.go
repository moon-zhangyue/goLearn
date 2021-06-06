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

	//const (
	//	Sunday = iota
	//	Monday
	//	Tuesday
	//	Wednesday
	//	Thursday
	//	Friday
	//	Saturday
	//	numberOfDays
	//)
	//fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, numberOfDays)

	//布尔类型
	//Go 语言中的布尔类型与其他主流编程语言差不多，类型关键字为 bool，可赋值且只可以赋值为预定义常量 true 和 false。示例代码如下：

	//var v1 bool
	//v1 = true
	//v2 := (1 == 2) // v2 也会被推导为 bool 类型

	//println(v1, v2)

	//不过通过表达式计算得到的布尔类型结果可以赋值给 Go 布尔类型变量：
	//var b bool
	//b = (1 != 0)              // 编译正确
	//fmt.Println("Result:", b) // 打印结果为Result: true

	//b := (false == 0)
	//println(b)

	//intValue := 8
	//if intValue == 8 {
	//	fmt.Printf("OK")
	//}

	//var intValueBit uint8
	//intValueBit = 255
	//intValueBit = ^intValueBit // 按位取反
	//fmt.Println(intValueBit)   // 0
	//intValueBit = 1
	//intValueBit = intValueBit << 3 // 左移 3 位，相当于乘以 2^3 = 8
	//fmt.Println(intValueBit)       // 8

	//floatValue4 := 0.1
	//floatValue5 := 0.7
	//floatValue6 := floatValue4 + floatValue5
	//
	//fmt.Println(floatValue6)
	//var floatValue1 float32
	//floatValue1 = 10
	//floatValue2 := 10.0
	//
	//p := 0.00001
	//// 判断 floatValue1 与 floatValue2 是否相等
	//if math.Dim(float64(floatValue1), floatValue2) < p {
	//	fmt.Println("floatValue1 和 floatValue2 相等")
	//}

	/*var str string      // 声明字符串变量
	str = "Hello World" // 变量初始化
	//str2 := "你好，学院君"    // 也可以同时进行声明和初始化
	ch := str[0] // 取字符串的第一个字符
	//还可以通过 Go 语言内置的len()函数获取指定字符串的长度，以及通过fmt包提供的Printf进行字符串格式化输出：
	fmt.Printf("The length of \"%s\" is %d \n", str, len(str))
	fmt.Printf("The first character of \"%s\" is %c.\n", str, ch)*/

	results := `Search results for "Golang":
		- Go
		- Golang
		Golang Programming
		`
	fmt.Printf("%s", results)

	//str := "hello, world"
	//str1 := str[:5]  // 获取索引5（不含）之前的子串
	//str2 := str[7:]  // 获取索引7（含）之后的子串
	//str3 := str[0:5] // 获取从索引0（含）到索引5（不含）之间的子串
	//fmt.Println("str1:", str1)
	//fmt.Println("str2:", str2)
	//fmt.Println("str3:", str3)

	str := "Hello, 世界"
	for i, ch := range str {
		fmt.Println(i, ch) // ch 的类型为 rune
	}
}
