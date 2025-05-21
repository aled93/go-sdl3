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
		externs := ldr()
		for _, extern := range externs {
			lib.BindProc(extern.Func, extern.Name)
		}
	}

	return nil
}

type externFunc struct {
	Func any
	Name string
}

var loaderFuncs []func() []externFunc

func registerLoaderFunc(f func() []externFunc) {
	loaderFuncs = append(loaderFuncs, f)
}

type library interface {
	BindProc(fn any, name string)
}
