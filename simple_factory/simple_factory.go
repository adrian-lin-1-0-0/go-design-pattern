package simplefactory

type Animal interface {
	Say() string
}

type Dog struct{}

func (dog *Dog) Say() string {
	return "woof"
}

type Cat struct{}

func (cat *Cat) Say() string {
	return "meow"
}

type Human struct{}

func (human *Human) Say() string {
	return "hi"
}

func NewAnimal(animalName string) Animal {
	switch animalName {
	case "dog":
		return &Dog{}
	case "cat":
		return &Cat{}
	default:
		return &Human{}
	}
}
