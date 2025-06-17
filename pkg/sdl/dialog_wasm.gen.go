//go:build wasm

package sdl

import (
  "runtime"
  "unsafe"
)



//go:wasmimport sdl3 SDL_ShowOpenFileDialog
func __SDL_ShowOpenFileDialog(uintptr, uintptr, uintptr, uintptr, int32, uintptr, int32)

//go:wasmimport sdl3 SDL_ShowSaveFileDialog
func __SDL_ShowSaveFileDialog(uintptr, uintptr, uintptr, uintptr, int32, uintptr)

//go:wasmimport sdl3 SDL_ShowOpenFolderDialog
func __SDL_ShowOpenFolderDialog(uintptr, uintptr, uintptr, uintptr, int32)

//go:wasmimport sdl3 SDL_ShowFileDialogWithProperties
func __SDL_ShowFileDialogWithProperties(int32, uintptr, uintptr, int32)



func __gowrap__SDL_ShowOpenFileDialog(callback DialogFileCallback, userdata uintptr, window Window, filters []DialogFileFilter, nfilters int32, default_location string, allow_many bool) {
	__SDL_ShowOpenFileDialog(0 /* TODO: callbacks */, uintptr(unsafe.Pointer(userdata)), uintptr(unsafe.Pointer(window)), uintptr(unsafe.Pointer(&filters[0])), *(*int32)(unsafe.Pointer(&nfilters)), ((*[2]uintptr)(unsafe.Pointer(&default_location)))[0], *(*int32)(unsafe.Pointer(&allow_many)))
	runtime.KeepAlive(filters)
	runtime.KeepAlive(default_location)
}

func __gowrap__SDL_ShowSaveFileDialog(callback DialogFileCallback, userdata uintptr, window Window, filters []DialogFileFilter, nfilters int32, default_location string) {
	__SDL_ShowSaveFileDialog(0 /* TODO: callbacks */, uintptr(unsafe.Pointer(userdata)), uintptr(unsafe.Pointer(window)), uintptr(unsafe.Pointer(&filters[0])), *(*int32)(unsafe.Pointer(&nfilters)), ((*[2]uintptr)(unsafe.Pointer(&default_location)))[0])
	runtime.KeepAlive(filters)
	runtime.KeepAlive(default_location)
}

func __gowrap__SDL_ShowOpenFolderDialog(callback DialogFileCallback, userdata uintptr, window Window, default_location string, allow_many bool) {
	__SDL_ShowOpenFolderDialog(0 /* TODO: callbacks */, uintptr(unsafe.Pointer(userdata)), uintptr(unsafe.Pointer(window)), ((*[2]uintptr)(unsafe.Pointer(&default_location)))[0], *(*int32)(unsafe.Pointer(&allow_many)))
	runtime.KeepAlive(default_location)
}

func __gowrap__SDL_ShowFileDialogWithProperties(type_ FileDialogType, callback DialogFileCallback, userdata uintptr, props PropertiesID) {
	__SDL_ShowFileDialogWithProperties(*(*int32)(unsafe.Pointer(&type_)), 0 /* TODO: callbacks */, uintptr(unsafe.Pointer(userdata)), *(*int32)(unsafe.Pointer(&props)))
}

 
func init() {
	ShowOpenFileDialog = __gowrap__SDL_ShowOpenFileDialog
	ShowSaveFileDialog = __gowrap__SDL_ShowSaveFileDialog
	ShowOpenFolderDialog = __gowrap__SDL_ShowOpenFolderDialog
	ShowFileDialogWithProperties = __gowrap__SDL_ShowFileDialogWithProperties
}
