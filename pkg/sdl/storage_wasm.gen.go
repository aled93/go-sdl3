//go:build wasm

package sdl

import (
  "runtime"
  "unsafe"
)



//go:wasmimport sdl3 SDL_OpenTitleStorage
func __SDL_OpenTitleStorage(uintptr, int32) uintptr

//go:wasmimport sdl3 SDL_OpenUserStorage
func __SDL_OpenUserStorage(uintptr, uintptr, int32) uintptr

//go:wasmimport sdl3 SDL_OpenFileStorage
func __SDL_OpenFileStorage(uintptr) uintptr

//go:wasmimport sdl3 SDL_OpenStorage
func __SDL_OpenStorage(uintptr, uintptr) uintptr

//go:wasmimport sdl3 SDL_CloseStorage
func __SDL_CloseStorage(uintptr) int32

//go:wasmimport sdl3 SDL_StorageReady
func __SDL_StorageReady(uintptr) int32

//go:wasmimport sdl3 SDL_ReadStorageFile
func __SDL_ReadStorageFile(uintptr, uintptr, uintptr, int64) int32

//go:wasmimport sdl3 SDL_WriteStorageFile
func __SDL_WriteStorageFile(uintptr, uintptr, uintptr, int64) int32

//go:wasmimport sdl3 SDL_CreateStorageDirectory
func __SDL_CreateStorageDirectory(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_EnumerateStorageDirectory
func __SDL_EnumerateStorageDirectory(uintptr, uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_RemoveStoragePath
func __SDL_RemoveStoragePath(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_RenameStoragePath
func __SDL_RenameStoragePath(uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_CopyStorageFile
func __SDL_CopyStorageFile(uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetStoragePathInfo
func __SDL_GetStoragePathInfo(uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetStorageSpaceRemaining
func __SDL_GetStorageSpaceRemaining(uintptr) int64

//go:wasmimport sdl3 SDL_GlobStorageDirectory
func __SDL_GlobStorageDirectory(uintptr, uintptr, uintptr, int32, uintptr) uintptr



func __gowrap__SDL_OpenTitleStorage(override string, props PropertiesID) (__result Storage) {
	__result = (Storage)(unsafe.Pointer(__SDL_OpenTitleStorage(((*[2]uintptr)(unsafe.Pointer(&override)))[0], *(*int32)(unsafe.Pointer(&props)))))
	runtime.KeepAlive(override)
	return
}

func __gowrap__SDL_OpenUserStorage(org string, app string, props PropertiesID) (__result Storage) {
	__result = (Storage)(unsafe.Pointer(__SDL_OpenUserStorage(((*[2]uintptr)(unsafe.Pointer(&org)))[0], ((*[2]uintptr)(unsafe.Pointer(&app)))[0], *(*int32)(unsafe.Pointer(&props)))))
	runtime.KeepAlive(org)
	runtime.KeepAlive(app)
	return
}

func __gowrap__SDL_OpenFileStorage(path string) (__result Storage) {
	__result = (Storage)(unsafe.Pointer(__SDL_OpenFileStorage(((*[2]uintptr)(unsafe.Pointer(&path)))[0])))
	runtime.KeepAlive(path)
	return
}

func __gowrap__SDL_OpenStorage(iface *StorageInterface, userdata uintptr) (__result Storage) {
	__result = (Storage)(unsafe.Pointer(__SDL_OpenStorage(uintptr(unsafe.Pointer(iface)), uintptr(unsafe.Pointer(userdata)))))
	return
}

func __gowrap__SDL_CloseStorage(storage Storage) (__result bool) {
	__result = (bool)(__SDL_CloseStorage(uintptr(unsafe.Pointer(storage))) != 0)
	return
}

func __gowrap__SDL_StorageReady(storage Storage) (__result bool) {
	__result = (bool)(__SDL_StorageReady(uintptr(unsafe.Pointer(storage))) != 0)
	return
}

func __gowrap__SDL_ReadStorageFile(storage Storage, path string, destination *byte, length uint64) (__result bool) {
	__result = (bool)(__SDL_ReadStorageFile(uintptr(unsafe.Pointer(storage)), ((*[2]uintptr)(unsafe.Pointer(&path)))[0], uintptr(unsafe.Pointer(destination)), *(*int64)(unsafe.Pointer(&length))) != 0)
	runtime.KeepAlive(path)
	return
}

func __gowrap__SDL_WriteStorageFile(storage Storage, path string, source string, length uint64) (__result bool) {
	__result = (bool)(__SDL_WriteStorageFile(uintptr(unsafe.Pointer(storage)), ((*[2]uintptr)(unsafe.Pointer(&path)))[0], ((*[2]uintptr)(unsafe.Pointer(&source)))[0], *(*int64)(unsafe.Pointer(&length))) != 0)
	runtime.KeepAlive(path)
	runtime.KeepAlive(source)
	return
}

func __gowrap__SDL_CreateStorageDirectory(storage Storage, path string) (__result bool) {
	__result = (bool)(__SDL_CreateStorageDirectory(uintptr(unsafe.Pointer(storage)), ((*[2]uintptr)(unsafe.Pointer(&path)))[0]) != 0)
	runtime.KeepAlive(path)
	return
}

func __gowrap__SDL_EnumerateStorageDirectory(storage Storage, path string, callback EnumerateDirectoryCallback, userdata uintptr) (__result bool) {
	__result = (bool)(__SDL_EnumerateStorageDirectory(uintptr(unsafe.Pointer(storage)), ((*[2]uintptr)(unsafe.Pointer(&path)))[0], 0 /* TODO: callbacks */, uintptr(unsafe.Pointer(userdata))) != 0)
	runtime.KeepAlive(path)
	return
}

func __gowrap__SDL_RemoveStoragePath(storage Storage, path string) (__result bool) {
	__result = (bool)(__SDL_RemoveStoragePath(uintptr(unsafe.Pointer(storage)), ((*[2]uintptr)(unsafe.Pointer(&path)))[0]) != 0)
	runtime.KeepAlive(path)
	return
}

func __gowrap__SDL_RenameStoragePath(storage Storage, oldpath string, newpath string) (__result bool) {
	__result = (bool)(__SDL_RenameStoragePath(uintptr(unsafe.Pointer(storage)), ((*[2]uintptr)(unsafe.Pointer(&oldpath)))[0], ((*[2]uintptr)(unsafe.Pointer(&newpath)))[0]) != 0)
	runtime.KeepAlive(oldpath)
	runtime.KeepAlive(newpath)
	return
}

func __gowrap__SDL_CopyStorageFile(storage Storage, oldpath string, newpath string) (__result bool) {
	__result = (bool)(__SDL_CopyStorageFile(uintptr(unsafe.Pointer(storage)), ((*[2]uintptr)(unsafe.Pointer(&oldpath)))[0], ((*[2]uintptr)(unsafe.Pointer(&newpath)))[0]) != 0)
	runtime.KeepAlive(oldpath)
	runtime.KeepAlive(newpath)
	return
}

func __gowrap__SDL_GetStoragePathInfo(storage Storage, path string, info *PathInfo) (__result bool) {
	__result = (bool)(__SDL_GetStoragePathInfo(uintptr(unsafe.Pointer(storage)), ((*[2]uintptr)(unsafe.Pointer(&path)))[0], uintptr(unsafe.Pointer(info))) != 0)
	runtime.KeepAlive(path)
	return
}

func __gowrap__SDL_GetStorageSpaceRemaining(storage Storage) (__result uint64) {
	__result = (uint64)(__SDL_GetStorageSpaceRemaining(uintptr(unsafe.Pointer(storage))))
	return
}

func __gowrap__SDL_GlobStorageDirectory(storage Storage, path string, pattern string, flags GlobFlags, count *int32) (__result *string) {
	__result = (*string)(unsafe.Pointer(__SDL_GlobStorageDirectory(uintptr(unsafe.Pointer(storage)), ((*[2]uintptr)(unsafe.Pointer(&path)))[0], ((*[2]uintptr)(unsafe.Pointer(&pattern)))[0], *(*int32)(unsafe.Pointer(&flags)), uintptr(unsafe.Pointer(count)))))
	runtime.KeepAlive(path)
	runtime.KeepAlive(pattern)
	return
}

 
func init() {
	OpenTitleStorage = __gowrap__SDL_OpenTitleStorage
	OpenUserStorage = __gowrap__SDL_OpenUserStorage
	OpenFileStorage = __gowrap__SDL_OpenFileStorage
	OpenStorage = __gowrap__SDL_OpenStorage
	CloseStorage = __gowrap__SDL_CloseStorage
	StorageReady = __gowrap__SDL_StorageReady
	ReadStorageFile = __gowrap__SDL_ReadStorageFile
	WriteStorageFile = __gowrap__SDL_WriteStorageFile
	CreateStorageDirectory = __gowrap__SDL_CreateStorageDirectory
	EnumerateStorageDirectory = __gowrap__SDL_EnumerateStorageDirectory
	RemoveStoragePath = __gowrap__SDL_RemoveStoragePath
	RenameStoragePath = __gowrap__SDL_RenameStoragePath
	CopyStorageFile = __gowrap__SDL_CopyStorageFile
	GetStoragePathInfo = __gowrap__SDL_GetStoragePathInfo
	GetStorageSpaceRemaining = __gowrap__SDL_GetStorageSpaceRemaining
	GlobStorageDirectory = __gowrap__SDL_GlobStorageDirectory
}
