//demo测试
package main // 声明 main 包
import "fmt"

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
func main() {
	mp1 := make(map[int][]int)
	mp2 := make(map[int]*[]int)

	a := []int{1, 2, 3}
	b := []int{1, 2, 4}

	mp1[0] = a
	mp1[1] = b

	fmt.Println(a)
	fmt.Println(mp1)
	fmt.Println(mp2)
}
