//go:build wasm

package sdl

import (
  "unsafe"
)



//go:wasmimport sdl3 SDL_SetiOSAnimationCallback
func __SDL_SetiOSAnimationCallback(uintptr, int32, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_SetiOSEventPump
func __SDL_SetiOSEventPump(int32)

//go:wasmimport sdl3 SDL_OnApplicationDidChangeStatusBarOrientation
func __SDL_OnApplicationDidChangeStatusBarOrientation()



func __gowrap__SDL_SetiOSAnimationCallback(window Window, interval int32, callback IOSAnimationCallback, callbackParam uintptr) (__result bool) {
	__result = (bool)(__SDL_SetiOSAnimationCallback(uintptr(unsafe.Pointer(window)), *(*int32)(unsafe.Pointer(&interval)), 0 /* TODO: callbacks */, uintptr(unsafe.Pointer(callbackParam))) != 0)
	return
}

func __gowrap__SDL_SetiOSEventPump(enabled bool) {
	__SDL_SetiOSEventPump(*(*int32)(unsafe.Pointer(&enabled)))
}

func __gowrap__SDL_OnApplicationDidChangeStatusBarOrientation() {
	__SDL_OnApplicationDidChangeStatusBarOrientation()
}

 
func init() {
	SetiOSAnimationCallback = __gowrap__SDL_SetiOSAnimationCallback
	SetiOSEventPump = __gowrap__SDL_SetiOSEventPump
	OnApplicationDidChangeStatusBarOrientation = __gowrap__SDL_OnApplicationDidChangeStatusBarOrientation
}
