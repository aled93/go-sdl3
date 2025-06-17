//go:build wasm

package sdl



//go:wasmimport sdl3 SDL_OutOfMemory
func __SDL_OutOfMemory() int32

//go:wasmimport sdl3 SDL_GetError
func __SDL_GetError() uintptr

//go:wasmimport sdl3 SDL_ClearError
func __SDL_ClearError() int32



func __gowrap__SDL_OutOfMemory() (__result bool) {
	__result = (bool)(__SDL_OutOfMemory() != 0)
	return
}

func __gowrap__SDL_GetError() (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetError()).Collect()))
	return
}

func __gowrap__SDL_ClearError() (__result bool) {
	__result = (bool)(__SDL_ClearError() != 0)
	return
}

 
func init() {
	OutOfMemory = __gowrap__SDL_OutOfMemory
	GetError = __gowrap__SDL_GetError
	ClearError = __gowrap__SDL_ClearError
}
