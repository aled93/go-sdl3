//go:build wasm

package sdl

import (
  "runtime"
  "unsafe"
)



//go:wasmimport sdl3 SDL_GetBasePath
func __SDL_GetBasePath() uintptr

//go:wasmimport sdl3 SDL_GetPrefPath
func __SDL_GetPrefPath(uintptr, uintptr) uintptr

//go:wasmimport sdl3 SDL_GetUserFolder
func __SDL_GetUserFolder(int32) uintptr

//go:wasmimport sdl3 SDL_CreateDirectory
func __SDL_CreateDirectory(uintptr) int32

//go:wasmimport sdl3 SDL_EnumerateDirectory
func __SDL_EnumerateDirectory(uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_RemovePath
func __SDL_RemovePath(uintptr) int32

//go:wasmimport sdl3 SDL_RenamePath
func __SDL_RenamePath(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_CopyFile
func __SDL_CopyFile(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetPathInfo
func __SDL_GetPathInfo(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GlobDirectory
func __SDL_GlobDirectory(uintptr, uintptr, int32, uintptr) uintptr

//go:wasmimport sdl3 SDL_GetCurrentDirectory
func __SDL_GetCurrentDirectory() uintptr



func __gowrap__SDL_GetBasePath() (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetBasePath()).Collect()))
	return
}

func __gowrap__SDL_GetPrefPath(org string, app string) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetPrefPath(((*[2]uintptr)(unsafe.Pointer(&org)))[0], ((*[2]uintptr)(unsafe.Pointer(&app)))[0])).Collect()))
	runtime.KeepAlive(org)
	runtime.KeepAlive(app)
	return
}

func __gowrap__SDL_GetUserFolder(folder Folder) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetUserFolder(*(*int32)(unsafe.Pointer(&folder)))).Collect()))
	return
}

func __gowrap__SDL_CreateDirectory(path string) (__result bool) {
	__result = (bool)(__SDL_CreateDirectory(((*[2]uintptr)(unsafe.Pointer(&path)))[0]) != 0)
	runtime.KeepAlive(path)
	return
}

func __gowrap__SDL_EnumerateDirectory(path string, callback EnumerateDirectoryCallback, userdata uintptr) (__result bool) {
	__result = (bool)(__SDL_EnumerateDirectory(((*[2]uintptr)(unsafe.Pointer(&path)))[0], 0 /* TODO: callbacks */, uintptr(unsafe.Pointer(userdata))) != 0)
	runtime.KeepAlive(path)
	return
}

func __gowrap__SDL_RemovePath(path string) (__result bool) {
	__result = (bool)(__SDL_RemovePath(((*[2]uintptr)(unsafe.Pointer(&path)))[0]) != 0)
	runtime.KeepAlive(path)
	return
}

func __gowrap__SDL_RenamePath(oldpath string, newpath string) (__result bool) {
	__result = (bool)(__SDL_RenamePath(((*[2]uintptr)(unsafe.Pointer(&oldpath)))[0], ((*[2]uintptr)(unsafe.Pointer(&newpath)))[0]) != 0)
	runtime.KeepAlive(oldpath)
	runtime.KeepAlive(newpath)
	return
}

func __gowrap__SDL_CopyFile(oldpath string, newpath string) (__result bool) {
	__result = (bool)(__SDL_CopyFile(((*[2]uintptr)(unsafe.Pointer(&oldpath)))[0], ((*[2]uintptr)(unsafe.Pointer(&newpath)))[0]) != 0)
	runtime.KeepAlive(oldpath)
	runtime.KeepAlive(newpath)
	return
}

func __gowrap__SDL_GetPathInfo(path string, info *PathInfo) (__result bool) {
	__result = (bool)(__SDL_GetPathInfo(((*[2]uintptr)(unsafe.Pointer(&path)))[0], uintptr(unsafe.Pointer(info))) != 0)
	runtime.KeepAlive(path)
	return
}

func __gowrap__SDL_GlobDirectory(path string, pattern string, flags GlobFlags, count *int32) (__result *string) {
	__result = (*string)(unsafe.Pointer(__SDL_GlobDirectory(((*[2]uintptr)(unsafe.Pointer(&path)))[0], ((*[2]uintptr)(unsafe.Pointer(&pattern)))[0], *(*int32)(unsafe.Pointer(&flags)), uintptr(unsafe.Pointer(count)))))
	runtime.KeepAlive(path)
	runtime.KeepAlive(pattern)
	return
}

func __gowrap__SDL_GetCurrentDirectory() (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetCurrentDirectory()).Collect()))
	return
}

 
func init() {
	GetBasePath = __gowrap__SDL_GetBasePath
	GetPrefPath = __gowrap__SDL_GetPrefPath
	GetUserFolder = __gowrap__SDL_GetUserFolder
	CreateDirectory = __gowrap__SDL_CreateDirectory
	EnumerateDirectory = __gowrap__SDL_EnumerateDirectory
	RemovePath = __gowrap__SDL_RemovePath
	RenamePath = __gowrap__SDL_RenamePath
	CopyFile = __gowrap__SDL_CopyFile
	GetPathInfo = __gowrap__SDL_GetPathInfo
	GlobDirectory = __gowrap__SDL_GlobDirectory
	GetCurrentDirectory = __gowrap__SDL_GetCurrentDirectory
}
