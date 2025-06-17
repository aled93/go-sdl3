//go:build wasm

package sdl

import (
  "runtime"
  "unsafe"
)



//go:wasmimport sdl3 SDL_HasKeyboard
func __SDL_HasKeyboard() int32

//go:wasmimport sdl3 SDL_GetKeyboards
func __SDL_GetKeyboards(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetKeyboardNameForID
func __SDL_GetKeyboardNameForID(int32) uintptr

//go:wasmimport sdl3 SDL_GetKeyboardFocus
func __SDL_GetKeyboardFocus() uintptr

//go:wasmimport sdl3 SDL_GetKeyboardState
func __SDL_GetKeyboardState(uintptr) uintptr

//go:wasmimport sdl3 SDL_ResetKeyboard
func __SDL_ResetKeyboard()

//go:wasmimport sdl3 SDL_GetModState
func __SDL_GetModState() int32

//go:wasmimport sdl3 SDL_SetModState
func __SDL_SetModState(int32)

//go:wasmimport sdl3 SDL_GetKeyFromScancode
func __SDL_GetKeyFromScancode(int32, int32, int32) int32

//go:wasmimport sdl3 SDL_GetScancodeFromKey
func __SDL_GetScancodeFromKey(int32, uintptr) int32

//go:wasmimport sdl3 SDL_SetScancodeName
func __SDL_SetScancodeName(int32, uintptr) int32

//go:wasmimport sdl3 SDL_GetScancodeName
func __SDL_GetScancodeName(int32) uintptr

//go:wasmimport sdl3 SDL_GetScancodeFromName
func __SDL_GetScancodeFromName(uintptr) int32

//go:wasmimport sdl3 SDL_GetKeyName
func __SDL_GetKeyName(int32) uintptr

//go:wasmimport sdl3 SDL_GetKeyFromName
func __SDL_GetKeyFromName(uintptr) int32

//go:wasmimport sdl3 SDL_StartTextInput
func __SDL_StartTextInput(uintptr) int32

//go:wasmimport sdl3 SDL_StartTextInputWithProperties
func __SDL_StartTextInputWithProperties(uintptr, int32) int32

//go:wasmimport sdl3 SDL_TextInputActive
func __SDL_TextInputActive(uintptr) int32

//go:wasmimport sdl3 SDL_StopTextInput
func __SDL_StopTextInput(uintptr) int32

//go:wasmimport sdl3 SDL_ClearComposition
func __SDL_ClearComposition(uintptr) int32

//go:wasmimport sdl3 SDL_SetTextInputArea
func __SDL_SetTextInputArea(uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_GetTextInputArea
func __SDL_GetTextInputArea(uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_HasScreenKeyboardSupport
func __SDL_HasScreenKeyboardSupport() int32

//go:wasmimport sdl3 SDL_ScreenKeyboardShown
func __SDL_ScreenKeyboardShown(uintptr) int32



func __gowrap__SDL_HasKeyboard() (__result bool) {
	__result = (bool)(__SDL_HasKeyboard() != 0)
	return
}

func __gowrap__SDL_GetKeyboards(count *int32) (__result *KeyboardID) {
	__result = (*KeyboardID)(unsafe.Pointer(__SDL_GetKeyboards(uintptr(unsafe.Pointer(count)))))
	return
}

func __gowrap__SDL_GetKeyboardNameForID(instance_id KeyboardID) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetKeyboardNameForID(*(*int32)(unsafe.Pointer(&instance_id)))).Collect()))
	return
}

func __gowrap__SDL_GetKeyboardFocus() (__result Window) {
	__result = (Window)(unsafe.Pointer(__SDL_GetKeyboardFocus()))
	return
}

func __gowrap__SDL_GetKeyboardState(numkeys *int32) (__result *bool) {
	__result = (*bool)(unsafe.Pointer(__SDL_GetKeyboardState(uintptr(unsafe.Pointer(numkeys)))))
	return
}

func __gowrap__SDL_ResetKeyboard() {
	__SDL_ResetKeyboard()
}

func __gowrap__SDL_GetModState() (__result Keymod) {
	__result = (Keymod)(__SDL_GetModState())
	return
}

func __gowrap__SDL_SetModState(modstate Keymod) {
	__SDL_SetModState(*(*int32)(unsafe.Pointer(&modstate)))
}

func __gowrap__SDL_GetKeyFromScancode(scancode Scancode, modstate Keymod, key_event bool) (__result Keycode) {
	__result = (Keycode)(__SDL_GetKeyFromScancode(*(*int32)(unsafe.Pointer(&scancode)), *(*int32)(unsafe.Pointer(&modstate)), *(*int32)(unsafe.Pointer(&key_event))))
	return
}

func __gowrap__SDL_GetScancodeFromKey(key Keycode, modstate *Keymod) (__result Scancode) {
	__result = (Scancode)(__SDL_GetScancodeFromKey(*(*int32)(unsafe.Pointer(&key)), uintptr(unsafe.Pointer(modstate))))
	return
}

func __gowrap__SDL_SetScancodeName(scancode Scancode, name string) (__result bool) {
	__result = (bool)(__SDL_SetScancodeName(*(*int32)(unsafe.Pointer(&scancode)), ((*[2]uintptr)(unsafe.Pointer(&name)))[0]) != 0)
	runtime.KeepAlive(name)
	return
}

func __gowrap__SDL_GetScancodeName(scancode Scancode) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetScancodeName(*(*int32)(unsafe.Pointer(&scancode)))).Collect()))
	return
}

func __gowrap__SDL_GetScancodeFromName(name string) (__result Scancode) {
	__result = (Scancode)(__SDL_GetScancodeFromName(((*[2]uintptr)(unsafe.Pointer(&name)))[0]))
	runtime.KeepAlive(name)
	return
}

func __gowrap__SDL_GetKeyName(key Keycode) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetKeyName(*(*int32)(unsafe.Pointer(&key)))).Collect()))
	return
}

func __gowrap__SDL_GetKeyFromName(name string) (__result Keycode) {
	__result = (Keycode)(__SDL_GetKeyFromName(((*[2]uintptr)(unsafe.Pointer(&name)))[0]))
	runtime.KeepAlive(name)
	return
}

func __gowrap__SDL_StartTextInput(window Window) (__result bool) {
	__result = (bool)(__SDL_StartTextInput(uintptr(unsafe.Pointer(window))) != 0)
	return
}

func __gowrap__SDL_StartTextInputWithProperties(window Window, props PropertiesID) (__result bool) {
	__result = (bool)(__SDL_StartTextInputWithProperties(uintptr(unsafe.Pointer(window)), *(*int32)(unsafe.Pointer(&props))) != 0)
	return
}

func __gowrap__SDL_TextInputActive(window Window) (__result bool) {
	__result = (bool)(__SDL_TextInputActive(uintptr(unsafe.Pointer(window))) != 0)
	return
}

func __gowrap__SDL_StopTextInput(window Window) (__result bool) {
	__result = (bool)(__SDL_StopTextInput(uintptr(unsafe.Pointer(window))) != 0)
	return
}

func __gowrap__SDL_ClearComposition(window Window) (__result bool) {
	__result = (bool)(__SDL_ClearComposition(uintptr(unsafe.Pointer(window))) != 0)
	return
}

func __gowrap__SDL_SetTextInputArea(window Window, rect *Rect, cursor int32) (__result bool) {
	__result = (bool)(__SDL_SetTextInputArea(uintptr(unsafe.Pointer(window)), uintptr(unsafe.Pointer(rect)), *(*int32)(unsafe.Pointer(&cursor))) != 0)
	return
}

func __gowrap__SDL_GetTextInputArea(window Window, rect *Rect, cursor *int32) (__result bool) {
	__result = (bool)(__SDL_GetTextInputArea(uintptr(unsafe.Pointer(window)), uintptr(unsafe.Pointer(rect)), uintptr(unsafe.Pointer(cursor))) != 0)
	return
}

func __gowrap__SDL_HasScreenKeyboardSupport() (__result bool) {
	__result = (bool)(__SDL_HasScreenKeyboardSupport() != 0)
	return
}

func __gowrap__SDL_ScreenKeyboardShown(window Window) (__result bool) {
	__result = (bool)(__SDL_ScreenKeyboardShown(uintptr(unsafe.Pointer(window))) != 0)
	return
}

 
func init() {
	HasKeyboard = __gowrap__SDL_HasKeyboard
	GetKeyboards = __gowrap__SDL_GetKeyboards
	GetKeyboardNameForID = __gowrap__SDL_GetKeyboardNameForID
	GetKeyboardFocus = __gowrap__SDL_GetKeyboardFocus
	GetKeyboardState = __gowrap__SDL_GetKeyboardState
	ResetKeyboard = __gowrap__SDL_ResetKeyboard
	GetModState = __gowrap__SDL_GetModState
	SetModState = __gowrap__SDL_SetModState
	GetKeyFromScancode = __gowrap__SDL_GetKeyFromScancode
	GetScancodeFromKey = __gowrap__SDL_GetScancodeFromKey
	SetScancodeName = __gowrap__SDL_SetScancodeName
	GetScancodeName = __gowrap__SDL_GetScancodeName
	GetScancodeFromName = __gowrap__SDL_GetScancodeFromName
	GetKeyName = __gowrap__SDL_GetKeyName
	GetKeyFromName = __gowrap__SDL_GetKeyFromName
	StartTextInput = __gowrap__SDL_StartTextInput
	StartTextInputWithProperties = __gowrap__SDL_StartTextInputWithProperties
	TextInputActive = __gowrap__SDL_TextInputActive
	StopTextInput = __gowrap__SDL_StopTextInput
	ClearComposition = __gowrap__SDL_ClearComposition
	SetTextInputArea = __gowrap__SDL_SetTextInputArea
	GetTextInputArea = __gowrap__SDL_GetTextInputArea
	HasScreenKeyboardSupport = __gowrap__SDL_HasScreenKeyboardSupport
	ScreenKeyboardShown = __gowrap__SDL_ScreenKeyboardShown
}
