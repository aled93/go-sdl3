//go:build wasm

package sdl

import (
  "runtime"
  "unsafe"
)



//go:wasmimport sdl3 SDL_CreateSurface
func __SDL_CreateSurface(int32, int32, int32) uintptr

//go:wasmimport sdl3 SDL_CreateSurfaceFrom
func __SDL_CreateSurfaceFrom(int32, int32, int32, uintptr, int32) uintptr

//go:wasmimport sdl3 SDL_DestroySurface
func __SDL_DestroySurface(uintptr)

//go:wasmimport sdl3 SDL_GetSurfaceProperties
func __SDL_GetSurfaceProperties(uintptr) int32

//go:wasmimport sdl3 SDL_SetSurfaceColorspace
func __SDL_SetSurfaceColorspace(uintptr, int32) int32

//go:wasmimport sdl3 SDL_GetSurfaceColorspace
func __SDL_GetSurfaceColorspace(uintptr) int32

//go:wasmimport sdl3 SDL_CreateSurfacePalette
func __SDL_CreateSurfacePalette(uintptr) uintptr

//go:wasmimport sdl3 SDL_SetSurfacePalette
func __SDL_SetSurfacePalette(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetSurfacePalette
func __SDL_GetSurfacePalette(uintptr) uintptr

//go:wasmimport sdl3 SDL_AddSurfaceAlternateImage
func __SDL_AddSurfaceAlternateImage(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_SurfaceHasAlternateImages
func __SDL_SurfaceHasAlternateImages(uintptr) int32

//go:wasmimport sdl3 SDL_RemoveSurfaceAlternateImages
func __SDL_RemoveSurfaceAlternateImages(uintptr)

//go:wasmimport sdl3 SDL_LockSurface
func __SDL_LockSurface(uintptr) int32

//go:wasmimport sdl3 SDL_UnlockSurface
func __SDL_UnlockSurface(uintptr)

//go:wasmimport sdl3 SDL_LoadBMP_IO
func __SDL_LoadBMP_IO(uintptr, int32) uintptr

//go:wasmimport sdl3 SDL_LoadBMP
func __SDL_LoadBMP(uintptr) uintptr

//go:wasmimport sdl3 SDL_SaveBMP_IO
func __SDL_SaveBMP_IO(uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_SaveBMP
func __SDL_SaveBMP(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_SetSurfaceRLE
func __SDL_SetSurfaceRLE(uintptr, int32) int32

//go:wasmimport sdl3 SDL_SurfaceHasRLE
func __SDL_SurfaceHasRLE(uintptr) int32

//go:wasmimport sdl3 SDL_SetSurfaceColorKey
func __SDL_SetSurfaceColorKey(uintptr, int32, int32) int32

//go:wasmimport sdl3 SDL_SurfaceHasColorKey
func __SDL_SurfaceHasColorKey(uintptr) int32

//go:wasmimport sdl3 SDL_GetSurfaceColorKey
func __SDL_GetSurfaceColorKey(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_SetSurfaceColorMod
func __SDL_SetSurfaceColorMod(uintptr, int32, int32, int32) int32

//go:wasmimport sdl3 SDL_GetSurfaceColorMod
func __SDL_GetSurfaceColorMod(uintptr, uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_SetSurfaceAlphaMod
func __SDL_SetSurfaceAlphaMod(uintptr, int32) int32

//go:wasmimport sdl3 SDL_GetSurfaceAlphaMod
func __SDL_GetSurfaceAlphaMod(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_SetSurfaceBlendMode
func __SDL_SetSurfaceBlendMode(uintptr, int32) int32

//go:wasmimport sdl3 SDL_GetSurfaceBlendMode
func __SDL_GetSurfaceBlendMode(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_SetSurfaceClipRect
func __SDL_SetSurfaceClipRect(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetSurfaceClipRect
func __SDL_GetSurfaceClipRect(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_FlipSurface
func __SDL_FlipSurface(uintptr, int32) int32

//go:wasmimport sdl3 SDL_DuplicateSurface
func __SDL_DuplicateSurface(uintptr) uintptr

//go:wasmimport sdl3 SDL_ScaleSurface
func __SDL_ScaleSurface(uintptr, int32, int32, int32) uintptr

//go:wasmimport sdl3 SDL_ConvertSurface
func __SDL_ConvertSurface(uintptr, int32) uintptr

//go:wasmimport sdl3 SDL_ConvertSurfaceAndColorspace
func __SDL_ConvertSurfaceAndColorspace(uintptr, int32, uintptr, int32, int32) uintptr

//go:wasmimport sdl3 SDL_ConvertPixels
func __SDL_ConvertPixels(int32, int32, int32, uintptr, int32, int32, uintptr, int32) int32

//go:wasmimport sdl3 SDL_ConvertPixelsAndColorspace
func __SDL_ConvertPixelsAndColorspace(int32, int32, int32, int32, int32, uintptr, int32, int32, int32, int32, uintptr, int32) int32

//go:wasmimport sdl3 SDL_PremultiplyAlpha
func __SDL_PremultiplyAlpha(int32, int32, int32, uintptr, int32, int32, uintptr, int32, int32) int32

//go:wasmimport sdl3 SDL_PremultiplySurfaceAlpha
func __SDL_PremultiplySurfaceAlpha(uintptr, int32) int32

//go:wasmimport sdl3 SDL_ClearSurface
func __SDL_ClearSurface(uintptr, float32, float32, float32, float32) int32

//go:wasmimport sdl3 SDL_FillSurfaceRect
func __SDL_FillSurfaceRect(uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_FillSurfaceRects
func __SDL_FillSurfaceRects(uintptr, uintptr, int32, int32) int32

//go:wasmimport sdl3 SDL_BlitSurface
func __SDL_BlitSurface(uintptr, uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_BlitSurfaceUnchecked
func __SDL_BlitSurfaceUnchecked(uintptr, uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_BlitSurfaceScaled
func __SDL_BlitSurfaceScaled(uintptr, uintptr, uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_BlitSurfaceUncheckedScaled
func __SDL_BlitSurfaceUncheckedScaled(uintptr, uintptr, uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_StretchSurface
func __SDL_StretchSurface(uintptr, uintptr, uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_BlitSurfaceTiled
func __SDL_BlitSurfaceTiled(uintptr, uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_BlitSurfaceTiledWithScale
func __SDL_BlitSurfaceTiledWithScale(uintptr, uintptr, float32, int32, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_BlitSurface9Grid
func __SDL_BlitSurface9Grid(uintptr, uintptr, int32, int32, int32, int32, float32, int32, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_MapSurfaceRGB
func __SDL_MapSurfaceRGB(uintptr, int32, int32, int32) int32

//go:wasmimport sdl3 SDL_MapSurfaceRGBA
func __SDL_MapSurfaceRGBA(uintptr, int32, int32, int32, int32) int32

//go:wasmimport sdl3 SDL_ReadSurfacePixel
func __SDL_ReadSurfacePixel(uintptr, int32, int32, uintptr, uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_ReadSurfacePixelFloat
func __SDL_ReadSurfacePixelFloat(uintptr, int32, int32, uintptr, uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_WriteSurfacePixel
func __SDL_WriteSurfacePixel(uintptr, int32, int32, int32, int32, int32, int32) int32

//go:wasmimport sdl3 SDL_WriteSurfacePixelFloat
func __SDL_WriteSurfacePixelFloat(uintptr, int32, int32, float32, float32, float32, float32) int32



func __gowrap__SDL_CreateSurface(width int32, height int32, format PixelFormat) (__result *Surface) {
	__result = (*Surface)(unsafe.Pointer(__SDL_CreateSurface(*(*int32)(unsafe.Pointer(&width)), *(*int32)(unsafe.Pointer(&height)), *(*int32)(unsafe.Pointer(&format)))))
	return
}

func __gowrap__SDL_CreateSurfaceFrom(width int32, height int32, format PixelFormat, pixels uintptr, pitch int32) (__result *Surface) {
	__result = (*Surface)(unsafe.Pointer(__SDL_CreateSurfaceFrom(*(*int32)(unsafe.Pointer(&width)), *(*int32)(unsafe.Pointer(&height)), *(*int32)(unsafe.Pointer(&format)), uintptr(unsafe.Pointer(pixels)), *(*int32)(unsafe.Pointer(&pitch)))))
	return
}

func __gowrap__SDL_DestroySurface(surface *Surface) {
	__SDL_DestroySurface(uintptr(unsafe.Pointer(surface)))
}

func __gowrap__SDL_GetSurfaceProperties(surface *Surface) (__result PropertiesID) {
	__result = (PropertiesID)(__SDL_GetSurfaceProperties(uintptr(unsafe.Pointer(surface))))
	return
}

func __gowrap__SDL_SetSurfaceColorspace(surface *Surface, colorspace Colorspace) (__result bool) {
	__result = (bool)(__SDL_SetSurfaceColorspace(uintptr(unsafe.Pointer(surface)), *(*int32)(unsafe.Pointer(&colorspace))) != 0)
	return
}

func __gowrap__SDL_GetSurfaceColorspace(surface *Surface) (__result Colorspace) {
	__result = (Colorspace)(__SDL_GetSurfaceColorspace(uintptr(unsafe.Pointer(surface))))
	return
}

func __gowrap__SDL_CreateSurfacePalette(surface *Surface) (__result *Palette) {
	__result = (*Palette)(unsafe.Pointer(__SDL_CreateSurfacePalette(uintptr(unsafe.Pointer(surface)))))
	return
}

func __gowrap__SDL_SetSurfacePalette(surface *Surface, palette *Palette) (__result bool) {
	__result = (bool)(__SDL_SetSurfacePalette(uintptr(unsafe.Pointer(surface)), uintptr(unsafe.Pointer(palette))) != 0)
	return
}

func __gowrap__SDL_GetSurfacePalette(surface *Surface) (__result *Palette) {
	__result = (*Palette)(unsafe.Pointer(__SDL_GetSurfacePalette(uintptr(unsafe.Pointer(surface)))))
	return
}

func __gowrap__SDL_AddSurfaceAlternateImage(surface *Surface, image *Surface) (__result bool) {
	__result = (bool)(__SDL_AddSurfaceAlternateImage(uintptr(unsafe.Pointer(surface)), uintptr(unsafe.Pointer(image))) != 0)
	return
}

func __gowrap__SDL_SurfaceHasAlternateImages(surface *Surface) (__result bool) {
	__result = (bool)(__SDL_SurfaceHasAlternateImages(uintptr(unsafe.Pointer(surface))) != 0)
	return
}

func __gowrap__SDL_RemoveSurfaceAlternateImages(surface *Surface) {
	__SDL_RemoveSurfaceAlternateImages(uintptr(unsafe.Pointer(surface)))
}

func __gowrap__SDL_LockSurface(surface *Surface) (__result bool) {
	__result = (bool)(__SDL_LockSurface(uintptr(unsafe.Pointer(surface))) != 0)
	return
}

func __gowrap__SDL_UnlockSurface(surface *Surface) {
	__SDL_UnlockSurface(uintptr(unsafe.Pointer(surface)))
}

func __gowrap__SDL_LoadBMP_IO(src *IOStream, closeio bool) (__result *Surface) {
	__result = (*Surface)(unsafe.Pointer(__SDL_LoadBMP_IO(uintptr(unsafe.Pointer(src)), *(*int32)(unsafe.Pointer(&closeio)))))
	return
}

func __gowrap__SDL_LoadBMP(file string) (__result *Surface) {
	__result = (*Surface)(unsafe.Pointer(__SDL_LoadBMP(((*[2]uintptr)(unsafe.Pointer(&file)))[0])))
	runtime.KeepAlive(file)
	return
}

func __gowrap__SDL_SaveBMP_IO(surface *Surface, dst *IOStream, closeio bool) (__result bool) {
	__result = (bool)(__SDL_SaveBMP_IO(uintptr(unsafe.Pointer(surface)), uintptr(unsafe.Pointer(dst)), *(*int32)(unsafe.Pointer(&closeio))) != 0)
	return
}

func __gowrap__SDL_SaveBMP(surface *Surface, file string) (__result bool) {
	__result = (bool)(__SDL_SaveBMP(uintptr(unsafe.Pointer(surface)), ((*[2]uintptr)(unsafe.Pointer(&file)))[0]) != 0)
	runtime.KeepAlive(file)
	return
}

func __gowrap__SDL_SetSurfaceRLE(surface *Surface, enabled bool) (__result bool) {
	__result = (bool)(__SDL_SetSurfaceRLE(uintptr(unsafe.Pointer(surface)), *(*int32)(unsafe.Pointer(&enabled))) != 0)
	return
}

func __gowrap__SDL_SurfaceHasRLE(surface *Surface) (__result bool) {
	__result = (bool)(__SDL_SurfaceHasRLE(uintptr(unsafe.Pointer(surface))) != 0)
	return
}

func __gowrap__SDL_SetSurfaceColorKey(surface *Surface, enabled bool, key uint32) (__result bool) {
	__result = (bool)(__SDL_SetSurfaceColorKey(uintptr(unsafe.Pointer(surface)), *(*int32)(unsafe.Pointer(&enabled)), *(*int32)(unsafe.Pointer(&key))) != 0)
	return
}

func __gowrap__SDL_SurfaceHasColorKey(surface *Surface) (__result bool) {
	__result = (bool)(__SDL_SurfaceHasColorKey(uintptr(unsafe.Pointer(surface))) != 0)
	return
}

func __gowrap__SDL_GetSurfaceColorKey(surface *Surface, key *uint32) (__result bool) {
	__result = (bool)(__SDL_GetSurfaceColorKey(uintptr(unsafe.Pointer(surface)), uintptr(unsafe.Pointer(key))) != 0)
	return
}

func __gowrap__SDL_SetSurfaceColorMod(surface *Surface, r uint8, g uint8, b uint8) (__result bool) {
	__result = (bool)(__SDL_SetSurfaceColorMod(uintptr(unsafe.Pointer(surface)), *(*int32)(unsafe.Pointer(&r)), *(*int32)(unsafe.Pointer(&g)), *(*int32)(unsafe.Pointer(&b))) != 0)
	return
}

func __gowrap__SDL_GetSurfaceColorMod(surface *Surface, r *uint8, g *uint8, b *uint8) (__result bool) {
	__result = (bool)(__SDL_GetSurfaceColorMod(uintptr(unsafe.Pointer(surface)), uintptr(unsafe.Pointer(r)), uintptr(unsafe.Pointer(g)), uintptr(unsafe.Pointer(b))) != 0)
	return
}

func __gowrap__SDL_SetSurfaceAlphaMod(surface *Surface, alpha uint8) (__result bool) {
	__result = (bool)(__SDL_SetSurfaceAlphaMod(uintptr(unsafe.Pointer(surface)), *(*int32)(unsafe.Pointer(&alpha))) != 0)
	return
}

func __gowrap__SDL_GetSurfaceAlphaMod(surface *Surface, alpha *uint8) (__result bool) {
	__result = (bool)(__SDL_GetSurfaceAlphaMod(uintptr(unsafe.Pointer(surface)), uintptr(unsafe.Pointer(alpha))) != 0)
	return
}

func __gowrap__SDL_SetSurfaceBlendMode(surface *Surface, blendMode BlendMode) (__result bool) {
	__result = (bool)(__SDL_SetSurfaceBlendMode(uintptr(unsafe.Pointer(surface)), *(*int32)(unsafe.Pointer(&blendMode))) != 0)
	return
}

func __gowrap__SDL_GetSurfaceBlendMode(surface *Surface, blendMode *BlendMode) (__result bool) {
	__result = (bool)(__SDL_GetSurfaceBlendMode(uintptr(unsafe.Pointer(surface)), uintptr(unsafe.Pointer(blendMode))) != 0)
	return
}

func __gowrap__SDL_SetSurfaceClipRect(surface *Surface, rect *Rect) (__result bool) {
	__result = (bool)(__SDL_SetSurfaceClipRect(uintptr(unsafe.Pointer(surface)), uintptr(unsafe.Pointer(rect))) != 0)
	return
}

func __gowrap__SDL_GetSurfaceClipRect(surface *Surface, rect *Rect) (__result bool) {
	__result = (bool)(__SDL_GetSurfaceClipRect(uintptr(unsafe.Pointer(surface)), uintptr(unsafe.Pointer(rect))) != 0)
	return
}

func __gowrap__SDL_FlipSurface(surface *Surface, flip FlipMode) (__result bool) {
	__result = (bool)(__SDL_FlipSurface(uintptr(unsafe.Pointer(surface)), *(*int32)(unsafe.Pointer(&flip))) != 0)
	return
}

func __gowrap__SDL_DuplicateSurface(surface *Surface) (__result *Surface) {
	__result = (*Surface)(unsafe.Pointer(__SDL_DuplicateSurface(uintptr(unsafe.Pointer(surface)))))
	return
}

func __gowrap__SDL_ScaleSurface(surface *Surface, width int32, height int32, scaleMode ScaleMode) (__result *Surface) {
	__result = (*Surface)(unsafe.Pointer(__SDL_ScaleSurface(uintptr(unsafe.Pointer(surface)), *(*int32)(unsafe.Pointer(&width)), *(*int32)(unsafe.Pointer(&height)), *(*int32)(unsafe.Pointer(&scaleMode)))))
	return
}

func __gowrap__SDL_ConvertSurface(surface *Surface, format PixelFormat) (__result *Surface) {
	__result = (*Surface)(unsafe.Pointer(__SDL_ConvertSurface(uintptr(unsafe.Pointer(surface)), *(*int32)(unsafe.Pointer(&format)))))
	return
}

func __gowrap__SDL_ConvertSurfaceAndColorspace(surface *Surface, format PixelFormat, palette *Palette, colorspace Colorspace, props PropertiesID) (__result *Surface) {
	__result = (*Surface)(unsafe.Pointer(__SDL_ConvertSurfaceAndColorspace(uintptr(unsafe.Pointer(surface)), *(*int32)(unsafe.Pointer(&format)), uintptr(unsafe.Pointer(palette)), *(*int32)(unsafe.Pointer(&colorspace)), *(*int32)(unsafe.Pointer(&props)))))
	return
}

func __gowrap__SDL_ConvertPixels(width int32, height int32, src_format PixelFormat, src uintptr, src_pitch int32, dst_format PixelFormat, dst uintptr, dst_pitch int32) (__result bool) {
	__result = (bool)(__SDL_ConvertPixels(*(*int32)(unsafe.Pointer(&width)), *(*int32)(unsafe.Pointer(&height)), *(*int32)(unsafe.Pointer(&src_format)), uintptr(unsafe.Pointer(src)), *(*int32)(unsafe.Pointer(&src_pitch)), *(*int32)(unsafe.Pointer(&dst_format)), uintptr(unsafe.Pointer(dst)), *(*int32)(unsafe.Pointer(&dst_pitch))) != 0)
	return
}

func __gowrap__SDL_ConvertPixelsAndColorspace(width int32, height int32, src_format PixelFormat, src_colorspace Colorspace, src_properties PropertiesID, src uintptr, src_pitch int32, dst_format PixelFormat, dst_colorspace Colorspace, dst_properties PropertiesID, dst uintptr, dst_pitch int32) (__result bool) {
	__result = (bool)(__SDL_ConvertPixelsAndColorspace(*(*int32)(unsafe.Pointer(&width)), *(*int32)(unsafe.Pointer(&height)), *(*int32)(unsafe.Pointer(&src_format)), *(*int32)(unsafe.Pointer(&src_colorspace)), *(*int32)(unsafe.Pointer(&src_properties)), uintptr(unsafe.Pointer(src)), *(*int32)(unsafe.Pointer(&src_pitch)), *(*int32)(unsafe.Pointer(&dst_format)), *(*int32)(unsafe.Pointer(&dst_colorspace)), *(*int32)(unsafe.Pointer(&dst_properties)), uintptr(unsafe.Pointer(dst)), *(*int32)(unsafe.Pointer(&dst_pitch))) != 0)
	return
}

func __gowrap__SDL_PremultiplyAlpha(width int32, height int32, src_format PixelFormat, src uintptr, src_pitch int32, dst_format PixelFormat, dst uintptr, dst_pitch int32, linear bool) (__result bool) {
	__result = (bool)(__SDL_PremultiplyAlpha(*(*int32)(unsafe.Pointer(&width)), *(*int32)(unsafe.Pointer(&height)), *(*int32)(unsafe.Pointer(&src_format)), uintptr(unsafe.Pointer(src)), *(*int32)(unsafe.Pointer(&src_pitch)), *(*int32)(unsafe.Pointer(&dst_format)), uintptr(unsafe.Pointer(dst)), *(*int32)(unsafe.Pointer(&dst_pitch)), *(*int32)(unsafe.Pointer(&linear))) != 0)
	return
}

func __gowrap__SDL_PremultiplySurfaceAlpha(surface *Surface, linear bool) (__result bool) {
	__result = (bool)(__SDL_PremultiplySurfaceAlpha(uintptr(unsafe.Pointer(surface)), *(*int32)(unsafe.Pointer(&linear))) != 0)
	return
}

func __gowrap__SDL_ClearSurface(surface *Surface, r float32, g float32, b float32, a float32) (__result bool) {
	__result = (bool)(__SDL_ClearSurface(uintptr(unsafe.Pointer(surface)), *(*float32)(unsafe.Pointer(&r)), *(*float32)(unsafe.Pointer(&g)), *(*float32)(unsafe.Pointer(&b)), *(*float32)(unsafe.Pointer(&a))) != 0)
	return
}

func __gowrap__SDL_FillSurfaceRect(dst *Surface, rect *Rect, color uint32) (__result bool) {
	__result = (bool)(__SDL_FillSurfaceRect(uintptr(unsafe.Pointer(dst)), uintptr(unsafe.Pointer(rect)), *(*int32)(unsafe.Pointer(&color))) != 0)
	return
}

func __gowrap__SDL_FillSurfaceRects(dst *Surface, rects *Rect, count int32, color uint32) (__result bool) {
	__result = (bool)(__SDL_FillSurfaceRects(uintptr(unsafe.Pointer(dst)), uintptr(unsafe.Pointer(rects)), *(*int32)(unsafe.Pointer(&count)), *(*int32)(unsafe.Pointer(&color))) != 0)
	return
}

func __gowrap__SDL_BlitSurface(src *Surface, srcrect *Rect, dst *Surface, dstrect *Rect) (__result bool) {
	__result = (bool)(__SDL_BlitSurface(uintptr(unsafe.Pointer(src)), uintptr(unsafe.Pointer(srcrect)), uintptr(unsafe.Pointer(dst)), uintptr(unsafe.Pointer(dstrect))) != 0)
	return
}

func __gowrap__SDL_BlitSurfaceUnchecked(src *Surface, srcrect *Rect, dst *Surface, dstrect *Rect) (__result bool) {
	__result = (bool)(__SDL_BlitSurfaceUnchecked(uintptr(unsafe.Pointer(src)), uintptr(unsafe.Pointer(srcrect)), uintptr(unsafe.Pointer(dst)), uintptr(unsafe.Pointer(dstrect))) != 0)
	return
}

func __gowrap__SDL_BlitSurfaceScaled(src *Surface, srcrect *Rect, dst *Surface, dstrect *Rect, scaleMode ScaleMode) (__result bool) {
	__result = (bool)(__SDL_BlitSurfaceScaled(uintptr(unsafe.Pointer(src)), uintptr(unsafe.Pointer(srcrect)), uintptr(unsafe.Pointer(dst)), uintptr(unsafe.Pointer(dstrect)), *(*int32)(unsafe.Pointer(&scaleMode))) != 0)
	return
}

func __gowrap__SDL_BlitSurfaceUncheckedScaled(src *Surface, srcrect *Rect, dst *Surface, dstrect *Rect, scaleMode ScaleMode) (__result bool) {
	__result = (bool)(__SDL_BlitSurfaceUncheckedScaled(uintptr(unsafe.Pointer(src)), uintptr(unsafe.Pointer(srcrect)), uintptr(unsafe.Pointer(dst)), uintptr(unsafe.Pointer(dstrect)), *(*int32)(unsafe.Pointer(&scaleMode))) != 0)
	return
}

func __gowrap__SDL_StretchSurface(src *Surface, srcrect *Rect, dst *Surface, dstrect *Rect, scaleMode ScaleMode) (__result bool) {
	__result = (bool)(__SDL_StretchSurface(uintptr(unsafe.Pointer(src)), uintptr(unsafe.Pointer(srcrect)), uintptr(unsafe.Pointer(dst)), uintptr(unsafe.Pointer(dstrect)), *(*int32)(unsafe.Pointer(&scaleMode))) != 0)
	return
}

func __gowrap__SDL_BlitSurfaceTiled(src *Surface, srcrect *Rect, dst *Surface, dstrect *Rect) (__result bool) {
	__result = (bool)(__SDL_BlitSurfaceTiled(uintptr(unsafe.Pointer(src)), uintptr(unsafe.Pointer(srcrect)), uintptr(unsafe.Pointer(dst)), uintptr(unsafe.Pointer(dstrect))) != 0)
	return
}

func __gowrap__SDL_BlitSurfaceTiledWithScale(src *Surface, srcrect *Rect, scale float32, scaleMode ScaleMode, dst *Surface, dstrect *Rect) (__result bool) {
	__result = (bool)(__SDL_BlitSurfaceTiledWithScale(uintptr(unsafe.Pointer(src)), uintptr(unsafe.Pointer(srcrect)), *(*float32)(unsafe.Pointer(&scale)), *(*int32)(unsafe.Pointer(&scaleMode)), uintptr(unsafe.Pointer(dst)), uintptr(unsafe.Pointer(dstrect))) != 0)
	return
}

func __gowrap__SDL_BlitSurface9Grid(src *Surface, srcrect *Rect, left_width int32, right_width int32, top_height int32, bottom_height int32, scale float32, scaleMode ScaleMode, dst *Surface, dstrect *Rect) (__result bool) {
	__result = (bool)(__SDL_BlitSurface9Grid(uintptr(unsafe.Pointer(src)), uintptr(unsafe.Pointer(srcrect)), *(*int32)(unsafe.Pointer(&left_width)), *(*int32)(unsafe.Pointer(&right_width)), *(*int32)(unsafe.Pointer(&top_height)), *(*int32)(unsafe.Pointer(&bottom_height)), *(*float32)(unsafe.Pointer(&scale)), *(*int32)(unsafe.Pointer(&scaleMode)), uintptr(unsafe.Pointer(dst)), uintptr(unsafe.Pointer(dstrect))) != 0)
	return
}

func __gowrap__SDL_MapSurfaceRGB(surface *Surface, r uint8, g uint8, b uint8) (__result uint32) {
	__result = (uint32)(__SDL_MapSurfaceRGB(uintptr(unsafe.Pointer(surface)), *(*int32)(unsafe.Pointer(&r)), *(*int32)(unsafe.Pointer(&g)), *(*int32)(unsafe.Pointer(&b))))
	return
}

func __gowrap__SDL_MapSurfaceRGBA(surface *Surface, r uint8, g uint8, b uint8, a uint8) (__result uint32) {
	__result = (uint32)(__SDL_MapSurfaceRGBA(uintptr(unsafe.Pointer(surface)), *(*int32)(unsafe.Pointer(&r)), *(*int32)(unsafe.Pointer(&g)), *(*int32)(unsafe.Pointer(&b)), *(*int32)(unsafe.Pointer(&a))))
	return
}

func __gowrap__SDL_ReadSurfacePixel(surface *Surface, x int32, y int32, r *uint8, g *uint8, b *uint8, a *uint8) (__result bool) {
	__result = (bool)(__SDL_ReadSurfacePixel(uintptr(unsafe.Pointer(surface)), *(*int32)(unsafe.Pointer(&x)), *(*int32)(unsafe.Pointer(&y)), uintptr(unsafe.Pointer(r)), uintptr(unsafe.Pointer(g)), uintptr(unsafe.Pointer(b)), uintptr(unsafe.Pointer(a))) != 0)
	return
}

func __gowrap__SDL_ReadSurfacePixelFloat(surface *Surface, x int32, y int32, r *float32, g *float32, b *float32, a *float32) (__result bool) {
	__result = (bool)(__SDL_ReadSurfacePixelFloat(uintptr(unsafe.Pointer(surface)), *(*int32)(unsafe.Pointer(&x)), *(*int32)(unsafe.Pointer(&y)), uintptr(unsafe.Pointer(r)), uintptr(unsafe.Pointer(g)), uintptr(unsafe.Pointer(b)), uintptr(unsafe.Pointer(a))) != 0)
	return
}

func __gowrap__SDL_WriteSurfacePixel(surface *Surface, x int32, y int32, r uint8, g uint8, b uint8, a uint8) (__result bool) {
	__result = (bool)(__SDL_WriteSurfacePixel(uintptr(unsafe.Pointer(surface)), *(*int32)(unsafe.Pointer(&x)), *(*int32)(unsafe.Pointer(&y)), *(*int32)(unsafe.Pointer(&r)), *(*int32)(unsafe.Pointer(&g)), *(*int32)(unsafe.Pointer(&b)), *(*int32)(unsafe.Pointer(&a))) != 0)
	return
}

func __gowrap__SDL_WriteSurfacePixelFloat(surface *Surface, x int32, y int32, r float32, g float32, b float32, a float32) (__result bool) {
	__result = (bool)(__SDL_WriteSurfacePixelFloat(uintptr(unsafe.Pointer(surface)), *(*int32)(unsafe.Pointer(&x)), *(*int32)(unsafe.Pointer(&y)), *(*float32)(unsafe.Pointer(&r)), *(*float32)(unsafe.Pointer(&g)), *(*float32)(unsafe.Pointer(&b)), *(*float32)(unsafe.Pointer(&a))) != 0)
	return
}

 
func init() {
	CreateSurface = __gowrap__SDL_CreateSurface
	CreateSurfaceFrom = __gowrap__SDL_CreateSurfaceFrom
	DestroySurface = __gowrap__SDL_DestroySurface
	GetSurfaceProperties = __gowrap__SDL_GetSurfaceProperties
	SetSurfaceColorspace = __gowrap__SDL_SetSurfaceColorspace
	GetSurfaceColorspace = __gowrap__SDL_GetSurfaceColorspace
	CreateSurfacePalette = __gowrap__SDL_CreateSurfacePalette
	SetSurfacePalette = __gowrap__SDL_SetSurfacePalette
	GetSurfacePalette = __gowrap__SDL_GetSurfacePalette
	AddSurfaceAlternateImage = __gowrap__SDL_AddSurfaceAlternateImage
	SurfaceHasAlternateImages = __gowrap__SDL_SurfaceHasAlternateImages
	RemoveSurfaceAlternateImages = __gowrap__SDL_RemoveSurfaceAlternateImages
	LockSurface = __gowrap__SDL_LockSurface
	UnlockSurface = __gowrap__SDL_UnlockSurface
	LoadBMP_IO = __gowrap__SDL_LoadBMP_IO
	LoadBMP = __gowrap__SDL_LoadBMP
	SaveBMP_IO = __gowrap__SDL_SaveBMP_IO
	SaveBMP = __gowrap__SDL_SaveBMP
	SetSurfaceRLE = __gowrap__SDL_SetSurfaceRLE
	SurfaceHasRLE = __gowrap__SDL_SurfaceHasRLE
	SetSurfaceColorKey = __gowrap__SDL_SetSurfaceColorKey
	SurfaceHasColorKey = __gowrap__SDL_SurfaceHasColorKey
	GetSurfaceColorKey = __gowrap__SDL_GetSurfaceColorKey
	SetSurfaceColorMod = __gowrap__SDL_SetSurfaceColorMod
	GetSurfaceColorMod = __gowrap__SDL_GetSurfaceColorMod
	SetSurfaceAlphaMod = __gowrap__SDL_SetSurfaceAlphaMod
	GetSurfaceAlphaMod = __gowrap__SDL_GetSurfaceAlphaMod
	SetSurfaceBlendMode = __gowrap__SDL_SetSurfaceBlendMode
	GetSurfaceBlendMode = __gowrap__SDL_GetSurfaceBlendMode
	SetSurfaceClipRect = __gowrap__SDL_SetSurfaceClipRect
	GetSurfaceClipRect = __gowrap__SDL_GetSurfaceClipRect
	FlipSurface = __gowrap__SDL_FlipSurface
	DuplicateSurface = __gowrap__SDL_DuplicateSurface
	ScaleSurface = __gowrap__SDL_ScaleSurface
	ConvertSurface = __gowrap__SDL_ConvertSurface
	ConvertSurfaceAndColorspace = __gowrap__SDL_ConvertSurfaceAndColorspace
	ConvertPixels = __gowrap__SDL_ConvertPixels
	ConvertPixelsAndColorspace = __gowrap__SDL_ConvertPixelsAndColorspace
	PremultiplyAlpha = __gowrap__SDL_PremultiplyAlpha
	PremultiplySurfaceAlpha = __gowrap__SDL_PremultiplySurfaceAlpha
	ClearSurface = __gowrap__SDL_ClearSurface
	FillSurfaceRect = __gowrap__SDL_FillSurfaceRect
	FillSurfaceRects = __gowrap__SDL_FillSurfaceRects
	BlitSurface = __gowrap__SDL_BlitSurface
	BlitSurfaceUnchecked = __gowrap__SDL_BlitSurfaceUnchecked
	BlitSurfaceScaled = __gowrap__SDL_BlitSurfaceScaled
	BlitSurfaceUncheckedScaled = __gowrap__SDL_BlitSurfaceUncheckedScaled
	StretchSurface = __gowrap__SDL_StretchSurface
	BlitSurfaceTiled = __gowrap__SDL_BlitSurfaceTiled
	BlitSurfaceTiledWithScale = __gowrap__SDL_BlitSurfaceTiledWithScale
	BlitSurface9Grid = __gowrap__SDL_BlitSurface9Grid
	MapSurfaceRGB = __gowrap__SDL_MapSurfaceRGB
	MapSurfaceRGBA = __gowrap__SDL_MapSurfaceRGBA
	ReadSurfacePixel = __gowrap__SDL_ReadSurfacePixel
	ReadSurfacePixelFloat = __gowrap__SDL_ReadSurfacePixelFloat
	WriteSurfacePixel = __gowrap__SDL_WriteSurfacePixel
	WriteSurfacePixelFloat = __gowrap__SDL_WriteSurfacePixelFloat
}
