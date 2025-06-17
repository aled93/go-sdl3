//go:build wasm

package sdl



//go:wasmimport sdl3 SDL_GetVersion
func __SDL_GetVersion() int32

//go:wasmimport sdl3 SDL_GetRevision
func __SDL_GetRevision() uintptr



func __gowrap__SDL_GetVersion() (__result int32) {
	__result = (int32)(__SDL_GetVersion())
	return
}

func __gowrap__SDL_GetRevision() (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetRevision()).Collect()))
	return
}

 
func init() {
	GetVersion = __gowrap__SDL_GetVersion
	GetRevision = __gowrap__SDL_GetRevision
}
