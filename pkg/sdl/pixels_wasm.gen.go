//go:build wasm

package sdl

import (
  "unsafe"
)



//go:wasmimport sdl3 SDL_GetPixelFormatName
func __SDL_GetPixelFormatName(int32) uintptr

//go:wasmimport sdl3 SDL_GetMasksForPixelFormat
func __SDL_GetMasksForPixelFormat(int32, uintptr, uintptr, uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetPixelFormatForMasks
func __SDL_GetPixelFormatForMasks(int32, int32, int32, int32, int32) int32

//go:wasmimport sdl3 SDL_GetPixelFormatDetails
func __SDL_GetPixelFormatDetails(int32) uintptr

//go:wasmimport sdl3 SDL_CreatePalette
func __SDL_CreatePalette(int32) uintptr

//go:wasmimport sdl3 SDL_SetPaletteColors
func __SDL_SetPaletteColors(uintptr, uintptr, int32, int32) int32

//go:wasmimport sdl3 SDL_DestroyPalette
func __SDL_DestroyPalette(uintptr)

//go:wasmimport sdl3 SDL_MapRGB
func __SDL_MapRGB(uintptr, uintptr, int32, int32, int32) int32

//go:wasmimport sdl3 SDL_MapRGBA
func __SDL_MapRGBA(uintptr, uintptr, int32, int32, int32, int32) int32

//go:wasmimport sdl3 SDL_GetRGB
func __SDL_GetRGB(int32, uintptr, uintptr, uintptr, uintptr, uintptr)

//go:wasmimport sdl3 SDL_GetRGBA
func __SDL_GetRGBA(int32, uintptr, uintptr, int32, int32, int32, int32)



func __gowrap__SDL_GetPixelFormatName(format PixelFormat) (__result *byte) {
	__result = (*byte)(unsafe.Pointer(__SDL_GetPixelFormatName(*(*int32)(unsafe.Pointer(&format)))))
	return
}

func __gowrap__SDL_GetMasksForPixelFormat(format PixelFormat, bpp *int32, Rmask *uint32, Gmask *uint32, Bmask *uint32, Amask *uint32) (__result bool) {
	__result = (bool)(__SDL_GetMasksForPixelFormat(*(*int32)(unsafe.Pointer(&format)), uintptr(unsafe.Pointer(bpp)), uintptr(unsafe.Pointer(Rmask)), uintptr(unsafe.Pointer(Gmask)), uintptr(unsafe.Pointer(Bmask)), uintptr(unsafe.Pointer(Amask))) != 0)
	return
}

func __gowrap__SDL_GetPixelFormatForMasks(bpp int32, Rmask uint32, Gmask uint32, Bmask uint32, Amask uint32) (__result PixelFormat) {
	__result = (PixelFormat)(__SDL_GetPixelFormatForMasks(*(*int32)(unsafe.Pointer(&bpp)), *(*int32)(unsafe.Pointer(&Rmask)), *(*int32)(unsafe.Pointer(&Gmask)), *(*int32)(unsafe.Pointer(&Bmask)), *(*int32)(unsafe.Pointer(&Amask))))
	return
}

func __gowrap__SDL_GetPixelFormatDetails(format PixelFormat) (__result *PixelFormatDetails) {
	__result = (*PixelFormatDetails)(unsafe.Pointer(__SDL_GetPixelFormatDetails(*(*int32)(unsafe.Pointer(&format)))))
	return
}

func __gowrap__SDL_CreatePalette(ncolors int32) (__result *Palette) {
	__result = (*Palette)(unsafe.Pointer(__SDL_CreatePalette(*(*int32)(unsafe.Pointer(&ncolors)))))
	return
}

func __gowrap__SDL_SetPaletteColors(palette *Palette, colors *Color, firstcolor int32, ncolors int32) (__result bool) {
	__result = (bool)(__SDL_SetPaletteColors(uintptr(unsafe.Pointer(palette)), uintptr(unsafe.Pointer(colors)), *(*int32)(unsafe.Pointer(&firstcolor)), *(*int32)(unsafe.Pointer(&ncolors))) != 0)
	return
}

func __gowrap__SDL_DestroyPalette(palette *Palette) {
	__SDL_DestroyPalette(uintptr(unsafe.Pointer(palette)))
}

func __gowrap__SDL_MapRGB(format *PixelFormatDetails, palette *Palette, r uint8, g uint8, b uint8) (__result uint32) {
	__result = (uint32)(__SDL_MapRGB(uintptr(unsafe.Pointer(format)), uintptr(unsafe.Pointer(palette)), *(*int32)(unsafe.Pointer(&r)), *(*int32)(unsafe.Pointer(&g)), *(*int32)(unsafe.Pointer(&b))))
	return
}

func __gowrap__SDL_MapRGBA(format *PixelFormatDetails, palette *Palette, r uint8, g uint8, b uint8, a uint8) (__result uint32) {
	__result = (uint32)(__SDL_MapRGBA(uintptr(unsafe.Pointer(format)), uintptr(unsafe.Pointer(palette)), *(*int32)(unsafe.Pointer(&r)), *(*int32)(unsafe.Pointer(&g)), *(*int32)(unsafe.Pointer(&b)), *(*int32)(unsafe.Pointer(&a))))
	return
}

func __gowrap__SDL_GetRGB(pixelvalue uint32, format *PixelFormatDetails, palette *Palette, r *uint8, g *uint8, b *uint8) {
	__SDL_GetRGB(*(*int32)(unsafe.Pointer(&pixelvalue)), uintptr(unsafe.Pointer(format)), uintptr(unsafe.Pointer(palette)), uintptr(unsafe.Pointer(r)), uintptr(unsafe.Pointer(g)), uintptr(unsafe.Pointer(b)))
}

func __gowrap__SDL_GetRGBA(pixelvalue uint32, format *PixelFormatDetails, palette *Palette, r uint8, g uint8, b uint8, a uint8) {
	__SDL_GetRGBA(*(*int32)(unsafe.Pointer(&pixelvalue)), uintptr(unsafe.Pointer(format)), uintptr(unsafe.Pointer(palette)), *(*int32)(unsafe.Pointer(&r)), *(*int32)(unsafe.Pointer(&g)), *(*int32)(unsafe.Pointer(&b)), *(*int32)(unsafe.Pointer(&a)))
}

 
func init() {
	GetPixelFormatName = __gowrap__SDL_GetPixelFormatName
	GetMasksForPixelFormat = __gowrap__SDL_GetMasksForPixelFormat
	GetPixelFormatForMasks = __gowrap__SDL_GetPixelFormatForMasks
	GetPixelFormatDetails = __gowrap__SDL_GetPixelFormatDetails
	CreatePalette = __gowrap__SDL_CreatePalette
	SetPaletteColors = __gowrap__SDL_SetPaletteColors
	DestroyPalette = __gowrap__SDL_DestroyPalette
	MapRGB = __gowrap__SDL_MapRGB
	MapRGBA = __gowrap__SDL_MapRGBA
	GetRGB = __gowrap__SDL_GetRGB
	GetRGBA = __gowrap__SDL_GetRGBA
}
