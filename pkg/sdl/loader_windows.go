package sdl

import "golang.org/x/sys/windows"

func loadLib(name string) (uintptr, error) {
	h, err := windows.LoadLibrary(name)
	return uintptr(h), err
}
