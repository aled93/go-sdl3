//go:build wasm && GDK

package sdl

import (
  "unsafe"
)



//go:wasmimport sdl3 SDL_GDKSuspendGPU
func __SDL_GDKSuspendGPU(uintptr)

//go:wasmimport sdl3 SDL_GDKResumeGPU
func __SDL_GDKResumeGPU(uintptr)



func __gowrap__SDL_GDKSuspendGPU(device GPUDevice) {
	__SDL_GDKSuspendGPU(uintptr(unsafe.Pointer(device)))
}

func __gowrap__SDL_GDKResumeGPU(device GPUDevice) {
	__SDL_GDKResumeGPU(uintptr(unsafe.Pointer(device)))
}

 
func init() {
	GDKSuspendGPU = __gowrap__SDL_GDKSuspendGPU
	GDKResumeGPU = __gowrap__SDL_GDKResumeGPU
}
