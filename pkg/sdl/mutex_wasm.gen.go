//go:build wasm

package sdl

import (
  "unsafe"
)



//go:wasmimport sdl3 SDL_ShouldInit
func __SDL_ShouldInit(uintptr) int32

//go:wasmimport sdl3 SDL_ShouldQuit
func __SDL_ShouldQuit(uintptr) int32

//go:wasmimport sdl3 SDL_SetInitialized
func __SDL_SetInitialized(uintptr, int32)



func __gowrap__SDL_ShouldInit(state *InitState) (__result bool) {
	__result = (bool)(__SDL_ShouldInit(uintptr(unsafe.Pointer(state))) != 0)
	return
}

func __gowrap__SDL_ShouldQuit(state *InitState) (__result bool) {
	__result = (bool)(__SDL_ShouldQuit(uintptr(unsafe.Pointer(state))) != 0)
	return
}

func __gowrap__SDL_SetInitialized(state *InitState, initialized bool) {
	__SDL_SetInitialized(uintptr(unsafe.Pointer(state)), *(*int32)(unsafe.Pointer(&initialized)))
}

 
func init() {
	ShouldInit = __gowrap__SDL_ShouldInit
	ShouldQuit = __gowrap__SDL_ShouldQuit
	SetInitialized = __gowrap__SDL_SetInitialized
}
