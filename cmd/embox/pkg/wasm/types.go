package wasm

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

type Function struct {
	Params  []Param
	Results []Result
}

type Param struct {
	Name string
	Type Type
}

type Result struct {
	Name string
	Type Type
}

type Limits struct {
	Min     uint32
	Max     uint32
	HaveMax bool
}

type Memory struct {
	Limits Limits
}

type Table struct {
	Limits  Limits
	Element RefType
}

type GlobalType struct {
	Type    ValueType
	Mutable bool
}

func (Function) isType()   {}
func (Memory) isType()     {}
func (Table) isType()      {}
func (GlobalType) isType() {}
