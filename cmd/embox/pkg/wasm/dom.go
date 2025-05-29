package wasm

type Module struct {
	Imports   []Import
	Exports   []Export
	Functions []Function
}

type Import struct {
	ModuleName string
	SymbolName string
	Type       Type
}

type Export struct {
	SymbolName string
	Type       Type
}

type Data struct {
	TargetMemory any
}

type Type interface {
	isType()
}
