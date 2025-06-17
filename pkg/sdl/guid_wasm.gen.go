//go:build wasm

package sdl

import (
  "runtime"
  "unsafe"
)



//go:wasmimport sdl3 SDL_GUIDToString
func __SDL_GUIDToString(uintptr, uintptr, int32)

//go:wasmimport sdl3 SDL_StringToGUID
func __SDL_StringToGUID(uintptr, uintptr)



func __gowrap__SDL_GUIDToString(guid *GUID, pszGUID []byte, cbGUID int32) {
	__SDL_GUIDToString(uintptr(unsafe.Pointer(guid)), uintptr(unsafe.Pointer(&pszGUID[0])), *(*int32)(unsafe.Pointer(&cbGUID)))
	runtime.KeepAlive(pszGUID)
}

func __gowrap__SDL_StringToGUID(pchGUID string) (__result GUID) {
	__SDL_StringToGUID(uintptr(unsafe.Pointer(&__result)), ((*[2]uintptr)(unsafe.Pointer(&pchGUID)))[0])
	runtime.KeepAlive(pchGUID)
	return
}

 
func init() {
	GUIDToString = __gowrap__SDL_GUIDToString
	StringToGUID = __gowrap__SDL_StringToGUID
}
