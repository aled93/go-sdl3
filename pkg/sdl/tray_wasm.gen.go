//go:build wasm

package sdl

import (
  "runtime"
  "unsafe"
)



//go:wasmimport sdl3 SDL_CreateTray
func __SDL_CreateTray(uintptr, uintptr) uintptr

//go:wasmimport sdl3 SDL_SetTrayIcon
func __SDL_SetTrayIcon(uintptr, uintptr)

//go:wasmimport sdl3 SDL_SetTrayTooltip
func __SDL_SetTrayTooltip(uintptr, uintptr)

//go:wasmimport sdl3 SDL_CreateTrayMenu
func __SDL_CreateTrayMenu(uintptr) uintptr

//go:wasmimport sdl3 SDL_CreateTraySubmenu
func __SDL_CreateTraySubmenu(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetTrayMenu
func __SDL_GetTrayMenu(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetTraySubmenu
func __SDL_GetTraySubmenu(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetTrayEntries
func __SDL_GetTrayEntries(uintptr, uintptr) uintptr

//go:wasmimport sdl3 SDL_RemoveTrayEntry
func __SDL_RemoveTrayEntry(uintptr)

//go:wasmimport sdl3 SDL_InsertTrayEntryAt
func __SDL_InsertTrayEntryAt(uintptr, int32, uintptr, int32) uintptr

//go:wasmimport sdl3 SDL_SetTrayEntryLabel
func __SDL_SetTrayEntryLabel(uintptr, uintptr)

//go:wasmimport sdl3 SDL_GetTrayEntryLabel
func __SDL_GetTrayEntryLabel(uintptr) uintptr

//go:wasmimport sdl3 SDL_SetTrayEntryChecked
func __SDL_SetTrayEntryChecked(uintptr, int32)

//go:wasmimport sdl3 SDL_GetTrayEntryChecked
func __SDL_GetTrayEntryChecked(uintptr) int32

//go:wasmimport sdl3 SDL_SetTrayEntryEnabled
func __SDL_SetTrayEntryEnabled(uintptr, int32)

//go:wasmimport sdl3 SDL_GetTrayEntryEnabled
func __SDL_GetTrayEntryEnabled(uintptr) int32

//go:wasmimport sdl3 SDL_SetTrayEntryCallback
func __SDL_SetTrayEntryCallback(uintptr, uintptr, uintptr)

//go:wasmimport sdl3 SDL_ClickTrayEntry
func __SDL_ClickTrayEntry(uintptr)

//go:wasmimport sdl3 SDL_DestroyTray
func __SDL_DestroyTray(uintptr)

//go:wasmimport sdl3 SDL_GetTrayEntryParent
func __SDL_GetTrayEntryParent(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetTrayMenuParentEntry
func __SDL_GetTrayMenuParentEntry(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetTrayMenuParentTray
func __SDL_GetTrayMenuParentTray(uintptr) uintptr

//go:wasmimport sdl3 SDL_UpdateTrays
func __SDL_UpdateTrays()



func __gowrap__SDL_CreateTray(icon *Surface, tooltip string) (__result Tray) {
	__result = (Tray)(unsafe.Pointer(__SDL_CreateTray(uintptr(unsafe.Pointer(icon)), ((*[2]uintptr)(unsafe.Pointer(&tooltip)))[0])))
	runtime.KeepAlive(tooltip)
	return
}

func __gowrap__SDL_SetTrayIcon(tray Tray, icon *Surface) {
	__SDL_SetTrayIcon(uintptr(unsafe.Pointer(tray)), uintptr(unsafe.Pointer(icon)))
}

func __gowrap__SDL_SetTrayTooltip(tray Tray, tooltip string) {
	__SDL_SetTrayTooltip(uintptr(unsafe.Pointer(tray)), ((*[2]uintptr)(unsafe.Pointer(&tooltip)))[0])
	runtime.KeepAlive(tooltip)
}

func __gowrap__SDL_CreateTrayMenu(tray Tray) (__result TrayMenu) {
	__result = (TrayMenu)(unsafe.Pointer(__SDL_CreateTrayMenu(uintptr(unsafe.Pointer(tray)))))
	return
}

func __gowrap__SDL_CreateTraySubmenu(entry TrayEntry) (__result TrayMenu) {
	__result = (TrayMenu)(unsafe.Pointer(__SDL_CreateTraySubmenu(uintptr(unsafe.Pointer(entry)))))
	return
}

func __gowrap__SDL_GetTrayMenu(tray Tray) (__result TrayMenu) {
	__result = (TrayMenu)(unsafe.Pointer(__SDL_GetTrayMenu(uintptr(unsafe.Pointer(tray)))))
	return
}

func __gowrap__SDL_GetTraySubmenu(entry TrayEntry) (__result TrayMenu) {
	__result = (TrayMenu)(unsafe.Pointer(__SDL_GetTraySubmenu(uintptr(unsafe.Pointer(entry)))))
	return
}

func __gowrap__SDL_GetTrayEntries(menu TrayMenu, count *int32) (__result *TrayEntry) {
	__result = (*TrayEntry)(unsafe.Pointer(__SDL_GetTrayEntries(uintptr(unsafe.Pointer(menu)), uintptr(unsafe.Pointer(count)))))
	return
}

func __gowrap__SDL_RemoveTrayEntry(entry TrayEntry) {
	__SDL_RemoveTrayEntry(uintptr(unsafe.Pointer(entry)))
}

func __gowrap__SDL_InsertTrayEntryAt(menu TrayMenu, pos int32, label string, flags TrayEntryFlags) (__result TrayEntry) {
	__result = (TrayEntry)(unsafe.Pointer(__SDL_InsertTrayEntryAt(uintptr(unsafe.Pointer(menu)), *(*int32)(unsafe.Pointer(&pos)), ((*[2]uintptr)(unsafe.Pointer(&label)))[0], *(*int32)(unsafe.Pointer(&flags)))))
	runtime.KeepAlive(label)
	return
}

func __gowrap__SDL_SetTrayEntryLabel(entry TrayEntry, label string) {
	__SDL_SetTrayEntryLabel(uintptr(unsafe.Pointer(entry)), ((*[2]uintptr)(unsafe.Pointer(&label)))[0])
	runtime.KeepAlive(label)
}

func __gowrap__SDL_GetTrayEntryLabel(entry TrayEntry) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetTrayEntryLabel(uintptr(unsafe.Pointer(entry)))).Collect()))
	return
}

func __gowrap__SDL_SetTrayEntryChecked(entry TrayEntry, checked bool) {
	__SDL_SetTrayEntryChecked(uintptr(unsafe.Pointer(entry)), *(*int32)(unsafe.Pointer(&checked)))
}

func __gowrap__SDL_GetTrayEntryChecked(entry TrayEntry) (__result bool) {
	__result = (bool)(__SDL_GetTrayEntryChecked(uintptr(unsafe.Pointer(entry))) != 0)
	return
}

func __gowrap__SDL_SetTrayEntryEnabled(entry TrayEntry, enabled bool) {
	__SDL_SetTrayEntryEnabled(uintptr(unsafe.Pointer(entry)), *(*int32)(unsafe.Pointer(&enabled)))
}

func __gowrap__SDL_GetTrayEntryEnabled(entry TrayEntry) (__result bool) {
	__result = (bool)(__SDL_GetTrayEntryEnabled(uintptr(unsafe.Pointer(entry))) != 0)
	return
}

func __gowrap__SDL_SetTrayEntryCallback(entry TrayEntry, callback TrayCallback, userdata uintptr) {
	__SDL_SetTrayEntryCallback(uintptr(unsafe.Pointer(entry)), 0 /* TODO: callbacks */, uintptr(unsafe.Pointer(userdata)))
}

func __gowrap__SDL_ClickTrayEntry(entry TrayEntry) {
	__SDL_ClickTrayEntry(uintptr(unsafe.Pointer(entry)))
}

func __gowrap__SDL_DestroyTray(tray Tray) {
	__SDL_DestroyTray(uintptr(unsafe.Pointer(tray)))
}

func __gowrap__SDL_GetTrayEntryParent(entry TrayEntry) (__result TrayMenu) {
	__result = (TrayMenu)(unsafe.Pointer(__SDL_GetTrayEntryParent(uintptr(unsafe.Pointer(entry)))))
	return
}

func __gowrap__SDL_GetTrayMenuParentEntry(menu TrayMenu) (__result TrayEntry) {
	__result = (TrayEntry)(unsafe.Pointer(__SDL_GetTrayMenuParentEntry(uintptr(unsafe.Pointer(menu)))))
	return
}

func __gowrap__SDL_GetTrayMenuParentTray(menu TrayMenu) (__result Tray) {
	__result = (Tray)(unsafe.Pointer(__SDL_GetTrayMenuParentTray(uintptr(unsafe.Pointer(menu)))))
	return
}

func __gowrap__SDL_UpdateTrays() {
	__SDL_UpdateTrays()
}

 
func init() {
	CreateTray = __gowrap__SDL_CreateTray
	SetTrayIcon = __gowrap__SDL_SetTrayIcon
	SetTrayTooltip = __gowrap__SDL_SetTrayTooltip
	CreateTrayMenu = __gowrap__SDL_CreateTrayMenu
	CreateTraySubmenu = __gowrap__SDL_CreateTraySubmenu
	GetTrayMenu = __gowrap__SDL_GetTrayMenu
	GetTraySubmenu = __gowrap__SDL_GetTraySubmenu
	GetTrayEntries = __gowrap__SDL_GetTrayEntries
	RemoveTrayEntry = __gowrap__SDL_RemoveTrayEntry
	InsertTrayEntryAt = __gowrap__SDL_InsertTrayEntryAt
	SetTrayEntryLabel = __gowrap__SDL_SetTrayEntryLabel
	GetTrayEntryLabel = __gowrap__SDL_GetTrayEntryLabel
	SetTrayEntryChecked = __gowrap__SDL_SetTrayEntryChecked
	GetTrayEntryChecked = __gowrap__SDL_GetTrayEntryChecked
	SetTrayEntryEnabled = __gowrap__SDL_SetTrayEntryEnabled
	GetTrayEntryEnabled = __gowrap__SDL_GetTrayEntryEnabled
	SetTrayEntryCallback = __gowrap__SDL_SetTrayEntryCallback
	ClickTrayEntry = __gowrap__SDL_ClickTrayEntry
	DestroyTray = __gowrap__SDL_DestroyTray
	GetTrayEntryParent = __gowrap__SDL_GetTrayEntryParent
	GetTrayMenuParentEntry = __gowrap__SDL_GetTrayMenuParentEntry
	GetTrayMenuParentTray = __gowrap__SDL_GetTrayMenuParentTray
	UpdateTrays = __gowrap__SDL_UpdateTrays
}
