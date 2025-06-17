//go:build wasm

package sdl



//go:wasmimport sdl3 SDL_GetNumLogicalCPUCores
func __SDL_GetNumLogicalCPUCores() int32

//go:wasmimport sdl3 SDL_GetCPUCacheLineSize
func __SDL_GetCPUCacheLineSize() int32

//go:wasmimport sdl3 SDL_HasAltiVec
func __SDL_HasAltiVec() int32

//go:wasmimport sdl3 SDL_HasMMX
func __SDL_HasMMX() int32

//go:wasmimport sdl3 SDL_HasSSE
func __SDL_HasSSE() int32

//go:wasmimport sdl3 SDL_HasSSE2
func __SDL_HasSSE2() int32

//go:wasmimport sdl3 SDL_HasSSE3
func __SDL_HasSSE3() int32

//go:wasmimport sdl3 SDL_HasSSE41
func __SDL_HasSSE41() int32

//go:wasmimport sdl3 SDL_HasSSE42
func __SDL_HasSSE42() int32

//go:wasmimport sdl3 SDL_HasAVX
func __SDL_HasAVX() int32

//go:wasmimport sdl3 SDL_HasAVX2
func __SDL_HasAVX2() int32

//go:wasmimport sdl3 SDL_HasAVX512F
func __SDL_HasAVX512F() int32

//go:wasmimport sdl3 SDL_HasARMSIMD
func __SDL_HasARMSIMD() int32

//go:wasmimport sdl3 SDL_HasNEON
func __SDL_HasNEON() int32

//go:wasmimport sdl3 SDL_HasLSX
func __SDL_HasLSX() int32

//go:wasmimport sdl3 SDL_HasLASX
func __SDL_HasLASX() int32

//go:wasmimport sdl3 SDL_GetSystemRAM
func __SDL_GetSystemRAM() int32

//go:wasmimport sdl3 SDL_GetSIMDAlignment
func __SDL_GetSIMDAlignment() int32



func __gowrap__SDL_GetNumLogicalCPUCores() (__result int32) {
	__result = (int32)(__SDL_GetNumLogicalCPUCores())
	return
}

func __gowrap__SDL_GetCPUCacheLineSize() (__result int32) {
	__result = (int32)(__SDL_GetCPUCacheLineSize())
	return
}

func __gowrap__SDL_HasAltiVec() (__result bool) {
	__result = (bool)(__SDL_HasAltiVec() != 0)
	return
}

func __gowrap__SDL_HasMMX() (__result bool) {
	__result = (bool)(__SDL_HasMMX() != 0)
	return
}

func __gowrap__SDL_HasSSE() (__result bool) {
	__result = (bool)(__SDL_HasSSE() != 0)
	return
}

func __gowrap__SDL_HasSSE2() (__result bool) {
	__result = (bool)(__SDL_HasSSE2() != 0)
	return
}

func __gowrap__SDL_HasSSE3() (__result bool) {
	__result = (bool)(__SDL_HasSSE3() != 0)
	return
}

func __gowrap__SDL_HasSSE41() (__result bool) {
	__result = (bool)(__SDL_HasSSE41() != 0)
	return
}

func __gowrap__SDL_HasSSE42() (__result bool) {
	__result = (bool)(__SDL_HasSSE42() != 0)
	return
}

func __gowrap__SDL_HasAVX() (__result bool) {
	__result = (bool)(__SDL_HasAVX() != 0)
	return
}

func __gowrap__SDL_HasAVX2() (__result bool) {
	__result = (bool)(__SDL_HasAVX2() != 0)
	return
}

func __gowrap__SDL_HasAVX512F() (__result bool) {
	__result = (bool)(__SDL_HasAVX512F() != 0)
	return
}

func __gowrap__SDL_HasARMSIMD() (__result bool) {
	__result = (bool)(__SDL_HasARMSIMD() != 0)
	return
}

func __gowrap__SDL_HasNEON() (__result bool) {
	__result = (bool)(__SDL_HasNEON() != 0)
	return
}

func __gowrap__SDL_HasLSX() (__result bool) {
	__result = (bool)(__SDL_HasLSX() != 0)
	return
}

func __gowrap__SDL_HasLASX() (__result bool) {
	__result = (bool)(__SDL_HasLASX() != 0)
	return
}

func __gowrap__SDL_GetSystemRAM() (__result int32) {
	__result = (int32)(__SDL_GetSystemRAM())
	return
}

func __gowrap__SDL_GetSIMDAlignment() (__result size_t) {
	__result = (size_t)(__SDL_GetSIMDAlignment())
	return
}

 
func init() {
	GetNumLogicalCPUCores = __gowrap__SDL_GetNumLogicalCPUCores
	GetCPUCacheLineSize = __gowrap__SDL_GetCPUCacheLineSize
	HasAltiVec = __gowrap__SDL_HasAltiVec
	HasMMX = __gowrap__SDL_HasMMX
	HasSSE = __gowrap__SDL_HasSSE
	HasSSE2 = __gowrap__SDL_HasSSE2
	HasSSE3 = __gowrap__SDL_HasSSE3
	HasSSE41 = __gowrap__SDL_HasSSE41
	HasSSE42 = __gowrap__SDL_HasSSE42
	HasAVX = __gowrap__SDL_HasAVX
	HasAVX2 = __gowrap__SDL_HasAVX2
	HasAVX512F = __gowrap__SDL_HasAVX512F
	HasARMSIMD = __gowrap__SDL_HasARMSIMD
	HasNEON = __gowrap__SDL_HasNEON
	HasLSX = __gowrap__SDL_HasLSX
	HasLASX = __gowrap__SDL_HasLASX
	GetSystemRAM = __gowrap__SDL_GetSystemRAM
	GetSIMDAlignment = __gowrap__SDL_GetSIMDAlignment
}
