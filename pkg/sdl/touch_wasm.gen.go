//go:build wasm

package sdl

import (
  "unsafe"
)



//go:wasmimport sdl3 SDL_GetTouchDevices
func __SDL_GetTouchDevices(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetTouchDeviceName
func __SDL_GetTouchDeviceName(int64) uintptr

//go:wasmimport sdl3 SDL_GetTouchDeviceType
func __SDL_GetTouchDeviceType(int64) int32

//go:wasmimport sdl3 SDL_GetTouchFingers
func __SDL_GetTouchFingers(int64, uintptr) uintptr



func __gowrap__SDL_GetTouchDevices(count *int32) (__result *TouchID) {
	__result = (*TouchID)(unsafe.Pointer(__SDL_GetTouchDevices(uintptr(unsafe.Pointer(count)))))
	return
}

func __gowrap__SDL_GetTouchDeviceName(touchID TouchID) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetTouchDeviceName(*(*int64)(unsafe.Pointer(&touchID)))).Collect()))
	return
}

func __gowrap__SDL_GetTouchDeviceType(touchID TouchID) (__result TouchDeviceType) {
	__result = (TouchDeviceType)(__SDL_GetTouchDeviceType(*(*int64)(unsafe.Pointer(&touchID))))
	return
}

func __gowrap__SDL_GetTouchFingers(touchID TouchID, count *int32) (__result **Finger) {
	__result = (**Finger)(unsafe.Pointer(__SDL_GetTouchFingers(*(*int64)(unsafe.Pointer(&touchID)), uintptr(unsafe.Pointer(count)))))
	return
}

 
func init() {
	GetTouchDevices = __gowrap__SDL_GetTouchDevices
	GetTouchDeviceName = __gowrap__SDL_GetTouchDeviceName
	GetTouchDeviceType = __gowrap__SDL_GetTouchDeviceType
	GetTouchFingers = __gowrap__SDL_GetTouchFingers
}
