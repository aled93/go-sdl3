//go:build wasm

package sdl

import (
  "runtime"
  "unsafe"
)



//go:wasmimport sdl3 SDL_OpenURL
func __SDL_OpenURL(uintptr) int32



func __gowrap__SDL_OpenURL(url string) (__result bool) {
	__result = (bool)(__SDL_OpenURL(((*[2]uintptr)(unsafe.Pointer(&url)))[0]) != 0)
	runtime.KeepAlive(url)
	return
}

 
func init() {
	OpenURL = __gowrap__SDL_OpenURL
}
