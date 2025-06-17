//go:build wasm

package sdl

import (
  "runtime"
  "unsafe"
)



//go:wasmimport sdl3 SDL_SetClipboardText
func __SDL_SetClipboardText(uintptr) int32

//go:wasmimport sdl3 SDL_GetClipboardText
func __SDL_GetClipboardText() uintptr

//go:wasmimport sdl3 SDL_HasClipboardText
func __SDL_HasClipboardText() int32

//go:wasmimport sdl3 SDL_SetPrimarySelectionText
func __SDL_SetPrimarySelectionText(uintptr) int32

//go:wasmimport sdl3 SDL_GetPrimarySelectionText
func __SDL_GetPrimarySelectionText() uintptr

//go:wasmimport sdl3 SDL_HasPrimarySelectionText
func __SDL_HasPrimarySelectionText() int32

//go:wasmimport sdl3 SDL_SetClipboardData
func __SDL_SetClipboardData(uintptr, uintptr, uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_ClearClipboardData
func __SDL_ClearClipboardData() int32

//go:wasmimport sdl3 SDL_GetClipboardData
func __SDL_GetClipboardData(uintptr, uintptr) uintptr

//go:wasmimport sdl3 SDL_HasClipboardData
func __SDL_HasClipboardData(uintptr) int32

//go:wasmimport sdl3 SDL_GetClipboardMimeTypes
func __SDL_GetClipboardMimeTypes(uintptr) uintptr



func __gowrap__SDL_SetClipboardText(text string) (__result bool) {
	__result = (bool)(__SDL_SetClipboardText(((*[2]uintptr)(unsafe.Pointer(&text)))[0]) != 0)
	runtime.KeepAlive(text)
	return
}

func __gowrap__SDL_GetClipboardText() (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetClipboardText()).Collect()))
	return
}

func __gowrap__SDL_HasClipboardText() (__result bool) {
	__result = (bool)(__SDL_HasClipboardText() != 0)
	return
}

func __gowrap__SDL_SetPrimarySelectionText(text string) (__result bool) {
	__result = (bool)(__SDL_SetPrimarySelectionText(((*[2]uintptr)(unsafe.Pointer(&text)))[0]) != 0)
	runtime.KeepAlive(text)
	return
}

func __gowrap__SDL_GetPrimarySelectionText() (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetPrimarySelectionText()).Collect()))
	return
}

func __gowrap__SDL_HasPrimarySelectionText() (__result bool) {
	__result = (bool)(__SDL_HasPrimarySelectionText() != 0)
	return
}

func __gowrap__SDL_SetClipboardData(callback ClipboardDataCallback, cleanup ClipboardCleanupCallback, userdata uintptr, mime_types []string, num_mime_types size_t) (__result bool) {
	__result = (bool)(__SDL_SetClipboardData(0 /* TODO: callbacks */, 0 /* TODO: callbacks */, uintptr(unsafe.Pointer(userdata)), uintptr(unsafe.Pointer(&mime_types[0])), *(*int32)(unsafe.Pointer(&num_mime_types))) != 0)
	runtime.KeepAlive(mime_types)
	return
}

func __gowrap__SDL_ClearClipboardData() (__result bool) {
	__result = (bool)(__SDL_ClearClipboardData() != 0)
	return
}

func __gowrap__SDL_GetClipboardData(mime_type string, size *size_t) (__result uintptr) {
	__result = (uintptr)(unsafe.Pointer(__SDL_GetClipboardData(((*[2]uintptr)(unsafe.Pointer(&mime_type)))[0], uintptr(unsafe.Pointer(size)))))
	runtime.KeepAlive(mime_type)
	return
}

func __gowrap__SDL_HasClipboardData(mime_type string) (__result bool) {
	__result = (bool)(__SDL_HasClipboardData(((*[2]uintptr)(unsafe.Pointer(&mime_type)))[0]) != 0)
	runtime.KeepAlive(mime_type)
	return
}

func __gowrap__SDL_GetClipboardMimeTypes(num_mime_types *size_t) (__result []string) {
	__result = ([]string)(newCSliceRaw[string](__SDL_GetClipboardMimeTypes(uintptr(unsafe.Pointer(num_mime_types)))).Collect())
	return
}

 
func init() {
	SetClipboardText = __gowrap__SDL_SetClipboardText
	GetClipboardText = __gowrap__SDL_GetClipboardText
	HasClipboardText = __gowrap__SDL_HasClipboardText
	SetPrimarySelectionText = __gowrap__SDL_SetPrimarySelectionText
	GetPrimarySelectionText = __gowrap__SDL_GetPrimarySelectionText
	HasPrimarySelectionText = __gowrap__SDL_HasPrimarySelectionText
	SetClipboardData = __gowrap__SDL_SetClipboardData
	ClearClipboardData = __gowrap__SDL_ClearClipboardData
	GetClipboardData = __gowrap__SDL_GetClipboardData
	HasClipboardData = __gowrap__SDL_HasClipboardData
	GetClipboardMimeTypes = __gowrap__SDL_GetClipboardMimeTypes
}
