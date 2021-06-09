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
	//
	//results := `Search results for "Golang":
	//	- Go
	//	- Golang
	//	Golang Programming
	//	`
	//fmt.Printf("%s", results)

	//str := "hello, world"
	//str1 := str[:5]  // 获取索引5（不含）之前的子串
	//str2 := str[7:]  // 获取索引7（含）之后的子串
	//str3 := str[0:5] // 获取从索引0（含）到索引5（不含）之间的子串
	//fmt.Println("str1:", str1)
	//fmt.Println("str2:", str2)
	//fmt.Println("str3:", str3)

	//str := "Hello, 世界"
	//for i, ch := range str {
	//	fmt.Println(i, string(ch)) // ch 的类型为 rune 将 Unicode 编码转化为可打印字符string
	//}

	//v1 := uint(-255)

	//v1 := uint(255)
	//v2 := int8(v1) // v2 = -1

	//v1 := 99.99
	//v2 := int(v1) // v2 = 99
	//
	//fmt.Println(v1)
	//fmt.Println(v2)
	////将整型转化为浮点型时，比较简单，直接调用对应的函数即可：
	//
	//v3 := 99
	//v4 := float64(v2) // v4 = 99
	//
	//fmt.Println(v3)
	//fmt.Println(v4)

	//v1 := 65
	//v2 := string(v1)  // v2 = A
	//v3 := 30028
	//v4 := string(v3)  // v4 = 界
	//
	//fmt.Println(v2)
	//fmt.Println(v4)

	//v1 := []byte{'h', 'e', 'l', 'l', 'o'}
	//v2 := string(v1) // v2 = hello
	//v3 := []rune{0x5b66, 0x9662, 0x541b}
	//v4 := string(v3) // v4 = 学院君
	//fmt.Println(v2)
	//fmt.Println(v4)

	//v1 := "100"
	//v2, _ := strconv.Atoi(v1) // 将字符串转化为整型，v2 = 100
	//v3 := 100
	//v4 := strconv.Itoa(v3) // 将整型转化为字符串, v4 = "100"
	//v5 := "true"
	//v6, _ := strconv.ParseBool(v5) // 将字符串转化为布尔型
	//v5 = strconv.FormatBool(v6)    // 将布尔值转化为字符串
	//v7 := "100"
	//v8, _ := strconv.ParseInt(v7, 10, 64)  // 将字符串转化为整型，第二个参数表示进制，第三个参数表示最大位数
	//v7 = strconv.FormatInt(v8, 10)         // 将整型转化为字符串，第二个参数表示进制
	//v9, _ := strconv.ParseUint(v7, 10, 64) // 将字符串转化为无符号整型，参数含义同 ParseInt
	//v7 = strconv.FormatUint(v9, 10)        // 将字符串转化为无符号整型，参数含义同 FormatInt
	//v10 := "99.99"
	//v11, _ := strconv.ParseFloat(v10, 64) // 将字符串转化为浮点型，第二个参数表示精度
	//v10 = strconv.FormatFloat(v11, 'E', -1, 64)
	//q := strconv.Quote("Hello, 世界")       // 为字符串加引号
	//q = strconv.QuoteToASCII("Hello, 世界") // 将字符串转化为 ASCII 编码
	//
	//fmt.Println(v2)
	//fmt.Println(v4)
	//fmt.Println(v5)
	//fmt.Println(v6)
	//fmt.Println(v7)
	//fmt.Println(v8)
	//fmt.Println(v9)
	//fmt.Println(v11)
	//fmt.Println(v10)
	//fmt.Println(q)

	//var a [8]byte           // 长度为8的数组，每个元素为一个字节
	//var b [3][3]int         // 二维数组（9宫格）
	//var c [3][3][3]float64  // 三维数组（立体的9宫格）
	//var d = [3]int{1, 2, 3} // 声明时初始化
	//var e = new([3]string)  // 通过 new 初始化
	//
	//m := [5]int{1, 2, 3}
	//fmt.Println(m)
	//
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)
	//fmt.Println(d)
	//fmt.Println(e)
	//
	//var arr = [5]int{1, 2, 3, 4, 5}
	//for i, v := range arr {
	//	fmt.Println("Element", i, "of arr is", v)
	//}

	// 通过二维数组生成九九乘法表
	//var multi [9][9]string
	//for j := 0; j < 9; j++ {
	//	for i := 0; i < 9; i++ {
	//		n1 := i + 1
	//		n2 := j + 1
	//		if n1 < n2 { // 摒除重复的记录
	//			continue
	//		}
	//		multi[i][j] = fmt.Sprintf("%dx%d=%d", n2, n1, n1*n2)
	//	}
	//}
	//// 打印九九乘法表
	//for _, v1 := range multi {
	//	for _, v2 := range v1 {
	//		fmt.Printf("%-8s", v2) // 位宽为8，左对齐
	//	}
	//	fmt.Println()
	//}

	//var slice []string = []string{"a", "b", "c"}

	//fmt.Println(slice)

	// 先定义一个数组
	//months := [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	//// 基于数组创建切片
	//q2 := months[3:6]     // 第二季度
	//summer := months[5:8] // 夏季
	//fmt.Println(q2)
	//fmt.Println(summer)
	//
	//fmt.Println(len(q2)) // 3
	//fmt.Println(cap(q2)) // 9
	//
	//firsthalf := months[:6]
	//q1 := firsthalf[:3] // 基于 firsthalf 的前 3 个元素构建新切片
	//
	//fmt.Println(q1)
	//fmt.Println(firsthalf)
	//
	//q3 := firsthalf[:9]
	//
	//fmt.Println(q3)
	//
	//mySlice := make([]int, 5)
	//mySlice1 := make([]int, 5, 10)
	//
	//fmt.Println(mySlice)
	//fmt.Println(mySlice1)
	//
	//aa := []int{1, 2, 3, 4, 5}
	//fmt.Println(aa)
	//
	//for i := 0; i < len(summer); i++ {
	//	fmt.Println("summer[", i, "]", summer[i])
	//}
	//
	//for i, v := range summer {
	//	fmt.Println(i, v)
	//}

	//var Slice = make([]int, 5, 10)
	//
	//fmt.Println((Slice))
	//fmt.Println(len(Slice))
	//fmt.Println(cap(Slice))
	//
	////newSlice := append(Slice, 1, 2, 3)
	////fmt.Println(newSlice)
	//
	//newSlice1 := []int{1, 2, 3, 4, 5}
	//newSlice := append(Slice, newSlice1...) //...不可省略
	//fmt.Println(newSlice)
	//
	//slice1 := []int{1, 2, 3, 4, 5}
	//slice2 := []int{5, 4, 3}
	//// 复制 slice1 到 slice 2
	////copy(slice2, slice1) // 只会复制 slice1 的前3个元素到 slice2 中
	//// slice2 结果: [1, 2, 3]
	//fmt.Println(slice2)
	//// 复制 slice2 到 slice 1
	//copy(slice1, slice2) // 只会复制 slice2 的 3 个元素到 slice1 的前 3 个位置
	//// slice1 结果：[5, 4, 3, 4, 5]
	//fmt.Println(slice1)

	//slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//slice3 = slice3[:len(slice3)-5] // 删除 slice3 尾部 5 个元素
	//slice3 = slice3[5:]             // 删除 slice3 头部 5 个元素
	//fmt.Println(slice3)
	//fmt.Println(len(slice3))
	//fmt.Println(cap(slice3))

	//slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//fmt.Println(slice3[:0])
	//slice4 := append(slice3[:0], slice3[3:]...) // 删除开头三个元素
	//slice5 := append(slice3[:1], slice3[4:]...) // 删除中间三个元素
	//slice6 := append(slice3[:0], slice3[:7]...) // 删除最后三个元素
	//slice7 := slice3[:copy(slice3, slice3[3:])] // 删除开头前三个元素

	//slice1 := []int{1, 2, 3, 4, 5}
	//slice2 := slice1[1:3]
	//slice2[1] = 6
	//fmt.Println("slice1:", slice1)
	//fmt.Println("slice2:", slice2)

	var testMap map[string]int
	testMap = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	k := "two"
	v, ok := testMap[k]
	if ok {
		fmt.Printf("The element of key %q: %d\n", k, v)
	} else {
		fmt.Println("Not found!")
	}
}
