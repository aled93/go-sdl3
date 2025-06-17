//go:build wasm

package sdl

import (
  "runtime"
  "unsafe"
)



//go:wasmimport sdl3 SDL_ShowMessageBox
func __SDL_ShowMessageBox(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_ShowSimpleMessageBox
func __SDL_ShowSimpleMessageBox(int32, uintptr, uintptr, uintptr) int32



func __gowrap__SDL_ShowMessageBox(messageboxdata *MessageBoxData, buttonid *int32) (__result bool) {
	__result = (bool)(__SDL_ShowMessageBox(uintptr(unsafe.Pointer(messageboxdata)), uintptr(unsafe.Pointer(buttonid))) != 0)
	return
}

func __gowrap__SDL_ShowSimpleMessageBox(flags MessageBoxFlags, title string, message string, window Window) (__result bool) {
	__result = (bool)(__SDL_ShowSimpleMessageBox(*(*int32)(unsafe.Pointer(&flags)), ((*[2]uintptr)(unsafe.Pointer(&title)))[0], ((*[2]uintptr)(unsafe.Pointer(&message)))[0], uintptr(unsafe.Pointer(window))) != 0)
	runtime.KeepAlive(title)
	runtime.KeepAlive(message)
	return
}

 
func init() {
	ShowMessageBox = __gowrap__SDL_ShowMessageBox
	ShowSimpleMessageBox = __gowrap__SDL_ShowSimpleMessageBox
}
