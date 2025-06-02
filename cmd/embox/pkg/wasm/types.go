package wasm

type Type interface {
	isType()
}

type ImportableType interface {
	isImportableType()
}

type ExportableType interface {
	isExportableType()
}

type NumberType uint8

const (
	Type_i32 NumberType = iota
	Type_i64
	Type_f32
	Type_f64
)

type VectorType uint8

const Type_v128 VectorType = iota

type ValueType interface {
	Type
	isValueType()
}

type RefType uint8

const (
	Type_FuncRef RefType = iota
	Type_ExternRef
)

func (NumberType) isType()      {}
func (NumberType) isValueType() {}
func (VectorType) isType()      {}
func (VectorType) isValueType() {}
func (RefType) isType()         {}
func (RefType) isValueType()    {}

type FuncType struct {
	FuncSignature
}

type Param struct {
	Name string
	Type Type
}

type Result struct {
	Type Type
}

func (FuncType) isType() {}

func (FuncType) isImportableType() {}
func (Memory) isImportableType()   {}
func (Table) isImportableType()    {}
func (Global) isImportableType()   {}

func (FuncType) isExportableType() {}
func (Memory) isExportableType()   {}
func (Table) isExportableType()    {}
func (Global) isExportableType()   {}
