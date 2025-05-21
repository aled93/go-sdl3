//go:build !cgo && !wasm

package sdl

import "github.com/ebitengine/purego"

type puregoLibrary uintptr

func (lib puregoLibrary) BindProc(fn any, name string) {
	purego.RegisterLibFunc(fn, uintptr(lib), name)
}
