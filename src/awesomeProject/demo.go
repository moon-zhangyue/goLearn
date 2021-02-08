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
type Data struct {
}

func dummy() *Data {
	// 实例化c为Data类型
	var c Data

	//返回函数局部变量地址
	return &c
}

func main() {
	fmt.Println(dummy())
}
