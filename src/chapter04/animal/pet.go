package animal

type Pet struct {
	Name string
}

func NewPet(name string) Pet {
	return Pet{name: name}
}

func (p Pet) GetName() string {
	return p.Name
}
