//go:build wasm

package sdl

import (
  "unsafe"
)



//go:wasmimport sdl3 SDL_GetPreferredLocales
func __SDL_GetPreferredLocales(uintptr) uintptr



func __gowrap__SDL_GetPreferredLocales(count *int32) (__result **Locale) {
	__result = (**Locale)(unsafe.Pointer(__SDL_GetPreferredLocales(uintptr(unsafe.Pointer(count)))))
	return
}

 
func init() {
	GetPreferredLocales = __gowrap__SDL_GetPreferredLocales
}
