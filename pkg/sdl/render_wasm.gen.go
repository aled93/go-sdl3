//go:build wasm

package sdl

import (
  "runtime"
  "unsafe"
)



//go:wasmimport sdl3 SDL_GetNumRenderDrivers
func __SDL_GetNumRenderDrivers() int32

//go:wasmimport sdl3 SDL_GetRenderDriver
func __SDL_GetRenderDriver(int32) uintptr

//go:wasmimport sdl3 SDL_CreateWindowAndRenderer
func __SDL_CreateWindowAndRenderer(uintptr, int32, int32, int64, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_CreateRenderer
func __SDL_CreateRenderer(uintptr, uintptr) uintptr

//go:wasmimport sdl3 SDL_CreateRendererWithProperties
func __SDL_CreateRendererWithProperties(int32) uintptr

//go:wasmimport sdl3 SDL_CreateSoftwareRenderer
func __SDL_CreateSoftwareRenderer(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetRenderer
func __SDL_GetRenderer(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetRenderWindow
func __SDL_GetRenderWindow(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetRendererName
func __SDL_GetRendererName(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetRendererProperties
func __SDL_GetRendererProperties(uintptr) int32

//go:wasmimport sdl3 SDL_GetRenderOutputSize
func __SDL_GetRenderOutputSize(uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetCurrentRenderOutputSize
func __SDL_GetCurrentRenderOutputSize(uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_CreateTexture
func __SDL_CreateTexture(uintptr, int32, int32, int32, int32) uintptr

//go:wasmimport sdl3 SDL_CreateTextureFromSurface
func __SDL_CreateTextureFromSurface(uintptr, uintptr) uintptr

//go:wasmimport sdl3 SDL_CreateTextureWithProperties
func __SDL_CreateTextureWithProperties(uintptr, int32) uintptr

//go:wasmimport sdl3 SDL_GetTextureProperties
func __SDL_GetTextureProperties(uintptr) int32

//go:wasmimport sdl3 SDL_GetRendererFromTexture
func __SDL_GetRendererFromTexture(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetTextureSize
func __SDL_GetTextureSize(uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_SetTextureColorMod
func __SDL_SetTextureColorMod(uintptr, int32, int32, int32) int32

//go:wasmimport sdl3 SDL_SetTextureColorModFloat
func __SDL_SetTextureColorModFloat(uintptr, float32, float32, float32) int32

//go:wasmimport sdl3 SDL_GetTextureColorMod
func __SDL_GetTextureColorMod(uintptr, uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetTextureColorModFloat
func __SDL_GetTextureColorModFloat(uintptr, uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_SetTextureAlphaMod
func __SDL_SetTextureAlphaMod(uintptr, int32) int32

//go:wasmimport sdl3 SDL_SetTextureAlphaModFloat
func __SDL_SetTextureAlphaModFloat(uintptr, float32) int32

//go:wasmimport sdl3 SDL_GetTextureAlphaMod
func __SDL_GetTextureAlphaMod(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetTextureAlphaModFloat
func __SDL_GetTextureAlphaModFloat(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_SetTextureBlendMode
func __SDL_SetTextureBlendMode(uintptr, int32) int32

//go:wasmimport sdl3 SDL_GetTextureBlendMode
func __SDL_GetTextureBlendMode(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_SetTextureScaleMode
func __SDL_SetTextureScaleMode(uintptr, int32) int32

//go:wasmimport sdl3 SDL_GetTextureScaleMode
func __SDL_GetTextureScaleMode(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_UpdateTexture
func __SDL_UpdateTexture(uintptr, uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_UpdateYUVTexture
func __SDL_UpdateYUVTexture(uintptr, uintptr, uintptr, int32, uintptr, int32, uintptr, int32) int32

//go:wasmimport sdl3 SDL_UpdateNVTexture
func __SDL_UpdateNVTexture(uintptr, uintptr, uintptr, int32, uintptr, int32) int32

//go:wasmimport sdl3 SDL_LockTexture
func __SDL_LockTexture(uintptr, uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_LockTextureToSurface
func __SDL_LockTextureToSurface(uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_UnlockTexture
func __SDL_UnlockTexture(uintptr)

//go:wasmimport sdl3 SDL_SetRenderTarget
func __SDL_SetRenderTarget(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetRenderTarget
func __SDL_GetRenderTarget(uintptr) uintptr

//go:wasmimport sdl3 SDL_SetRenderLogicalPresentation
func __SDL_SetRenderLogicalPresentation(uintptr, int32, int32, int32) int32

//go:wasmimport sdl3 SDL_GetRenderLogicalPresentation
func __SDL_GetRenderLogicalPresentation(uintptr, uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_GetRenderLogicalPresentationRect
func __SDL_GetRenderLogicalPresentationRect(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_RenderCoordinatesFromWindow
func __SDL_RenderCoordinatesFromWindow(uintptr, float32, float32, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_RenderCoordinatesToWindow
func __SDL_RenderCoordinatesToWindow(uintptr, float32, float32, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_ConvertEventToRenderCoordinates
func __SDL_ConvertEventToRenderCoordinates(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_SetRenderViewport
func __SDL_SetRenderViewport(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetRenderViewport
func __SDL_GetRenderViewport(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_RenderViewportSet
func __SDL_RenderViewportSet(uintptr) int32

//go:wasmimport sdl3 SDL_GetRenderSafeArea
func __SDL_GetRenderSafeArea(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_SetRenderClipRect
func __SDL_SetRenderClipRect(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetRenderClipRect
func __SDL_GetRenderClipRect(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_RenderClipEnabled
func __SDL_RenderClipEnabled(uintptr) int32

//go:wasmimport sdl3 SDL_SetRenderScale
func __SDL_SetRenderScale(uintptr, float32, float32) int32

//go:wasmimport sdl3 SDL_GetRenderScale
func __SDL_GetRenderScale(uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_SetRenderDrawColor
func __SDL_SetRenderDrawColor(uintptr, int32, int32, int32, int32) int32

//go:wasmimport sdl3 SDL_SetRenderDrawColorFloat
func __SDL_SetRenderDrawColorFloat(uintptr, float32, float32, float32, float32) int32

//go:wasmimport sdl3 SDL_GetRenderDrawColor
func __SDL_GetRenderDrawColor(uintptr, uintptr, uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetRenderDrawColorFloat
func __SDL_GetRenderDrawColorFloat(uintptr, uintptr, uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_SetRenderColorScale
func __SDL_SetRenderColorScale(uintptr, float32) int32

//go:wasmimport sdl3 SDL_GetRenderColorScale
func __SDL_GetRenderColorScale(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_SetRenderDrawBlendMode
func __SDL_SetRenderDrawBlendMode(uintptr, int32) int32

//go:wasmimport sdl3 SDL_GetRenderDrawBlendMode
func __SDL_GetRenderDrawBlendMode(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_RenderClear
func __SDL_RenderClear(uintptr) int32

//go:wasmimport sdl3 SDL_RenderPoint
func __SDL_RenderPoint(uintptr, float32, float32) int32

//go:wasmimport sdl3 SDL_RenderPoints
func __SDL_RenderPoints(uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_RenderLine
func __SDL_RenderLine(uintptr, float32, float32, float32, float32) int32

//go:wasmimport sdl3 SDL_RenderLines
func __SDL_RenderLines(uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_RenderRect
func __SDL_RenderRect(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_RenderRects
func __SDL_RenderRects(uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_RenderFillRect
func __SDL_RenderFillRect(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_RenderFillRects
func __SDL_RenderFillRects(uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_RenderTexture
func __SDL_RenderTexture(uintptr, uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_RenderTextureRotated
func __SDL_RenderTextureRotated(uintptr, uintptr, uintptr, uintptr, float64, uintptr, int32) int32

//go:wasmimport sdl3 SDL_RenderTextureAffine
func __SDL_RenderTextureAffine(uintptr, uintptr, uintptr, uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_RenderTextureTiled
func __SDL_RenderTextureTiled(uintptr, uintptr, uintptr, float32, uintptr) int32

//go:wasmimport sdl3 SDL_RenderTexture9Grid
func __SDL_RenderTexture9Grid(uintptr, uintptr, uintptr, float32, float32, float32, float32, float32, uintptr) int32

//go:wasmimport sdl3 SDL_RenderTexture9GridTiled
func __SDL_RenderTexture9GridTiled(uintptr, uintptr, uintptr, float32, float32, float32, float32, float32, uintptr, float32) int32

//go:wasmimport sdl3 SDL_RenderGeometry
func __SDL_RenderGeometry(uintptr, uintptr, uintptr, int32, uintptr, int32) int32

//go:wasmimport sdl3 SDL_RenderGeometryRaw
func __SDL_RenderGeometryRaw(uintptr, uintptr, uintptr, int32, uintptr, int32, uintptr, int32, int32, uintptr, int32, int32) int32

//go:wasmimport sdl3 SDL_SetRenderTextureAddressMode
func __SDL_SetRenderTextureAddressMode(uintptr, int32, int32) int32

//go:wasmimport sdl3 SDL_GetRenderTextureAddressMode
func __SDL_GetRenderTextureAddressMode(uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_RenderReadPixels
func __SDL_RenderReadPixels(uintptr, uintptr) uintptr

//go:wasmimport sdl3 SDL_RenderPresent
func __SDL_RenderPresent(uintptr) int32

//go:wasmimport sdl3 SDL_DestroyTexture
func __SDL_DestroyTexture(uintptr)

//go:wasmimport sdl3 SDL_DestroyRenderer
func __SDL_DestroyRenderer(uintptr)

//go:wasmimport sdl3 SDL_FlushRenderer
func __SDL_FlushRenderer(uintptr) int32

//go:wasmimport sdl3 SDL_GetRenderMetalLayer
func __SDL_GetRenderMetalLayer(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetRenderMetalCommandEncoder
func __SDL_GetRenderMetalCommandEncoder(uintptr) uintptr

//go:wasmimport sdl3 SDL_AddVulkanRenderSemaphores
func __SDL_AddVulkanRenderSemaphores(uintptr, int32, int64, int64) int32

//go:wasmimport sdl3 SDL_SetRenderVSync
func __SDL_SetRenderVSync(uintptr, int32) int32

//go:wasmimport sdl3 SDL_GetRenderVSync
func __SDL_GetRenderVSync(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_RenderDebugText
func __SDL_RenderDebugText(uintptr, float32, float32, uintptr) int32

//go:wasmimport sdl3 SDL_SetDefaultTextureScaleMode
func __SDL_SetDefaultTextureScaleMode(uintptr, int32) int32

//go:wasmimport sdl3 SDL_GetDefaultTextureScaleMode
func __SDL_GetDefaultTextureScaleMode(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_CreateGPURenderState
func __SDL_CreateGPURenderState(uintptr, uintptr) uintptr

//go:wasmimport sdl3 SDL_SetGPURenderStateFragmentUniforms
func __SDL_SetGPURenderStateFragmentUniforms(uintptr, int32, uintptr, int32) int32

//go:wasmimport sdl3 SDL_SetRenderGPUState
func __SDL_SetRenderGPUState(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_DestroyGPURenderState
func __SDL_DestroyGPURenderState(uintptr)



func __gowrap__SDL_GetNumRenderDrivers() (__result int32) {
	__result = (int32)(__SDL_GetNumRenderDrivers())
	return
}

func __gowrap__SDL_GetRenderDriver(index int32) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetRenderDriver(*(*int32)(unsafe.Pointer(&index)))).Collect()))
	return
}

func __gowrap__SDL_CreateWindowAndRenderer(title string, width int32, height int32, window_flags WindowFlags, window *Window, renderer *Renderer) (__result bool) {
	__result = (bool)(__SDL_CreateWindowAndRenderer(((*[2]uintptr)(unsafe.Pointer(&title)))[0], *(*int32)(unsafe.Pointer(&width)), *(*int32)(unsafe.Pointer(&height)), *(*int64)(unsafe.Pointer(&window_flags)), uintptr(unsafe.Pointer(window)), uintptr(unsafe.Pointer(renderer))) != 0)
	runtime.KeepAlive(title)
	return
}

func __gowrap__SDL_CreateRenderer(window Window, name string) (__result Renderer) {
	__result = (Renderer)(unsafe.Pointer(__SDL_CreateRenderer(uintptr(unsafe.Pointer(window)), ((*[2]uintptr)(unsafe.Pointer(&name)))[0])))
	runtime.KeepAlive(name)
	return
}

func __gowrap__SDL_CreateRendererWithProperties(props PropertiesID) (__result Renderer) {
	__result = (Renderer)(unsafe.Pointer(__SDL_CreateRendererWithProperties(*(*int32)(unsafe.Pointer(&props)))))
	return
}

func __gowrap__SDL_CreateSoftwareRenderer(surface *Surface) (__result Renderer) {
	__result = (Renderer)(unsafe.Pointer(__SDL_CreateSoftwareRenderer(uintptr(unsafe.Pointer(surface)))))
	return
}

func __gowrap__SDL_GetRenderer(window Window) (__result Renderer) {
	__result = (Renderer)(unsafe.Pointer(__SDL_GetRenderer(uintptr(unsafe.Pointer(window)))))
	return
}

func __gowrap__SDL_GetRenderWindow(renderer Renderer) (__result Window) {
	__result = (Window)(unsafe.Pointer(__SDL_GetRenderWindow(uintptr(unsafe.Pointer(renderer)))))
	return
}

func __gowrap__SDL_GetRendererName(renderer Renderer) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetRendererName(uintptr(unsafe.Pointer(renderer)))).Collect()))
	return
}

func __gowrap__SDL_GetRendererProperties(renderer Renderer) (__result PropertiesID) {
	__result = (PropertiesID)(__SDL_GetRendererProperties(uintptr(unsafe.Pointer(renderer))))
	return
}

func __gowrap__SDL_GetRenderOutputSize(renderer Renderer, w *int32, h *int32) (__result bool) {
	__result = (bool)(__SDL_GetRenderOutputSize(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(w)), uintptr(unsafe.Pointer(h))) != 0)
	return
}

func __gowrap__SDL_GetCurrentRenderOutputSize(renderer Renderer, w *int32, h *int32) (__result bool) {
	__result = (bool)(__SDL_GetCurrentRenderOutputSize(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(w)), uintptr(unsafe.Pointer(h))) != 0)
	return
}

func __gowrap__SDL_CreateTexture(renderer Renderer, format PixelFormat, access TextureAccess, w int32, h int32) (__result *Texture) {
	__result = (*Texture)(unsafe.Pointer(__SDL_CreateTexture(uintptr(unsafe.Pointer(renderer)), *(*int32)(unsafe.Pointer(&format)), *(*int32)(unsafe.Pointer(&access)), *(*int32)(unsafe.Pointer(&w)), *(*int32)(unsafe.Pointer(&h)))))
	return
}

func __gowrap__SDL_CreateTextureFromSurface(renderer Renderer, surface *Surface) (__result *Texture) {
	__result = (*Texture)(unsafe.Pointer(__SDL_CreateTextureFromSurface(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(surface)))))
	return
}

func __gowrap__SDL_CreateTextureWithProperties(renderer Renderer, props PropertiesID) (__result *Texture) {
	__result = (*Texture)(unsafe.Pointer(__SDL_CreateTextureWithProperties(uintptr(unsafe.Pointer(renderer)), *(*int32)(unsafe.Pointer(&props)))))
	return
}

func __gowrap__SDL_GetTextureProperties(texture *Texture) (__result PropertiesID) {
	__result = (PropertiesID)(__SDL_GetTextureProperties(uintptr(unsafe.Pointer(texture))))
	return
}

func __gowrap__SDL_GetRendererFromTexture(texture *Texture) (__result Renderer) {
	__result = (Renderer)(unsafe.Pointer(__SDL_GetRendererFromTexture(uintptr(unsafe.Pointer(texture)))))
	return
}

func __gowrap__SDL_GetTextureSize(texture *Texture, w *float32, h *float32) (__result bool) {
	__result = (bool)(__SDL_GetTextureSize(uintptr(unsafe.Pointer(texture)), uintptr(unsafe.Pointer(w)), uintptr(unsafe.Pointer(h))) != 0)
	return
}

func __gowrap__SDL_SetTextureColorMod(texture *Texture, r uint8, g uint8, b uint8) (__result bool) {
	__result = (bool)(__SDL_SetTextureColorMod(uintptr(unsafe.Pointer(texture)), *(*int32)(unsafe.Pointer(&r)), *(*int32)(unsafe.Pointer(&g)), *(*int32)(unsafe.Pointer(&b))) != 0)
	return
}

func __gowrap__SDL_SetTextureColorModFloat(texture *Texture, r float32, g float32, b float32) (__result bool) {
	__result = (bool)(__SDL_SetTextureColorModFloat(uintptr(unsafe.Pointer(texture)), *(*float32)(unsafe.Pointer(&r)), *(*float32)(unsafe.Pointer(&g)), *(*float32)(unsafe.Pointer(&b))) != 0)
	return
}

func __gowrap__SDL_GetTextureColorMod(texture *Texture, r *uint8, g *uint8, b *uint8) (__result bool) {
	__result = (bool)(__SDL_GetTextureColorMod(uintptr(unsafe.Pointer(texture)), uintptr(unsafe.Pointer(r)), uintptr(unsafe.Pointer(g)), uintptr(unsafe.Pointer(b))) != 0)
	return
}

func __gowrap__SDL_GetTextureColorModFloat(texture *Texture, r *float32, g *float32, b *float32) (__result bool) {
	__result = (bool)(__SDL_GetTextureColorModFloat(uintptr(unsafe.Pointer(texture)), uintptr(unsafe.Pointer(r)), uintptr(unsafe.Pointer(g)), uintptr(unsafe.Pointer(b))) != 0)
	return
}

func __gowrap__SDL_SetTextureAlphaMod(texture *Texture, alpha uint8) (__result bool) {
	__result = (bool)(__SDL_SetTextureAlphaMod(uintptr(unsafe.Pointer(texture)), *(*int32)(unsafe.Pointer(&alpha))) != 0)
	return
}

func __gowrap__SDL_SetTextureAlphaModFloat(texture *Texture, alpha float32) (__result bool) {
	__result = (bool)(__SDL_SetTextureAlphaModFloat(uintptr(unsafe.Pointer(texture)), *(*float32)(unsafe.Pointer(&alpha))) != 0)
	return
}

func __gowrap__SDL_GetTextureAlphaMod(texture *Texture, alpha *uint8) (__result bool) {
	__result = (bool)(__SDL_GetTextureAlphaMod(uintptr(unsafe.Pointer(texture)), uintptr(unsafe.Pointer(alpha))) != 0)
	return
}

func __gowrap__SDL_GetTextureAlphaModFloat(texture *Texture, alpha *float32) (__result bool) {
	__result = (bool)(__SDL_GetTextureAlphaModFloat(uintptr(unsafe.Pointer(texture)), uintptr(unsafe.Pointer(alpha))) != 0)
	return
}

func __gowrap__SDL_SetTextureBlendMode(texture *Texture, blendMode BlendMode) (__result bool) {
	__result = (bool)(__SDL_SetTextureBlendMode(uintptr(unsafe.Pointer(texture)), *(*int32)(unsafe.Pointer(&blendMode))) != 0)
	return
}

func __gowrap__SDL_GetTextureBlendMode(texture *Texture, blendMode *BlendMode) (__result bool) {
	__result = (bool)(__SDL_GetTextureBlendMode(uintptr(unsafe.Pointer(texture)), uintptr(unsafe.Pointer(blendMode))) != 0)
	return
}

func __gowrap__SDL_SetTextureScaleMode(texture *Texture, scaleMode ScaleMode) (__result bool) {
	__result = (bool)(__SDL_SetTextureScaleMode(uintptr(unsafe.Pointer(texture)), *(*int32)(unsafe.Pointer(&scaleMode))) != 0)
	return
}

func __gowrap__SDL_GetTextureScaleMode(texture *Texture, scaleMode *ScaleMode) (__result bool) {
	__result = (bool)(__SDL_GetTextureScaleMode(uintptr(unsafe.Pointer(texture)), uintptr(unsafe.Pointer(scaleMode))) != 0)
	return
}

func __gowrap__SDL_UpdateTexture(texture *Texture, rect *Rect, pixels []byte, pitch int32) (__result bool) {
	__result = (bool)(__SDL_UpdateTexture(uintptr(unsafe.Pointer(texture)), uintptr(unsafe.Pointer(rect)), uintptr(unsafe.Pointer(&pixels[0])), *(*int32)(unsafe.Pointer(&pitch))) != 0)
	runtime.KeepAlive(pixels)
	return
}

func __gowrap__SDL_UpdateYUVTexture(texture *Texture, rect *Rect, Yplane *uint8, Ypitch int32, Uplane *uint8, Upitch int32, Vplane *uint8, Vpitch int32) (__result bool) {
	__result = (bool)(__SDL_UpdateYUVTexture(uintptr(unsafe.Pointer(texture)), uintptr(unsafe.Pointer(rect)), uintptr(unsafe.Pointer(Yplane)), *(*int32)(unsafe.Pointer(&Ypitch)), uintptr(unsafe.Pointer(Uplane)), *(*int32)(unsafe.Pointer(&Upitch)), uintptr(unsafe.Pointer(Vplane)), *(*int32)(unsafe.Pointer(&Vpitch))) != 0)
	return
}

func __gowrap__SDL_UpdateNVTexture(texture *Texture, rect *Rect, Yplane *uint8, Ypitch int32, UVplane *uint8, UVpitch int32) (__result bool) {
	__result = (bool)(__SDL_UpdateNVTexture(uintptr(unsafe.Pointer(texture)), uintptr(unsafe.Pointer(rect)), uintptr(unsafe.Pointer(Yplane)), *(*int32)(unsafe.Pointer(&Ypitch)), uintptr(unsafe.Pointer(UVplane)), *(*int32)(unsafe.Pointer(&UVpitch))) != 0)
	return
}

func __gowrap__SDL_LockTexture(texture *Texture, rect *Rect, pixels *uintptr, pitch *int32) (__result bool) {
	__result = (bool)(__SDL_LockTexture(uintptr(unsafe.Pointer(texture)), uintptr(unsafe.Pointer(rect)), uintptr(unsafe.Pointer(pixels)), uintptr(unsafe.Pointer(pitch))) != 0)
	return
}

func __gowrap__SDL_LockTextureToSurface(texture *Texture, rect *Rect, surface **Surface) (__result bool) {
	__result = (bool)(__SDL_LockTextureToSurface(uintptr(unsafe.Pointer(texture)), uintptr(unsafe.Pointer(rect)), uintptr(unsafe.Pointer(surface))) != 0)
	return
}

func __gowrap__SDL_UnlockTexture(texture *Texture) {
	__SDL_UnlockTexture(uintptr(unsafe.Pointer(texture)))
}

func __gowrap__SDL_SetRenderTarget(renderer Renderer, texture *Texture) (__result bool) {
	__result = (bool)(__SDL_SetRenderTarget(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(texture))) != 0)
	return
}

func __gowrap__SDL_GetRenderTarget(renderer Renderer) (__result *Texture) {
	__result = (*Texture)(unsafe.Pointer(__SDL_GetRenderTarget(uintptr(unsafe.Pointer(renderer)))))
	return
}

func __gowrap__SDL_SetRenderLogicalPresentation(renderer Renderer, w int32, h int32, mode RendererLogicalPresentation) (__result bool) {
	__result = (bool)(__SDL_SetRenderLogicalPresentation(uintptr(unsafe.Pointer(renderer)), *(*int32)(unsafe.Pointer(&w)), *(*int32)(unsafe.Pointer(&h)), *(*int32)(unsafe.Pointer(&mode))) != 0)
	return
}

func __gowrap__SDL_GetRenderLogicalPresentation(renderer Renderer, w *int32, h *int32, mode RendererLogicalPresentation) (__result bool) {
	__result = (bool)(__SDL_GetRenderLogicalPresentation(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(w)), uintptr(unsafe.Pointer(h)), *(*int32)(unsafe.Pointer(&mode))) != 0)
	return
}

func __gowrap__SDL_GetRenderLogicalPresentationRect(renderer Renderer, rect *FRect) (__result bool) {
	__result = (bool)(__SDL_GetRenderLogicalPresentationRect(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(rect))) != 0)
	return
}

func __gowrap__SDL_RenderCoordinatesFromWindow(renderer Renderer, window_x float32, window_y float32, x *float32, y *float32) (__result bool) {
	__result = (bool)(__SDL_RenderCoordinatesFromWindow(uintptr(unsafe.Pointer(renderer)), *(*float32)(unsafe.Pointer(&window_x)), *(*float32)(unsafe.Pointer(&window_y)), uintptr(unsafe.Pointer(x)), uintptr(unsafe.Pointer(y))) != 0)
	return
}

func __gowrap__SDL_RenderCoordinatesToWindow(renderer Renderer, x float32, y float32, window_x *float32, window_y *float32) (__result bool) {
	__result = (bool)(__SDL_RenderCoordinatesToWindow(uintptr(unsafe.Pointer(renderer)), *(*float32)(unsafe.Pointer(&x)), *(*float32)(unsafe.Pointer(&y)), uintptr(unsafe.Pointer(window_x)), uintptr(unsafe.Pointer(window_y))) != 0)
	return
}

func __gowrap__SDL_ConvertEventToRenderCoordinates(renderer Renderer, event *Event) (__result bool) {
	__result = (bool)(__SDL_ConvertEventToRenderCoordinates(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(event))) != 0)
	return
}

func __gowrap__SDL_SetRenderViewport(renderer Renderer, rect *Rect) (__result bool) {
	__result = (bool)(__SDL_SetRenderViewport(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(rect))) != 0)
	return
}

func __gowrap__SDL_GetRenderViewport(renderer Renderer, rect *Rect) (__result bool) {
	__result = (bool)(__SDL_GetRenderViewport(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(rect))) != 0)
	return
}

func __gowrap__SDL_RenderViewportSet(renderer Renderer) (__result bool) {
	__result = (bool)(__SDL_RenderViewportSet(uintptr(unsafe.Pointer(renderer))) != 0)
	return
}

func __gowrap__SDL_GetRenderSafeArea(renderer Renderer, rect *Rect) (__result bool) {
	__result = (bool)(__SDL_GetRenderSafeArea(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(rect))) != 0)
	return
}

func __gowrap__SDL_SetRenderClipRect(renderer Renderer, rect *Rect) (__result bool) {
	__result = (bool)(__SDL_SetRenderClipRect(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(rect))) != 0)
	return
}

func __gowrap__SDL_GetRenderClipRect(renderer Renderer, rect *Rect) (__result bool) {
	__result = (bool)(__SDL_GetRenderClipRect(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(rect))) != 0)
	return
}

func __gowrap__SDL_RenderClipEnabled(renderer Renderer) (__result bool) {
	__result = (bool)(__SDL_RenderClipEnabled(uintptr(unsafe.Pointer(renderer))) != 0)
	return
}

func __gowrap__SDL_SetRenderScale(renderer Renderer, scaleX float32, scaleY float32) (__result bool) {
	__result = (bool)(__SDL_SetRenderScale(uintptr(unsafe.Pointer(renderer)), *(*float32)(unsafe.Pointer(&scaleX)), *(*float32)(unsafe.Pointer(&scaleY))) != 0)
	return
}

func __gowrap__SDL_GetRenderScale(renderer Renderer, scaleX *float32, scaleY *float32) (__result bool) {
	__result = (bool)(__SDL_GetRenderScale(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(scaleX)), uintptr(unsafe.Pointer(scaleY))) != 0)
	return
}

func __gowrap__SDL_SetRenderDrawColor(renderer Renderer, r uint8, g uint8, b uint8, a uint8) (__result bool) {
	__result = (bool)(__SDL_SetRenderDrawColor(uintptr(unsafe.Pointer(renderer)), *(*int32)(unsafe.Pointer(&r)), *(*int32)(unsafe.Pointer(&g)), *(*int32)(unsafe.Pointer(&b)), *(*int32)(unsafe.Pointer(&a))) != 0)
	return
}

func __gowrap__SDL_SetRenderDrawColorFloat(renderer Renderer, r float32, g float32, b float32, a float32) (__result bool) {
	__result = (bool)(__SDL_SetRenderDrawColorFloat(uintptr(unsafe.Pointer(renderer)), *(*float32)(unsafe.Pointer(&r)), *(*float32)(unsafe.Pointer(&g)), *(*float32)(unsafe.Pointer(&b)), *(*float32)(unsafe.Pointer(&a))) != 0)
	return
}

func __gowrap__SDL_GetRenderDrawColor(renderer Renderer, r *uint8, g *uint8, b *uint8, a *uint8) (__result bool) {
	__result = (bool)(__SDL_GetRenderDrawColor(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(r)), uintptr(unsafe.Pointer(g)), uintptr(unsafe.Pointer(b)), uintptr(unsafe.Pointer(a))) != 0)
	return
}

func __gowrap__SDL_GetRenderDrawColorFloat(renderer Renderer, r *float32, g *float32, b *float32, a *float32) (__result bool) {
	__result = (bool)(__SDL_GetRenderDrawColorFloat(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(r)), uintptr(unsafe.Pointer(g)), uintptr(unsafe.Pointer(b)), uintptr(unsafe.Pointer(a))) != 0)
	return
}

func __gowrap__SDL_SetRenderColorScale(renderer Renderer, scale float32) (__result bool) {
	__result = (bool)(__SDL_SetRenderColorScale(uintptr(unsafe.Pointer(renderer)), *(*float32)(unsafe.Pointer(&scale))) != 0)
	return
}

func __gowrap__SDL_GetRenderColorScale(renderer Renderer, scale *float32) (__result bool) {
	__result = (bool)(__SDL_GetRenderColorScale(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(scale))) != 0)
	return
}

func __gowrap__SDL_SetRenderDrawBlendMode(renderer Renderer, blendMode BlendMode) (__result bool) {
	__result = (bool)(__SDL_SetRenderDrawBlendMode(uintptr(unsafe.Pointer(renderer)), *(*int32)(unsafe.Pointer(&blendMode))) != 0)
	return
}

func __gowrap__SDL_GetRenderDrawBlendMode(renderer Renderer, blendMode *BlendMode) (__result bool) {
	__result = (bool)(__SDL_GetRenderDrawBlendMode(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(blendMode))) != 0)
	return
}

func __gowrap__SDL_RenderClear(renderer Renderer) (__result bool) {
	__result = (bool)(__SDL_RenderClear(uintptr(unsafe.Pointer(renderer))) != 0)
	return
}

func __gowrap__SDL_RenderPoint(renderer Renderer, x float32, y float32) (__result bool) {
	__result = (bool)(__SDL_RenderPoint(uintptr(unsafe.Pointer(renderer)), *(*float32)(unsafe.Pointer(&x)), *(*float32)(unsafe.Pointer(&y))) != 0)
	return
}

func __gowrap__SDL_RenderPoints(renderer Renderer, points []FPoint, count int32) (__result bool) {
	__result = (bool)(__SDL_RenderPoints(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(&points[0])), *(*int32)(unsafe.Pointer(&count))) != 0)
	runtime.KeepAlive(points)
	return
}

func __gowrap__SDL_RenderLine(renderer Renderer, x1 float32, y1 float32, x2 float32, y2 float32) (__result bool) {
	__result = (bool)(__SDL_RenderLine(uintptr(unsafe.Pointer(renderer)), *(*float32)(unsafe.Pointer(&x1)), *(*float32)(unsafe.Pointer(&y1)), *(*float32)(unsafe.Pointer(&x2)), *(*float32)(unsafe.Pointer(&y2))) != 0)
	return
}

func __gowrap__SDL_RenderLines(renderer Renderer, points []FPoint, count int32) (__result bool) {
	__result = (bool)(__SDL_RenderLines(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(&points[0])), *(*int32)(unsafe.Pointer(&count))) != 0)
	runtime.KeepAlive(points)
	return
}

func __gowrap__SDL_RenderRect(renderer Renderer, rect *FRect) (__result bool) {
	__result = (bool)(__SDL_RenderRect(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(rect))) != 0)
	return
}

func __gowrap__SDL_RenderRects(renderer Renderer, rects []FRect, count int32) (__result bool) {
	__result = (bool)(__SDL_RenderRects(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(&rects[0])), *(*int32)(unsafe.Pointer(&count))) != 0)
	runtime.KeepAlive(rects)
	return
}

func __gowrap__SDL_RenderFillRect(renderer Renderer, rect *FRect) (__result bool) {
	__result = (bool)(__SDL_RenderFillRect(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(rect))) != 0)
	return
}

func __gowrap__SDL_RenderFillRects(renderer Renderer, rects []FRect, count int32) (__result bool) {
	__result = (bool)(__SDL_RenderFillRects(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(&rects[0])), *(*int32)(unsafe.Pointer(&count))) != 0)
	runtime.KeepAlive(rects)
	return
}

func __gowrap__SDL_RenderTexture(renderer Renderer, texture *Texture, srcrect *FRect, dstrect *FRect) (__result bool) {
	__result = (bool)(__SDL_RenderTexture(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(texture)), uintptr(unsafe.Pointer(srcrect)), uintptr(unsafe.Pointer(dstrect))) != 0)
	return
}

func __gowrap__SDL_RenderTextureRotated(renderer Renderer, texture *Texture, srcrect *FRect, dstrect *FRect, angle float64, center *FPoint, flip FlipMode) (__result bool) {
	__result = (bool)(__SDL_RenderTextureRotated(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(texture)), uintptr(unsafe.Pointer(srcrect)), uintptr(unsafe.Pointer(dstrect)), *(*float64)(unsafe.Pointer(&angle)), uintptr(unsafe.Pointer(center)), *(*int32)(unsafe.Pointer(&flip))) != 0)
	return
}

func __gowrap__SDL_RenderTextureAffine(renderer Renderer, texture *Texture, srcrect *FRect, origin *FPoint, right *FPoint, down *FPoint) (__result bool) {
	__result = (bool)(__SDL_RenderTextureAffine(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(texture)), uintptr(unsafe.Pointer(srcrect)), uintptr(unsafe.Pointer(origin)), uintptr(unsafe.Pointer(right)), uintptr(unsafe.Pointer(down))) != 0)
	return
}

func __gowrap__SDL_RenderTextureTiled(renderer Renderer, texture *Texture, srcrect *FRect, scale float32, dstrect *FRect) (__result bool) {
	__result = (bool)(__SDL_RenderTextureTiled(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(texture)), uintptr(unsafe.Pointer(srcrect)), *(*float32)(unsafe.Pointer(&scale)), uintptr(unsafe.Pointer(dstrect))) != 0)
	return
}

func __gowrap__SDL_RenderTexture9Grid(renderer Renderer, texture *Texture, srcrect *FRect, left_width float32, right_width float32, top_height float32, bottom_height float32, scale float32, dstrect *FRect) (__result bool) {
	__result = (bool)(__SDL_RenderTexture9Grid(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(texture)), uintptr(unsafe.Pointer(srcrect)), *(*float32)(unsafe.Pointer(&left_width)), *(*float32)(unsafe.Pointer(&right_width)), *(*float32)(unsafe.Pointer(&top_height)), *(*float32)(unsafe.Pointer(&bottom_height)), *(*float32)(unsafe.Pointer(&scale)), uintptr(unsafe.Pointer(dstrect))) != 0)
	return
}

func __gowrap__SDL_RenderTexture9GridTiled(renderer Renderer, texture *Texture, srcrect *FRect, left_width float32, right_width float32, top_height float32, bottom_height float32, scale float32, dstrect *FRect, tileScale float32) (__result bool) {
	__result = (bool)(__SDL_RenderTexture9GridTiled(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(texture)), uintptr(unsafe.Pointer(srcrect)), *(*float32)(unsafe.Pointer(&left_width)), *(*float32)(unsafe.Pointer(&right_width)), *(*float32)(unsafe.Pointer(&top_height)), *(*float32)(unsafe.Pointer(&bottom_height)), *(*float32)(unsafe.Pointer(&scale)), uintptr(unsafe.Pointer(dstrect)), *(*float32)(unsafe.Pointer(&tileScale))) != 0)
	return
}

func __gowrap__SDL_RenderGeometry(renderer Renderer, texture *Texture, vertices []Vertex, num_vertices int32, indices []int32, num_indices int32) (__result bool) {
	__result = (bool)(__SDL_RenderGeometry(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(texture)), uintptr(unsafe.Pointer(&vertices[0])), *(*int32)(unsafe.Pointer(&num_vertices)), uintptr(unsafe.Pointer(&indices[0])), *(*int32)(unsafe.Pointer(&num_indices))) != 0)
	runtime.KeepAlive(vertices)
	runtime.KeepAlive(indices)
	return
}

func __gowrap__SDL_RenderGeometryRaw(renderer Renderer, texture *Texture, xy []float32, xy_stride int32, color *FColor, color_stride int32, uv []float32, uv_stride int32, num_vertices int32, indices uintptr, num_indices int32, size_indices int32) (__result bool) {
	__result = (bool)(__SDL_RenderGeometryRaw(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(texture)), uintptr(unsafe.Pointer(&xy[0])), *(*int32)(unsafe.Pointer(&xy_stride)), uintptr(unsafe.Pointer(color)), *(*int32)(unsafe.Pointer(&color_stride)), uintptr(unsafe.Pointer(&uv[0])), *(*int32)(unsafe.Pointer(&uv_stride)), *(*int32)(unsafe.Pointer(&num_vertices)), uintptr(unsafe.Pointer(indices)), *(*int32)(unsafe.Pointer(&num_indices)), *(*int32)(unsafe.Pointer(&size_indices))) != 0)
	runtime.KeepAlive(xy)
	runtime.KeepAlive(uv)
	return
}

func __gowrap__SDL_SetRenderTextureAddressMode(renderer Renderer, u_mode TextureAddressMode, v_mode TextureAddressMode) (__result bool) {
	__result = (bool)(__SDL_SetRenderTextureAddressMode(uintptr(unsafe.Pointer(renderer)), *(*int32)(unsafe.Pointer(&u_mode)), *(*int32)(unsafe.Pointer(&v_mode))) != 0)
	return
}

func __gowrap__SDL_GetRenderTextureAddressMode(renderer Renderer, u_mode *TextureAddressMode, v_mode *TextureAddressMode) (__result bool) {
	__result = (bool)(__SDL_GetRenderTextureAddressMode(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(u_mode)), uintptr(unsafe.Pointer(v_mode))) != 0)
	return
}

func __gowrap__SDL_RenderReadPixels(renderer Renderer, rect *Rect) (__result *Surface) {
	__result = (*Surface)(unsafe.Pointer(__SDL_RenderReadPixels(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(rect)))))
	return
}

func __gowrap__SDL_RenderPresent(renderer Renderer) (__result bool) {
	__result = (bool)(__SDL_RenderPresent(uintptr(unsafe.Pointer(renderer))) != 0)
	return
}

func __gowrap__SDL_DestroyTexture(texture *Texture) {
	__SDL_DestroyTexture(uintptr(unsafe.Pointer(texture)))
}

func __gowrap__SDL_DestroyRenderer(renderer Renderer) {
	__SDL_DestroyRenderer(uintptr(unsafe.Pointer(renderer)))
}

func __gowrap__SDL_FlushRenderer(renderer Renderer) (__result bool) {
	__result = (bool)(__SDL_FlushRenderer(uintptr(unsafe.Pointer(renderer))) != 0)
	return
}

func __gowrap__SDL_GetRenderMetalLayer(renderer Renderer) (__result uintptr) {
	__result = (uintptr)(unsafe.Pointer(__SDL_GetRenderMetalLayer(uintptr(unsafe.Pointer(renderer)))))
	return
}

func __gowrap__SDL_GetRenderMetalCommandEncoder(renderer Renderer) (__result uintptr) {
	__result = (uintptr)(unsafe.Pointer(__SDL_GetRenderMetalCommandEncoder(uintptr(unsafe.Pointer(renderer)))))
	return
}

func __gowrap__SDL_AddVulkanRenderSemaphores(renderer Renderer, wait_stage_mask uint32, wait_semaphore int64, signal_semaphore int64) (__result bool) {
	__result = (bool)(__SDL_AddVulkanRenderSemaphores(uintptr(unsafe.Pointer(renderer)), *(*int32)(unsafe.Pointer(&wait_stage_mask)), *(*int64)(unsafe.Pointer(&wait_semaphore)), *(*int64)(unsafe.Pointer(&signal_semaphore))) != 0)
	return
}

func __gowrap__SDL_SetRenderVSync(renderer Renderer, vsync VSync) (__result bool) {
	__result = (bool)(__SDL_SetRenderVSync(uintptr(unsafe.Pointer(renderer)), *(*int32)(unsafe.Pointer(&vsync))) != 0)
	return
}

func __gowrap__SDL_GetRenderVSync(renderer Renderer, vsync *int32) (__result bool) {
	__result = (bool)(__SDL_GetRenderVSync(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(vsync))) != 0)
	return
}

func __gowrap__SDL_RenderDebugText(renderer Renderer, x float32, y float32, str string) (__result bool) {
	__result = (bool)(__SDL_RenderDebugText(uintptr(unsafe.Pointer(renderer)), *(*float32)(unsafe.Pointer(&x)), *(*float32)(unsafe.Pointer(&y)), ((*[2]uintptr)(unsafe.Pointer(&str)))[0]) != 0)
	runtime.KeepAlive(str)
	return
}

func __gowrap__SDL_SetDefaultTextureScaleMode(renderer Renderer, scale_mode ScaleMode) (__result bool) {
	__result = (bool)(__SDL_SetDefaultTextureScaleMode(uintptr(unsafe.Pointer(renderer)), *(*int32)(unsafe.Pointer(&scale_mode))) != 0)
	return
}

func __gowrap__SDL_GetDefaultTextureScaleMode(renderer Renderer, scale_mode *ScaleMode) (__result bool) {
	__result = (bool)(__SDL_GetDefaultTextureScaleMode(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(scale_mode))) != 0)
	return
}

func __gowrap__SDL_CreateGPURenderState(renderer Renderer, desc *GPURenderStateDesc) (__result *GPURenderState) {
	__result = (*GPURenderState)(unsafe.Pointer(__SDL_CreateGPURenderState(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(desc)))))
	return
}

func __gowrap__SDL_SetGPURenderStateFragmentUniforms(state *GPURenderState, slot_index int32, data []byte, length int32) (__result bool) {
	__result = (bool)(__SDL_SetGPURenderStateFragmentUniforms(uintptr(unsafe.Pointer(state)), *(*int32)(unsafe.Pointer(&slot_index)), uintptr(unsafe.Pointer(&data[0])), *(*int32)(unsafe.Pointer(&length))) != 0)
	runtime.KeepAlive(data)
	return
}

func __gowrap__SDL_SetRenderGPUState(renderer Renderer, state *GPURenderState) (__result bool) {
	__result = (bool)(__SDL_SetRenderGPUState(uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(state))) != 0)
	return
}

func __gowrap__SDL_DestroyGPURenderState(state *GPURenderState) {
	__SDL_DestroyGPURenderState(uintptr(unsafe.Pointer(state)))
}

 
func init() {
	GetNumRenderDrivers = __gowrap__SDL_GetNumRenderDrivers
	GetRenderDriver = __gowrap__SDL_GetRenderDriver
	CreateWindowAndRenderer = __gowrap__SDL_CreateWindowAndRenderer
	CreateRenderer = __gowrap__SDL_CreateRenderer
	CreateRendererWithProperties = __gowrap__SDL_CreateRendererWithProperties
	CreateSoftwareRenderer = __gowrap__SDL_CreateSoftwareRenderer
	GetRenderer = __gowrap__SDL_GetRenderer
	GetRenderWindow = __gowrap__SDL_GetRenderWindow
	GetRendererName = __gowrap__SDL_GetRendererName
	GetRendererProperties = __gowrap__SDL_GetRendererProperties
	GetRenderOutputSize = __gowrap__SDL_GetRenderOutputSize
	GetCurrentRenderOutputSize = __gowrap__SDL_GetCurrentRenderOutputSize
	CreateTexture = __gowrap__SDL_CreateTexture
	CreateTextureFromSurface = __gowrap__SDL_CreateTextureFromSurface
	CreateTextureWithProperties = __gowrap__SDL_CreateTextureWithProperties
	GetTextureProperties = __gowrap__SDL_GetTextureProperties
	GetRendererFromTexture = __gowrap__SDL_GetRendererFromTexture
	GetTextureSize = __gowrap__SDL_GetTextureSize
	SetTextureColorMod = __gowrap__SDL_SetTextureColorMod
	SetTextureColorModFloat = __gowrap__SDL_SetTextureColorModFloat
	GetTextureColorMod = __gowrap__SDL_GetTextureColorMod
	GetTextureColorModFloat = __gowrap__SDL_GetTextureColorModFloat
	SetTextureAlphaMod = __gowrap__SDL_SetTextureAlphaMod
	SetTextureAlphaModFloat = __gowrap__SDL_SetTextureAlphaModFloat
	GetTextureAlphaMod = __gowrap__SDL_GetTextureAlphaMod
	GetTextureAlphaModFloat = __gowrap__SDL_GetTextureAlphaModFloat
	SetTextureBlendMode = __gowrap__SDL_SetTextureBlendMode
	GetTextureBlendMode = __gowrap__SDL_GetTextureBlendMode
	SetTextureScaleMode = __gowrap__SDL_SetTextureScaleMode
	GetTextureScaleMode = __gowrap__SDL_GetTextureScaleMode
	UpdateTexture = __gowrap__SDL_UpdateTexture
	UpdateYUVTexture = __gowrap__SDL_UpdateYUVTexture
	UpdateNVTexture = __gowrap__SDL_UpdateNVTexture
	LockTexture = __gowrap__SDL_LockTexture
	LockTextureToSurface = __gowrap__SDL_LockTextureToSurface
	UnlockTexture = __gowrap__SDL_UnlockTexture
	SetRenderTarget = __gowrap__SDL_SetRenderTarget
	GetRenderTarget = __gowrap__SDL_GetRenderTarget
	SetRenderLogicalPresentation = __gowrap__SDL_SetRenderLogicalPresentation
	GetRenderLogicalPresentation = __gowrap__SDL_GetRenderLogicalPresentation
	GetRenderLogicalPresentationRect = __gowrap__SDL_GetRenderLogicalPresentationRect
	RenderCoordinatesFromWindow = __gowrap__SDL_RenderCoordinatesFromWindow
	RenderCoordinatesToWindow = __gowrap__SDL_RenderCoordinatesToWindow
	ConvertEventToRenderCoordinates = __gowrap__SDL_ConvertEventToRenderCoordinates
	SetRenderViewport = __gowrap__SDL_SetRenderViewport
	GetRenderViewport = __gowrap__SDL_GetRenderViewport
	RenderViewportSet = __gowrap__SDL_RenderViewportSet
	GetRenderSafeArea = __gowrap__SDL_GetRenderSafeArea
	SetRenderClipRect = __gowrap__SDL_SetRenderClipRect
	GetRenderClipRect = __gowrap__SDL_GetRenderClipRect
	RenderClipEnabled = __gowrap__SDL_RenderClipEnabled
	SetRenderScale = __gowrap__SDL_SetRenderScale
	GetRenderScale = __gowrap__SDL_GetRenderScale
	SetRenderDrawColor = __gowrap__SDL_SetRenderDrawColor
	SetRenderDrawColorFloat = __gowrap__SDL_SetRenderDrawColorFloat
	GetRenderDrawColor = __gowrap__SDL_GetRenderDrawColor
	GetRenderDrawColorFloat = __gowrap__SDL_GetRenderDrawColorFloat
	SetRenderColorScale = __gowrap__SDL_SetRenderColorScale
	GetRenderColorScale = __gowrap__SDL_GetRenderColorScale
	SetRenderDrawBlendMode = __gowrap__SDL_SetRenderDrawBlendMode
	GetRenderDrawBlendMode = __gowrap__SDL_GetRenderDrawBlendMode
	RenderClear = __gowrap__SDL_RenderClear
	RenderPoint = __gowrap__SDL_RenderPoint
	RenderPoints = __gowrap__SDL_RenderPoints
	RenderLine = __gowrap__SDL_RenderLine
	RenderLines = __gowrap__SDL_RenderLines
	RenderRect = __gowrap__SDL_RenderRect
	RenderRects = __gowrap__SDL_RenderRects
	RenderFillRect = __gowrap__SDL_RenderFillRect
	RenderFillRects = __gowrap__SDL_RenderFillRects
	RenderTexture = __gowrap__SDL_RenderTexture
	RenderTextureRotated = __gowrap__SDL_RenderTextureRotated
	RenderTextureAffine = __gowrap__SDL_RenderTextureAffine
	RenderTextureTiled = __gowrap__SDL_RenderTextureTiled
	RenderTexture9Grid = __gowrap__SDL_RenderTexture9Grid
	RenderTexture9GridTiled = __gowrap__SDL_RenderTexture9GridTiled
	RenderGeometry = __gowrap__SDL_RenderGeometry
	RenderGeometryRaw = __gowrap__SDL_RenderGeometryRaw
	SetRenderTextureAddressMode = __gowrap__SDL_SetRenderTextureAddressMode
	GetRenderTextureAddressMode = __gowrap__SDL_GetRenderTextureAddressMode
	RenderReadPixels = __gowrap__SDL_RenderReadPixels
	RenderPresent = __gowrap__SDL_RenderPresent
	DestroyTexture = __gowrap__SDL_DestroyTexture
	DestroyRenderer = __gowrap__SDL_DestroyRenderer
	FlushRenderer = __gowrap__SDL_FlushRenderer
	GetRenderMetalLayer = __gowrap__SDL_GetRenderMetalLayer
	GetRenderMetalCommandEncoder = __gowrap__SDL_GetRenderMetalCommandEncoder
	AddVulkanRenderSemaphores = __gowrap__SDL_AddVulkanRenderSemaphores
	SetRenderVSync = __gowrap__SDL_SetRenderVSync
	GetRenderVSync = __gowrap__SDL_GetRenderVSync
	RenderDebugText = __gowrap__SDL_RenderDebugText
	SetDefaultTextureScaleMode = __gowrap__SDL_SetDefaultTextureScaleMode
	GetDefaultTextureScaleMode = __gowrap__SDL_GetDefaultTextureScaleMode
	CreateGPURenderState = __gowrap__SDL_CreateGPURenderState
	SetGPURenderStateFragmentUniforms = __gowrap__SDL_SetGPURenderStateFragmentUniforms
	SetRenderGPUState = __gowrap__SDL_SetRenderGPUState
	DestroyGPURenderState = __gowrap__SDL_DestroyGPURenderState
}
