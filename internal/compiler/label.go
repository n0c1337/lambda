package compiler

import "fmt"

type Label struct {
	Name string
}

func NewLabel(name string) Label {
	return Label{Name: ".label" + name}
}

func (l Label) Compile() string {
	return fmt.Sprintf("%s:", l.Name)
}
