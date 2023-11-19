package compiler

type Variable struct {
	Name string
}

func NewVariable(name string, global bool) Variable {
	return Variable{Name: name}
}
