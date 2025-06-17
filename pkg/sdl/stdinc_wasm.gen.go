//go:build wasm

package sdl

import (
  "unsafe"
)



//go:wasmimport sdl3 SDL_malloc
func __SDL_malloc(int64) uintptr

//go:wasmimport sdl3 SDL_calloc
func __SDL_calloc(int64, int64) uintptr

//go:wasmimport sdl3 SDL_realloc
func __SDL_realloc(uintptr, int64) uintptr

//go:wasmimport sdl3 SDL_free
func __SDL_free(uintptr)

//go:wasmimport sdl3 SDL_GetOriginalMemoryFunctions
func __SDL_GetOriginalMemoryFunctions(uintptr, uintptr, uintptr, uintptr)

//go:wasmimport sdl3 SDL_GetMemoryFunctions
func __SDL_GetMemoryFunctions(uintptr, uintptr, uintptr, uintptr)

//go:wasmimport sdl3 SDL_SetMemoryFunctions
func __SDL_SetMemoryFunctions(uintptr, uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_aligned_alloc
func __SDL_aligned_alloc(int64, int64) uintptr

//go:wasmimport sdl3 SDL_aligned_free
func __SDL_aligned_free(uintptr)

//go:wasmimport sdl3 SDL_GetNumAllocations
func __SDL_GetNumAllocations() int32

//go:wasmimport sdl3 SDL_GetEnvironment
func __SDL_GetEnvironment() uintptr

//go:wasmimport sdl3 SDL_CreateEnvironment
func __SDL_CreateEnvironment(int32) uintptr

//go:wasmimport sdl3 SDL_GetEnvironmentVariable
func __SDL_GetEnvironmentVariable(uintptr, uintptr) uintptr

//go:wasmimport sdl3 SDL_SetEnvironmentVariable
func __SDL_SetEnvironmentVariable(uintptr, uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_UnsetEnvironmentVariable
func __SDL_UnsetEnvironmentVariable(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_DestroyEnvironment
func __SDL_DestroyEnvironment(uintptr)

//go:wasmimport sdl3 SDL_getenv
func __SDL_getenv(uintptr) uintptr

//go:wasmimport sdl3 SDL_getenv_unsafe
func __SDL_getenv_unsafe(uintptr) uintptr

//go:wasmimport sdl3 SDL_setenv_unsafe
func __SDL_setenv_unsafe(uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_unsetenv_unsafe
func __SDL_unsetenv_unsafe(uintptr) int32



func __gowrap__SDL_malloc(size uint) (__result uintptr) {
	__result = (uintptr)(unsafe.Pointer(__SDL_malloc(*(*int64)(unsafe.Pointer(&size)))))
	return
}

func __gowrap__SDL_calloc(nmemb uint, size uint) (__result uintptr) {
	__result = (uintptr)(unsafe.Pointer(__SDL_calloc(*(*int64)(unsafe.Pointer(&nmemb)), *(*int64)(unsafe.Pointer(&size)))))
	return
}

func __gowrap__SDL_realloc(mem uintptr, size uint) (__result uintptr) {
	__result = (uintptr)(unsafe.Pointer(__SDL_realloc(uintptr(unsafe.Pointer(mem)), *(*int64)(unsafe.Pointer(&size)))))
	return
}

func __gowrap__SDL_free(mem uintptr) {
	__SDL_free(uintptr(unsafe.Pointer(mem)))
}

func __gowrap__SDL_GetOriginalMemoryFunctions(malloc_func MallocFunc, calloc_func CallocFunc, realloc_func ReallocFunc, free_func FreeFunc) {
	__SDL_GetOriginalMemoryFunctions(0 /* TODO: callbacks */, 0 /* TODO: callbacks */, 0 /* TODO: callbacks */, 0 /* TODO: callbacks */)
}

func __gowrap__SDL_GetMemoryFunctions(malloc_func *MallocFunc, calloc_func *CallocFunc, realloc_func *ReallocFunc, free_func *FreeFunc) {
	__SDL_GetMemoryFunctions(uintptr(unsafe.Pointer(malloc_func)), uintptr(unsafe.Pointer(calloc_func)), uintptr(unsafe.Pointer(realloc_func)), uintptr(unsafe.Pointer(free_func)))
}

func __gowrap__SDL_SetMemoryFunctions(malloc_func MallocFunc, calloc_func CallocFunc, realloc_func ReallocFunc, free_func FreeFunc) (__result bool) {
	__result = (bool)(__SDL_SetMemoryFunctions(0 /* TODO: callbacks */, 0 /* TODO: callbacks */, 0 /* TODO: callbacks */, 0 /* TODO: callbacks */) != 0)
	return
}

func __gowrap__SDL_aligned_alloc(alignment uint, size uint) (__result uintptr) {
	__result = (uintptr)(unsafe.Pointer(__SDL_aligned_alloc(*(*int64)(unsafe.Pointer(&alignment)), *(*int64)(unsafe.Pointer(&size)))))
	return
}

func __gowrap__SDL_aligned_free(mem uintptr) {
	__SDL_aligned_free(uintptr(unsafe.Pointer(mem)))
}

func __gowrap__SDL_GetNumAllocations() (__result int32) {
	__result = (int32)(__SDL_GetNumAllocations())
	return
}

func __gowrap__SDL_GetEnvironment() (__result Environment) {
	__result = (Environment)(unsafe.Pointer(__SDL_GetEnvironment()))
	return
}

func __gowrap__SDL_CreateEnvironment(populated bool) (__result Environment) {
	__result = (Environment)(unsafe.Pointer(__SDL_CreateEnvironment(*(*int32)(unsafe.Pointer(&populated)))))
	return
}

func __gowrap__SDL_GetEnvironmentVariable(env Environment, name uintptr) (__result uintptr) {
	__result = (uintptr)(unsafe.Pointer(__SDL_GetEnvironmentVariable(uintptr(unsafe.Pointer(env)), uintptr(unsafe.Pointer(name)))))
	return
}

func __gowrap__SDL_SetEnvironmentVariable(env Environment, name uintptr, value uintptr, overwrite bool) (__result bool) {
	__result = (bool)(__SDL_SetEnvironmentVariable(uintptr(unsafe.Pointer(env)), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(value)), *(*int32)(unsafe.Pointer(&overwrite))) != 0)
	return
}

func __gowrap__SDL_UnsetEnvironmentVariable(env Environment, name uintptr) (__result bool) {
	__result = (bool)(__SDL_UnsetEnvironmentVariable(uintptr(unsafe.Pointer(env)), uintptr(unsafe.Pointer(name))) != 0)
	return
}

func __gowrap__SDL_DestroyEnvironment(env Environment) {
	__SDL_DestroyEnvironment(uintptr(unsafe.Pointer(env)))
}

func __gowrap__SDL_getenv(name uintptr) (__result uintptr) {
	__result = (uintptr)(unsafe.Pointer(__SDL_getenv(uintptr(unsafe.Pointer(name)))))
	return
}

func __gowrap__SDL_getenv_unsafe(name uintptr) (__result uintptr) {
	__result = (uintptr)(unsafe.Pointer(__SDL_getenv_unsafe(uintptr(unsafe.Pointer(name)))))
	return
}

func __gowrap__SDL_setenv_unsafe(name uintptr, value uintptr, overwrite int32) (__result int32) {
	__result = (int32)(__SDL_setenv_unsafe(uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(value)), *(*int32)(unsafe.Pointer(&overwrite))))
	return
}

func __gowrap__SDL_unsetenv_unsafe(name uintptr) (__result int32) {
	__result = (int32)(__SDL_unsetenv_unsafe(uintptr(unsafe.Pointer(name))))
	return
}

 
func init() {
	Malloc = __gowrap__SDL_malloc
	Calloc = __gowrap__SDL_calloc
	Realloc = __gowrap__SDL_realloc
	Free = __gowrap__SDL_free
	GetOriginalMemoryFunctions = __gowrap__SDL_GetOriginalMemoryFunctions
	GetMemoryFunctions = __gowrap__SDL_GetMemoryFunctions
	SetMemoryFunctions = __gowrap__SDL_SetMemoryFunctions
	AlignedAlloc = __gowrap__SDL_aligned_alloc
	AlignedFree = __gowrap__SDL_aligned_free
	GetNumAllocations = __gowrap__SDL_GetNumAllocations
	GetEnvironment = __gowrap__SDL_GetEnvironment
	CreateEnvironment = __gowrap__SDL_CreateEnvironment
	GetEnvironmentVariable = __gowrap__SDL_GetEnvironmentVariable
	SetEnvironmentVariable = __gowrap__SDL_SetEnvironmentVariable
	UnsetEnvironmentVariable = __gowrap__SDL_UnsetEnvironmentVariable
	DestroyEnvironment = __gowrap__SDL_DestroyEnvironment
	Getenv = __gowrap__SDL_getenv
	GetenvUnsafe = __gowrap__SDL_getenv_unsafe
	SetenvUnsafe = __gowrap__SDL_setenv_unsafe
	UnsetenvUnsafe = __gowrap__SDL_unsetenv_unsafe
}
