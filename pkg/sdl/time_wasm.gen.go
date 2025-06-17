//go:build wasm

package sdl

import (
  "unsafe"
)



//go:wasmimport sdl3 SDL_GetDateTimeLocalePreferences
func __SDL_GetDateTimeLocalePreferences(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetCurrentTime
func __SDL_GetCurrentTime(uintptr) int32

//go:wasmimport sdl3 SDL_TimeToDateTime
func __SDL_TimeToDateTime(int64, uintptr, int32) int32

//go:wasmimport sdl3 SDL_DateTimeToTime
func __SDL_DateTimeToTime(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_TimeToWindows
func __SDL_TimeToWindows(int64, uintptr, uintptr)

//go:wasmimport sdl3 SDL_TimeFromWindows
func __SDL_TimeFromWindows(int32, int32) int64

//go:wasmimport sdl3 SDL_GetDaysInMonth
func __SDL_GetDaysInMonth(int32, int32) int32

//go:wasmimport sdl3 SDL_GetDayOfYear
func __SDL_GetDayOfYear(int32, int32, int32) int32

//go:wasmimport sdl3 SDL_GetDayOfWeek
func __SDL_GetDayOfWeek(int32, int32, int32) int32



func __gowrap__SDL_GetDateTimeLocalePreferences(dateFormat *DateFormat, timeFormat *TimeFormat) (__result bool) {
	__result = (bool)(__SDL_GetDateTimeLocalePreferences(uintptr(unsafe.Pointer(dateFormat)), uintptr(unsafe.Pointer(timeFormat))) != 0)
	return
}

func __gowrap__SDL_GetCurrentTime(ticks *Time) (__result bool) {
	__result = (bool)(__SDL_GetCurrentTime(uintptr(unsafe.Pointer(ticks))) != 0)
	return
}

func __gowrap__SDL_TimeToDateTime(ticks Time, dt *DateTime, localTime bool) (__result bool) {
	__result = (bool)(__SDL_TimeToDateTime(*(*int64)(unsafe.Pointer(&ticks)), uintptr(unsafe.Pointer(dt)), *(*int32)(unsafe.Pointer(&localTime))) != 0)
	return
}

func __gowrap__SDL_DateTimeToTime(dt *DateTime, ticks *Time) (__result bool) {
	__result = (bool)(__SDL_DateTimeToTime(uintptr(unsafe.Pointer(dt)), uintptr(unsafe.Pointer(ticks))) != 0)
	return
}

func __gowrap__SDL_TimeToWindows(ticks Time, dwLowDateTime *uint32, dwHighDateTime *uint32) {
	__SDL_TimeToWindows(*(*int64)(unsafe.Pointer(&ticks)), uintptr(unsafe.Pointer(dwLowDateTime)), uintptr(unsafe.Pointer(dwHighDateTime)))
}

func __gowrap__SDL_TimeFromWindows(dwLowDateTime uint32, dwHighDateTime uint32) (__result Time) {
	__result = (Time)(__SDL_TimeFromWindows(*(*int32)(unsafe.Pointer(&dwLowDateTime)), *(*int32)(unsafe.Pointer(&dwHighDateTime))))
	return
}

func __gowrap__SDL_GetDaysInMonth(year int32, month int32) (__result int32) {
	__result = (int32)(__SDL_GetDaysInMonth(*(*int32)(unsafe.Pointer(&year)), *(*int32)(unsafe.Pointer(&month))))
	return
}

func __gowrap__SDL_GetDayOfYear(year int32, month int32, day int32) (__result int32) {
	__result = (int32)(__SDL_GetDayOfYear(*(*int32)(unsafe.Pointer(&year)), *(*int32)(unsafe.Pointer(&month)), *(*int32)(unsafe.Pointer(&day))))
	return
}

func __gowrap__SDL_GetDayOfWeek(year int32, month int32, day int32) (__result int32) {
	__result = (int32)(__SDL_GetDayOfWeek(*(*int32)(unsafe.Pointer(&year)), *(*int32)(unsafe.Pointer(&month)), *(*int32)(unsafe.Pointer(&day))))
	return
}

 
func init() {
	GetDateTimeLocalePreferences = __gowrap__SDL_GetDateTimeLocalePreferences
	GetCurrentTime = __gowrap__SDL_GetCurrentTime
	TimeToDateTime = __gowrap__SDL_TimeToDateTime
	DateTimeToTime = __gowrap__SDL_DateTimeToTime
	TimeToWindows = __gowrap__SDL_TimeToWindows
	TimeFromWindows = __gowrap__SDL_TimeFromWindows
	GetDaysInMonth = __gowrap__SDL_GetDaysInMonth
	GetDayOfYear = __gowrap__SDL_GetDayOfYear
	GetDayOfWeek = __gowrap__SDL_GetDayOfWeek
}
