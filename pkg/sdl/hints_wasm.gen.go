//go:build wasm

package sdl

import (
  "runtime"
  "unsafe"
)



//go:wasmimport sdl3 SDL_SetHintWithPriority
func __SDL_SetHintWithPriority(uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_SetHint
func __SDL_SetHint(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_ResetHint
func __SDL_ResetHint(uintptr) int32

//go:wasmimport sdl3 SDL_ResetHints
func __SDL_ResetHints()

//go:wasmimport sdl3 SDL_GetHint
func __SDL_GetHint(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetHintBoolean
func __SDL_GetHintBoolean(uintptr, int32) int32

//go:wasmimport sdl3 SDL_AddHintCallback
func __SDL_AddHintCallback(uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_RemoveHintCallback
func __SDL_RemoveHintCallback(uintptr, uintptr, uintptr)



func __gowrap__SDL_SetHintWithPriority(name string, value string, priority HintPriority) (__result bool) {
	__result = (bool)(__SDL_SetHintWithPriority(((*[2]uintptr)(unsafe.Pointer(&name)))[0], ((*[2]uintptr)(unsafe.Pointer(&value)))[0], *(*int32)(unsafe.Pointer(&priority))) != 0)
	runtime.KeepAlive(name)
	runtime.KeepAlive(value)
	return
}

func __gowrap__SDL_SetHint(name string, value string) (__result bool) {
	__result = (bool)(__SDL_SetHint(((*[2]uintptr)(unsafe.Pointer(&name)))[0], ((*[2]uintptr)(unsafe.Pointer(&value)))[0]) != 0)
	runtime.KeepAlive(name)
	runtime.KeepAlive(value)
	return
}

func __gowrap__SDL_ResetHint(name string) (__result bool) {
	__result = (bool)(__SDL_ResetHint(((*[2]uintptr)(unsafe.Pointer(&name)))[0]) != 0)
	runtime.KeepAlive(name)
	return
}

func __gowrap__SDL_ResetHints() {
	__SDL_ResetHints()
}

func __gowrap__SDL_GetHint(name string) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetHint(((*[2]uintptr)(unsafe.Pointer(&name)))[0])).Collect()))
	runtime.KeepAlive(name)
	return
}

func __gowrap__SDL_GetHintBoolean(name string, default_value bool) (__result bool) {
	__result = (bool)(__SDL_GetHintBoolean(((*[2]uintptr)(unsafe.Pointer(&name)))[0], *(*int32)(unsafe.Pointer(&default_value))) != 0)
	runtime.KeepAlive(name)
	return
}

func __gowrap__SDL_AddHintCallback(name string, callback HintCallback, userdata uintptr) (__result bool) {
	__result = (bool)(__SDL_AddHintCallback(((*[2]uintptr)(unsafe.Pointer(&name)))[0], 0 /* TODO: callbacks */, uintptr(unsafe.Pointer(userdata))) != 0)
	runtime.KeepAlive(name)
	return
}

func __gowrap__SDL_RemoveHintCallback(name string, callback HintCallback, userdata uintptr) {
	__SDL_RemoveHintCallback(((*[2]uintptr)(unsafe.Pointer(&name)))[0], 0 /* TODO: callbacks */, uintptr(unsafe.Pointer(userdata)))
	runtime.KeepAlive(name)
}

 
func init() {
	SetHintWithPriority = __gowrap__SDL_SetHintWithPriority
	SetHint = __gowrap__SDL_SetHint
	ResetHint = __gowrap__SDL_ResetHint
	ResetHints = __gowrap__SDL_ResetHints
	GetHint = __gowrap__SDL_GetHint
	GetHintBoolean = __gowrap__SDL_GetHintBoolean
	AddHintCallback = __gowrap__SDL_AddHintCallback
	RemoveHintCallback = __gowrap__SDL_RemoveHintCallback
}
