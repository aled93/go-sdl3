//go:build wasm

package sdl

import (
  "unsafe"
)



//go:wasmimport sdl3 SDL_GetNumVideoDrivers
func __SDL_GetNumVideoDrivers() int32

//go:wasmimport sdl3 SDL_GetVideoDriver
func __SDL_GetVideoDriver(int32) uintptr

//go:wasmimport sdl3 SDL_GetCurrentVideoDriver
func __SDL_GetCurrentVideoDriver() uintptr

//go:wasmimport sdl3 SDL_GetSystemTheme
func __SDL_GetSystemTheme() int32

//go:wasmimport sdl3 SDL_GetDisplays
func __SDL_GetDisplays(uintptr) int32

//go:wasmimport sdl3 SDL_GetPrimaryDisplay
func __SDL_GetPrimaryDisplay() int32

//go:wasmimport sdl3 SDL_GetDisplayProperties
func __SDL_GetDisplayProperties(int32) int32

//go:wasmimport sdl3 SDL_GetDisplayName
func __SDL_GetDisplayName(int32) uintptr

//go:wasmimport sdl3 SDL_GetDisplayBounds
func __SDL_GetDisplayBounds(int32, uintptr) int32

//go:wasmimport sdl3 SDL_GetDisplayUsableBounds
func __SDL_GetDisplayUsableBounds(int32, uintptr) int32

//go:wasmimport sdl3 SDL_GetNaturalDisplayOrientation
func __SDL_GetNaturalDisplayOrientation(int32) int32

//go:wasmimport sdl3 SDL_GetCurrentDisplayOrientation
func __SDL_GetCurrentDisplayOrientation(int32) int32

//go:wasmimport sdl3 SDL_GetDisplayContentScale
func __SDL_GetDisplayContentScale(int32) float32

//go:wasmimport sdl3 SDL_GetFullscreenDisplayModes
func __SDL_GetFullscreenDisplayModes(int32, uintptr) uintptr

//go:wasmimport sdl3 SDL_GetClosestFullscreenDisplayMode
func __SDL_GetClosestFullscreenDisplayMode(int32, int32, int32, float32, int32, uintptr) int32

//go:wasmimport sdl3 SDL_GetDesktopDisplayMode
func __SDL_GetDesktopDisplayMode(int32) uintptr

//go:wasmimport sdl3 SDL_GetCurrentDisplayMode
func __SDL_GetCurrentDisplayMode(int32) uintptr

//go:wasmimport sdl3 SDL_GetDisplayForPoint
func __SDL_GetDisplayForPoint(uintptr) int32

//go:wasmimport sdl3 SDL_GetDisplayForRect
func __SDL_GetDisplayForRect(uintptr) int32

//go:wasmimport sdl3 SDL_GetDisplayForWindow
func __SDL_GetDisplayForWindow(uintptr) int32

//go:wasmimport sdl3 SDL_GetWindowPixelDensity
func __SDL_GetWindowPixelDensity(uintptr) float32

//go:wasmimport sdl3 SDL_GetWindowDisplayScale
func __SDL_GetWindowDisplayScale(uintptr) float32

//go:wasmimport sdl3 SDL_SetWindowFullscreenMode
func __SDL_SetWindowFullscreenMode(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetWindowFullscreenMode
func __SDL_GetWindowFullscreenMode(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetWindowICCProfile
func __SDL_GetWindowICCProfile(uintptr, uintptr) uintptr

//go:wasmimport sdl3 SDL_GetWindowPixelFormat
func __SDL_GetWindowPixelFormat(uintptr) int32

//go:wasmimport sdl3 SDL_GetWindows
func __SDL_GetWindows(uintptr) uintptr

//go:wasmimport sdl3 SDL_CreateWindow
func __SDL_CreateWindow(uintptr, int32, int32, int64) uintptr

//go:wasmimport sdl3 SDL_CreatePopupWindow
func __SDL_CreatePopupWindow(uintptr, int32, int32, int32, int32, int64) uintptr

//go:wasmimport sdl3 SDL_CreateWindowWithProperties
func __SDL_CreateWindowWithProperties(int32) uintptr

//go:wasmimport sdl3 SDL_GetWindowID
func __SDL_GetWindowID(uintptr) int32

//go:wasmimport sdl3 SDL_GetWindowFromID
func __SDL_GetWindowFromID(int32) uintptr

//go:wasmimport sdl3 SDL_GetWindowParent
func __SDL_GetWindowParent(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetWindowProperties
func __SDL_GetWindowProperties(uintptr) int32

//go:wasmimport sdl3 SDL_GetWindowFlags
func __SDL_GetWindowFlags(uintptr) int64

//go:wasmimport sdl3 SDL_SetWindowTitle
func __SDL_SetWindowTitle(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetWindowTitle
func __SDL_GetWindowTitle(uintptr) uintptr

//go:wasmimport sdl3 SDL_SetWindowIcon
func __SDL_SetWindowIcon(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_SetWindowPosition
func __SDL_SetWindowPosition(uintptr, int32, int32) int32

//go:wasmimport sdl3 SDL_GetWindowPosition
func __SDL_GetWindowPosition(uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_SetWindowSize
func __SDL_SetWindowSize(uintptr, int32, int32) int32

//go:wasmimport sdl3 SDL_GetWindowSize
func __SDL_GetWindowSize(uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetWindowSafeArea
func __SDL_GetWindowSafeArea(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_SetWindowAspectRatio
func __SDL_SetWindowAspectRatio(uintptr, float32, float32) int32

//go:wasmimport sdl3 SDL_GetWindowAspectRatio
func __SDL_GetWindowAspectRatio(uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetWindowBordersSize
func __SDL_GetWindowBordersSize(uintptr, uintptr, uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetWindowSizeInPixels
func __SDL_GetWindowSizeInPixels(uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_SetWindowMinimumSize
func __SDL_SetWindowMinimumSize(uintptr, int32, int32) int32

//go:wasmimport sdl3 SDL_GetWindowMinimumSize
func __SDL_GetWindowMinimumSize(uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_SetWindowMaximumSize
func __SDL_SetWindowMaximumSize(uintptr, int32, int32) int32

//go:wasmimport sdl3 SDL_GetWindowMaximumSize
func __SDL_GetWindowMaximumSize(uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_SetWindowBordered
func __SDL_SetWindowBordered(uintptr, int32) int32

//go:wasmimport sdl3 SDL_SetWindowResizable
func __SDL_SetWindowResizable(uintptr, int32) int32

//go:wasmimport sdl3 SDL_SetWindowAlwaysOnTop
func __SDL_SetWindowAlwaysOnTop(uintptr, int32) int32

//go:wasmimport sdl3 SDL_ShowWindow
func __SDL_ShowWindow(uintptr) int32

//go:wasmimport sdl3 SDL_HideWindow
func __SDL_HideWindow(uintptr) int32

//go:wasmimport sdl3 SDL_RaiseWindow
func __SDL_RaiseWindow(uintptr) int32

//go:wasmimport sdl3 SDL_MaximizeWindow
func __SDL_MaximizeWindow(uintptr) int32

//go:wasmimport sdl3 SDL_MinimizeWindow
func __SDL_MinimizeWindow(uintptr) int32

//go:wasmimport sdl3 SDL_RestoreWindow
func __SDL_RestoreWindow(uintptr) int32

//go:wasmimport sdl3 SDL_SetWindowFullscreen
func __SDL_SetWindowFullscreen(uintptr, int32) int32

//go:wasmimport sdl3 SDL_SyncWindow
func __SDL_SyncWindow(uintptr) int32

//go:wasmimport sdl3 SDL_WindowHasSurface
func __SDL_WindowHasSurface(uintptr) int32

//go:wasmimport sdl3 SDL_GetWindowSurface
func __SDL_GetWindowSurface(uintptr) uintptr

//go:wasmimport sdl3 SDL_SetWindowSurfaceVSync
func __SDL_SetWindowSurfaceVSync(uintptr, int32) int32

//go:wasmimport sdl3 SDL_GetWindowSurfaceVSync
func __SDL_GetWindowSurfaceVSync(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_UpdateWindowSurface
func __SDL_UpdateWindowSurface(uintptr) int32

//go:wasmimport sdl3 SDL_UpdateWindowSurfaceRects
func __SDL_UpdateWindowSurfaceRects(uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_DestroyWindowSurface
func __SDL_DestroyWindowSurface(uintptr) int32

//go:wasmimport sdl3 SDL_SetWindowKeyboardGrab
func __SDL_SetWindowKeyboardGrab(uintptr, int32) int32

//go:wasmimport sdl3 SDL_SetWindowMouseGrab
func __SDL_SetWindowMouseGrab(uintptr, int32) int32

//go:wasmimport sdl3 SDL_GetWindowKeyboardGrab
func __SDL_GetWindowKeyboardGrab(uintptr) int32

//go:wasmimport sdl3 SDL_GetWindowMouseGrab
func __SDL_GetWindowMouseGrab(uintptr) int32

//go:wasmimport sdl3 SDL_GetGrabbedWindow
func __SDL_GetGrabbedWindow() uintptr

//go:wasmimport sdl3 SDL_SetWindowMouseRect
func __SDL_SetWindowMouseRect(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetWindowMouseRect
func __SDL_GetWindowMouseRect(uintptr) uintptr

//go:wasmimport sdl3 SDL_SetWindowOpacity
func __SDL_SetWindowOpacity(uintptr, float32) int32

//go:wasmimport sdl3 SDL_GetWindowOpacity
func __SDL_GetWindowOpacity(uintptr) float32

//go:wasmimport sdl3 SDL_SetWindowParent
func __SDL_SetWindowParent(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_SetWindowModal
func __SDL_SetWindowModal(uintptr, int32) int32

//go:wasmimport sdl3 SDL_SetWindowFocusable
func __SDL_SetWindowFocusable(uintptr, int32) int32

//go:wasmimport sdl3 SDL_ShowWindowSystemMenu
func __SDL_ShowWindowSystemMenu(uintptr, int32, int32) int32

//go:wasmimport sdl3 SDL_SetWindowHitTest
func __SDL_SetWindowHitTest(uintptr, int32, uintptr) int32

//go:wasmimport sdl3 SDL_SetWindowShape
func __SDL_SetWindowShape(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_FlashWindow
func __SDL_FlashWindow(uintptr, int32) int32

//go:wasmimport sdl3 SDL_SetWindowProgressState
func __SDL_SetWindowProgressState(uintptr, int32) int32

//go:wasmimport sdl3 SDL_GetWindowProgressState
func __SDL_GetWindowProgressState(uintptr) int32

//go:wasmimport sdl3 SDL_SetWindowProgressValue
func __SDL_SetWindowProgressValue(uintptr, float32) int32

//go:wasmimport sdl3 SDL_GetWindowProgressValue
func __SDL_GetWindowProgressValue(uintptr) float32

//go:wasmimport sdl3 SDL_DestroyWindow
func __SDL_DestroyWindow(uintptr)

//go:wasmimport sdl3 SDL_ScreenSaverEnabled
func __SDL_ScreenSaverEnabled() int32

//go:wasmimport sdl3 SDL_EnableScreenSaver
func __SDL_EnableScreenSaver() int32

//go:wasmimport sdl3 SDL_DisableScreenSaver
func __SDL_DisableScreenSaver() int32

//go:wasmimport sdl3 SDL_GL_LoadLibrary
func __SDL_GL_LoadLibrary(uintptr) int32

//go:wasmimport sdl3 SDL_GL_GetProcAddress
func __SDL_GL_GetProcAddress(uintptr) uintptr

//go:wasmimport sdl3 SDL_EGL_GetProcAddress
func __SDL_EGL_GetProcAddress(uintptr) uintptr

//go:wasmimport sdl3 SDL_GL_UnloadLibrary
func __SDL_GL_UnloadLibrary()

//go:wasmimport sdl3 SDL_GL_ExtensionSupported
func __SDL_GL_ExtensionSupported(uintptr) int32

//go:wasmimport sdl3 SDL_GL_ResetAttributes
func __SDL_GL_ResetAttributes()

//go:wasmimport sdl3 SDL_GL_SetAttribute
func __SDL_GL_SetAttribute(int32, int32) int32

//go:wasmimport sdl3 SDL_GL_GetAttribute
func __SDL_GL_GetAttribute(int32, uintptr) int32

//go:wasmimport sdl3 SDL_GL_CreateContext
func __SDL_GL_CreateContext(uintptr) uintptr

//go:wasmimport sdl3 SDL_GL_MakeCurrent
func __SDL_GL_MakeCurrent(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GL_GetCurrentWindow
func __SDL_GL_GetCurrentWindow() uintptr

//go:wasmimport sdl3 SDL_GL_GetCurrentContext
func __SDL_GL_GetCurrentContext() uintptr

//go:wasmimport sdl3 SDL_EGL_GetCurrentDisplay
func __SDL_EGL_GetCurrentDisplay() uintptr

//go:wasmimport sdl3 SDL_EGL_GetCurrentConfig
func __SDL_EGL_GetCurrentConfig() uintptr

//go:wasmimport sdl3 SDL_EGL_GetWindowSurface
func __SDL_EGL_GetWindowSurface(uintptr) uintptr

//go:wasmimport sdl3 SDL_EGL_SetAttributeCallbacks
func __SDL_EGL_SetAttributeCallbacks(uintptr, uintptr, uintptr, uintptr)

//go:wasmimport sdl3 SDL_GL_SetSwapInterval
func __SDL_GL_SetSwapInterval(int32) int32

//go:wasmimport sdl3 SDL_GL_GetSwapInterval
func __SDL_GL_GetSwapInterval(uintptr) int32

//go:wasmimport sdl3 SDL_GL_SwapWindow
func __SDL_GL_SwapWindow(uintptr) int32

//go:wasmimport sdl3 SDL_GL_DestroyContext
func __SDL_GL_DestroyContext(uintptr) int32



func __gowrap__SDL_GetNumVideoDrivers() (__result int32) {
	__result = (int32)(__SDL_GetNumVideoDrivers())
	return
}

func __gowrap__SDL_GetVideoDriver(index int32) (__result *byte) {
	__result = (*byte)(unsafe.Pointer(__SDL_GetVideoDriver(*(*int32)(unsafe.Pointer(&index)))))
	return
}

func __gowrap__SDL_GetCurrentVideoDriver() (__result *byte) {
	__result = (*byte)(unsafe.Pointer(__SDL_GetCurrentVideoDriver()))
	return
}

func __gowrap__SDL_GetSystemTheme() (__result SystemTheme) {
	__result = (SystemTheme)(__SDL_GetSystemTheme())
	return
}

func __gowrap__SDL_GetDisplays(count *int32) (__result DisplayID) {
	__result = (DisplayID)(__SDL_GetDisplays(uintptr(unsafe.Pointer(count))))
	return
}

func __gowrap__SDL_GetPrimaryDisplay() (__result DisplayID) {
	__result = (DisplayID)(__SDL_GetPrimaryDisplay())
	return
}

func __gowrap__SDL_GetDisplayProperties(displayID DisplayID) (__result PropertiesID) {
	__result = (PropertiesID)(__SDL_GetDisplayProperties(*(*int32)(unsafe.Pointer(&displayID))))
	return
}

func __gowrap__SDL_GetDisplayName(displayID DisplayID) (__result *byte) {
	__result = (*byte)(unsafe.Pointer(__SDL_GetDisplayName(*(*int32)(unsafe.Pointer(&displayID)))))
	return
}

func __gowrap__SDL_GetDisplayBounds(displayID DisplayID, rect *Rect) (__result bool) {
	__result = (bool)(__SDL_GetDisplayBounds(*(*int32)(unsafe.Pointer(&displayID)), uintptr(unsafe.Pointer(rect))) != 0)
	return
}

func __gowrap__SDL_GetDisplayUsableBounds(displayID DisplayID, rect *Rect) (__result bool) {
	__result = (bool)(__SDL_GetDisplayUsableBounds(*(*int32)(unsafe.Pointer(&displayID)), uintptr(unsafe.Pointer(rect))) != 0)
	return
}

func __gowrap__SDL_GetNaturalDisplayOrientation(displayID DisplayID) (__result DisplayOrientation) {
	__result = (DisplayOrientation)(__SDL_GetNaturalDisplayOrientation(*(*int32)(unsafe.Pointer(&displayID))))
	return
}

func __gowrap__SDL_GetCurrentDisplayOrientation(displayID DisplayID) (__result DisplayOrientation) {
	__result = (DisplayOrientation)(__SDL_GetCurrentDisplayOrientation(*(*int32)(unsafe.Pointer(&displayID))))
	return
}

func __gowrap__SDL_GetDisplayContentScale(displayID DisplayID) (__result float32) {
	__result = (float32)(__SDL_GetDisplayContentScale(*(*int32)(unsafe.Pointer(&displayID))))
	return
}

func __gowrap__SDL_GetFullscreenDisplayModes(displayID DisplayID, count *int32) (__result **DisplayMode) {
	__result = (**DisplayMode)(unsafe.Pointer(__SDL_GetFullscreenDisplayModes(*(*int32)(unsafe.Pointer(&displayID)), uintptr(unsafe.Pointer(count)))))
	return
}

func __gowrap__SDL_GetClosestFullscreenDisplayMode(displayID DisplayID, w int32, h int32, refresh_rate float32, include_high_density_modes bool, closest *DisplayMode) (__result bool) {
	__result = (bool)(__SDL_GetClosestFullscreenDisplayMode(*(*int32)(unsafe.Pointer(&displayID)), *(*int32)(unsafe.Pointer(&w)), *(*int32)(unsafe.Pointer(&h)), *(*float32)(unsafe.Pointer(&refresh_rate)), *(*int32)(unsafe.Pointer(&include_high_density_modes)), uintptr(unsafe.Pointer(closest))) != 0)
	return
}

func __gowrap__SDL_GetDesktopDisplayMode(displayID DisplayID) (__result *DisplayMode) {
	__result = (*DisplayMode)(unsafe.Pointer(__SDL_GetDesktopDisplayMode(*(*int32)(unsafe.Pointer(&displayID)))))
	return
}

func __gowrap__SDL_GetCurrentDisplayMode(displayID DisplayID) (__result *DisplayMode) {
	__result = (*DisplayMode)(unsafe.Pointer(__SDL_GetCurrentDisplayMode(*(*int32)(unsafe.Pointer(&displayID)))))
	return
}

func __gowrap__SDL_GetDisplayForPoint(point *Point) (__result DisplayID) {
	__result = (DisplayID)(__SDL_GetDisplayForPoint(uintptr(unsafe.Pointer(point))))
	return
}

func __gowrap__SDL_GetDisplayForRect(rect *Rect) (__result DisplayID) {
	__result = (DisplayID)(__SDL_GetDisplayForRect(uintptr(unsafe.Pointer(rect))))
	return
}

func __gowrap__SDL_GetDisplayForWindow(window Window) (__result DisplayID) {
	__result = (DisplayID)(__SDL_GetDisplayForWindow(uintptr(unsafe.Pointer(window))))
	return
}

func __gowrap__SDL_GetWindowPixelDensity(window Window) (__result float32) {
	__result = (float32)(__SDL_GetWindowPixelDensity(uintptr(unsafe.Pointer(window))))
	return
}

func __gowrap__SDL_GetWindowDisplayScale(window Window) (__result float32) {
	__result = (float32)(__SDL_GetWindowDisplayScale(uintptr(unsafe.Pointer(window))))
	return
}

func __gowrap__SDL_SetWindowFullscreenMode(window Window, mode *DisplayMode) (__result bool) {
	__result = (bool)(__SDL_SetWindowFullscreenMode(uintptr(unsafe.Pointer(window)), uintptr(unsafe.Pointer(mode))) != 0)
	return
}

func __gowrap__SDL_GetWindowFullscreenMode(window Window) (__result *DisplayMode) {
	__result = (*DisplayMode)(unsafe.Pointer(__SDL_GetWindowFullscreenMode(uintptr(unsafe.Pointer(window)))))
	return
}

func __gowrap__SDL_GetWindowICCProfile(window Window, size *uint) (__result uintptr) {
	__result = (uintptr)(unsafe.Pointer(__SDL_GetWindowICCProfile(uintptr(unsafe.Pointer(window)), uintptr(unsafe.Pointer(size)))))
	return
}

func __gowrap__SDL_GetWindowPixelFormat(window Window) (__result PixelFormat) {
	__result = (PixelFormat)(__SDL_GetWindowPixelFormat(uintptr(unsafe.Pointer(window))))
	return
}

func __gowrap__SDL_GetWindows(count *int32) (__result *Window) {
	__result = (*Window)(unsafe.Pointer(__SDL_GetWindows(uintptr(unsafe.Pointer(count)))))
	return
}

func __gowrap__SDL_CreateWindow(title *byte, w int32, h int32, flags WindowFlags) (__result Window) {
	__result = (Window)(unsafe.Pointer(__SDL_CreateWindow(uintptr(unsafe.Pointer(title)), *(*int32)(unsafe.Pointer(&w)), *(*int32)(unsafe.Pointer(&h)), *(*int64)(unsafe.Pointer(&flags)))))
	return
}

func __gowrap__SDL_CreatePopupWindow(parent Window, offset_x int32, offset_y int32, w int32, h int32, flags WindowFlags) (__result Window) {
	__result = (Window)(unsafe.Pointer(__SDL_CreatePopupWindow(uintptr(unsafe.Pointer(parent)), *(*int32)(unsafe.Pointer(&offset_x)), *(*int32)(unsafe.Pointer(&offset_y)), *(*int32)(unsafe.Pointer(&w)), *(*int32)(unsafe.Pointer(&h)), *(*int64)(unsafe.Pointer(&flags)))))
	return
}

func __gowrap__SDL_CreateWindowWithProperties(props PropertiesID) (__result Window) {
	__result = (Window)(unsafe.Pointer(__SDL_CreateWindowWithProperties(*(*int32)(unsafe.Pointer(&props)))))
	return
}

func __gowrap__SDL_GetWindowID(window Window) (__result WindowID) {
	__result = (WindowID)(__SDL_GetWindowID(uintptr(unsafe.Pointer(window))))
	return
}

func __gowrap__SDL_GetWindowFromID(id WindowID) (__result Window) {
	__result = (Window)(unsafe.Pointer(__SDL_GetWindowFromID(*(*int32)(unsafe.Pointer(&id)))))
	return
}

func __gowrap__SDL_GetWindowParent(window Window) (__result Window) {
	__result = (Window)(unsafe.Pointer(__SDL_GetWindowParent(uintptr(unsafe.Pointer(window)))))
	return
}

func __gowrap__SDL_GetWindowProperties(window Window) (__result PropertiesID) {
	__result = (PropertiesID)(__SDL_GetWindowProperties(uintptr(unsafe.Pointer(window))))
	return
}

func __gowrap__SDL_GetWindowFlags(window Window) (__result WindowFlags) {
	__result = (WindowFlags)(__SDL_GetWindowFlags(uintptr(unsafe.Pointer(window))))
	return
}

func __gowrap__SDL_SetWindowTitle(window Window, title *byte) (__result bool) {
	__result = (bool)(__SDL_SetWindowTitle(uintptr(unsafe.Pointer(window)), uintptr(unsafe.Pointer(title))) != 0)
	return
}

func __gowrap__SDL_GetWindowTitle(window Window) (__result *byte) {
	__result = (*byte)(unsafe.Pointer(__SDL_GetWindowTitle(uintptr(unsafe.Pointer(window)))))
	return
}

func __gowrap__SDL_SetWindowIcon(window Window, icon *Surface) (__result bool) {
	__result = (bool)(__SDL_SetWindowIcon(uintptr(unsafe.Pointer(window)), uintptr(unsafe.Pointer(icon))) != 0)
	return
}

func __gowrap__SDL_SetWindowPosition(window Window, x int32, y int32) (__result bool) {
	__result = (bool)(__SDL_SetWindowPosition(uintptr(unsafe.Pointer(window)), *(*int32)(unsafe.Pointer(&x)), *(*int32)(unsafe.Pointer(&y))) != 0)
	return
}

func __gowrap__SDL_GetWindowPosition(window Window, x *int32, y *int32) (__result bool) {
	__result = (bool)(__SDL_GetWindowPosition(uintptr(unsafe.Pointer(window)), uintptr(unsafe.Pointer(x)), uintptr(unsafe.Pointer(y))) != 0)
	return
}

func __gowrap__SDL_SetWindowSize(window Window, w int32, h int32) (__result bool) {
	__result = (bool)(__SDL_SetWindowSize(uintptr(unsafe.Pointer(window)), *(*int32)(unsafe.Pointer(&w)), *(*int32)(unsafe.Pointer(&h))) != 0)
	return
}

func __gowrap__SDL_GetWindowSize(window Window, w *int32, h *int32) (__result bool) {
	__result = (bool)(__SDL_GetWindowSize(uintptr(unsafe.Pointer(window)), uintptr(unsafe.Pointer(w)), uintptr(unsafe.Pointer(h))) != 0)
	return
}

func __gowrap__SDL_GetWindowSafeArea(window Window, SDL_Rect *Rect) (__result bool) {
	__result = (bool)(__SDL_GetWindowSafeArea(uintptr(unsafe.Pointer(window)), uintptr(unsafe.Pointer(SDL_Rect))) != 0)
	return
}

func __gowrap__SDL_SetWindowAspectRatio(window Window, min_aspect float32, max_aspect float32) (__result bool) {
	__result = (bool)(__SDL_SetWindowAspectRatio(uintptr(unsafe.Pointer(window)), *(*float32)(unsafe.Pointer(&min_aspect)), *(*float32)(unsafe.Pointer(&max_aspect))) != 0)
	return
}

func __gowrap__SDL_GetWindowAspectRatio(window Window, min_aspect *float32, max_aspect *float32) (__result bool) {
	__result = (bool)(__SDL_GetWindowAspectRatio(uintptr(unsafe.Pointer(window)), uintptr(unsafe.Pointer(min_aspect)), uintptr(unsafe.Pointer(max_aspect))) != 0)
	return
}

func __gowrap__SDL_GetWindowBordersSize(window Window, top *int32, left *int32, bottom *int32, right *int32) (__result bool) {
	__result = (bool)(__SDL_GetWindowBordersSize(uintptr(unsafe.Pointer(window)), uintptr(unsafe.Pointer(top)), uintptr(unsafe.Pointer(left)), uintptr(unsafe.Pointer(bottom)), uintptr(unsafe.Pointer(right))) != 0)
	return
}

func __gowrap__SDL_GetWindowSizeInPixels(window Window, w *int32, h *int32) (__result bool) {
	__result = (bool)(__SDL_GetWindowSizeInPixels(uintptr(unsafe.Pointer(window)), uintptr(unsafe.Pointer(w)), uintptr(unsafe.Pointer(h))) != 0)
	return
}

func __gowrap__SDL_SetWindowMinimumSize(window Window, min_w int32, min_h int32) (__result bool) {
	__result = (bool)(__SDL_SetWindowMinimumSize(uintptr(unsafe.Pointer(window)), *(*int32)(unsafe.Pointer(&min_w)), *(*int32)(unsafe.Pointer(&min_h))) != 0)
	return
}

func __gowrap__SDL_GetWindowMinimumSize(window Window, w *int32, h *int32) (__result bool) {
	__result = (bool)(__SDL_GetWindowMinimumSize(uintptr(unsafe.Pointer(window)), uintptr(unsafe.Pointer(w)), uintptr(unsafe.Pointer(h))) != 0)
	return
}

func __gowrap__SDL_SetWindowMaximumSize(window Window, max_w int32, max_h int32) (__result bool) {
	__result = (bool)(__SDL_SetWindowMaximumSize(uintptr(unsafe.Pointer(window)), *(*int32)(unsafe.Pointer(&max_w)), *(*int32)(unsafe.Pointer(&max_h))) != 0)
	return
}

func __gowrap__SDL_GetWindowMaximumSize(window Window, w *int32, h *int32) (__result bool) {
	__result = (bool)(__SDL_GetWindowMaximumSize(uintptr(unsafe.Pointer(window)), uintptr(unsafe.Pointer(w)), uintptr(unsafe.Pointer(h))) != 0)
	return
}

func __gowrap__SDL_SetWindowBordered(window Window, bordered bool) (__result bool) {
	__result = (bool)(__SDL_SetWindowBordered(uintptr(unsafe.Pointer(window)), *(*int32)(unsafe.Pointer(&bordered))) != 0)
	return
}

func __gowrap__SDL_SetWindowResizable(window Window, resizable bool) (__result bool) {
	__result = (bool)(__SDL_SetWindowResizable(uintptr(unsafe.Pointer(window)), *(*int32)(unsafe.Pointer(&resizable))) != 0)
	return
}

func __gowrap__SDL_SetWindowAlwaysOnTop(window Window, on_top bool) (__result bool) {
	__result = (bool)(__SDL_SetWindowAlwaysOnTop(uintptr(unsafe.Pointer(window)), *(*int32)(unsafe.Pointer(&on_top))) != 0)
	return
}

func __gowrap__SDL_ShowWindow(window Window) (__result bool) {
	__result = (bool)(__SDL_ShowWindow(uintptr(unsafe.Pointer(window))) != 0)
	return
}

func __gowrap__SDL_HideWindow(window Window) (__result bool) {
	__result = (bool)(__SDL_HideWindow(uintptr(unsafe.Pointer(window))) != 0)
	return
}

func __gowrap__SDL_RaiseWindow(window Window) (__result bool) {
	__result = (bool)(__SDL_RaiseWindow(uintptr(unsafe.Pointer(window))) != 0)
	return
}

func __gowrap__SDL_MaximizeWindow(window Window) (__result bool) {
	__result = (bool)(__SDL_MaximizeWindow(uintptr(unsafe.Pointer(window))) != 0)
	return
}

func __gowrap__SDL_MinimizeWindow(window Window) (__result bool) {
	__result = (bool)(__SDL_MinimizeWindow(uintptr(unsafe.Pointer(window))) != 0)
	return
}

func __gowrap__SDL_RestoreWindow(window Window) (__result bool) {
	__result = (bool)(__SDL_RestoreWindow(uintptr(unsafe.Pointer(window))) != 0)
	return
}

func __gowrap__SDL_SetWindowFullscreen(window Window, fullscreen bool) (__result bool) {
	__result = (bool)(__SDL_SetWindowFullscreen(uintptr(unsafe.Pointer(window)), *(*int32)(unsafe.Pointer(&fullscreen))) != 0)
	return
}

func __gowrap__SDL_SyncWindow(window Window) (__result bool) {
	__result = (bool)(__SDL_SyncWindow(uintptr(unsafe.Pointer(window))) != 0)
	return
}

func __gowrap__SDL_WindowHasSurface(window Window) (__result bool) {
	__result = (bool)(__SDL_WindowHasSurface(uintptr(unsafe.Pointer(window))) != 0)
	return
}

func __gowrap__SDL_GetWindowSurface(window Window) (__result *Surface) {
	__result = (*Surface)(unsafe.Pointer(__SDL_GetWindowSurface(uintptr(unsafe.Pointer(window)))))
	return
}

func __gowrap__SDL_SetWindowSurfaceVSync(window Window, vsync VSync) (__result bool) {
	__result = (bool)(__SDL_SetWindowSurfaceVSync(uintptr(unsafe.Pointer(window)), *(*int32)(unsafe.Pointer(&vsync))) != 0)
	return
}

func __gowrap__SDL_GetWindowSurfaceVSync(window Window, vsync *VSync) (__result bool) {
	__result = (bool)(__SDL_GetWindowSurfaceVSync(uintptr(unsafe.Pointer(window)), uintptr(unsafe.Pointer(vsync))) != 0)
	return
}

func __gowrap__SDL_UpdateWindowSurface(window Window) (__result bool) {
	__result = (bool)(__SDL_UpdateWindowSurface(uintptr(unsafe.Pointer(window))) != 0)
	return
}

func __gowrap__SDL_UpdateWindowSurfaceRects(window Window, rects *Rect, numrects int32) (__result bool) {
	__result = (bool)(__SDL_UpdateWindowSurfaceRects(uintptr(unsafe.Pointer(window)), uintptr(unsafe.Pointer(rects)), *(*int32)(unsafe.Pointer(&numrects))) != 0)
	return
}

func __gowrap__SDL_DestroyWindowSurface(window Window) (__result bool) {
	__result = (bool)(__SDL_DestroyWindowSurface(uintptr(unsafe.Pointer(window))) != 0)
	return
}

func __gowrap__SDL_SetWindowKeyboardGrab(window Window, grabbed bool) (__result bool) {
	__result = (bool)(__SDL_SetWindowKeyboardGrab(uintptr(unsafe.Pointer(window)), *(*int32)(unsafe.Pointer(&grabbed))) != 0)
	return
}

func __gowrap__SDL_SetWindowMouseGrab(window Window, grabbed bool) (__result bool) {
	__result = (bool)(__SDL_SetWindowMouseGrab(uintptr(unsafe.Pointer(window)), *(*int32)(unsafe.Pointer(&grabbed))) != 0)
	return
}

func __gowrap__SDL_GetWindowKeyboardGrab(window Window) (__result bool) {
	__result = (bool)(__SDL_GetWindowKeyboardGrab(uintptr(unsafe.Pointer(window))) != 0)
	return
}

func __gowrap__SDL_GetWindowMouseGrab(window Window) (__result bool) {
	__result = (bool)(__SDL_GetWindowMouseGrab(uintptr(unsafe.Pointer(window))) != 0)
	return
}

func __gowrap__SDL_GetGrabbedWindow() (__result Window) {
	__result = (Window)(unsafe.Pointer(__SDL_GetGrabbedWindow()))
	return
}

func __gowrap__SDL_SetWindowMouseRect(window Window, rect *Rect) (__result bool) {
	__result = (bool)(__SDL_SetWindowMouseRect(uintptr(unsafe.Pointer(window)), uintptr(unsafe.Pointer(rect))) != 0)
	return
}

func __gowrap__SDL_GetWindowMouseRect(window Window) (__result *Rect) {
	__result = (*Rect)(unsafe.Pointer(__SDL_GetWindowMouseRect(uintptr(unsafe.Pointer(window)))))
	return
}

func __gowrap__SDL_SetWindowOpacity(window Window, opacity float32) (__result bool) {
	__result = (bool)(__SDL_SetWindowOpacity(uintptr(unsafe.Pointer(window)), *(*float32)(unsafe.Pointer(&opacity))) != 0)
	return
}

func __gowrap__SDL_GetWindowOpacity(window Window) (__result float32) {
	__result = (float32)(__SDL_GetWindowOpacity(uintptr(unsafe.Pointer(window))))
	return
}

func __gowrap__SDL_SetWindowParent(window Window, parent Window) (__result bool) {
	__result = (bool)(__SDL_SetWindowParent(uintptr(unsafe.Pointer(window)), uintptr(unsafe.Pointer(parent))) != 0)
	return
}

func __gowrap__SDL_SetWindowModal(window Window, modal bool) (__result bool) {
	__result = (bool)(__SDL_SetWindowModal(uintptr(unsafe.Pointer(window)), *(*int32)(unsafe.Pointer(&modal))) != 0)
	return
}

func __gowrap__SDL_SetWindowFocusable(window Window, focusable bool) (__result bool) {
	__result = (bool)(__SDL_SetWindowFocusable(uintptr(unsafe.Pointer(window)), *(*int32)(unsafe.Pointer(&focusable))) != 0)
	return
}

func __gowrap__SDL_ShowWindowSystemMenu(window Window, x int32, y int32) (__result bool) {
	__result = (bool)(__SDL_ShowWindowSystemMenu(uintptr(unsafe.Pointer(window)), *(*int32)(unsafe.Pointer(&x)), *(*int32)(unsafe.Pointer(&y))) != 0)
	return
}

func __gowrap__SDL_SetWindowHitTest(window Window, callback HitTest, callback_data uintptr) (__result bool) {
	__result = (bool)(__SDL_SetWindowHitTest(uintptr(unsafe.Pointer(window)), *(*int32)(unsafe.Pointer(&callback)), uintptr(unsafe.Pointer(callback_data))) != 0)
	return
}

func __gowrap__SDL_SetWindowShape(window Window, shape *Surface) (__result bool) {
	__result = (bool)(__SDL_SetWindowShape(uintptr(unsafe.Pointer(window)), uintptr(unsafe.Pointer(shape))) != 0)
	return
}

func __gowrap__SDL_FlashWindow(window Window, operation FlashOperation) (__result bool) {
	__result = (bool)(__SDL_FlashWindow(uintptr(unsafe.Pointer(window)), *(*int32)(unsafe.Pointer(&operation))) != 0)
	return
}

func __gowrap__SDL_SetWindowProgressState(window Window, state ProgressState) (__result bool) {
	__result = (bool)(__SDL_SetWindowProgressState(uintptr(unsafe.Pointer(window)), *(*int32)(unsafe.Pointer(&state))) != 0)
	return
}

func __gowrap__SDL_GetWindowProgressState(window Window) (__result ProgressState) {
	__result = (ProgressState)(__SDL_GetWindowProgressState(uintptr(unsafe.Pointer(window))))
	return
}

func __gowrap__SDL_SetWindowProgressValue(window Window, value float32) (__result bool) {
	__result = (bool)(__SDL_SetWindowProgressValue(uintptr(unsafe.Pointer(window)), *(*float32)(unsafe.Pointer(&value))) != 0)
	return
}

func __gowrap__SDL_GetWindowProgressValue(window Window) (__result float32) {
	__result = (float32)(__SDL_GetWindowProgressValue(uintptr(unsafe.Pointer(window))))
	return
}

func __gowrap__SDL_DestroyWindow(window Window) {
	__SDL_DestroyWindow(uintptr(unsafe.Pointer(window)))
}

func __gowrap__SDL_ScreenSaverEnabled() (__result bool) {
	__result = (bool)(__SDL_ScreenSaverEnabled() != 0)
	return
}

func __gowrap__SDL_EnableScreenSaver() (__result bool) {
	__result = (bool)(__SDL_EnableScreenSaver() != 0)
	return
}

func __gowrap__SDL_DisableScreenSaver() (__result bool) {
	__result = (bool)(__SDL_DisableScreenSaver() != 0)
	return
}

func __gowrap__SDL_GL_LoadLibrary(path *byte) (__result bool) {
	__result = (bool)(__SDL_GL_LoadLibrary(uintptr(unsafe.Pointer(path))) != 0)
	return
}

func __gowrap__SDL_GL_GetProcAddress(proc *byte) (__result uintptr) {
	__result = (uintptr)(unsafe.Pointer(__SDL_GL_GetProcAddress(uintptr(unsafe.Pointer(proc)))))
	return
}

func __gowrap__SDL_EGL_GetProcAddress(proc *byte) (__result uintptr) {
	__result = (uintptr)(unsafe.Pointer(__SDL_EGL_GetProcAddress(uintptr(unsafe.Pointer(proc)))))
	return
}

func __gowrap__SDL_GL_UnloadLibrary() {
	__SDL_GL_UnloadLibrary()
}

func __gowrap__SDL_GL_ExtensionSupported(extension *byte) (__result bool) {
	__result = (bool)(__SDL_GL_ExtensionSupported(uintptr(unsafe.Pointer(extension))) != 0)
	return
}

func __gowrap__SDL_GL_ResetAttributes() {
	__SDL_GL_ResetAttributes()
}

func __gowrap__SDL_GL_SetAttribute(attr GLAttr, value int32) (__result bool) {
	__result = (bool)(__SDL_GL_SetAttribute(*(*int32)(unsafe.Pointer(&attr)), *(*int32)(unsafe.Pointer(&value))) != 0)
	return
}

func __gowrap__SDL_GL_GetAttribute(attr GLAttr, value *int32) (__result bool) {
	__result = (bool)(__SDL_GL_GetAttribute(*(*int32)(unsafe.Pointer(&attr)), uintptr(unsafe.Pointer(value))) != 0)
	return
}

func __gowrap__SDL_GL_CreateContext(window Window) (__result GLContext) {
	__result = (GLContext)(unsafe.Pointer(__SDL_GL_CreateContext(uintptr(unsafe.Pointer(window)))))
	return
}

func __gowrap__SDL_GL_MakeCurrent(window Window, context GLContext) (__result bool) {
	__result = (bool)(__SDL_GL_MakeCurrent(uintptr(unsafe.Pointer(window)), uintptr(unsafe.Pointer(context))) != 0)
	return
}

func __gowrap__SDL_GL_GetCurrentWindow() (__result Window) {
	__result = (Window)(unsafe.Pointer(__SDL_GL_GetCurrentWindow()))
	return
}

func __gowrap__SDL_GL_GetCurrentContext() (__result GLContext) {
	__result = (GLContext)(unsafe.Pointer(__SDL_GL_GetCurrentContext()))
	return
}

func __gowrap__SDL_EGL_GetCurrentDisplay() (__result EGLDisplay) {
	__result = (EGLDisplay)(unsafe.Pointer(__SDL_EGL_GetCurrentDisplay()))
	return
}

func __gowrap__SDL_EGL_GetCurrentConfig() (__result EGLConfig) {
	__result = (EGLConfig)(unsafe.Pointer(__SDL_EGL_GetCurrentConfig()))
	return
}

func __gowrap__SDL_EGL_GetWindowSurface(window Window) (__result EGLSurface) {
	__result = (EGLSurface)(unsafe.Pointer(__SDL_EGL_GetWindowSurface(uintptr(unsafe.Pointer(window)))))
	return
}

func __gowrap__SDL_EGL_SetAttributeCallbacks(platformAttribCallback EGLAttribArrayCallback, surfaceAttribCallback EGLIntArrayCallback, contextAttribCallback EGLIntArrayCallback, userdata uintptr) {
	__SDL_EGL_SetAttributeCallbacks(0 /* TODO: callbacks */, 0 /* TODO: callbacks */, 0 /* TODO: callbacks */, uintptr(unsafe.Pointer(userdata)))
}

func __gowrap__SDL_GL_SetSwapInterval(interval int32) (__result bool) {
	__result = (bool)(__SDL_GL_SetSwapInterval(*(*int32)(unsafe.Pointer(&interval))) != 0)
	return
}

func __gowrap__SDL_GL_GetSwapInterval(interval *int32) (__result bool) {
	__result = (bool)(__SDL_GL_GetSwapInterval(uintptr(unsafe.Pointer(interval))) != 0)
	return
}

func __gowrap__SDL_GL_SwapWindow(window Window) (__result bool) {
	__result = (bool)(__SDL_GL_SwapWindow(uintptr(unsafe.Pointer(window))) != 0)
	return
}

func __gowrap__SDL_GL_DestroyContext(context GLContext) (__result bool) {
	__result = (bool)(__SDL_GL_DestroyContext(uintptr(unsafe.Pointer(context))) != 0)
	return
}

 
func init() {
	GetNumVideoDrivers = __gowrap__SDL_GetNumVideoDrivers
	GetVideoDriver = __gowrap__SDL_GetVideoDriver
	GetCurrentVideoDriver = __gowrap__SDL_GetCurrentVideoDriver
	GetSystemTheme = __gowrap__SDL_GetSystemTheme
	GetDisplays = __gowrap__SDL_GetDisplays
	GetPrimaryDisplay = __gowrap__SDL_GetPrimaryDisplay
	GetDisplayProperties = __gowrap__SDL_GetDisplayProperties
	GetDisplayName = __gowrap__SDL_GetDisplayName
	GetDisplayBounds = __gowrap__SDL_GetDisplayBounds
	GetDisplayUsableBounds = __gowrap__SDL_GetDisplayUsableBounds
	GetNaturalDisplayOrientation = __gowrap__SDL_GetNaturalDisplayOrientation
	GetCurrentDisplayOrientation = __gowrap__SDL_GetCurrentDisplayOrientation
	GetDisplayContentScale = __gowrap__SDL_GetDisplayContentScale
	GetFullscreenDisplayModes = __gowrap__SDL_GetFullscreenDisplayModes
	GetClosestFullscreenDisplayMode = __gowrap__SDL_GetClosestFullscreenDisplayMode
	GetDesktopDisplayMode = __gowrap__SDL_GetDesktopDisplayMode
	GetCurrentDisplayMode = __gowrap__SDL_GetCurrentDisplayMode
	GetDisplayForPoint = __gowrap__SDL_GetDisplayForPoint
	GetDisplayForRect = __gowrap__SDL_GetDisplayForRect
	GetDisplayForWindow = __gowrap__SDL_GetDisplayForWindow
	GetWindowPixelDensity = __gowrap__SDL_GetWindowPixelDensity
	GetWindowDisplayScale = __gowrap__SDL_GetWindowDisplayScale
	SetWindowFullscreenMode = __gowrap__SDL_SetWindowFullscreenMode
	GetWindowFullscreenMode = __gowrap__SDL_GetWindowFullscreenMode
	GetWindowICCProfile = __gowrap__SDL_GetWindowICCProfile
	GetWindowPixelFormat = __gowrap__SDL_GetWindowPixelFormat
	GetWindows = __gowrap__SDL_GetWindows
	CreateWindow = __gowrap__SDL_CreateWindow
	CreatePopupWindow = __gowrap__SDL_CreatePopupWindow
	CreateWindowWithProperties = __gowrap__SDL_CreateWindowWithProperties
	GetWindowID = __gowrap__SDL_GetWindowID
	GetWindowFromID = __gowrap__SDL_GetWindowFromID
	GetWindowParent = __gowrap__SDL_GetWindowParent
	GetWindowProperties = __gowrap__SDL_GetWindowProperties
	GetWindowFlags = __gowrap__SDL_GetWindowFlags
	SetWindowTitle = __gowrap__SDL_SetWindowTitle
	GetWindowTitle = __gowrap__SDL_GetWindowTitle
	SetWindowIcon = __gowrap__SDL_SetWindowIcon
	SetWindowPosition = __gowrap__SDL_SetWindowPosition
	GetWindowPosition = __gowrap__SDL_GetWindowPosition
	SetWindowSize = __gowrap__SDL_SetWindowSize
	GetWindowSize = __gowrap__SDL_GetWindowSize
	GetWindowSafeArea = __gowrap__SDL_GetWindowSafeArea
	SetWindowAspectRatio = __gowrap__SDL_SetWindowAspectRatio
	GetWindowAspectRatio = __gowrap__SDL_GetWindowAspectRatio
	GetWindowBordersSize = __gowrap__SDL_GetWindowBordersSize
	GetWindowSizeInPixels = __gowrap__SDL_GetWindowSizeInPixels
	SetWindowMinimumSize = __gowrap__SDL_SetWindowMinimumSize
	GetWindowMinimumSize = __gowrap__SDL_GetWindowMinimumSize
	SetWindowMaximumSize = __gowrap__SDL_SetWindowMaximumSize
	GetWindowMaximumSize = __gowrap__SDL_GetWindowMaximumSize
	SetWindowBordered = __gowrap__SDL_SetWindowBordered
	SetWindowResizable = __gowrap__SDL_SetWindowResizable
	SetWindowAlwaysOnTop = __gowrap__SDL_SetWindowAlwaysOnTop
	ShowWindow = __gowrap__SDL_ShowWindow
	HideWindow = __gowrap__SDL_HideWindow
	RaiseWindow = __gowrap__SDL_RaiseWindow
	MaximizeWindow = __gowrap__SDL_MaximizeWindow
	MinimizeWindow = __gowrap__SDL_MinimizeWindow
	RestoreWindow = __gowrap__SDL_RestoreWindow
	SetWindowFullscreen = __gowrap__SDL_SetWindowFullscreen
	SyncWindow = __gowrap__SDL_SyncWindow
	WindowHasSurface = __gowrap__SDL_WindowHasSurface
	GetWindowSurface = __gowrap__SDL_GetWindowSurface
	SetWindowSurfaceVSync = __gowrap__SDL_SetWindowSurfaceVSync
	GetWindowSurfaceVSync = __gowrap__SDL_GetWindowSurfaceVSync
	UpdateWindowSurface = __gowrap__SDL_UpdateWindowSurface
	UpdateWindowSurfaceRects = __gowrap__SDL_UpdateWindowSurfaceRects
	DestroyWindowSurface = __gowrap__SDL_DestroyWindowSurface
	SetWindowKeyboardGrab = __gowrap__SDL_SetWindowKeyboardGrab
	SetWindowMouseGrab = __gowrap__SDL_SetWindowMouseGrab
	GetWindowKeyboardGrab = __gowrap__SDL_GetWindowKeyboardGrab
	GetWindowMouseGrab = __gowrap__SDL_GetWindowMouseGrab
	GetGrabbedWindow = __gowrap__SDL_GetGrabbedWindow
	SetWindowMouseRect = __gowrap__SDL_SetWindowMouseRect
	GetWindowMouseRect = __gowrap__SDL_GetWindowMouseRect
	SetWindowOpacity = __gowrap__SDL_SetWindowOpacity
	GetWindowOpacity = __gowrap__SDL_GetWindowOpacity
	SetWindowParent = __gowrap__SDL_SetWindowParent
	SetWindowModal = __gowrap__SDL_SetWindowModal
	SetWindowFocusable = __gowrap__SDL_SetWindowFocusable
	ShowWindowSystemMenu = __gowrap__SDL_ShowWindowSystemMenu
	SetWindowHitTest = __gowrap__SDL_SetWindowHitTest
	SetWindowShape = __gowrap__SDL_SetWindowShape
	FlashWindow = __gowrap__SDL_FlashWindow
	SetWindowProgressState = __gowrap__SDL_SetWindowProgressState
	GetWindowProgressState = __gowrap__SDL_GetWindowProgressState
	SetWindowProgressValue = __gowrap__SDL_SetWindowProgressValue
	GetWindowProgressValue = __gowrap__SDL_GetWindowProgressValue
	DestroyWindow = __gowrap__SDL_DestroyWindow
	ScreenSaverEnabled = __gowrap__SDL_ScreenSaverEnabled
	EnableScreenSaver = __gowrap__SDL_EnableScreenSaver
	DisableScreenSaver = __gowrap__SDL_DisableScreenSaver
	GL_LoadLibrary = __gowrap__SDL_GL_LoadLibrary
	GL_GetProcAddress = __gowrap__SDL_GL_GetProcAddress
	EGL_GetProcAddress = __gowrap__SDL_EGL_GetProcAddress
	GL_UnloadLibrary = __gowrap__SDL_GL_UnloadLibrary
	GL_ExtensionSupported = __gowrap__SDL_GL_ExtensionSupported
	GL_ResetAttributes = __gowrap__SDL_GL_ResetAttributes
	GL_SetAttribute = __gowrap__SDL_GL_SetAttribute
	GL_GetAttribute = __gowrap__SDL_GL_GetAttribute
	GL_CreateContext = __gowrap__SDL_GL_CreateContext
	GL_MakeCurrent = __gowrap__SDL_GL_MakeCurrent
	GL_GetCurrentWindow = __gowrap__SDL_GL_GetCurrentWindow
	GL_GetCurrentContext = __gowrap__SDL_GL_GetCurrentContext
	EGL_GetCurrentDisplay = __gowrap__SDL_EGL_GetCurrentDisplay
	EGL_GetCurrentConfig = __gowrap__SDL_EGL_GetCurrentConfig
	EGL_GetWindowSurface = __gowrap__SDL_EGL_GetWindowSurface
	EGL_SetAttributeCallbacks = __gowrap__SDL_EGL_SetAttributeCallbacks
	GL_SetSwapInterval = __gowrap__SDL_GL_SetSwapInterval
	GL_GetSwapInterval = __gowrap__SDL_GL_GetSwapInterval
	GL_SwapWindow = __gowrap__SDL_GL_SwapWindow
	GL_DestroyContext = __gowrap__SDL_GL_DestroyContext
}
