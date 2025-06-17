//go:build wasm

package sdl

import (
  "unsafe"
)



//go:wasmimport sdl3 SDL_SetX11EventHook
func __SDL_SetX11EventHook(uintptr, uintptr)

//go:wasmimport sdl3 SDL_SetLinuxThreadPriority
func __SDL_SetLinuxThreadPriority(int64, int32) int32

//go:wasmimport sdl3 SDL_SetLinuxThreadPriorityAndPolicy
func __SDL_SetLinuxThreadPriorityAndPolicy(int64, int32, int32) int32



func __gowrap__SDL_SetX11EventHook(callback X11EventHook, userdata uintptr) {
	__SDL_SetX11EventHook(0 /* TODO: callbacks */, uintptr(unsafe.Pointer(userdata)))
}

func __gowrap__SDL_SetLinuxThreadPriority(threadID int64, priority int32) (__result bool) {
	__result = (bool)(__SDL_SetLinuxThreadPriority(*(*int64)(unsafe.Pointer(&threadID)), *(*int32)(unsafe.Pointer(&priority))) != 0)
	return
}

func __gowrap__SDL_SetLinuxThreadPriorityAndPolicy(threadID int64, sdlPriority int32, schedPolicy int32) (__result bool) {
	__result = (bool)(__SDL_SetLinuxThreadPriorityAndPolicy(*(*int64)(unsafe.Pointer(&threadID)), *(*int32)(unsafe.Pointer(&sdlPriority)), *(*int32)(unsafe.Pointer(&schedPolicy))) != 0)
	return
}

 
func init() {
	SetX11EventHook = __gowrap__SDL_SetX11EventHook
	SetLinuxThreadPriority = __gowrap__SDL_SetLinuxThreadPriority
	SetLinuxThreadPriorityAndPolicy = __gowrap__SDL_SetLinuxThreadPriorityAndPolicy
}
