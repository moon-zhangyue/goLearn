package animal

//Dog 类需要在 animal 包以外的地方进行初始化，所以需要将其属性名首字母都都替换成大写字母
type Dog struct {
	Animal *Animal
	Pet    Pet
}

func (d Dog) FavorFood() string {
	return "骨头"
}
func (d Dog) Call() string {
	return "汪汪汪"
}
