//go:build wasm

package sdl

import (
  "unsafe"
)



//go:wasmimport sdl3 SDL_Metal_CreateView
func __SDL_Metal_CreateView(uintptr) uintptr

//go:wasmimport sdl3 SDL_Metal_DestroyView
func __SDL_Metal_DestroyView(uintptr)

//go:wasmimport sdl3 SDL_Metal_GetLayer
func __SDL_Metal_GetLayer(uintptr) uintptr



func __gowrap__SDL_Metal_CreateView(window Window) (__result MetalView) {
	__result = (MetalView)(unsafe.Pointer(__SDL_Metal_CreateView(uintptr(unsafe.Pointer(window)))))
	return
}

func __gowrap__SDL_Metal_DestroyView(view MetalView) {
	__SDL_Metal_DestroyView(uintptr(unsafe.Pointer(view)))
}

func __gowrap__SDL_Metal_GetLayer(view MetalView) (__result uintptr) {
	__result = (uintptr)(unsafe.Pointer(__SDL_Metal_GetLayer(uintptr(unsafe.Pointer(view)))))
	return
}

 
func init() {
	Metal_CreateView = __gowrap__SDL_Metal_CreateView
	Metal_DestroyView = __gowrap__SDL_Metal_DestroyView
	Metal_GetLayer = __gowrap__SDL_Metal_GetLayer
}
