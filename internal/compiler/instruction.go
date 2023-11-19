package compiler

type Instruction struct {
	Mnemonic string
}

func NewInstruction(mnemonic string) Instruction {
	return Instruction{Mnemonic: mnemonic}
}

func (i Instruction) Compile() string {
	return i.Mnemonic
}
