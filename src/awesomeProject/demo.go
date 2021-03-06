//demo测试
package main // 声明 main 包
import (
	. "awesomeProject/animal"
	"fmt"
)

//func main() { // 声明 main 主函数
//fmt.Println("Hello World!") // 打印 Hello World!

/*var attack = 40
var defence = 20
var damageRate float32 = 0.17
var damage = float32(attack-defence) * damageRate
fmt.Println(damage)*/

/*var a int = 100
var b int = 200

b, a = a, b

fmt.Println(a, b)*/
//}

/*func GetData() (int, int) {
	return 100, 200
}
func main(){
	//匿名变量，可赋值，但会被抛弃，无法使用，不占用空间
	a, _ := GetData()
	_, b := GetData()
	fmt.Println(a, b)
}*/

//局部变量
/*func main() {
	//声明局部变量 a和b并赋值
	var a int = 3
	var b int = 4

	//声明局部变量 c 并计算 a 和 b 的和
	c := a + b
	fmt.Print("a=%d,b=%d,c=%d\n", a, b, c)
}*/

//全局变量
//声明全局变量
/*var c int

func main() {
	//声明局部变量
	var a, b int

	//初始化参数
	a = 3
	b = 4
	c = a + b

	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
}*/

//形参
//全局变量
/*var a int = 13

func main() {
	//局部变量 a和b
	var a int = 3
	var b int = 4

	fmt.Printf("main()函数中 a = %d\n", a)
	fmt.Printf("main()函数中 b = %d\n", b)
	c := sum(a, b)
	fmt.Printf("main() 函数中 c = %d\n", c)
}

func sum(a, b int) int {
	fmt.Printf("sum() 函数中 a = %d\n", a)
	fmt.Printf("sum() 函数中 b = %d\n", b)
	num := a + b
	return num
}*/

/*//复数
func main() {
	var x complex128 = complex(1, 2) // 1+2i
	var y complex128 = complex(3, 4) // 3+4i
	fmt.Println(x * y)               // "(-5+10i)"
	fmt.Println(real(x * y))         // "-5"  实部
	fmt.Println(imag(x * y))         // "10"  虚部
}*/

//正弦函数图像
/*func main() {

	// 图片大小
	const size = 300
	// 根据给定大小创建灰度图
	pic := image.NewGray(image.Rect(0, 0, size, size))

	// 遍历每个像素
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			// 填充为白色
			pic.SetGray(x, y, color.Gray{255})
		}
	}

	// 从0到最大像素生成x坐标
	for x := 0; x < size; x++ {

		// 让sin的值的范围在0~2Pi之间
		s := float64(x) * 2 * math.Pi / size

		// sin的幅度为一半的像素。向下偏移一半像素并翻转
		y := size/2 - math.Sin(s)*size/2

		// 用黑色绘制sin轨迹
		pic.SetGray(x, int(y), color.Gray{0})
	}

	// 创建文件
	file, err := os.Create("sin.png")

	if err != nil {
		log.Fatal(err)
	}
	// 使用png格式将数据写入文件
	png.Encode(file, pic) //将image信息写入文件中

	// 关闭文件
	file.Close()
}*/

//拼接字符串
/*func main() {
	s := "hel" + "lo,"
	s += "world!"
	fmt.Println(s) //输出 “hello, world!”
}*/

/*func main() {

	// 输出各数值范围
	fmt.Println("int8 range:", math.MinInt8, math.MaxInt8)
	fmt.Println("int16 range:", math.MinInt16, math.MaxInt16)
	fmt.Println("int32 range:", math.MinInt32, math.MaxInt32)
	fmt.Println("int64 range:", math.MinInt64, math.MaxInt64)

	// 初始化一个32位整型值
	var a int32 = 1047483647
	// 输出变量的十六进制形式和十进制值
	fmt.Printf("int32: 0x%x %d\n", a, a)

	// 将a变量数值转换为十六进制, 发生数值截断
	b := int16(a)
	// 输出变量的十六进制形式和十进制值
	fmt.Printf("int16: 0x%x %d\n", b, b)

	// 将常量保存为float32类型
	var c float32 = math.Pi
	// 转换为int类型, 浮点发生精度丢失
	fmt.Println(int(c))
}*/

//指针
/*func main() {
	var cat int = 1
	var str string = "banana"
	fmt.Printf("%p %p", &cat, &str)
}*/

/*func main() {

	// 准备一个字符串类型
	var house = "Malibu Point 10880, 90265"

	// 对字符串取地址, ptr类型为*string
	ptr := &house

	// 打印ptr的类型
	fmt.Printf("ptr type: %T\n", ptr)

	// 打印ptr的指针地址
	fmt.Printf("address: %p\n", ptr)

	// 对指针进行取值操作
	value := *ptr

	// 取值后的类型
	fmt.Printf("value type: %T\n", value)

	// 指针取值后就是指向变量的值
	fmt.Printf("value: %s\n", value)

}*/

// 交换函数
/*func swap(a, b *int) {

	// 取a指针的值, 赋给临时变量t
	t := *a

	// 取b指针的值, 赋给a指针指向的变量
	*a = *b

	// 将a指针的值赋给b指针指向的变量
	*b = t
}

func main() {

	// 准备两个变量, 赋值1和2
	x, y := 1, 2

	// 交换变量值
	swap(&x, &y)

	// 输出变量值
	fmt.Println(x, y)
}*/

//在swap() 函数中交换操作的是指针值
/*func swap(a, b *int) {
	b, a = a, b
	fmt.Println(a, b)
}

func main() {
	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x, y)
}*/

// 定义命令行参数
/*var mode = flag.String("mode", "", "process mode")

func main() {
	// 解析命令行参数
	flag.Parse()
	// 输出命令行参数
	fmt.Println(*mode)
}*/
//go run main.go --mode=HAHA

//new() 函数可以创建一个对应类型的指针，创建过程会分配内存，被创建的指针指向默认值。
/*func main() {
	str := new(string)
	*str = "Go语言教程"
	fmt.Println(*str)
}*/

/*func calc(a, b int) int {
	var c int
	c = a * b

	var x int
	x = c * 10

	return x
}
func main() {
	res := calc(1, 2)
	fmt.Print(res)
}*/

//变量逃逸
// 本函数测试入口参数和返回值情况
/*func dummy(b int) int {

	// 声明一个变量c并赋值
	var c int
	c = b

	return c
}

// 空函数, 什么也不做
func void() {
}

func main() {

	// 声明a变量并打印
	var a int

	// 调用void()函数
	void()

	// 打印a变量的值和dummy()函数返回
	fmt.Println(a, dummy(0))
}*/

// 声明空结构体测试结构体逃逸情况
/*type Data struct {
}

func dummy() *Data {
	// 实例化c为Data类型
	var c Data

	//返回函数局部变量地址
	return &c
}

func main() {
	fmt.Println(dummy())
}*/

//const关键字
/*func main() {
	//%后的副词[1]告知Printf重复使用第一个操作数。
	//const noDelay time.Duration = 0
	//const timeout = 5 * time.Minute
	//fmt.Printf("%T %[1]v\n", noDelay)     // "time.Duration 0"
	//fmt.Printf("%T %[1]v\n", timeout)     // "time.Duration 5m0s"
	//fmt.Printf("%T %[1]v\n", time.Minute) // "time.Duration 1m0s"

	//iota 常量生成器
	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Println(a, b, c, d) // "1 1 2 2"
}*/

//GO语言模拟枚举
/*func main() {
	//type Weapon int
	//
	//const (
	//	Arrow Weapon = iota // 开始生成枚举值, 默认为0
		Shuriken
		SniperRifle
 		Rifle
	    Blower
	//)
	//
	//// 输出所有枚举值
	//fmt.Println(Arrow, Shuriken, SniperRifle, Rifle, Blower)
	//
	//// 使用枚举类型并赋初值
	//var weapon Weapon = Blower
	//fmt.Println(weapon)

	//2进制 左移一位
	const (
		FlagNone = 1 << iota
		FlagRed
		FlagGreen
		FlagBlue
	)

	fmt.Printf("%d %d %d\n", FlagRed, FlagGreen, FlagBlue)
	fmt.Printf("%b %b %b\n", FlagRed, FlagGreen, FlagBlue)
}*/

// 声明芯片类型
/*type ChipType int

const (
	Memory ChipType = iota
	CPU             // 中央处理器
	GPU             // 图形处理器
)

func (c ChipType) String() string {
	switch c {
	case Memory:
		return "Memory"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}

	return "N/A"
}

func main() {
	// 输出CPU的值并以整型格式显示
	fmt.Printf("%s %d", Memory, Memory)
}*/

//type关键字
//将NewInt定义为int类型
/*type NewInt int

//将int取一个别名叫IntAlias
type IntAlias = int

func main() {

	//将a声明为NewInt类型
	var a NewInt
	//查看a的类型名
	fmt.Printf("a type: %T\n", a)

	//将a2声明为IntAlias类型
	var a2 IntAlias
	//查看a2的类型名
	fmt.Printf("a2 type: %T\n", a2)
}*/

// 定义time.Duration的别名为MyDuration
/*type MyDuration = time.Duration

// 为MyDuration添加一个函数
func (m MyDuration) EasySet(a string) {
}
func main() {
}*/

/*
编译器提示：不能在一个非本地的类型 time.Duration 上定义新方法，非本地类型指的就是 time.Duration 不是在 main 包中定义的，而是在 time 包中定义的，与 main 包不在同一个包中，因此不能为不在一个包中的类型定义方法。

解决这个问题有下面两种方法：
将第 8 行修改为 type MyDuration time.Duration，也就是将 MyDuration 从别名改为类型；
将 MyDuration 的别名定义放在 time 包中。
*/

/*
在结构体成员嵌入时使用别名
*/
// 定义商标结构
/*type Brand struct {
}

// 为商标结构添加Show()方法
func (t Brand) Show() {
}

// 为Brand定义一个别名FakeBrand
type FakeBrand = Brand

// 定义车辆结构
type Vehicle struct {
	// 嵌入两个结构
	FakeBrand
	Brand
}

func main() {
	// 声明变量a为车辆类型
	var a Vehicle

	// 指定调用FakeBrand的Show
	a.FakeBrand.Show()

	// 取a的类型反射对象
	ta := reflect.TypeOf(a)

	// 遍历a的所有成员
	for i := 0; i < ta.NumField(); i++ {
		// a的成员信息
		f := ta.Field(i)

		// 打印成员的字段名和类型
		fmt.Printf("FieldName: %v, FieldType: %v\n", f.Name, f.Type.
			Name())
	}
}*/

//Go语言 词法元素5种：标识符（identifier）、关键字（keyword）、操作符（operator）、分隔符（delimiter）、字面量（literal）

//字符串和数字类型转换
//整型转字符串
/*func main() {
	num := 100
	str := strconv.Itoa(num)
	fmt.Printf("type:%T value:%#v\n", str, str)
}*/

//字符串转整型
/*func main() {
	str1 := "110"
	str2 := "s100"
	num1, err := strconv.Atoi(str1)
	if err != nil {
		fmt.Printf("%v 转换失败！", str1)
	} else {
		fmt.Printf("type:%T value:%#v\n", num1, num1)
	}
	num2, err := strconv.Atoi(str2)
	if err != nil {
		fmt.Printf("%v 转换失败！", str2)
	} else {
		fmt.Printf("type:%T value:%#v\n", num2, num2)
	}
}*/

//将字符串转换为bool，接受 1、0、t、f、T、F、true、false、True、False、TRUE、FALSE，其它的值均返回错误
/*func main() {
	str1 := "110"
	boo1, err := strconv.ParseBool(str1)
	if err != nil {
		fmt.Printf("str1: %v\n", err)
	} else {
		fmt.Println(boo1)
	}
	str2 := "t"
	boo2, err := strconv.ParseBool(str2)
	if err != nil {
		fmt.Printf("str2: %v\n", err)
	} else {
		fmt.Println(boo2)
	}
}*/

//ParseInt() 函数用于返回字符串表示的整数值（可以包含正负号）
/*func main() {
	str := "-11"
	num, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num)
	}
}*/

//ParseUint() 函数的功能类似于 ParseInt() 函数，但 ParseUint() 函数不接受正负号，用于无符号整型
/*func main() {
	str := "11"
	num, err := strconv.ParseUint(str, 10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num)
	}
}*/

/*func main() {
	str := "3.1415926"
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num)
	}
}*/

/*func main() {
	// 声明一个slice
	b10 := []byte("int (base 10):")

	// 将转换为10进制的string，追加到slice中
	b10 = strconv.AppendInt(b10, -42, 10)

	fmt.Println(string(b10))
	b16 := []byte("int (base 16):")
	b16 = strconv.AppendInt(b16, -42, 16)
	fmt.Println(string(b16))
}*/

/*
定义数组
var 数组变量名 [元素数量]Type
*/
/*func main() {
	//var a [3]int // 定义三个整数的数组
	//
	//fmt.Println(a[0])        // 打印第一个元素
	//fmt.Println(a[len(a)-1]) // 打印最后一个元素
	//
	//// 打印索引和元素
	//for i, v := range a {
	//	fmt.Printf("%d %d\n", i, v)
	//}
	//
	//// 仅打印元素
	//for _, v := range a {
	//	fmt.Printf("%d\n", v)
	//}

	//var q [3]int = [3]int{1, 2, 3}
	//var r [3]int = [3]int{1, 2}
	//fmt.Println(r[2]) // "0"
	//fmt.Println(q[2]) // "3"

	//m := [...]int{1, 2, 3}
	//fmt.Printf("%T\n", m) // "[3]int"

	//错误
	//q := [3]int{1, 2, 3}
	//q = [4]int{1, 2, 3, 4} // 编译错误：无法将 [4]int 赋给 [3]int
	//
	//fmt.Printf("%T\n", q) // "[3]int"

	//a := [2]int{1, 2}
	//b := [...]int{1, 2}
	//c := [2]int{1, 3}
	//
	//fmt.Println(a == b, a == c, b == c) // "true false false"
	//d := [3]int{1, 2}
	//fmt.Println(a == d) // 编译错误：无法比较 [2]int == [3]int

	var team [3]string
	team[0] = "hammer"
	team[1] = "soldier"
	team[2] = "mum"

	for k, v := range team {
		fmt.Println(k, v)
	}
}*/

//多维数组
/*func main() {
	// 声明一个二维整型数组，两个维度的长度分别是 4 和 2
	var array [4][2]int
	// 使用数组字面量来声明并初始化一个二维整型数组
	array = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	fmt.Printf("%v", array)
	// 声明并初始化数组中索引为 1 和 3 的元素
	array = [4][2]int{1: {20, 21}, 3: {40, 41}}
	fmt.Printf("%v", array)
	// 声明并初始化数组中指定的元素
	array = [4][2]int{1: {0: 20}, 3: {1: 41}}
	fmt.Printf("%v", array)
}*/

/*func main() {
	// 声明一个 2×2 的二维整型数组
	var array [2][2]int
	// 设置每个元素的整型值
	array[0][0] = 10
	array[0][1] = 20
	array[1][0] = 30
	array[1][1] = 40
}*/

//切片
/*func main() {
	//var a = [3]int{1, 2, 3}

	//fmt.Println(a, a[1:2])

	//var highRiseBuilding [30]int

	//for i := 0; i < 30; i++ {
	//	highRiseBuilding[i] = i + 1
	//}

	// 区间
	//fmt.Println(highRiseBuilding[10:15])

	// 中间到尾部的所有元素
	//fmt.Println(highRiseBuilding[20:])

	// 开头到中间指定位置的所有元素
	//fmt.Println(highRiseBuilding[:2])

	// 声明字符串切片
	//var strList []string

	// 声明整型切片
	//var numList []int

	// 声明一个空切片
	//var numListEmpty = []int{}

	// 输出3个切片
	//fmt.Println(strList, numList, numListEmpty)

	// 输出3个切片大小
	//fmt.Println(len(strList), len(numList), len(numListEmpty))

	// 切片判定空的结果
	//fmt.Println(strList == nil)
	//fmt.Println(numList == nil)
	//fmt.Println(numListEmpty == nil)

	//使用 make() 函数生成的切片一定发生了内存分配操作，但给定开始与结束位置（包括切片复位）的切片只是将新的切片结构指向已经分配好的内存区域，设定开始与结束位置，不会发生内存分配操作。
	a := make([]int, 2)
	b := make([]int, 2, 10)

	fmt.Println(a, b)
	fmt.Println(len(a), len(b))
}*/

//append
/*func main() {
	//var a []int
	//a = append(a, 1) // 追加1个元素
	//fmt.Println(a)
	//a = append(a, 1, 2, 3) // 追加多个元素, 手写解包方式
	//fmt.Println(a)
	//a = append(a, []int{1, 2, 3}...) // 追加一个切片, 切片需要解包
	//fmt.Println(a)

	//切片在扩容时，容量的扩展规律是按容量的 2 倍数进行扩充，例如 1、2、4、8、16……
	//var numbers []int

	//for i := 0; i < 10; i++ {
	//	numbers = append(numbers, i)
	//	fmt.Printf("len: %d  cap: %d pointer: %p\n", len(numbers), cap(numbers), numbers)
	//}

	//在切片开头添加元素-性能较差（内容会重新分配，已有元素全部被复制1次）
	var a = []int{1, 2, 3}
	a = append([]int{0}, a...) // 在开头添加1个元素
	fmt.Println(a)
	a = append([]int{-3, -2, -1}, a...) // 在开头添加1个切片
	fmt.Println(a)
}*/

//copy复制
/*func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	fmt.Println(slice2)
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Println(slice1)
}*/

//切片操作
/*func main() {
	// 设置元素数量为1000
	//const elementCount = 1000
	//
	//// 预分配足够多的元素切片
	//srcData := make([]int, elementCount)
	//
	//// 将切片赋值
	//for i := 0; i < elementCount; i++ {
	//	srcData[i] = i
	//}
	//
	//// 引用切片数据
	//refData := srcData
	//
	//// 预分配足够多的元素切片
	//copyData := make([]int, elementCount)
	//// 将数据复制到新的切片空间中
	//copy(copyData, srcData)
	//
	//// 修改原始数据的第一个元素
	//srcData[0] = 999
	//
	//// 打印引用切片的第一个元素
	//fmt.Println(refData[0])
	//
	//// 打印复制切片的第一个和最后一个元素
	//fmt.Println(copyData[0], copyData[elementCount-1])
	//
	//// 复制原始数据从4到6(不包含)
	//copy(copyData, srcData[4:6])
	//
	//for i := 0; i < 5; i++ {
	//	fmt.Printf("%d ", copyData[i])
	//}

	//删除开头元素直接移动指针数据
	//var a = []int{1, 2, 3, 4, 5}
	//a = a[1:] //删除开头1个元素
	//fmt.Println(a)
	//a = a[3:] //删除开头3个元素
	//fmt.Println(a)

	//可以不移动数据指针，但是将后面的数据向开头移动，可以用 append 原地完成（所谓原地完成是指在原有的切片数据对应的内存区间内完成，不会导致内存空间结构的变化）：
	//var a = []int{1, 2, 3, 4, 5, 6}
	//a = append(a[:0], a[1:]...) // 删除开头1个元素
	//fmt.Println(a)
	//a = append(a[:0], a[3:]...) // 删除开头N个元素
	//fmt.Println(a)

	//用 copy() 函数来删除开头的元素：
	//var a = []int{1, 2, 3, 4, 5}
	//var b = []int{0, 0, 0, 0}
	//copy(b, a[1:])
	//fmt.Println(b)
	//fmt.Println(a[1:])
	//a = a[:copy(a, a[1:])] // 删除开头1个元素
	//fmt.Println(a)
	//a = a[:copy(a, a[3:])] // 删除开头N个元素
	//fmt.Println(a)

	seq := []string{"a", "b", "c", "d", "e"}

	// 指定删除位置
	index := 2

	// 查看删除位置之前的元素和之后的元素
	fmt.Println(seq[:index], seq[index+1:])

	// 将删除点前后的元素连接起来
	seq = append(seq[:index], seq[index+1:]...)

	fmt.Println(seq)
}*/

//range，它可以配合关键字 for 来迭代切片里的每一个元素  从头部开始迭代
/*func main() {
	// 创建一个整型切片，并赋值
	slice := []int{10, 20, 30, 40}
	// 迭代每一个元素，并显示其值
	for index, value := range slice {
		fmt.Printf("Index: %d Value: %d\n", index, value)
	}

	//range 返回的是每个元素的副本，而不是直接返回对该元素的引用
	// 创建一个整型切片，并赋值
	slices := []int{10, 20, 30, 40}
	// 迭代每个元素，并显示值和地址
	for index, value := range slices {
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n", value, &value, &slice[index])
	}
	//迭代返回的变量是一个在迭代过程中根据切片依次赋值的新变量，所以 value 的地址总是相同的，要想获取每个元素的地址，需要使用切片变量和索引值（例如上面代码中的 &slice[index]）
}*/

//多维切片--切片有几个维度就需要几个[ ]
/*func main() {
	//声明一个二维切片
	var slice [][]int
	//为二维切片赋值
	slice = [][]int{{10}, {100, 200}}
	fmt.Println(slice)

	// 声明一个二维整型切片并赋值
	slices := [][]int{{10}, {100, 200}}
	fmt.Println(slices)

	//声明一个二维性切片并赋值
	newSlice := [][]int{{10}, {20, 30}}
	fmt.Println(newSlice)

	//为第一个切片追加20的元素
	newSlice[0] = append(newSlice[0], 20)
	fmt.Println(newSlice)
}*/

//map-引用类型（一种元素对（pair）的无序集合，pair 对应一个 key（索引）和一个 value（值），所以这个结构也称为关联数组或字典，这是一种能够快速寻找值的理想结构，给定 key，就可以迅速找到对应的 value。map 这种数据结构在其他编程语言中也称为字典（Python）、hash 和 HashTable 等。）
/*func main() {
	var mapList map[string]int
	//var mapCreated map[string]float32
	var mapAssigned map[string]int

	mapList = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32) //可以使用 make()，但不能使用 new() 来构造 map，如果错误的使用 new() 分配了一个引用对象，会获得一个空引用的指针，相当于声明了一个未初始化的变量并且取了它的地址
	mapAssigned = mapList

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3

	fmt.Printf("Map literal at \"one\" is: %d\n", mapList["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapList["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapList["ten"])
}*/

//用切片作为 map 的值 -既然一个 key 只能对应一个 value，而 value 又是一个原始类型，那么如果一个 key 要对应多个值怎么办？通过将 value 定义为 []int 类型或者其他类型的切片，就可以优雅的解决这个问题
/*func main() {
	mp1 := make(map[int][]int)
	mp2 := make(map[int]*[]int)

	a := []int{1, 2, 3}
	b := []int{1, 2, 4}

	mp1[0] = a
	mp1[1] = b

	fmt.Println(a)
	fmt.Println(mp1)
	fmt.Println(mp2)
}*/

/*func main() {
	scene := make(map[string]int)

	//单引号里面都是字符，只能有一个字符 双引号里面是字符串
	scene["route"] = 55
	scene["china"] = 53
	scene["now"] = 52

	fmt.Println(scene)

	for k, v := range scene {
		fmt.Println(k, v)
	}

	// 声明一个切片保存map数据
	var sceneList []string

	// 将map数据遍历复制到切片中
	for k := range scene {
		sceneList = append(sceneList, k)
	}

	// 对切片进行排序
	sort.Strings(sceneList)

	// 输出
	fmt.Println(sceneList)

	//删除
	delete(scene, "route")
	fmt.Println(scene)
}*/

/*
基于哈希值的多键索引及查询
传统的数据索引过程是将输入的数据做特征值。这种特征值有几种常见做法：
将特征使用某种算法转为整数，即哈希值，使用整型值做索引。
将特征转为字符串，使用字符串做索引。

数据都基于特征值构建好索引后，就可以进行查询。查询时，重复这个过程，将查询条件转为特征值，使用特征值进行查询得到结果。
*/
// 人员档案
/*type Profile struct {
	Name    string // 名字
	Age     int    // 年龄
	Married bool   // 已婚
}

func main() {
	list := []*Profile{
		{Name: "张三", Age: 30, Married: true},
		{Name: "李四", Age: 21},
		{Name: "王麻子", Age: 21},
	}
	buildIndex(list)
	queryData("张三", 30)
}*/

/*func simpleHash(str string) (ret int) {
	// 遍历字符串中的每一个ASCII字符
	for i := 0; i < len(str); i++ {
		// 取出字符
		c := str[i]
		// 将字符的ASCII码相加
		ret += int(c)
	}
	return
}

// 查询键
type classicQueryKey struct {
	Name string // 要查询的名字
	Age  int    // 要查询的年龄
}

// 计算查询键的哈希值
func (c *classicQueryKey) hash() int {
	// 将名字的Hash和年龄哈希合并
	return simpleHash(c.Name) + c.Age*1000000
}

// 创建哈希值到数据的索引关系
var mapper = make(map[int][]*Profile)

// 构建数据索引
func buildIndex(list []*Profile) {

	// 遍历所有的数据
	for _, profile := range list {

		// 构建数据的查询索引
		key := classicQueryKey{profile.Name, profile.Age}

		// 计算数据的哈希值, 取出已经存在的记录
		existValue := mapper[key.hash()]

		// 将当前数据添加到已经存在的记录切片中
		existValue = append(existValue, profile)

		// 将切片重新设置到映射中
		mapper[key.hash()] = existValue
	}
}

//查询逻辑 给定查询条件（名字、年龄）。 根据查询条件构建查询键。 查询键生成哈希值。 根据哈希值在索引中查找数据集合。 遍历数据集合逐个与条件比对。 获得结果。
func queryData(name string, age int) {

	// 根据给定查询条件构建查询键
	keyToQuery := classicQueryKey{name, age}

	// 计算查询键的哈希值并查询, 获得相同哈希值的所有结果集合
	resultList := mapper[keyToQuery.hash()]

	// 遍历结果集合
	for _, result := range resultList {

		// 与查询结果比对, 确认找到打印结果
		if result.Name == name && result.Age == age {
			fmt.Println(result)
			return
		}
	}

	// 没有查询到时, 打印结果
	fmt.Println("no found")

}*/

//利用 map 特性的多键索引及查询
//使用结构体进行多键索引和查询比传统的写法更为简单，最主要的区别是无须准备哈希函数及相应的字段无须做哈希合并

/*// 查询键
type queryKey struct {
	Name string
	Age  int
}

// 创建查询键到数据的映射
var mapper = make(map[queryKey]*Profile)

// 构建查询索引
func buildIndex(list []*Profile) {

	// 遍历所有数据
	for _, profile := range list {

		// 构建查询键
		key := queryKey{
			Name: profile.Name,
			Age:  profile.Age,
		}

		// 保存查询键
		mapper[key] = profile
	}
}

// 根据条件查询数据
func queryData(name string, age int) {

	// 根据查询条件构建查询键
	key := queryKey{name, age}

	// 根据键值查询数据
	result, ok := mapper[key]

	// 找到数据打印出来
	if ok {
		fmt.Println(result)
	} else {
		fmt.Println("no found")
	}
}*/

/*
代码量大大减少的关键是：Go语言的底层会为 map 的键自动构建哈希值。能够构建哈希值的类型必须是非动态类型、非指针、函数、闭包。
非动态类型：可用数组，不能用切片。
非指针：每个指针数值都不同，失去哈希意义。
函数、闭包不能作为 map 的键。
*/

/*
sync.Map 有以下特性：
无须初始化，直接声明即可。
sync.Map 不能使用 map 的方式进行取值和设置等操作，而是使用 sync.Map 的方法进行调用，Store 表示存储，Load 表示获取，Delete 表示删除。 sync.Map 不能使用 make 创建
使用 Range 配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值，Range 参数中回调函数的返回值在需要继续迭代遍历时，返回 true，终止迭代遍历时，返回 false。
*/
//sync.Map 没有提供获取 map 数量的方法，替代方法是在获取 sync.Map 时遍历自行计算数量，sync.Map 为了保证并发安全有一些性能损失，因此在非并发情况下，使用 map 相比使用 sync.Map 会有更好的性能。
/*func main() {
	var scene sync.Map

	// 将键值对保存到sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)

	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))

	// 根据键删除对应的键值对
	scene.Delete("london")

	// 遍历所有sync.Map中的键值对
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}*/

//list列表
/*func main() {
	l := list.New()

	l.PushBack("first")
	l.PushFront(123)

	fmt.Printf("v%", l)
	fmt.Println(l)
}*/

/*func main() {
	l := list.New()

	//尾部添加
	l.PushBack("tail")

	//头部添加
	l.PushFront("head")

	//尾部添加后保存元素句柄
	element := l.PushBack("first")

	//在first之后添加high
	l.InsertAfter("high", element)

	//在first之前添加noon
	l.InsertBefore("lalala", element)

	fmt.Println(l)

	//使用
	l.Remove(element)

	ll := list.New()

	//尾部添加
	ll.PushBack("dudu")

	//头部添加
	ll.PushFront("haha")

	for i := ll.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}*/

//布尔类型的零值（初始值）为 false，数值类型的零值为 0，字符串类型的零值为空字符串""，而指针、切片、映射、通道、函数和接口的零值则是 nil。
/*func main() {
	//nil 是 map、slice、pointer、channel、func、interface 的零值

	//var m map[int]string
	//var ptr *int
	//var c chan int
	//var sl []int
	//var f func()
	//var i interface{}
	//
	//fmt.Printf("%#v\n", m)
	//fmt.Printf("%#v\n", ptr)
	//fmt.Printf("%#v\n", c)
	//fmt.Printf("%#v\n", sl)
	//fmt.Printf("%#v\n", f)
	//fmt.Printf("%#v\n", i)

	var p *struct{}
	fmt.Println(unsafe.Sizeof(p)) // 8

	var s []int
	fmt.Println(unsafe.Sizeof(s)) // 24

	var m map[int]bool
	fmt.Println(unsafe.Sizeof(m)) // 8

	var c chan string
	fmt.Println(unsafe.Sizeof(c)) // 8

	var f func()
	fmt.Println(unsafe.Sizeof(f)) // 8

	var i interface{}
	fmt.Println(unsafe.Sizeof(i)) // 16
}*/

//make和new关键字的区别及实现原理  返回的永远是类型的指针，指针指向分配类型的内存地址
/*func main() {
	var sum *int
	sum = new(int) //分配空间
	*sum = 98

	fmt.Println(*sum)

	type Student struct {
		name string
		age  int
	}

	var s *Student
	s = new(Student) //分配空间
	s.name = "dequan"

	fmt.Println(s)
}*/

/*
Go语言中的 new 和 make 主要区别如下：
make 只能用来分配及初始化类型为 slice、map、chan 的数据。new 可以分配任意类型的数据；
new 分配返回的是指针，即类型 *Type。make 返回引用，即 Type；
new 分配的空间被清零。make 分配空间后，会进行初始化。
总结
最后，简单总结一下Go语言中 make 和 new 关键字的实现原理，make 关键字的主要作用是创建 slice、map 和 Channel 等内置的数据结构，而 new 的主要作用是为类型申请一片内存空间，并返回指向这片内存的指针。
*/

/*func main() {
	var ten int = 11
	if ten > 10 {
		fmt.Println(">10")
	} else {
		fmt.Println("<=10")
	}
}*/

//if 还有一种特殊的写法，可以在 if 表达式之前添加一个执行语句，再根据变量值进行判断，代码如下：
//纯文本复制
/*func main() {
	if err1, err2 := Connect(); err1 == nil {
		fmt.Println(err1)
		fmt.Println(err2)
		return
	}
}
func Connect() (error, string) {
	return nil, "error"
}*/

/*Connect 是一个带有返回值的函数，err := Connect() 是一个语句，执行 Connect 后，将错误保存到 err 变量中。
err != nil 才是 if 的判断表达式，当 err 不为空时，打印错误并返回。
这种写法可以将返回值与判断放在一行进行处理，而且返回值的作用范围被限制在 if 、 else 语句组合中。*/

/*func main() {
	//sum := 0
	//for i := 1; i < 10; i++ {
	//	sum += i
	//	fmt.Println(sum)
	//}

	//结束循环时带可执行语句的无限循环
	//var i int
	//
	//for ; ; i++ {
	//	fmt.Println(i)
	//	if i > 10 {
	//		break
	//	}
	//}

	//var i int
	//for {
	//	if i > 10 {
	//		break
	//	}
	//	i++
	//	fmt.Println(i)
	//}

	//类似while
	var i int
	for i <= 10 {
		i++
	}
}*/

//9*9 乘法表
/*func main() {
	for y := 1; y < 10; y++ {
		for x := 1; x <= y; x++ {
			fmt.Printf("%d*%d=%d ", x, y, x*y)
		}
		// 手动生成回车
		fmt.Println()
	}
}*/

//键值循环-for range
//for range 可以遍历数组、切片、字符串、map 及通道（channel）  类似 foreach
/*func main() {
	//遍历切片
	for key, value := range []int{1, 2, 3, 4} {
		fmt.Printf("key:%d  value:%d\n", key, value)
	}

	//遍历字符串
	var str = "hello 你好"
	for key, value := range str {
		fmt.Printf("key:%d value:0x%x\n", key, value)
	}

	//遍历map-对 map 遍历时，遍历输出的键值是无序的，如果需要有序的键值对输出，需要对结果进行排序。
	m := map[string]int{
		"hello": 100,
		"world": 200,
	}
	for key, value := range m {
		fmt.Println(key, value)
	}

	//遍历通道（channel）——接收通道数据
	//for range 可以遍历通道（channel），但是通道在遍历时，只输出一个值，即管道内的类型对应的数据。
	c := make(chan int)

	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}


		第 1 行创建了一个整型类型的通道。
		第 3 行启动了一个 goroutine，其逻辑的实现体现在第 5～8 行，实现功能是往通道中推送数据 1、2、3，然后结束并关闭通道。
		这段 goroutine 在声明结束后，在第 9 行马上被执行。
		从第 11 行开始，使用 for range 对通道 c 进行遍历，其实就是不断地从通道中取数据，直到通道被关闭。


	map1 := map[string]int{
		"hello": 100,
		"world": 200,
	}
	for _, value := range map1 {
		fmt.Println(value)
	}
}*/

/*func main() {
	var a = "hello"

	switch a {
	case "hello":
		fmt.Println(1)
	case "haha":
		fmt.Println(2)
	default:
		fmt.Println(3)
	}

	var b = "mum"
	switch b {
	case "mum", "daddy":
		fmt.Println("family")
	}

	var s = "hello"
	switch {
	case s == "hello":
		fmt.Println("hello")
		fallthrough
	case s != "world":
		fmt.Println("world")
	}
}*/

/*func main() {

	//for x := 0; x < 10; x++ {
	//
	//	for y := 0; y < 10; y++ {
	//
	//		if y == 2 {
	//			// 跳转到标签
	//			goto breakHere
	//		}
	//
	//	}
	//}

	// 手动返回, 避免执行进入标签
	//return

	// 标签
	//breakHere:
	//	fmt.Println("done")

OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				break OuterLoop
			case 3:
				fmt.Println(i, j)
				break OuterLoop
			}
		}
	}
}*/

/*func main() {
	// 准备从标准输入读取数据。
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your name:")
	// 读取数据直到碰到 \n 为止。
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
		// 异常退出。
		os.Exit(1)
	} else {
		// 用切片操作删除最后的 \n 。
		name := input[:len(input)-2]
		fmt.Printf("Hello, %s! What can I do for you?\n", name)
	}
	for {
		input, err = inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred: %s\n", err)
			continue
		}
		input = input[:len(input)-1]
		// 全部转换为小写。
		input = strings.ToLower(input)
		fmt.Println(input)
		switch input {
		case "":
			continue
		case "nothing", "bye":
			fmt.Println("Bye!")
			// 正常退出。
			os.Exit(0)
		default:
			fmt.Println("Sorry, I didn't catch you.")
		}
	}
}*/

//词频统计
/*func main() {
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s <file1> [<file2> [... <fileN>]]\n",
			filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	frequencyForWord := map[string]int{} // 与:make(map[string]int)相同
	for _, filename := range commandLineFiles(os.Args[1:]) {
		updateFrequencies(filename, frequencyForWord)
	}
	reportByWords(frequencyForWord)
	wordsForFrequency := invertStringIntMap(frequencyForWord)
	reportByFrequency(wordsForFrequency)
}

func commandLineFiles(files []string) []string {
	if runtime.GOOS == "windows" {
		args := make([]string, 0, len(files))
		for _, name := range files {
			if matches, err := filepath.Glob(name); err != nil {
				args = append(args, name) // 无效模式
			} else if matches != nil {
				args = append(args, matches...)
			}
		}
		return args
	}
	return files
}

func updateFrequencies(filename string, frequencyForWord map[string]int) {
	var file *os.File
	var err error
	if file, err = os.Open(filename); err != nil {
		log.Println("failed to open the file: ", err)
		return
	}
	defer file.Close()
	readAndUpdateFrequencies(bufio.NewReader(file), frequencyForWord)
}

func readAndUpdateFrequencies(reader *bufio.Reader,
	frequencyForWord map[string]int) {
	for {
		line, err := reader.ReadString('\n')
		for _, word := range SplitOnNonLetters(strings.TrimSpace(line)) {
			if len(word) > utf8.UTFMax ||
				utf8.RuneCountInString(word) > 1 {
				frequencyForWord[strings.ToLower(word)] += 1
			}
		}
		if err != nil {
			if err != io.EOF {
				log.Println("failed to finish reading the file: ", err)
			}
			break
		}
	}
}

func SplitOnNonLetters(s string) []string {
	notALetter := func(char rune) bool { return !unicode.IsLetter(char) }
	return strings.FieldsFunc(s, notALetter)
}

func invertStringIntMap(intForString map[string]int) map[int][]string {
	stringsForInt := make(map[int][]string, len(intForString))
	for key, value := range intForString {
		stringsForInt[value] = append(stringsForInt[value], key)
	}
	return stringsForInt
}

func reportByWords(frequencyForWord map[string]int) {
	words := make([]string, 0, len(frequencyForWord))
	wordWidth, frequencyWidth := 0, 0
	for word, frequency := range frequencyForWord {
		words = append(words, word)
		if width := utf8.RuneCountInString(word); width > wordWidth {
			wordWidth = width
		}
		if width := len(fmt.Sprint(frequency)); width > frequencyWidth {
			frequencyWidth = width
		}
	}
	sort.Strings(words)
	gap := wordWidth + frequencyWidth - len("Word") - len("Frequency")
	fmt.Printf("Word %*s%s\n", gap, " ", "Frequency")
	for _, word := range words {
		fmt.Printf("%-*s %*d\n", wordWidth, word, frequencyWidth,
			frequencyForWord[word])
	}
}

func reportByFrequency(wordsForFrequency map[int][]string) {
	frequencies := make([]int, 0, len(wordsForFrequency))
	for frequency := range wordsForFrequency {
		frequencies = append(frequencies, frequency)
	}
	sort.Ints(frequencies)
	width := len(fmt.Sprint(frequencies[len(frequencies)-1]))
	fmt.Println("Frequency → Words")
	for _, frequency := range frequencies {
		words := wordsForFrequency[frequency]
		sort.Strings(words)
		fmt.Printf("%*d %s\n", width, frequency, strings.Join(words, ", "))
	}
}*/

//缩进排序
/*var original = []string{
	"Nonmetals",
	"    Hydrogen",
	"    Carbon",
	"    Nitrogen",
	"    Oxygen",
	"Inner Transitionals",
	"    Lanthanides",
	"        Europium",
	"        Cerium",
	"    Actinides",
	"        Uranium",
	"        Plutonium",
	"        Curium",
	"Alkali Metals",
	"    Lithium",
	"    Sodium",
	"    Potassium",
}

func main() {
	fmt.Println("|     Original      |       Sorted      |")
	fmt.Println("|-------------------|-------------------|")
	sorted := SortedIndentedStrings(original) // 最初是 []string
	for i := range original {                 // 在全局变量中设置
		fmt.Printf("|%-19s|%-19s|\n", original[i], sorted[i])
	}
}

func SortedIndentedStrings(slice []string) []string {
	entries := populateEntries(slice)
	return sortedEntries(entries)
}

func populateEntries(slice []string) Entries {
	indent, indentSize := computeIndent(slice)
	entries := make(Entries, 0)
	for _, item := range slice {
		i, level := 0, 0
		for strings.HasPrefix(item[i:], indent) {
			i += indentSize
			level++
		}
		key := strings.ToLower(strings.TrimSpace(item))
		addEntry(level, key, item, &entries)
	}
	return entries
}

func computeIndent(slice []string) (string, int) {
	for _, item := range slice {
		if len(item) > 0 && (item[0] == ' ' || item[0] == '\t') {
			whitespace := rune(item[0])
			for i, char := range item[1:] {
				if char != whitespace {
					i++
					return strings.Repeat(string(whitespace), i), i
				}
			}
		}
	}
	return "", 0
}

func addEntry(level int, key, value string, entries *Entries) {
	if level == 0 {
		*entries = append(*entries, Entry{key, value, make(Entries, 0)})
	} else {
		addEntry(level-1, key, value,
			&((*entries)[entries.Len()-1].children))
	}
}

func sortedEntries(entries Entries) []string {
	var indentedSlice []string
	sort.Sort(entries)
	for _, entry := range entries {
		populateIndentedStrings(entry, &indentedSlice)
	}
	return indentedSlice
}

func populateIndentedStrings(entry Entry, indentedSlice *[]string) {
	*indentedSlice = append(*indentedSlice, entry.value)
	sort.Sort(entry.children)
	for _, child := range entry.children {
		populateIndentedStrings(child, indentedSlice)
	}
}

type Entry struct {
	key      string
	value    string
	children Entries
}
type Entries []Entry

func (entries Entries) Len() int { return len(entries) }

func (entries Entries) Less(i, j int) bool {
	return entries[i].key < entries[j].key
}
func (entries Entries) Swap(i, j int) {
	entries[i], entries[j] = entries[j], entries[i]
}*/

//Go语言二分查找法
//二分查找函数 //假设有序数组的顺序是从小到大（很关键，决定左右方向）
/*func BinaryFind(arr *[]int, leftIndex int, rightIndex int, findVal int) {
	//判断 leftIndex是否大于rightIndex
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}
	//先找到 中间的下标
	middle := (leftIndex + rightIndex) / 2

	if (*arr)[middle] > findVal {
		//要查找的数，范围应该在 leftIndex 到 middle+1
		BinaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		//要查找的数，范围应该在 middle+1 到 rightIndex
		BinaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Printf("找到了，下标为：%v \n", middle)
	}
}

func main() {
	//定义一个数组
	arr := []int{1, 2, 5, 7, 15, 25, 30, 36, 39, 51, 67, 78, 80, 82, 85, 91, 92, 97}
	BinaryFind(&arr, 0, len(arr)-1, 30)
	fmt.Println("main arr=", arr)
}*/

/**
冒泡排序
*/
/*func main() {
	arr := [...]int{21, 32, 12, 33, 14, 6, 87, 24}
	var n = len(arr)
	fmt.Println("--------没排序前--------\n", arr)
	for i := 0; i <= n-1; i++ {
		fmt.Println("--------第", i+1, "次冒泡--------")
		for j := i; j <= n-1; j++ {
			if arr[i] > arr[j] {
				t := arr[i]
				arr[i] = arr[j]
				arr[j] = t
			}
			fmt.Println(arr)

		}
	}
	fmt.Println("--------最终结果--------\n", arr)
}*/

//go语言分布式id
//① 使用snowflak
/*func main() {
	//需要安装snowflake  go get github.com/bwmarrin/snowflake
	n, err := snowflake.NewNode(1)
	if err != nil {
		println(err)
		os.Exit(1)
	}
	for i := 0; i < 3; i++ {
		id := n.Generate()
		fmt.Println("id", id)
		fmt.Println(
			"node: ", id.Node(),
			"step: ", id.Step(),
			"time: ", id.Time(),
			"\n",
		)
	}
}
*/

//② 使用sonyflake
/*func getMachineID() (uint16, error) {
	var machineID uint16
	var err error
	machineID = readMachineIDFromLocalFile()
	if machineID == 0 {
		machineID, err = generateMachineID()
		if err != nil {
			return 0, err
		}
	}
	return machineID, nil
}
func checkMachineID(machineID uint16) bool {
	saddResult, err := saddMachineIDToRedisSet()
	if err != nil || saddResult == 0 {
		return true
	}
	err := saveMachineIDToLocalFile(machineID)
	if err != nil {
		return true
	}
	return false
}
func main() {
	t, _ := time.Parse("2006-01-02", "2018-01-01")
	settings := sonyflake.Settings{
		StartTime:      t,
		MachineID:      getMachineID,
		CheckMachineID: checkMachineID,
	}
	sf := sonyflake.NewSonyflake(settings)
	id, err := sf.NextID()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(id)
}*/

/*func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}
func add(x int, y int) int   { return x + y }
func sub(x, y int) (z int)   { z = x - y; return }
func first(x int, _ int) int { return x }
func zero(int, int) int      { return 0 }

func main() {
	fmt.Println(hypot(3, 4))  // "5"fmt.Printf("%T\n", add) // "func(int, int) int"
	fmt.Printf("%T\n", add) // "func(int, int) int"
	fmt.Printf("%T\n", sub)   // "func(int, int) int"
	fmt.Printf("%T\n", first) // "func(int, int) int"
	fmt.Printf("%T\n", zero)  // "func(int, int) int"
}*/
/*func typedTwoValues() (int, int) {
	return 1, 2
}

func namedRetValues() (a, b int) {

	a = 1
	b = 2

	return
}

func main() {
	a, b := typedTwoValues()
	c, d := namedRetValues()
	fmt.Println(a, b)
	fmt.Println(c, d)
}*/

//将秒转化为具体时间
/*const (
	// 定义每分钟的秒数
	SecondsPerMinute = 60

	// 定义每小时的秒数
	SecondsPerHour = SecondsPerMinute * 60

	// 定义每天的秒数
	SecondsPerDay = SecondsPerHour * 24
)

// 将传入的“秒”解析为3种时间单位
func resolveTime(seconds int) (day int, hour int, minute int) {
	day = seconds / SecondsPerDay
	hour = seconds / SecondsPerHour
	minute = seconds / SecondsPerMinute

	return
}

func main() {
	// 将返回值作为打印参数
	fmt.Println(resolveTime(1000)) //fmt.Println() 使用了可变参数，可以接收不定量的参数。

	// 只获取小时和分钟
	_, hour, minute := resolveTime(18000)
	fmt.Println(hour, minute)

	// 只获取天
	day, _, _ := resolveTime(90000)
	fmt.Println(day)
}*/

//参数传递
// 用于测试值传递效果的结构体
//将 Data 声明为结构体类型，结构体是拥有多个字段的复杂结构。
/*type Data struct {
	complax []int // 测试切片在参数传递中的效果 切片是一种动态类型，内部以指针存在。

	instance InnerData // 实例分配的innerData

	ptr *InnerData // 将ptr声明为InnerData的指针类型
}

// 代表各种结构体字段
type InnerData struct {
	a int
}

// 值传递测试函数
func passByValue(inFunc Data) Data {

	// 输出参数的成员情况
	fmt.Printf("inFunc value: %+v\n", inFunc)

	// 打印inFunc的指针
	fmt.Printf("inFunc ptr: %p\n", &inFunc)

	return inFunc
}

func main() {
	// 准备传入函数的结构
	in := Data{
		complax: []int{1, 2, 3},
		instance: InnerData{
			5,
		},

		ptr: &InnerData{1},
	}

	// 输入结构的成员情况
	fmt.Printf("in value: %+v\n", in)

	// 输入结构的指针地址
	fmt.Printf("in ptr: %p\n", &in)

	// 传入结构体，返回同类型的结构体
	out := passByValue(in)

	// 输出结构的成员情况
	fmt.Printf("out value: %+v\n", out)

	// 输出结构的指针地址
	fmt.Printf("out ptr: %p\n", &out)

	//所有的 Data 结构的指针地址都发生了变化，意味着所有的结构都是一块新的内存，无论是将 Data 结构传入函数内部，还是通过函数返回值传回 Data 都会发生复制行为。
	//所有的 Data 结构中的成员值都没有发生变化，原样传递，意味着所有参数都是值传递。
	//Data 结构的 ptr 成员在传递过程中保持一致，表示指针在函数参数值传递中传递的只是指针值，不会复制指针指向的部分。
}*/

//GO函数 将函数作为值保存到变量中

/*func fire() {
	fmt.Println("fire")
}
func main() {
	var f func()
	f = fire
	f()
}*/

//字符串的链式处理——操作与数据分离的设计技巧
// 字符串处理函数，传入字符串切片和处理链
/*func StringProccess(list []string, chain []func(string) string) {

	// 遍历每一个字符串
	for index, str := range list {

		// 第一个需要处理的字符串
		result := str

		// 遍历每一个处理链
		for _, proc := range chain {

			// 输入一个字符串进行处理，返回数据作为下一个处理链的输入。
			result = proc(result)
		}

		// 将结果放回切片
		list[index] = result
	}
}

// 自定义的移除前缀的处理函数
func removePrefix(str string) string {

	return strings.TrimPrefix(str, "go")
}

func main() {

	// 待处理的字符串列表
	list := []string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formater",
	}

	// 处理函数链
	chain := []func(string) string{
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}

	// 处理字符串
	StringProccess(list, chain)

	// 输出处理好的字符串
	for _, str := range list {
		fmt.Println(str)
	}

}*/

//匿名函数
//使用f()调用
/*func main() {
	// 将匿名函数体保存到f()中
	f := func(data int) {
		fmt.Println("hello", data)
	}

	f(100)
	f(100)
}*/

// 遍历切片的每个元素, 通过给定函数进行元素访问
//func visit(list []int, f func(int)) {
//	for _, v := range list {
//		f(v)
//	}
//}
//
//func main() {
//	// 使用匿名函数打印切片内容
//	visit([]int{1, 2, 3, 4}, func(v int) {
//		fmt.Println(v)
//	})
//}

//定义命令行参数 skill，从命令行输入 --skill 可以将=后的字符串传入 skillParam 指针变量。
//var skillParam = flag.String("skill", "", "skill to perform")
//
//func main() {
//
//	flag.Parse()
//
//	//定义一个从字符串映射到 func() 的 map，然后填充这个 map。
//	var skill = map[string]func(){
//		"fire": func() {
//			fmt.Println("chicken fire")
//		},
//		"run": func() {
//			fmt.Println("soldier run")
//		},
//		"fly": func() {
//			fmt.Println("angel fly")
//		},
//	}
//
//	//skillParam 是一个 *string 类型的指针变量，使用 *skillParam 获取到命令行传过来的值，并在 map 中查找对应命令行参数指定的字符串的函数。
//	if f, ok := skill[*skillParam]; ok {
//		fmt.Println(skill[*skillParam])
//		f()
//	} else {
//		fmt.Println("skill not found")
//	}
//	//go run demo.go --skill=fly
//}

//Go语言函数类型实现接口——把函数作为接口来调用

// 调用器接口
//type Invoker interface {
//	// 需要实现一个Call方法
//	Call(interface{})
//}
//
//// 结构体类型
//type Struct struct {
//}
//
//// 实现Invoker的Call
//func (s *Struct) Call(p interface{}) {
//	fmt.Println("from struct", p)
//}
//
//// 函数定义为类型
//type FuncCaller func(interface{})
//
//// 实现Invoker的Call
//func (f FuncCaller) Call(p interface{}) {
//
//	// 调用f函数本体
//	f(p)
//}
//
//func main() {
//
//	// 声明接口变量
//	var invoker Invoker
//
//	// 实例化结构体
//	s := new(Struct)
//
//	// 将实例化的结构体赋值到接口
//	invoker = s
//
//	// 使用接口调用实例化结构体的方法Struct.Call
//	invoker.Call("hello")
//
//	// 将匿名函数转为FuncCaller类型，再赋值给接口
//	invoker = FuncCaller(func(v interface{}) {
//		fmt.Println("from function", v)
//	})
//
//	// 使用接口调用FuncCaller.Call，内部会调用函数本体
//	invoker.Call("hello")
//}

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
