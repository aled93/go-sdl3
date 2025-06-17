//go:build wasm

package sdl

import (
  "unsafe"
)



//go:wasmimport sdl3 SDL_GetTicks
func __SDL_GetTicks() int64

//go:wasmimport sdl3 SDL_GetTicksNS
func __SDL_GetTicksNS() int64

//go:wasmimport sdl3 SDL_GetPerformanceCounter
func __SDL_GetPerformanceCounter() int64

//go:wasmimport sdl3 SDL_GetPerformanceFrequency
func __SDL_GetPerformanceFrequency() int64

//go:wasmimport sdl3 SDL_Delay
func __SDL_Delay(int32)

//go:wasmimport sdl3 SDL_DelayNS
func __SDL_DelayNS(int64)

//go:wasmimport sdl3 SDL_DelayPrecise
func __SDL_DelayPrecise(int64)

//go:wasmimport sdl3 SDL_AddTimer
func __SDL_AddTimer(int32, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_AddTimerNS
func __SDL_AddTimerNS(int64, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_RemoveTimer
func __SDL_RemoveTimer(int32) int32



func __gowrap__SDL_GetTicks() (__result uint64) {
	__result = (uint64)(__SDL_GetTicks())
	return
}

func __gowrap__SDL_GetTicksNS() (__result uint64) {
	__result = (uint64)(__SDL_GetTicksNS())
	return
}

func __gowrap__SDL_GetPerformanceCounter() (__result uint64) {
	__result = (uint64)(__SDL_GetPerformanceCounter())
	return
}

func __gowrap__SDL_GetPerformanceFrequency() (__result uint64) {
	__result = (uint64)(__SDL_GetPerformanceFrequency())
	return
}

func __gowrap__SDL_Delay(ms uint32) {
	__SDL_Delay(*(*int32)(unsafe.Pointer(&ms)))
}

func __gowrap__SDL_DelayNS(ns uint64) {
	__SDL_DelayNS(*(*int64)(unsafe.Pointer(&ns)))
}

func __gowrap__SDL_DelayPrecise(ns uint64) {
	__SDL_DelayPrecise(*(*int64)(unsafe.Pointer(&ns)))
}

func __gowrap__SDL_AddTimer(interval uint32, callback TimerCallback, userdata uintptr) (__result TimerID) {
	__result = (TimerID)(__SDL_AddTimer(*(*int32)(unsafe.Pointer(&interval)), 0 /* TODO: callbacks */, uintptr(unsafe.Pointer(userdata))))
	return
}

func __gowrap__SDL_AddTimerNS(interval uint64, callback NSTimerCallback, userdata uintptr) (__result TimerID) {
	__result = (TimerID)(__SDL_AddTimerNS(*(*int64)(unsafe.Pointer(&interval)), 0 /* TODO: callbacks */, uintptr(unsafe.Pointer(userdata))))
	return
}

func __gowrap__SDL_RemoveTimer(id TimerID) (__result bool) {
	__result = (bool)(__SDL_RemoveTimer(*(*int32)(unsafe.Pointer(&id))) != 0)
	return
}

 
func init() {
	GetTicks = __gowrap__SDL_GetTicks
	GetTicksNS = __gowrap__SDL_GetTicksNS
	GetPerformanceCounter = __gowrap__SDL_GetPerformanceCounter
	GetPerformanceFrequency = __gowrap__SDL_GetPerformanceFrequency
	Delay = __gowrap__SDL_Delay
	DelayNS = __gowrap__SDL_DelayNS
	DelayPrecise = __gowrap__SDL_DelayPrecise
	AddTimer = __gowrap__SDL_AddTimer
	AddTimerNS = __gowrap__SDL_AddTimerNS
	RemoveTimer = __gowrap__SDL_RemoveTimer
}
