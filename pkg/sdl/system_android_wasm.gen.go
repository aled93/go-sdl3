//go:build wasm

package sdl

import (
  "unsafe"
)



//go:wasmimport sdl3 SDL_GetAndroidJNIEnv
func __SDL_GetAndroidJNIEnv() uintptr

//go:wasmimport sdl3 SDL_GetAndroidActivity
func __SDL_GetAndroidActivity() uintptr

//go:wasmimport sdl3 SDL_GetAndroidSDKVersion
func __SDL_GetAndroidSDKVersion() int32

//go:wasmimport sdl3 SDL_IsChromebook
func __SDL_IsChromebook() int32

//go:wasmimport sdl3 SDL_IsDeXMode
func __SDL_IsDeXMode() int32

//go:wasmimport sdl3 SDL_SendAndroidBackButton
func __SDL_SendAndroidBackButton()

//go:wasmimport sdl3 SDL_GetAndroidInternalStoragePath
func __SDL_GetAndroidInternalStoragePath() uintptr

//go:wasmimport sdl3 SDL_GetAndroidExternalStorageState
func __SDL_GetAndroidExternalStorageState() int32

//go:wasmimport sdl3 SDL_GetAndroidExternalStoragePath
func __SDL_GetAndroidExternalStoragePath() uintptr

//go:wasmimport sdl3 SDL_GetAndroidCachePath
func __SDL_GetAndroidCachePath() uintptr

//go:wasmimport sdl3 SDL_RequestAndroidPermission
func __SDL_RequestAndroidPermission(uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_ShowAndroidToast
func __SDL_ShowAndroidToast(uintptr, int32, int32, int32, int32) int32

//go:wasmimport sdl3 SDL_SendAndroidMessage
func __SDL_SendAndroidMessage(int32, int32) int32



func __gowrap__SDL_GetAndroidJNIEnv() (__result uintptr) {
	__result = (uintptr)(unsafe.Pointer(__SDL_GetAndroidJNIEnv()))
	return
}

func __gowrap__SDL_GetAndroidActivity() (__result uintptr) {
	__result = (uintptr)(unsafe.Pointer(__SDL_GetAndroidActivity()))
	return
}

func __gowrap__SDL_GetAndroidSDKVersion() (__result int32) {
	__result = (int32)(__SDL_GetAndroidSDKVersion())
	return
}

func __gowrap__SDL_IsChromebook() (__result bool) {
	__result = (bool)(__SDL_IsChromebook() != 0)
	return
}

func __gowrap__SDL_IsDeXMode() (__result bool) {
	__result = (bool)(__SDL_IsDeXMode() != 0)
	return
}

func __gowrap__SDL_SendAndroidBackButton() {
	__SDL_SendAndroidBackButton()
}

func __gowrap__SDL_GetAndroidInternalStoragePath() (__result *byte) {
	__result = (*byte)(unsafe.Pointer(__SDL_GetAndroidInternalStoragePath()))
	return
}

func __gowrap__SDL_GetAndroidExternalStorageState() (__result uint32) {
	__result = (uint32)(__SDL_GetAndroidExternalStorageState())
	return
}

func __gowrap__SDL_GetAndroidExternalStoragePath() (__result *byte) {
	__result = (*byte)(unsafe.Pointer(__SDL_GetAndroidExternalStoragePath()))
	return
}

func __gowrap__SDL_GetAndroidCachePath() (__result *byte) {
	__result = (*byte)(unsafe.Pointer(__SDL_GetAndroidCachePath()))
	return
}

func __gowrap__SDL_RequestAndroidPermission(permission *byte, cb RequestAndroidPermissionCallback, userdata uintptr) (__result bool) {
	__result = (bool)(__SDL_RequestAndroidPermission(uintptr(unsafe.Pointer(permission)), 0 /* TODO: callbacks */, uintptr(unsafe.Pointer(userdata))) != 0)
	return
}

func __gowrap__SDL_ShowAndroidToast(message *byte, duration int32, gravity int32, xoffset int32, yoffset int32) (__result bool) {
	__result = (bool)(__SDL_ShowAndroidToast(uintptr(unsafe.Pointer(message)), *(*int32)(unsafe.Pointer(&duration)), *(*int32)(unsafe.Pointer(&gravity)), *(*int32)(unsafe.Pointer(&xoffset)), *(*int32)(unsafe.Pointer(&yoffset))) != 0)
	return
}

func __gowrap__SDL_SendAndroidMessage(command uint32, param int32) (__result bool) {
	__result = (bool)(__SDL_SendAndroidMessage(*(*int32)(unsafe.Pointer(&command)), *(*int32)(unsafe.Pointer(&param))) != 0)
	return
}

 
func init() {
	GetAndroidJNIEnv = __gowrap__SDL_GetAndroidJNIEnv
	GetAndroidActivity = __gowrap__SDL_GetAndroidActivity
	GetAndroidSDKVersion = __gowrap__SDL_GetAndroidSDKVersion
	IsChromebook = __gowrap__SDL_IsChromebook
	IsDeXMode = __gowrap__SDL_IsDeXMode
	SendAndroidBackButton = __gowrap__SDL_SendAndroidBackButton
	GetAndroidInternalStoragePath = __gowrap__SDL_GetAndroidInternalStoragePath
	GetAndroidExternalStorageState = __gowrap__SDL_GetAndroidExternalStorageState
	GetAndroidExternalStoragePath = __gowrap__SDL_GetAndroidExternalStoragePath
	GetAndroidCachePath = __gowrap__SDL_GetAndroidCachePath
	RequestAndroidPermission = __gowrap__SDL_RequestAndroidPermission
	ShowAndroidToast = __gowrap__SDL_ShowAndroidToast
	SendAndroidMessage = __gowrap__SDL_SendAndroidMessage
}
