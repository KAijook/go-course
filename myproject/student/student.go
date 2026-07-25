package student

type Student struct {
	Name string
	Age  int
}

func (s Student) GetName() string {
	return s.Name
}

func (s Student) GetAge() int {
	return s.Age
}

type StudentInterface interface {
	GetName() string
	GetAge() int
}
