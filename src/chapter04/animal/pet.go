package animal

type Pet struct {
	Name string
}

func (p Pet) GetName() string {
	return p.Name
}
