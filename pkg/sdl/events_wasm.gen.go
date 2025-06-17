//go:build wasm

package sdl

import (
  "runtime"
  "unsafe"
)



//go:wasmimport sdl3 SDL_PumpEvents
func __SDL_PumpEvents()

//go:wasmimport sdl3 SDL_PeepEvents
func __SDL_PeepEvents(uintptr, int32, int32, int32, int32) int32

//go:wasmimport sdl3 SDL_HasEvent
func __SDL_HasEvent(int32) int32

//go:wasmimport sdl3 SDL_HasEvents
func __SDL_HasEvents(int32, int32) int32

//go:wasmimport sdl3 SDL_FlushEvent
func __SDL_FlushEvent(int32)

//go:wasmimport sdl3 SDL_FlushEvents
func __SDL_FlushEvents(int32, int32)

//go:wasmimport sdl3 SDL_PollEvent
func __SDL_PollEvent(uintptr) int32

//go:wasmimport sdl3 SDL_WaitEvent
func __SDL_WaitEvent(uintptr) int32

//go:wasmimport sdl3 SDL_WaitEventTimeout
func __SDL_WaitEventTimeout(uintptr, int32) int32

//go:wasmimport sdl3 SDL_PushEvent
func __SDL_PushEvent(uintptr) int32

//go:wasmimport sdl3 SDL_SetEventFilter
func __SDL_SetEventFilter(uintptr, uintptr)

//go:wasmimport sdl3 SDL_GetEventFilter
func __SDL_GetEventFilter(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_AddEventWatch
func __SDL_AddEventWatch(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_RemoveEventWatch
func __SDL_RemoveEventWatch(uintptr, uintptr)

//go:wasmimport sdl3 SDL_FilterEvents
func __SDL_FilterEvents(uintptr, uintptr)

//go:wasmimport sdl3 SDL_SetEventEnabled
func __SDL_SetEventEnabled(int32, int32)

//go:wasmimport sdl3 SDL_EventEnabled
func __SDL_EventEnabled(int32) int32

//go:wasmimport sdl3 SDL_RegisterEvents
func __SDL_RegisterEvents(int32) int32

//go:wasmimport sdl3 SDL_GetWindowFromEvent
func __SDL_GetWindowFromEvent(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetEventDescription
func __SDL_GetEventDescription(uintptr, uintptr, int32) int32



func __gowrap__SDL_PumpEvents() {
	__SDL_PumpEvents()
}

func __gowrap__SDL_PeepEvents(events *Event, numevents int32, action EventAction, minType uint32, maxType uint32) (__result int32) {
	__result = (int32)(__SDL_PeepEvents(uintptr(unsafe.Pointer(events)), *(*int32)(unsafe.Pointer(&numevents)), *(*int32)(unsafe.Pointer(&action)), *(*int32)(unsafe.Pointer(&minType)), *(*int32)(unsafe.Pointer(&maxType))))
	return
}

func __gowrap__SDL_HasEvent(eventType EventType) (__result bool) {
	__result = (bool)(__SDL_HasEvent(*(*int32)(unsafe.Pointer(&eventType))) != 0)
	return
}

func __gowrap__SDL_HasEvents(minType EventType, maxType EventType) (__result bool) {
	__result = (bool)(__SDL_HasEvents(*(*int32)(unsafe.Pointer(&minType)), *(*int32)(unsafe.Pointer(&maxType))) != 0)
	return
}

func __gowrap__SDL_FlushEvent(eventType EventType) {
	__SDL_FlushEvent(*(*int32)(unsafe.Pointer(&eventType)))
}

func __gowrap__SDL_FlushEvents(minType EventType, maxType EventType) {
	__SDL_FlushEvents(*(*int32)(unsafe.Pointer(&minType)), *(*int32)(unsafe.Pointer(&maxType)))
}

func __gowrap__SDL_PollEvent(event *Event) (__result bool) {
	__result = (bool)(__SDL_PollEvent(uintptr(unsafe.Pointer(event))) != 0)
	return
}

func __gowrap__SDL_WaitEvent(event *Event) (__result bool) {
	__result = (bool)(__SDL_WaitEvent(uintptr(unsafe.Pointer(event))) != 0)
	return
}

func __gowrap__SDL_WaitEventTimeout(event *Event, timeoutMS int32) (__result bool) {
	__result = (bool)(__SDL_WaitEventTimeout(uintptr(unsafe.Pointer(event)), *(*int32)(unsafe.Pointer(&timeoutMS))) != 0)
	return
}

func __gowrap__SDL_PushEvent(event *Event) (__result bool) {
	__result = (bool)(__SDL_PushEvent(uintptr(unsafe.Pointer(event))) != 0)
	return
}

func __gowrap__SDL_SetEventFilter(filter EventFilter, userdata uintptr) {
	__SDL_SetEventFilter(0 /* TODO: callbacks */, uintptr(unsafe.Pointer(userdata)))
}

func __gowrap__SDL_GetEventFilter(filter EventFilter, userdata *uintptr) (__result bool) {
	__result = (bool)(__SDL_GetEventFilter(0 /* TODO: callbacks */, uintptr(unsafe.Pointer(userdata))) != 0)
	return
}

func __gowrap__SDL_AddEventWatch(filter EventFilter, userdata uintptr) (__result bool) {
	__result = (bool)(__SDL_AddEventWatch(0 /* TODO: callbacks */, uintptr(unsafe.Pointer(userdata))) != 0)
	return
}

func __gowrap__SDL_RemoveEventWatch(filter EventFilter, userdata uintptr) {
	__SDL_RemoveEventWatch(0 /* TODO: callbacks */, uintptr(unsafe.Pointer(userdata)))
}

func __gowrap__SDL_FilterEvents(filter EventFilter, userdata uintptr) {
	__SDL_FilterEvents(0 /* TODO: callbacks */, uintptr(unsafe.Pointer(userdata)))
}

func __gowrap__SDL_SetEventEnabled(eventType EventType, enabled bool) {
	__SDL_SetEventEnabled(*(*int32)(unsafe.Pointer(&eventType)), *(*int32)(unsafe.Pointer(&enabled)))
}

func __gowrap__SDL_EventEnabled(eventType EventType) (__result bool) {
	__result = (bool)(__SDL_EventEnabled(*(*int32)(unsafe.Pointer(&eventType))) != 0)
	return
}

func __gowrap__SDL_RegisterEvents(numevents int32) (__result uint32) {
	__result = (uint32)(__SDL_RegisterEvents(*(*int32)(unsafe.Pointer(&numevents))))
	return
}

func __gowrap__SDL_GetWindowFromEvent(event *Event) (__result Window) {
	__result = (Window)(unsafe.Pointer(__SDL_GetWindowFromEvent(uintptr(unsafe.Pointer(event)))))
	return
}

func __gowrap__SDL_GetEventDescription(event *Event, buf []byte, buflen int32) (__result int32) {
	__result = (int32)(__SDL_GetEventDescription(uintptr(unsafe.Pointer(event)), uintptr(unsafe.Pointer(&buf[0])), *(*int32)(unsafe.Pointer(&buflen))))
	runtime.KeepAlive(buf)
	return
}

 
func init() {
	PumpEvents = __gowrap__SDL_PumpEvents
	PeepEvents = __gowrap__SDL_PeepEvents
	HasEvent = __gowrap__SDL_HasEvent
	HasEvents = __gowrap__SDL_HasEvents
	FlushEvent = __gowrap__SDL_FlushEvent
	FlushEvents = __gowrap__SDL_FlushEvents
	PollEvent = __gowrap__SDL_PollEvent
	WaitEvent = __gowrap__SDL_WaitEvent
	WaitEventTimeout = __gowrap__SDL_WaitEventTimeout
	PushEvent = __gowrap__SDL_PushEvent
	SetEventFilter = __gowrap__SDL_SetEventFilter
	GetEventFilter = __gowrap__SDL_GetEventFilter
	AddEventWatch = __gowrap__SDL_AddEventWatch
	RemoveEventWatch = __gowrap__SDL_RemoveEventWatch
	FilterEvents = __gowrap__SDL_FilterEvents
	SetEventEnabled = __gowrap__SDL_SetEventEnabled
	EventEnabled = __gowrap__SDL_EventEnabled
	RegisterEvents = __gowrap__SDL_RegisterEvents
	GetWindowFromEvent = __gowrap__SDL_GetWindowFromEvent
	GetEventDescription = __gowrap__SDL_GetEventDescription
}
