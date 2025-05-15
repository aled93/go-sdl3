//go:build !windows

package sdl

import "github.com/ebitengine/purego"

func loadLib(name string) (uintptr, error) {
	return purego.Dlopen(name, purego.RTLD_NOW|purego.RTLD_GLOBAL)
}
