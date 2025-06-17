//go:build wasm && windows || WINGDK

package sdl

import (
  "unsafe"
)



//go:wasmimport sdl3 SDL_GetDirect3D9AdapterIndex
func __SDL_GetDirect3D9AdapterIndex(int32)

//go:wasmimport sdl3 SDL_GetDXGIOutputInfo
func __SDL_GetDXGIOutputInfo(int32, uintptr, uintptr) int32



func __gowrap__SDL_GetDirect3D9AdapterIndex(displayID DisplayID) {
	__SDL_GetDirect3D9AdapterIndex(*(*int32)(unsafe.Pointer(&displayID)))
}

func __gowrap__SDL_GetDXGIOutputInfo(displayID DisplayID, adapterIndex *int32, outputIndex *int32) (__result bool) {
	__result = (bool)(__SDL_GetDXGIOutputInfo(*(*int32)(unsafe.Pointer(&displayID)), uintptr(unsafe.Pointer(adapterIndex)), uintptr(unsafe.Pointer(outputIndex))) != 0)
	return
}

 
func init() {
	GetDirect3D9AdapterIndex = __gowrap__SDL_GetDirect3D9AdapterIndex
	GetDXGIOutputInfo = __gowrap__SDL_GetDXGIOutputInfo
}
