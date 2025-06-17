//go:build wasm

package sdl

import (
  "unsafe"
)



//go:wasmimport sdl3 SDL_HasRectIntersection
func __SDL_HasRectIntersection(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetRectIntersection
func __SDL_GetRectIntersection(uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetRectUnion
func __SDL_GetRectUnion(uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetRectEnclosingPoints
func __SDL_GetRectEnclosingPoints(uintptr, int32, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetRectAndLineIntersection
func __SDL_GetRectAndLineIntersection(uintptr, uintptr, uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_HasRectIntersectionFloat
func __SDL_HasRectIntersectionFloat(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetRectIntersectionFloat
func __SDL_GetRectIntersectionFloat(uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetRectUnionFloat
func __SDL_GetRectUnionFloat(uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetRectEnclosingPointsFloat
func __SDL_GetRectEnclosingPointsFloat(uintptr, int32, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetRectAndLineIntersectionFloat
func __SDL_GetRectAndLineIntersectionFloat(uintptr, uintptr, uintptr, uintptr, uintptr) int32



func __gowrap__SDL_HasRectIntersection(A *Rect, B *Rect) (__result bool) {
	__result = (bool)(__SDL_HasRectIntersection(uintptr(unsafe.Pointer(A)), uintptr(unsafe.Pointer(B))) != 0)
	return
}

func __gowrap__SDL_GetRectIntersection(A *Rect, B *Rect, result *Rect) (__result bool) {
	__result = (bool)(__SDL_GetRectIntersection(uintptr(unsafe.Pointer(A)), uintptr(unsafe.Pointer(B)), uintptr(unsafe.Pointer(result))) != 0)
	return
}

func __gowrap__SDL_GetRectUnion(A *Rect, B *Rect, result *Rect) (__result bool) {
	__result = (bool)(__SDL_GetRectUnion(uintptr(unsafe.Pointer(A)), uintptr(unsafe.Pointer(B)), uintptr(unsafe.Pointer(result))) != 0)
	return
}

func __gowrap__SDL_GetRectEnclosingPoints(points *Point, count int32, clip *Rect, result *Rect) (__result bool) {
	__result = (bool)(__SDL_GetRectEnclosingPoints(uintptr(unsafe.Pointer(points)), *(*int32)(unsafe.Pointer(&count)), uintptr(unsafe.Pointer(clip)), uintptr(unsafe.Pointer(result))) != 0)
	return
}

func __gowrap__SDL_GetRectAndLineIntersection(rect *Rect, X1 *int32, Y1 *int32, X2 *int32, Y2 *int32) (__result bool) {
	__result = (bool)(__SDL_GetRectAndLineIntersection(uintptr(unsafe.Pointer(rect)), uintptr(unsafe.Pointer(X1)), uintptr(unsafe.Pointer(Y1)), uintptr(unsafe.Pointer(X2)), uintptr(unsafe.Pointer(Y2))) != 0)
	return
}

func __gowrap__SDL_HasRectIntersectionFloat(A *FRect, B *FRect) (__result bool) {
	__result = (bool)(__SDL_HasRectIntersectionFloat(uintptr(unsafe.Pointer(A)), uintptr(unsafe.Pointer(B))) != 0)
	return
}

func __gowrap__SDL_GetRectIntersectionFloat(A *FRect, B *FRect, result *FRect) (__result bool) {
	__result = (bool)(__SDL_GetRectIntersectionFloat(uintptr(unsafe.Pointer(A)), uintptr(unsafe.Pointer(B)), uintptr(unsafe.Pointer(result))) != 0)
	return
}

func __gowrap__SDL_GetRectUnionFloat(A *FRect, B *FRect, result *FRect) (__result bool) {
	__result = (bool)(__SDL_GetRectUnionFloat(uintptr(unsafe.Pointer(A)), uintptr(unsafe.Pointer(B)), uintptr(unsafe.Pointer(result))) != 0)
	return
}

func __gowrap__SDL_GetRectEnclosingPointsFloat(points *FPoint, count int32, clip *FRect, result *FRect) (__result bool) {
	__result = (bool)(__SDL_GetRectEnclosingPointsFloat(uintptr(unsafe.Pointer(points)), *(*int32)(unsafe.Pointer(&count)), uintptr(unsafe.Pointer(clip)), uintptr(unsafe.Pointer(result))) != 0)
	return
}

func __gowrap__SDL_GetRectAndLineIntersectionFloat(rect *FRect, X1 *float32, Y1 *float32, X2 *float32, Y2 *float32) (__result bool) {
	__result = (bool)(__SDL_GetRectAndLineIntersectionFloat(uintptr(unsafe.Pointer(rect)), uintptr(unsafe.Pointer(X1)), uintptr(unsafe.Pointer(Y1)), uintptr(unsafe.Pointer(X2)), uintptr(unsafe.Pointer(Y2))) != 0)
	return
}

 
func init() {
	HasRectIntersection = __gowrap__SDL_HasRectIntersection
	GetRectIntersection = __gowrap__SDL_GetRectIntersection
	GetRectUnion = __gowrap__SDL_GetRectUnion
	GetRectEnclosingPoints = __gowrap__SDL_GetRectEnclosingPoints
	GetRectAndLineIntersection = __gowrap__SDL_GetRectAndLineIntersection
	HasRectIntersectionFloat = __gowrap__SDL_HasRectIntersectionFloat
	GetRectIntersectionFloat = __gowrap__SDL_GetRectIntersectionFloat
	GetRectUnionFloat = __gowrap__SDL_GetRectUnionFloat
	GetRectEnclosingPointsFloat = __gowrap__SDL_GetRectEnclosingPointsFloat
	GetRectAndLineIntersectionFloat = __gowrap__SDL_GetRectAndLineIntersectionFloat
}
