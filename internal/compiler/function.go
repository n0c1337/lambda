package compiler

import "fmt"

type Function struct {
	Name string
}

func NewFunction(name string) Function {
	return Function{Name: name}
}

func (f Function) Compile() string {
	return fmt.Sprintf("%s:", f.Name)
}
