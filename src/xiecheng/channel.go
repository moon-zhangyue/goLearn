package main

//func add(a, b int, ch chan int) {
//	c := a + b
//	fmt.Printf("%d + %d = %d\n", a, b, c)
//	ch <- 1
//}
//
//func main() {
//	start := time.Now()
//	chs := make([]chan int, 10)
//
//	for i := 0; i < 10; i++ {
//		chs[i] = make(chan int)
//		go add(1, i, chs[i])
//	}
//
//	for _, ch := range chs {
//		<-ch
//	}
//
//	end := time.Now()
//	consume := end.Sub(start).Seconds()
//	fmt.Println("程序执行耗时(s)：", consume)
//}

//缓冲通道
//func test(ch chan int) {
//	for i := 0; i < 100; i++ {
//		ch <- i
//	}
//	close(ch) //关闭通道
//}
//
//func main() {
//	start := time.Now()
//	ch := make(chan int, 20) //缓冲通道
//	//ch := make(chan int) //非缓冲通道
//	go test(ch)
//	for i := range ch {
//		fmt.Println("接收到的数据:", i)
//	}
//	end := time.Now()
//	consume := end.Sub(start).Seconds()
//	fmt.Println("程序执行耗时(s)：", consume)
//}

//单向通道
//func test(ch chan<- int) {
//	//func test(ch <-chan int) {
//	for i := 0; i < 100; i++ {
//		ch <- i
//	}
//	close(ch)
//}
//
//单项只读通道
//func test1() <-chan int {
//	ch := make(chan int, 20)
//	for i := 0; i < 10; i++ {
//		ch <- i
//	}
//	close(ch)
//	return ch
//}
//
//func main() {
//	start := time.Now()
//	//ch := make(chan int, 20)
//	//go test(ch)
//	ch1 := <-test1()
//
//	//for i := range ch {
//	//	fmt.Println("接收到的数据:", i)
//	//}
//	end := time.Now()
//	consume := end.Sub(start).Seconds()
//	fmt.Println("程序执行耗时(s)：", consume)
//
//	fmt.Printf("ch值为%#v \n", ch1)
//}

// 单向通道：一般只在参数中限制通道方向是取值还是输入
// 现在返回值第二个只写，这会导致运用这个函数的时候，能获取到channel的内存地址，但是获取不到其中存着的值
//func danxiangtongdao(i chan<- int, j <-chan int) (<-chan int, chan<- int) {
//
//	fmt.Println("dxtd通道内存量：", len(i))
//
//	//i只能写，不能取
//	//x := <-i
//	//通道为引用类型，i和j指向的都是同一个通道
//	i <- 20
//	fmt.Println("dxtd通道存入一个20")
//	fmt.Println("dxtd通道内存量：", len(i))
//
//	//========从通道内第一次取值
//	y := <-j
//	//通道为引用类型，此次获取的值为函数外传入通道的值10
//	fmt.Println("函数内第一次通道取值：", y)
//	y++
//	fmt.Println("dxtd通道内存量：", len(i))
//
//	//创建p通道并写入
//	var p = make(chan int, 1)
//	p <- y
//
//	//========从通道内第二次取值
//	y = <-j
//	//由于通道是引用类型，此次获取的值为本函数内传入的20
//	fmt.Println("函数内第二次通道取值：", y)
//	y--
//	fmt.Println("dxtd通道内存量：", len(i))
//
//	//创建q通道并写入
//	var q = make(chan int, 1)
//	q <- y
//
//	//========关闭通道，不是销毁，不影响获取已有的值
//	close(p)
//	close(q)
//
//	//========返回两个通道，状态一个可读，一个可写
//	return p, q
//}
//
//func main() {
//	fmt.Println("==============单向通道=============")
//
//	//创建dxtd通道，允许有两个值
//	//因为本例中此通道传入函数后首先就做了一次写入操作，如果不是2，则会因通道存量不够而提示死锁deadlock
//	var dxtd = make(chan int, 2)
//	//通道存值
//	dxtd <- 10
//	fmt.Println("dxtd通道存入一个10")
//	//由于通道是引用类型，如果此处关闭通道则函数中将不能修改本通道
//	//close(dxtd)
//
//	//通道作为参数传给函数danxiangtongdao()，并返回两个方向不同的单向通道
//	dxtd1, dxtd2 := danxiangtongdao(dxtd, dxtd)
//	//获取dxtd1通道内的值
//	dxtd3 := <-dxtd1
//	//由于dxtd2通道只写，因此取不到dxtd2内的值
//	//dxtd4 := <-dxtd2
//
//	//获取dxtd通道内现在的存量
//	dxtdnow := len(dxtd)
//	//关闭dxtd通道
//	close(dxtd)
//
//	fmt.Printf("原dxtd为%#v【其存量为%#v】\n函数的两个返回通道如下\n第一个：%#v【只读，存有的值为%#v】\n第二个：%#v【只写，取不到】\n", dxtd, dxtdnow, dxtd1, dxtd3, dxtd2)
//}

// 打印内容：
// ==============单向通道=============
// dxtd通道存入一个10
// dxtd通道内存量： 1
// dxtd通道存入一个20
// dxtd通道内存量： 2
// 函数内第一次通道取值： 10
// dxtd通道内存量： 1
// 函数内第二次通道取值： 20
// dxtd通道内存量： 0
// 原dxtd为(chan int)(0xc00001a0e0)【其存量为0】
// 函数的两个返回通道如下
// 第一个：(<-chan int)(0xc00001a150)【只读，存有的值为11】
// 第二个：(chan<- int)(0xc00001a1c0)【只写，取不到】

//select
//func main() {
//	chs := [3]chan int{
//		make(chan int, 1),
//		make(chan int, 1),
//		make(chan int, 1),
//	}
//
//	rand.Seed(time.Now().Unix())
//	index := rand.Intn(3) // 随机生成0-2之间的数字
//	fmt.Printf("随机索引/数值: %d\n", index)
//	chs[index] <- index // 向通道发送随机数字
//
//	// 哪一个通道中有值，哪个对应的分支就会被执行
//	select {
//	case <-chs[0]:
//		fmt.Println("第一个条件分支被选中")
//	case <-chs[1]:
//		fmt.Println("第二个条件分支被选中")
//	case num := <-chs[2]:
//		fmt.Println("第三个条件分支被选中:", num)
//	default:
//		fmt.Println("没有分支被选中")
//	}
//}

//func main() {
//	chs := [3]chan int{
//		make(chan int, 3),
//		make(chan int, 3),
//		make(chan int, 3),
//	}
//
//	rand.Seed(time.Now().UnixNano())
//	index1 := rand.Intn(3) // 随机生成0-2之间的数字
//	fmt.Printf("随机索引/数值: %d\n", index1)
//	chs[index1] <- rand.Int() // 向通道发送随机数字
//
//	rand.Seed(time.Now().UnixNano())
//	index2 := rand.Intn(3)
//	fmt.Printf("随机索引/数值: %d\n", index2)
//	chs[index2] <- rand.Int()
//
//	rand.Seed(time.Now().UnixNano())
//	index3 := rand.Intn(3)
//	fmt.Printf("随机索引/数值: %d\n", index3)
//	chs[index3] <- rand.Int()
//
//	// 哪一个通道中有值，哪个对应的分支就会被执行
//	for i := 0; i < 3; i++ {
//		select {
//		case num, ok := <-chs[0]:
//			if !ok {
//				break
//			}
//			fmt.Println("第一个条件分支被选中: chs[0]=>", num)
//		case num, ok := <-chs[1]:
//			if !ok {
//				break
//			}
//			fmt.Println("第二个条件分支被选中: chs[1]=>", num)
//		case num, ok := <-chs[2]:
//			if !ok {
//				break
//			}
//			fmt.Println("第三个条件分支被选中: chs[2]=>", num)
//		default:
//			fmt.Println("没有分支被选中")
//		}
//	}
//}
