//go:build wasm && GDK

package sdl

import (
  "unsafe"
)



//go:wasmimport sdl3 SDL_GetGDKTaskQueue
func __SDL_GetGDKTaskQueue(uintptr) int32

//go:wasmimport sdl3 SDL_GetGDKDefaultUser
func __SDL_GetGDKDefaultUser(uintptr) int32



func __gowrap__SDL_GetGDKTaskQueue(outTaskQueue *XTaskQueueHandle) (__result bool) {
	__result = (bool)(__SDL_GetGDKTaskQueue(uintptr(unsafe.Pointer(outTaskQueue))) != 0)
	return
}

func __gowrap__SDL_GetGDKDefaultUser(outUserHandle *XUserHandle) (__result bool) {
	__result = (bool)(__SDL_GetGDKDefaultUser(uintptr(unsafe.Pointer(outUserHandle))) != 0)
	return
}

 
func init() {
	GetGDKTaskQueue = __gowrap__SDL_GetGDKTaskQueue
	GetGDKDefaultUser = __gowrap__SDL_GetGDKDefaultUser
}
