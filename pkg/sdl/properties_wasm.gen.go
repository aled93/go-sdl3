//go:build wasm

package sdl

import (
  "unsafe"
)



//go:wasmimport sdl3 SDL_GetGlobalProperties
func __SDL_GetGlobalProperties() int32

//go:wasmimport sdl3 SDL_CreateProperties
func __SDL_CreateProperties() int32

//go:wasmimport sdl3 SDL_CopyProperties
func __SDL_CopyProperties(int32, int32) int32

//go:wasmimport sdl3 SDL_LockProperties
func __SDL_LockProperties(int32) int32

//go:wasmimport sdl3 SDL_UnlockProperties
func __SDL_UnlockProperties(int32)

//go:wasmimport sdl3 SDL_SetPointerPropertyWithCleanup
func __SDL_SetPointerPropertyWithCleanup(int32, uintptr, uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_SetPointerProperty
func __SDL_SetPointerProperty(int32, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_SetStringProperty
func __SDL_SetStringProperty(int32, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_SetNumberProperty
func __SDL_SetNumberProperty(int32, uintptr, int64) int32

//go:wasmimport sdl3 SDL_SetFloatProperty
func __SDL_SetFloatProperty(int32, uintptr, float32) int32

//go:wasmimport sdl3 SDL_SetBooleanProperty
func __SDL_SetBooleanProperty(int32, uintptr, int32) int32

//go:wasmimport sdl3 SDL_HasProperty
func __SDL_HasProperty(int32, uintptr) int32

//go:wasmimport sdl3 SDL_GetPropertyType
func __SDL_GetPropertyType(int32, uintptr) int32

//go:wasmimport sdl3 SDL_GetPointerProperty
func __SDL_GetPointerProperty(int32, uintptr, uintptr) uintptr

//go:wasmimport sdl3 SDL_GetStringProperty
func __SDL_GetStringProperty(int32, uintptr, uintptr) uintptr

//go:wasmimport sdl3 SDL_GetNumberProperty
func __SDL_GetNumberProperty(int32, uintptr, int64) int64

//go:wasmimport sdl3 SDL_GetFloatProperty
func __SDL_GetFloatProperty(int32, uintptr, float32) float32

//go:wasmimport sdl3 SDL_GetBooleanProperty
func __SDL_GetBooleanProperty(int32, uintptr, int32) int32

//go:wasmimport sdl3 SDL_ClearProperty
func __SDL_ClearProperty(int32, uintptr) int32

//go:wasmimport sdl3 SDL_EnumerateProperties
func __SDL_EnumerateProperties(int32, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_DestroyProperties
func __SDL_DestroyProperties(int32)



func __gowrap__SDL_GetGlobalProperties() (__result PropertiesID) {
	__result = (PropertiesID)(__SDL_GetGlobalProperties())
	return
}

func __gowrap__SDL_CreateProperties() (__result PropertiesID) {
	__result = (PropertiesID)(__SDL_CreateProperties())
	return
}

func __gowrap__SDL_CopyProperties(src PropertiesID, dst PropertiesID) (__result bool) {
	__result = (bool)(__SDL_CopyProperties(*(*int32)(unsafe.Pointer(&src)), *(*int32)(unsafe.Pointer(&dst))) != 0)
	return
}

func __gowrap__SDL_LockProperties(props PropertiesID) (__result bool) {
	__result = (bool)(__SDL_LockProperties(*(*int32)(unsafe.Pointer(&props))) != 0)
	return
}

func __gowrap__SDL_UnlockProperties(props PropertiesID) {
	__SDL_UnlockProperties(*(*int32)(unsafe.Pointer(&props)))
}

func __gowrap__SDL_SetPointerPropertyWithCleanup(props PropertiesID, name *byte, value uintptr, cleanup CleanupPropertyCallback, userdata uintptr) (__result bool) {
	__result = (bool)(__SDL_SetPointerPropertyWithCleanup(*(*int32)(unsafe.Pointer(&props)), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(value)), 0 /* TODO: callbacks */, uintptr(unsafe.Pointer(userdata))) != 0)
	return
}

func __gowrap__SDL_SetPointerProperty(props PropertiesID, name *byte, value uintptr) (__result bool) {
	__result = (bool)(__SDL_SetPointerProperty(*(*int32)(unsafe.Pointer(&props)), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(value))) != 0)
	return
}

func __gowrap__SDL_SetStringProperty(props PropertiesID, name *byte, value *byte) (__result bool) {
	__result = (bool)(__SDL_SetStringProperty(*(*int32)(unsafe.Pointer(&props)), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(value))) != 0)
	return
}

func __gowrap__SDL_SetNumberProperty(props PropertiesID, name *byte, value int64) (__result bool) {
	__result = (bool)(__SDL_SetNumberProperty(*(*int32)(unsafe.Pointer(&props)), uintptr(unsafe.Pointer(name)), *(*int64)(unsafe.Pointer(&value))) != 0)
	return
}

func __gowrap__SDL_SetFloatProperty(props PropertiesID, name *byte, value float32) (__result bool) {
	__result = (bool)(__SDL_SetFloatProperty(*(*int32)(unsafe.Pointer(&props)), uintptr(unsafe.Pointer(name)), *(*float32)(unsafe.Pointer(&value))) != 0)
	return
}

func __gowrap__SDL_SetBooleanProperty(props PropertiesID, name *byte, value bool) (__result bool) {
	__result = (bool)(__SDL_SetBooleanProperty(*(*int32)(unsafe.Pointer(&props)), uintptr(unsafe.Pointer(name)), *(*int32)(unsafe.Pointer(&value))) != 0)
	return
}

func __gowrap__SDL_HasProperty(props PropertiesID, name *byte) (__result bool) {
	__result = (bool)(__SDL_HasProperty(*(*int32)(unsafe.Pointer(&props)), uintptr(unsafe.Pointer(name))) != 0)
	return
}

func __gowrap__SDL_GetPropertyType(props PropertiesID, name *byte) (__result PropertyType) {
	__result = (PropertyType)(__SDL_GetPropertyType(*(*int32)(unsafe.Pointer(&props)), uintptr(unsafe.Pointer(name))))
	return
}

func __gowrap__SDL_GetPointerProperty(props PropertiesID, name *byte, default_value uintptr) (__result uintptr) {
	__result = (uintptr)(unsafe.Pointer(__SDL_GetPointerProperty(*(*int32)(unsafe.Pointer(&props)), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(default_value)))))
	return
}

func __gowrap__SDL_GetStringProperty(props PropertiesID, name *byte, default_value *byte) (__result *byte) {
	__result = (*byte)(unsafe.Pointer(__SDL_GetStringProperty(*(*int32)(unsafe.Pointer(&props)), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(default_value)))))
	return
}

func __gowrap__SDL_GetNumberProperty(props PropertiesID, name *byte, default_value int64) (__result int64) {
	__result = (int64)(__SDL_GetNumberProperty(*(*int32)(unsafe.Pointer(&props)), uintptr(unsafe.Pointer(name)), *(*int64)(unsafe.Pointer(&default_value))))
	return
}

func __gowrap__SDL_GetFloatProperty(props PropertiesID, name *byte, default_value float32) (__result float32) {
	__result = (float32)(__SDL_GetFloatProperty(*(*int32)(unsafe.Pointer(&props)), uintptr(unsafe.Pointer(name)), *(*float32)(unsafe.Pointer(&default_value))))
	return
}

func __gowrap__SDL_GetBooleanProperty(props PropertiesID, name *byte, default_value bool) (__result bool) {
	__result = (bool)(__SDL_GetBooleanProperty(*(*int32)(unsafe.Pointer(&props)), uintptr(unsafe.Pointer(name)), *(*int32)(unsafe.Pointer(&default_value))) != 0)
	return
}

func __gowrap__SDL_ClearProperty(props PropertiesID, name *byte) (__result bool) {
	__result = (bool)(__SDL_ClearProperty(*(*int32)(unsafe.Pointer(&props)), uintptr(unsafe.Pointer(name))) != 0)
	return
}

func __gowrap__SDL_EnumerateProperties(props PropertiesID, callback EnumeratePropertiesCallback, userdata uintptr) (__result bool) {
	__result = (bool)(__SDL_EnumerateProperties(*(*int32)(unsafe.Pointer(&props)), 0 /* TODO: callbacks */, uintptr(unsafe.Pointer(userdata))) != 0)
	return
}

func __gowrap__SDL_DestroyProperties(props PropertiesID) {
	__SDL_DestroyProperties(*(*int32)(unsafe.Pointer(&props)))
}

 
func init() {
	GetGlobalProperties = __gowrap__SDL_GetGlobalProperties
	CreateProperties = __gowrap__SDL_CreateProperties
	CopyProperties = __gowrap__SDL_CopyProperties
	LockProperties = __gowrap__SDL_LockProperties
	UnlockProperties = __gowrap__SDL_UnlockProperties
	SetPointerPropertyWithCleanup = __gowrap__SDL_SetPointerPropertyWithCleanup
	SetPointerProperty = __gowrap__SDL_SetPointerProperty
	SetStringProperty = __gowrap__SDL_SetStringProperty
	SetNumberProperty = __gowrap__SDL_SetNumberProperty
	SetFloatProperty = __gowrap__SDL_SetFloatProperty
	SetBooleanProperty = __gowrap__SDL_SetBooleanProperty
	HasProperty = __gowrap__SDL_HasProperty
	GetPropertyType = __gowrap__SDL_GetPropertyType
	GetPointerProperty = __gowrap__SDL_GetPointerProperty
	GetStringProperty = __gowrap__SDL_GetStringProperty
	GetNumberProperty = __gowrap__SDL_GetNumberProperty
	GetFloatProperty = __gowrap__SDL_GetFloatProperty
	GetBooleanProperty = __gowrap__SDL_GetBooleanProperty
	ClearProperty = __gowrap__SDL_ClearProperty
	EnumerateProperties = __gowrap__SDL_EnumerateProperties
	DestroyProperties = __gowrap__SDL_DestroyProperties
}
