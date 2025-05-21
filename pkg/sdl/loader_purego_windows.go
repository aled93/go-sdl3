//go:build !cgo && windows

package sdl

import "golang.org/x/sys/windows"

func loadLib(name string) (library, error) {
	h, err := windows.LoadLibrary(name)
	return puregoLibrary(h), err
}
