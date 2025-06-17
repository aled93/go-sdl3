//go:build wasm

package sdl

import (
  "runtime"
  "unsafe"
)



//go:wasmimport sdl3 SDL_HasMouse
func __SDL_HasMouse() int32

//go:wasmimport sdl3 SDL_GetMice
func __SDL_GetMice(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetMouseNameForID
func __SDL_GetMouseNameForID(int32) uintptr

//go:wasmimport sdl3 SDL_GetMouseFocus
func __SDL_GetMouseFocus() uintptr

//go:wasmimport sdl3 SDL_GetMouseState
func __SDL_GetMouseState(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetGlobalMouseState
func __SDL_GetGlobalMouseState(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetRelativeMouseState
func __SDL_GetRelativeMouseState(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_WarpMouseInWindow
func __SDL_WarpMouseInWindow(uintptr, float32, float32)

//go:wasmimport sdl3 SDL_WarpMouseGlobal
func __SDL_WarpMouseGlobal(float32, float32) int32

//go:wasmimport sdl3 SDL_SetRelativeMouseTransform
func __SDL_SetRelativeMouseTransform(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_SetWindowRelativeMouseMode
func __SDL_SetWindowRelativeMouseMode(uintptr, int32) int32

//go:wasmimport sdl3 SDL_GetWindowRelativeMouseMode
func __SDL_GetWindowRelativeMouseMode(uintptr) int32

//go:wasmimport sdl3 SDL_CaptureMouse
func __SDL_CaptureMouse(int32) int32

//go:wasmimport sdl3 SDL_CreateCursor
func __SDL_CreateCursor(uintptr, uintptr, int32, int32, int32, int32) uintptr

//go:wasmimport sdl3 SDL_CreateColorCursor
func __SDL_CreateColorCursor(uintptr, int32, int32) uintptr

//go:wasmimport sdl3 SDL_CreateSystemCursor
func __SDL_CreateSystemCursor(int32) uintptr

//go:wasmimport sdl3 SDL_SetCursor
func __SDL_SetCursor(uintptr) int32

//go:wasmimport sdl3 SDL_GetCursor
func __SDL_GetCursor() uintptr

//go:wasmimport sdl3 SDL_GetDefaultCursor
func __SDL_GetDefaultCursor() uintptr

//go:wasmimport sdl3 SDL_DestroyCursor
func __SDL_DestroyCursor(uintptr)

//go:wasmimport sdl3 SDL_ShowCursor
func __SDL_ShowCursor() int32

//go:wasmimport sdl3 SDL_HideCursor
func __SDL_HideCursor() int32

//go:wasmimport sdl3 SDL_CursorVisible
func __SDL_CursorVisible() int32



func __gowrap__SDL_HasMouse() (__result bool) {
	__result = (bool)(__SDL_HasMouse() != 0)
	return
}

func __gowrap__SDL_GetMice(count *int32) (__result *MouseID) {
	__result = (*MouseID)(unsafe.Pointer(__SDL_GetMice(uintptr(unsafe.Pointer(count)))))
	return
}

func __gowrap__SDL_GetMouseNameForID(instance_id MouseID) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetMouseNameForID(*(*int32)(unsafe.Pointer(&instance_id)))).Collect()))
	return
}

func __gowrap__SDL_GetMouseFocus() (__result Window) {
	__result = (Window)(unsafe.Pointer(__SDL_GetMouseFocus()))
	return
}

func __gowrap__SDL_GetMouseState(x *float32, y *float32) (__result MouseButtonFlags) {
	__result = (MouseButtonFlags)(__SDL_GetMouseState(uintptr(unsafe.Pointer(x)), uintptr(unsafe.Pointer(y))))
	return
}

func __gowrap__SDL_GetGlobalMouseState(x *float32, y *float32) (__result MouseButtonFlags) {
	__result = (MouseButtonFlags)(__SDL_GetGlobalMouseState(uintptr(unsafe.Pointer(x)), uintptr(unsafe.Pointer(y))))
	return
}

func __gowrap__SDL_GetRelativeMouseState(x *float32, y *float32) (__result MouseButtonFlags) {
	__result = (MouseButtonFlags)(__SDL_GetRelativeMouseState(uintptr(unsafe.Pointer(x)), uintptr(unsafe.Pointer(y))))
	return
}

func __gowrap__SDL_WarpMouseInWindow(window Window, x float32, y float32) {
	__SDL_WarpMouseInWindow(uintptr(unsafe.Pointer(window)), *(*float32)(unsafe.Pointer(&x)), *(*float32)(unsafe.Pointer(&y)))
}

func __gowrap__SDL_WarpMouseGlobal(x float32, y float32) (__result bool) {
	__result = (bool)(__SDL_WarpMouseGlobal(*(*float32)(unsafe.Pointer(&x)), *(*float32)(unsafe.Pointer(&y))) != 0)
	return
}

func __gowrap__SDL_SetRelativeMouseTransform(callback MouseMotionTransformCallback, userdata uintptr) (__result bool) {
	__result = (bool)(__SDL_SetRelativeMouseTransform(0 /* TODO: callbacks */, uintptr(unsafe.Pointer(userdata))) != 0)
	return
}

func __gowrap__SDL_SetWindowRelativeMouseMode(window Window, enabled bool) (__result bool) {
	__result = (bool)(__SDL_SetWindowRelativeMouseMode(uintptr(unsafe.Pointer(window)), *(*int32)(unsafe.Pointer(&enabled))) != 0)
	return
}

func __gowrap__SDL_GetWindowRelativeMouseMode(window Window) (__result bool) {
	__result = (bool)(__SDL_GetWindowRelativeMouseMode(uintptr(unsafe.Pointer(window))) != 0)
	return
}

func __gowrap__SDL_CaptureMouse(enabled bool) (__result bool) {
	__result = (bool)(__SDL_CaptureMouse(*(*int32)(unsafe.Pointer(&enabled))) != 0)
	return
}

func __gowrap__SDL_CreateCursor(data []byte, mask []byte, w int32, h int32, hot_x int32, hot_y int32) (__result Cursor) {
	__result = (Cursor)(unsafe.Pointer(__SDL_CreateCursor(uintptr(unsafe.Pointer(&data[0])), uintptr(unsafe.Pointer(&mask[0])), *(*int32)(unsafe.Pointer(&w)), *(*int32)(unsafe.Pointer(&h)), *(*int32)(unsafe.Pointer(&hot_x)), *(*int32)(unsafe.Pointer(&hot_y)))))
	runtime.KeepAlive(data)
	runtime.KeepAlive(mask)
	return
}

func __gowrap__SDL_CreateColorCursor(surface *Surface, hot_x int32, hot_y int32) (__result Cursor) {
	__result = (Cursor)(unsafe.Pointer(__SDL_CreateColorCursor(uintptr(unsafe.Pointer(surface)), *(*int32)(unsafe.Pointer(&hot_x)), *(*int32)(unsafe.Pointer(&hot_y)))))
	return
}

func __gowrap__SDL_CreateSystemCursor(id SystemCursor) (__result Cursor) {
	__result = (Cursor)(unsafe.Pointer(__SDL_CreateSystemCursor(*(*int32)(unsafe.Pointer(&id)))))
	return
}

func __gowrap__SDL_SetCursor(cursor Cursor) (__result bool) {
	__result = (bool)(__SDL_SetCursor(uintptr(unsafe.Pointer(cursor))) != 0)
	return
}

func __gowrap__SDL_GetCursor() (__result Cursor) {
	__result = (Cursor)(unsafe.Pointer(__SDL_GetCursor()))
	return
}

func __gowrap__SDL_GetDefaultCursor() (__result Cursor) {
	__result = (Cursor)(unsafe.Pointer(__SDL_GetDefaultCursor()))
	return
}

func __gowrap__SDL_DestroyCursor(cursor Cursor) {
	__SDL_DestroyCursor(uintptr(unsafe.Pointer(cursor)))
}

func __gowrap__SDL_ShowCursor() (__result bool) {
	__result = (bool)(__SDL_ShowCursor() != 0)
	return
}

func __gowrap__SDL_HideCursor() (__result bool) {
	__result = (bool)(__SDL_HideCursor() != 0)
	return
}

func __gowrap__SDL_CursorVisible() (__result bool) {
	__result = (bool)(__SDL_CursorVisible() != 0)
	return
}

 
func init() {
	HasMouse = __gowrap__SDL_HasMouse
	GetMice = __gowrap__SDL_GetMice
	GetMouseNameForID = __gowrap__SDL_GetMouseNameForID
	GetMouseFocus = __gowrap__SDL_GetMouseFocus
	GetMouseState = __gowrap__SDL_GetMouseState
	GetGlobalMouseState = __gowrap__SDL_GetGlobalMouseState
	GetRelativeMouseState = __gowrap__SDL_GetRelativeMouseState
	WarpMouseInWindow = __gowrap__SDL_WarpMouseInWindow
	WarpMouseGlobal = __gowrap__SDL_WarpMouseGlobal
	SetRelativeMouseTransform = __gowrap__SDL_SetRelativeMouseTransform
	SetWindowRelativeMouseMode = __gowrap__SDL_SetWindowRelativeMouseMode
	GetWindowRelativeMouseMode = __gowrap__SDL_GetWindowRelativeMouseMode
	CaptureMouse = __gowrap__SDL_CaptureMouse
	CreateCursor = __gowrap__SDL_CreateCursor
	CreateColorCursor = __gowrap__SDL_CreateColorCursor
	CreateSystemCursor = __gowrap__SDL_CreateSystemCursor
	SetCursor = __gowrap__SDL_SetCursor
	GetCursor = __gowrap__SDL_GetCursor
	GetDefaultCursor = __gowrap__SDL_GetDefaultCursor
	DestroyCursor = __gowrap__SDL_DestroyCursor
	ShowCursor = __gowrap__SDL_ShowCursor
	HideCursor = __gowrap__SDL_HideCursor
	CursorVisible = __gowrap__SDL_CursorVisible
}
