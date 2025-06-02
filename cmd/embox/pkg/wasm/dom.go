package wasm

type ElementId struct {
	HasValue bool
	Name     string
	Index    uint32
}

type Module struct {
	Typedefs   []TypeDef
	Imports    []Import
	Functions  []Function
	Tables     []Table
	Memories   []Memory
	Globals    []Global
	Exports    []Export
	StartFunc  *ElementId
	TableElems []TableElem
	Datas      []Data
}

type Import struct {
	ModuleName string
	SymbolName string
	Type       ImportableType
}

type Export struct {
	SymbolName string
	Type       ExportableType
}

type FuncSignature struct {
	Id      ElementId
	Params  []Param
	Results []Result
}

type Function struct {
	FuncSignature
	Body Expression
}

type Expression struct {
	Instructions []Instruction
}

type Instruction struct {
	Opcode    uint32
	Operands  [4]uint32
	BrTargets []uint32
}

type TypeDef struct {
	Id       ElementId
	FuncType FuncSignature
}

type Table struct {
	Id      ElementId
	Limits  Limits
	Element RefType
}

type TableElem struct {
	Id            ElementId
	IsDeclarative bool
	Offset        Expression
	TableId       ElementId
	RefType       RefType
	Items         []Expression
}

type Limits struct {
	Min     uint32
	Max     uint32
	HaveMax bool
}

type Memory struct {
	Limits Limits
}

type Global struct {
	Type    ValueType
	Mutable bool
	Value   Expression
}

type Data struct {
	Id     ElementId
	Memory ElementId
	Offset Expression
	Value  string
}
