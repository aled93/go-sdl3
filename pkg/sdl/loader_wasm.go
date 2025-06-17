//go:build wasm

package sdl

func LoadLibrary() error {
	return LoadLibraryFromFile("sdl3")
}

func LoadLibraryFromFile(libpath string) error {
	// in wasm we "statically" link sdl3
	return nil
}

type externFunc struct {
	Func any
	Name string
}

func registerLoaderFunc(f func() []externFunc) {
	//
}
