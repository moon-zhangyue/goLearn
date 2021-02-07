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

func GetData() (int, int) {
	return 100, 200
}
func main(){
	//匿名变量，可赋值，但会被抛弃，无法使用，不占用空间
	a, _ := GetData()
	_, b := GetData()
	fmt.Println(a, b)
}
