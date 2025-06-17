//go:build wasm

package sdl

import (
  "runtime"
  "unsafe"
)



//go:wasmimport sdl3 SDL_GetHaptics
func __SDL_GetHaptics(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetHapticNameForID
func __SDL_GetHapticNameForID(int32) uintptr

//go:wasmimport sdl3 SDL_OpenHaptic
func __SDL_OpenHaptic(int32) uintptr

//go:wasmimport sdl3 SDL_GetHapticFromID
func __SDL_GetHapticFromID(int32) uintptr

//go:wasmimport sdl3 SDL_GetHapticID
func __SDL_GetHapticID(uintptr) int32

//go:wasmimport sdl3 SDL_GetHapticName
func __SDL_GetHapticName(uintptr) uintptr

//go:wasmimport sdl3 SDL_IsMouseHaptic
func __SDL_IsMouseHaptic() int32

//go:wasmimport sdl3 SDL_OpenHapticFromMouse
func __SDL_OpenHapticFromMouse() uintptr

//go:wasmimport sdl3 SDL_IsJoystickHaptic
func __SDL_IsJoystickHaptic(uintptr) int32

//go:wasmimport sdl3 SDL_OpenHapticFromJoystick
func __SDL_OpenHapticFromJoystick(uintptr) uintptr

//go:wasmimport sdl3 SDL_CloseHaptic
func __SDL_CloseHaptic(uintptr)

//go:wasmimport sdl3 SDL_GetMaxHapticEffects
func __SDL_GetMaxHapticEffects(uintptr) int32

//go:wasmimport sdl3 SDL_GetMaxHapticEffectsPlaying
func __SDL_GetMaxHapticEffectsPlaying(uintptr) int32

//go:wasmimport sdl3 SDL_GetHapticFeatures
func __SDL_GetHapticFeatures(uintptr) int32

//go:wasmimport sdl3 SDL_GetNumHapticAxes
func __SDL_GetNumHapticAxes(uintptr) int32

//go:wasmimport sdl3 SDL_HapticEffectSupported
func __SDL_HapticEffectSupported(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_CreateHapticEffect
func __SDL_CreateHapticEffect(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_UpdateHapticEffect
func __SDL_UpdateHapticEffect(uintptr, int32, uintptr) int32

//go:wasmimport sdl3 SDL_RunHapticEffect
func __SDL_RunHapticEffect(uintptr, int32, int32) int32

//go:wasmimport sdl3 SDL_StopHapticEffect
func __SDL_StopHapticEffect(uintptr, int32) int32

//go:wasmimport sdl3 SDL_DestroyHapticEffect
func __SDL_DestroyHapticEffect(uintptr, int32)

//go:wasmimport sdl3 SDL_GetHapticEffectStatus
func __SDL_GetHapticEffectStatus(uintptr, int32) int32

//go:wasmimport sdl3 SDL_SetHapticGain
func __SDL_SetHapticGain(uintptr, int32) int32

//go:wasmimport sdl3 SDL_SetHapticAutocenter
func __SDL_SetHapticAutocenter(uintptr, int32) int32

//go:wasmimport sdl3 SDL_PauseHaptic
func __SDL_PauseHaptic(uintptr) int32

//go:wasmimport sdl3 SDL_ResumeHaptic
func __SDL_ResumeHaptic(uintptr) int32

//go:wasmimport sdl3 SDL_StopHapticEffects
func __SDL_StopHapticEffects(uintptr) int32

//go:wasmimport sdl3 SDL_HapticRumbleSupported
func __SDL_HapticRumbleSupported(uintptr) int32

//go:wasmimport sdl3 SDL_InitHapticRumble
func __SDL_InitHapticRumble(uintptr) int32

//go:wasmimport sdl3 SDL_PlayHapticRumble
func __SDL_PlayHapticRumble(uintptr, float32, int32) int32

//go:wasmimport sdl3 SDL_StopHapticRumble
func __SDL_StopHapticRumble(uintptr) int32



func __gowrap__SDL_GetHaptics(count *int32) (__result *HapticID) {
	__result = (*HapticID)(unsafe.Pointer(__SDL_GetHaptics(uintptr(unsafe.Pointer(count)))))
	return
}

func __gowrap__SDL_GetHapticNameForID(instance_id HapticID) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetHapticNameForID(*(*int32)(unsafe.Pointer(&instance_id)))).Collect()))
	return
}

func __gowrap__SDL_OpenHaptic(instance_id HapticID) (__result Haptic) {
	__result = (Haptic)(unsafe.Pointer(__SDL_OpenHaptic(*(*int32)(unsafe.Pointer(&instance_id)))))
	return
}

func __gowrap__SDL_GetHapticFromID(instance_id HapticID) (__result Haptic) {
	__result = (Haptic)(unsafe.Pointer(__SDL_GetHapticFromID(*(*int32)(unsafe.Pointer(&instance_id)))))
	return
}

func __gowrap__SDL_GetHapticID(haptic Haptic) (__result HapticID) {
	__result = (HapticID)(__SDL_GetHapticID(uintptr(unsafe.Pointer(haptic))))
	return
}

func __gowrap__SDL_GetHapticName(haptic Haptic) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetHapticName(uintptr(unsafe.Pointer(haptic)))).Collect()))
	return
}

func __gowrap__SDL_IsMouseHaptic() (__result bool) {
	__result = (bool)(__SDL_IsMouseHaptic() != 0)
	return
}

func __gowrap__SDL_OpenHapticFromMouse() (__result Haptic) {
	__result = (Haptic)(unsafe.Pointer(__SDL_OpenHapticFromMouse()))
	return
}

func __gowrap__SDL_IsJoystickHaptic(joystick Joystick) (__result bool) {
	__result = (bool)(__SDL_IsJoystickHaptic(uintptr(unsafe.Pointer(joystick))) != 0)
	return
}

func __gowrap__SDL_OpenHapticFromJoystick(joystick Joystick) (__result Haptic) {
	__result = (Haptic)(unsafe.Pointer(__SDL_OpenHapticFromJoystick(uintptr(unsafe.Pointer(joystick)))))
	return
}

func __gowrap__SDL_CloseHaptic(haptic Haptic) {
	__SDL_CloseHaptic(uintptr(unsafe.Pointer(haptic)))
}

func __gowrap__SDL_GetMaxHapticEffects(haptic Haptic) (__result int32) {
	__result = (int32)(__SDL_GetMaxHapticEffects(uintptr(unsafe.Pointer(haptic))))
	return
}

func __gowrap__SDL_GetMaxHapticEffectsPlaying(haptic Haptic) (__result int32) {
	__result = (int32)(__SDL_GetMaxHapticEffectsPlaying(uintptr(unsafe.Pointer(haptic))))
	return
}

func __gowrap__SDL_GetHapticFeatures(haptic Haptic) (__result uint32) {
	__result = (uint32)(__SDL_GetHapticFeatures(uintptr(unsafe.Pointer(haptic))))
	return
}

func __gowrap__SDL_GetNumHapticAxes(haptic Haptic) (__result int32) {
	__result = (int32)(__SDL_GetNumHapticAxes(uintptr(unsafe.Pointer(haptic))))
	return
}

func __gowrap__SDL_HapticEffectSupported(haptic Haptic, effect []HapticEffect) (__result bool) {
	__result = (bool)(__SDL_HapticEffectSupported(uintptr(unsafe.Pointer(haptic)), uintptr(unsafe.Pointer(&effect[0]))) != 0)
	runtime.KeepAlive(effect)
	return
}

func __gowrap__SDL_CreateHapticEffect(haptic Haptic, effect []HapticEffect) (__result HapticEffectID) {
	__result = (HapticEffectID)(__SDL_CreateHapticEffect(uintptr(unsafe.Pointer(haptic)), uintptr(unsafe.Pointer(&effect[0]))))
	runtime.KeepAlive(effect)
	return
}

func __gowrap__SDL_UpdateHapticEffect(haptic Haptic, effect HapticEffectID, data []HapticEffect) (__result bool) {
	__result = (bool)(__SDL_UpdateHapticEffect(uintptr(unsafe.Pointer(haptic)), *(*int32)(unsafe.Pointer(&effect)), uintptr(unsafe.Pointer(&data[0]))) != 0)
	runtime.KeepAlive(data)
	return
}

func __gowrap__SDL_RunHapticEffect(haptic Haptic, effect HapticEffectID, iterations uint32) (__result bool) {
	__result = (bool)(__SDL_RunHapticEffect(uintptr(unsafe.Pointer(haptic)), *(*int32)(unsafe.Pointer(&effect)), *(*int32)(unsafe.Pointer(&iterations))) != 0)
	return
}

func __gowrap__SDL_StopHapticEffect(haptic Haptic, effect HapticEffectID) (__result bool) {
	__result = (bool)(__SDL_StopHapticEffect(uintptr(unsafe.Pointer(haptic)), *(*int32)(unsafe.Pointer(&effect))) != 0)
	return
}

func __gowrap__SDL_DestroyHapticEffect(haptic Haptic, effect HapticEffectID) {
	__SDL_DestroyHapticEffect(uintptr(unsafe.Pointer(haptic)), *(*int32)(unsafe.Pointer(&effect)))
}

func __gowrap__SDL_GetHapticEffectStatus(haptic Haptic, effect HapticEffectID) (__result bool) {
	__result = (bool)(__SDL_GetHapticEffectStatus(uintptr(unsafe.Pointer(haptic)), *(*int32)(unsafe.Pointer(&effect))) != 0)
	return
}

func __gowrap__SDL_SetHapticGain(haptic Haptic, gain int32) (__result bool) {
	__result = (bool)(__SDL_SetHapticGain(uintptr(unsafe.Pointer(haptic)), *(*int32)(unsafe.Pointer(&gain))) != 0)
	return
}

func __gowrap__SDL_SetHapticAutocenter(haptic Haptic, autocenter int32) (__result bool) {
	__result = (bool)(__SDL_SetHapticAutocenter(uintptr(unsafe.Pointer(haptic)), *(*int32)(unsafe.Pointer(&autocenter))) != 0)
	return
}

func __gowrap__SDL_PauseHaptic(haptic Haptic) (__result bool) {
	__result = (bool)(__SDL_PauseHaptic(uintptr(unsafe.Pointer(haptic))) != 0)
	return
}

func __gowrap__SDL_ResumeHaptic(haptic Haptic) (__result bool) {
	__result = (bool)(__SDL_ResumeHaptic(uintptr(unsafe.Pointer(haptic))) != 0)
	return
}

func __gowrap__SDL_StopHapticEffects(haptic Haptic) (__result bool) {
	__result = (bool)(__SDL_StopHapticEffects(uintptr(unsafe.Pointer(haptic))) != 0)
	return
}

func __gowrap__SDL_HapticRumbleSupported(haptic Haptic) (__result bool) {
	__result = (bool)(__SDL_HapticRumbleSupported(uintptr(unsafe.Pointer(haptic))) != 0)
	return
}

func __gowrap__SDL_InitHapticRumble(haptic Haptic) (__result bool) {
	__result = (bool)(__SDL_InitHapticRumble(uintptr(unsafe.Pointer(haptic))) != 0)
	return
}

func __gowrap__SDL_PlayHapticRumble(haptic Haptic, strength float32, length uint32) (__result bool) {
	__result = (bool)(__SDL_PlayHapticRumble(uintptr(unsafe.Pointer(haptic)), *(*float32)(unsafe.Pointer(&strength)), *(*int32)(unsafe.Pointer(&length))) != 0)
	return
}

func __gowrap__SDL_StopHapticRumble(haptic Haptic) (__result bool) {
	__result = (bool)(__SDL_StopHapticRumble(uintptr(unsafe.Pointer(haptic))) != 0)
	return
}

 
func init() {
	GetHaptics = __gowrap__SDL_GetHaptics
	GetHapticNameForID = __gowrap__SDL_GetHapticNameForID
	OpenHaptic = __gowrap__SDL_OpenHaptic
	GetHapticFromID = __gowrap__SDL_GetHapticFromID
	GetHapticID = __gowrap__SDL_GetHapticID
	GetHapticName = __gowrap__SDL_GetHapticName
	IsMouseHaptic = __gowrap__SDL_IsMouseHaptic
	OpenHapticFromMouse = __gowrap__SDL_OpenHapticFromMouse
	IsJoystickHaptic = __gowrap__SDL_IsJoystickHaptic
	OpenHapticFromJoystick = __gowrap__SDL_OpenHapticFromJoystick
	CloseHaptic = __gowrap__SDL_CloseHaptic
	GetMaxHapticEffects = __gowrap__SDL_GetMaxHapticEffects
	GetMaxHapticEffectsPlaying = __gowrap__SDL_GetMaxHapticEffectsPlaying
	GetHapticFeatures = __gowrap__SDL_GetHapticFeatures
	GetNumHapticAxes = __gowrap__SDL_GetNumHapticAxes
	HapticEffectSupported = __gowrap__SDL_HapticEffectSupported
	CreateHapticEffect = __gowrap__SDL_CreateHapticEffect
	UpdateHapticEffect = __gowrap__SDL_UpdateHapticEffect
	RunHapticEffect = __gowrap__SDL_RunHapticEffect
	StopHapticEffect = __gowrap__SDL_StopHapticEffect
	DestroyHapticEffect = __gowrap__SDL_DestroyHapticEffect
	GetHapticEffectStatus = __gowrap__SDL_GetHapticEffectStatus
	SetHapticGain = __gowrap__SDL_SetHapticGain
	SetHapticAutocenter = __gowrap__SDL_SetHapticAutocenter
	PauseHaptic = __gowrap__SDL_PauseHaptic
	ResumeHaptic = __gowrap__SDL_ResumeHaptic
	StopHapticEffects = __gowrap__SDL_StopHapticEffects
	HapticRumbleSupported = __gowrap__SDL_HapticRumbleSupported
	InitHapticRumble = __gowrap__SDL_InitHapticRumble
	PlayHapticRumble = __gowrap__SDL_PlayHapticRumble
	StopHapticRumble = __gowrap__SDL_StopHapticRumble
}
