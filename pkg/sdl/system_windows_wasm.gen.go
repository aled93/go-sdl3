//go:build wasm

package sdl

import (
  "unsafe"
)



//go:wasmimport sdl3 SDL_SetWindowsMessageHook
func __SDL_SetWindowsMessageHook(uintptr, uintptr)



func __gowrap__SDL_SetWindowsMessageHook(callback WindowsMessageHook, userdata uintptr) {
	__SDL_SetWindowsMessageHook(0 /* TODO: callbacks */, uintptr(unsafe.Pointer(userdata)))
}

 
func init() {
	SetWindowsMessageHook = __gowrap__SDL_SetWindowsMessageHook
}
