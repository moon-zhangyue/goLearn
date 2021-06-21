package animal

type Animal struct {
	Name string
}

//构造函数私有化
func NewAnimal(name string) Animal {
	return Animal{name: name}
}

func (a Animal) Call() string {
	return "动物的叫声..."
}
func (a Animal) FavorFood() string {
	return "爱吃的食物..."
}
func (a Animal) GetName() string {
	return a.Name
}
