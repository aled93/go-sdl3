//go:build wasm

package sdl

import (
  "runtime"
  "unsafe"
)



//go:wasmimport sdl3 SDL_Init
func __SDL_Init(int32) int32

//go:wasmimport sdl3 SDL_InitSubSystem
func __SDL_InitSubSystem(int32) int32

//go:wasmimport sdl3 SDL_QuitSubSystem
func __SDL_QuitSubSystem(int32)

//go:wasmimport sdl3 SDL_WasInit
func __SDL_WasInit(int32) int32

//go:wasmimport sdl3 SDL_Quit
func __SDL_Quit()

//go:wasmimport sdl3 SDL_IsMainThread
func __SDL_IsMainThread() int32

//go:wasmimport sdl3 SDL_RunOnMainThread
func __SDL_RunOnMainThread(uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_SetAppMetadata
func __SDL_SetAppMetadata(uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_SetAppMetadataProperty
func __SDL_SetAppMetadataProperty(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetAppMetadataProperty
func __SDL_GetAppMetadataProperty(uintptr) uintptr



func __gowrap__SDL_Init(flags InitFlags) (__result bool) {
	__result = (bool)(__SDL_Init(*(*int32)(unsafe.Pointer(&flags))) != 0)
	return
}

func __gowrap__SDL_InitSubSystem(flags InitFlags) (__result bool) {
	__result = (bool)(__SDL_InitSubSystem(*(*int32)(unsafe.Pointer(&flags))) != 0)
	return
}

func __gowrap__SDL_QuitSubSystem(flags InitFlags) {
	__SDL_QuitSubSystem(*(*int32)(unsafe.Pointer(&flags)))
}

func __gowrap__SDL_WasInit(flags InitFlags) (__result InitFlags) {
	__result = (InitFlags)(__SDL_WasInit(*(*int32)(unsafe.Pointer(&flags))))
	return
}

func __gowrap__SDL_Quit() {
	__SDL_Quit()
}

func __gowrap__SDL_IsMainThread() (__result bool) {
	__result = (bool)(__SDL_IsMainThread() != 0)
	return
}

func __gowrap__SDL_RunOnMainThread(callback MainThreadCallback, userdata uintptr, wait_complete bool) (__result bool) {
	__result = (bool)(__SDL_RunOnMainThread(0 /* TODO: callbacks */, uintptr(unsafe.Pointer(userdata)), *(*int32)(unsafe.Pointer(&wait_complete))) != 0)
	return
}

func __gowrap__SDL_SetAppMetadata(appname string, appversion string, appidentifier string) (__result bool) {
	__result = (bool)(__SDL_SetAppMetadata(((*[2]uintptr)(unsafe.Pointer(&appname)))[0], ((*[2]uintptr)(unsafe.Pointer(&appversion)))[0], ((*[2]uintptr)(unsafe.Pointer(&appidentifier)))[0]) != 0)
	runtime.KeepAlive(appname)
	runtime.KeepAlive(appversion)
	runtime.KeepAlive(appidentifier)
	return
}

func __gowrap__SDL_SetAppMetadataProperty(name PropertyName, value string) (__result bool) {
	__result = (bool)(__SDL_SetAppMetadataProperty(((*[2]uintptr)(unsafe.Pointer(&name)))[0], ((*[2]uintptr)(unsafe.Pointer(&value)))[0]) != 0)
	runtime.KeepAlive(name)
	runtime.KeepAlive(value)
	return
}

func __gowrap__SDL_GetAppMetadataProperty(name PropertyName) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetAppMetadataProperty(((*[2]uintptr)(unsafe.Pointer(&name)))[0])).Collect()))
	runtime.KeepAlive(name)
	return
}

 
func init() {
	Init = __gowrap__SDL_Init
	InitSubSystem = __gowrap__SDL_InitSubSystem
	QuitSubSystem = __gowrap__SDL_QuitSubSystem
	WasInit = __gowrap__SDL_WasInit
	Quit = __gowrap__SDL_Quit
	IsMainThread = __gowrap__SDL_IsMainThread
	RunOnMainThread = __gowrap__SDL_RunOnMainThread
	SetAppMetadata = __gowrap__SDL_SetAppMetadata
	SetAppMetadataProperty = __gowrap__SDL_SetAppMetadataProperty
	GetAppMetadataProperty = __gowrap__SDL_GetAppMetadataProperty
}
