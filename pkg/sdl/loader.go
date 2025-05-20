package sdl

import (
	"errors"
	"runtime"
)

func LoadLibrary() error {
	var libname string
	switch runtime.GOOS {
	case "windows":
		libname = "SDL3.dll"
	case "linux", "freebsd":
		libname = "libSDL3.so.0"
	case "darwin":
		libname = "libSDL3.dylib"
	default:
		return errors.New("unsupported os")
	}

	return LoadLibraryFromFile(libname)
}

func LoadLibraryFromFile(libpath string) error {
	lib, err := loadLib(libpath)
	if err != nil {
		return err
	}

	for _, ldr := range loaderFuncs {
		ldr(lib)
	}

	return nil
}

var loaderFuncs []func(uintptr)

func registerLoaderFunc(f func(uintptr)) {
	loaderFuncs = append(loaderFuncs, f)
}
