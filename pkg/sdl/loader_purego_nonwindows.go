//go:build !cgo && !windows

package sdl

import "github.com/ebitengine/purego"

func loadLib(name string) (library, error) {
	lib, err := purego.Dlopen(name, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		return nil, err
	}

	return puregoLibrary(lib), nil
}
