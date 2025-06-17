//go:build wasm

package sdl

import (
  "runtime"
  "unsafe"
)



//go:wasmimport sdl3 SDL_LockJoysticks
func __SDL_LockJoysticks()

//go:wasmimport sdl3 SDL_UnlockJoysticks
func __SDL_UnlockJoysticks()

//go:wasmimport sdl3 SDL_HasJoystick
func __SDL_HasJoystick() int32

//go:wasmimport sdl3 SDL_GetJoysticks
func __SDL_GetJoysticks(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetJoystickNameForID
func __SDL_GetJoystickNameForID(int32) uintptr

//go:wasmimport sdl3 SDL_GetJoystickPathForID
func __SDL_GetJoystickPathForID(int32) uintptr

//go:wasmimport sdl3 SDL_GetJoystickPlayerIndexForID
func __SDL_GetJoystickPlayerIndexForID(int32) int32

//go:wasmimport sdl3 SDL_GetJoystickGUIDForID
func __SDL_GetJoystickGUIDForID(uintptr, int32)

//go:wasmimport sdl3 SDL_GetJoystickVendorForID
func __SDL_GetJoystickVendorForID(int32) int32

//go:wasmimport sdl3 SDL_GetJoystickProductForID
func __SDL_GetJoystickProductForID(int32) int32

//go:wasmimport sdl3 SDL_GetJoystickProductVersionForID
func __SDL_GetJoystickProductVersionForID(int32) int32

//go:wasmimport sdl3 SDL_GetJoystickTypeForID
func __SDL_GetJoystickTypeForID(int32) int32

//go:wasmimport sdl3 SDL_OpenJoystick
func __SDL_OpenJoystick(int32) uintptr

//go:wasmimport sdl3 SDL_GetJoystickFromID
func __SDL_GetJoystickFromID(int32) uintptr

//go:wasmimport sdl3 SDL_GetJoystickFromPlayerIndex
func __SDL_GetJoystickFromPlayerIndex(int32) uintptr

//go:wasmimport sdl3 SDL_AttachVirtualJoystick
func __SDL_AttachVirtualJoystick(uintptr) int32

//go:wasmimport sdl3 SDL_DetachVirtualJoystick
func __SDL_DetachVirtualJoystick(int32) int32

//go:wasmimport sdl3 SDL_IsJoystickVirtual
func __SDL_IsJoystickVirtual(int32) int32

//go:wasmimport sdl3 SDL_SetJoystickVirtualAxis
func __SDL_SetJoystickVirtualAxis(uintptr, int32, int32) int32

//go:wasmimport sdl3 SDL_SetJoystickVirtualBall
func __SDL_SetJoystickVirtualBall(uintptr, int32, int32, int32) int32

//go:wasmimport sdl3 SDL_SetJoystickVirtualButton
func __SDL_SetJoystickVirtualButton(uintptr, int32, int32) int32

//go:wasmimport sdl3 SDL_SetJoystickVirtualHat
func __SDL_SetJoystickVirtualHat(uintptr, int32, int32) int32

//go:wasmimport sdl3 SDL_SetJoystickVirtualTouchpad
func __SDL_SetJoystickVirtualTouchpad(uintptr, int32, int32, int32, float32, float32, float32) int32

//go:wasmimport sdl3 SDL_SendJoystickVirtualSensorData
func __SDL_SendJoystickVirtualSensorData(uintptr, int32, int64, uintptr, int32) int32

//go:wasmimport sdl3 SDL_GetJoystickProperties
func __SDL_GetJoystickProperties(uintptr) int32

//go:wasmimport sdl3 SDL_GetJoystickName
func __SDL_GetJoystickName(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetJoystickPath
func __SDL_GetJoystickPath(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetJoystickPlayerIndex
func __SDL_GetJoystickPlayerIndex(uintptr) int32

//go:wasmimport sdl3 SDL_SetJoystickPlayerIndex
func __SDL_SetJoystickPlayerIndex(uintptr, int32) int32

//go:wasmimport sdl3 SDL_GetJoystickGUID
func __SDL_GetJoystickGUID(uintptr, uintptr)

//go:wasmimport sdl3 SDL_GetJoystickVendor
func __SDL_GetJoystickVendor(uintptr) int32

//go:wasmimport sdl3 SDL_GetJoystickProduct
func __SDL_GetJoystickProduct(uintptr) int32

//go:wasmimport sdl3 SDL_GetJoystickProductVersion
func __SDL_GetJoystickProductVersion(uintptr) int32

//go:wasmimport sdl3 SDL_GetJoystickFirmwareVersion
func __SDL_GetJoystickFirmwareVersion(uintptr) int32

//go:wasmimport sdl3 SDL_GetJoystickSerial
func __SDL_GetJoystickSerial(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetJoystickType
func __SDL_GetJoystickType(uintptr) int32

//go:wasmimport sdl3 SDL_GetJoystickGUIDInfo
func __SDL_GetJoystickGUIDInfo(uintptr, uintptr, uintptr, uintptr, uintptr)

//go:wasmimport sdl3 SDL_JoystickConnected
func __SDL_JoystickConnected(uintptr) int32

//go:wasmimport sdl3 SDL_GetJoystickID
func __SDL_GetJoystickID(uintptr) int32

//go:wasmimport sdl3 SDL_GetNumJoystickAxes
func __SDL_GetNumJoystickAxes(uintptr) int32

//go:wasmimport sdl3 SDL_GetNumJoystickBalls
func __SDL_GetNumJoystickBalls(uintptr) int32

//go:wasmimport sdl3 SDL_GetNumJoystickHats
func __SDL_GetNumJoystickHats(uintptr) int32

//go:wasmimport sdl3 SDL_GetNumJoystickButtons
func __SDL_GetNumJoystickButtons(uintptr) int32

//go:wasmimport sdl3 SDL_SetJoystickEventsEnabled
func __SDL_SetJoystickEventsEnabled(int32)

//go:wasmimport sdl3 SDL_JoystickEventsEnabled
func __SDL_JoystickEventsEnabled() int32

//go:wasmimport sdl3 SDL_UpdateJoysticks
func __SDL_UpdateJoysticks()

//go:wasmimport sdl3 SDL_GetJoystickAxis
func __SDL_GetJoystickAxis(uintptr, int32) int32

//go:wasmimport sdl3 SDL_GetJoystickAxisInitialState
func __SDL_GetJoystickAxisInitialState(uintptr, int32, uintptr) int32

//go:wasmimport sdl3 SDL_GetJoystickBall
func __SDL_GetJoystickBall(uintptr, int32, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetJoystickHat
func __SDL_GetJoystickHat(uintptr, int32) int32

//go:wasmimport sdl3 SDL_GetJoystickButton
func __SDL_GetJoystickButton(uintptr, int32) int32

//go:wasmimport sdl3 SDL_RumbleJoystick
func __SDL_RumbleJoystick(uintptr, int32, int32, int32) int32

//go:wasmimport sdl3 SDL_RumbleJoystickTriggers
func __SDL_RumbleJoystickTriggers(uintptr, int32, int32, int32) int32

//go:wasmimport sdl3 SDL_SetJoystickLED
func __SDL_SetJoystickLED(uintptr, int32, int32, int32) int32

//go:wasmimport sdl3 SDL_SendJoystickEffect
func __SDL_SendJoystickEffect(uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_CloseJoystick
func __SDL_CloseJoystick(uintptr)

//go:wasmimport sdl3 SDL_GetJoystickConnectionState
func __SDL_GetJoystickConnectionState(uintptr) int32

//go:wasmimport sdl3 SDL_GetJoystickPowerInfo
func __SDL_GetJoystickPowerInfo(uintptr, uintptr) int32



func __gowrap__SDL_LockJoysticks() {
	__SDL_LockJoysticks()
}

func __gowrap__SDL_UnlockJoysticks() {
	__SDL_UnlockJoysticks()
}

func __gowrap__SDL_HasJoystick() (__result bool) {
	__result = (bool)(__SDL_HasJoystick() != 0)
	return
}

func __gowrap__SDL_GetJoysticks(count *int32) (__result *JoystickID) {
	__result = (*JoystickID)(unsafe.Pointer(__SDL_GetJoysticks(uintptr(unsafe.Pointer(count)))))
	return
}

func __gowrap__SDL_GetJoystickNameForID(instance_id JoystickID) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetJoystickNameForID(*(*int32)(unsafe.Pointer(&instance_id)))).Collect()))
	return
}

func __gowrap__SDL_GetJoystickPathForID(instance_id JoystickID) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetJoystickPathForID(*(*int32)(unsafe.Pointer(&instance_id)))).Collect()))
	return
}

func __gowrap__SDL_GetJoystickPlayerIndexForID(instance_id JoystickID) (__result int32) {
	__result = (int32)(__SDL_GetJoystickPlayerIndexForID(*(*int32)(unsafe.Pointer(&instance_id))))
	return
}

func __gowrap__SDL_GetJoystickGUIDForID(instance_id JoystickID) (__result GUID) {
	__SDL_GetJoystickGUIDForID(uintptr(unsafe.Pointer(&__result)), *(*int32)(unsafe.Pointer(&instance_id)))
	return
}

func __gowrap__SDL_GetJoystickVendorForID(instance_id JoystickID) (__result uint16) {
	__result = (uint16)(__SDL_GetJoystickVendorForID(*(*int32)(unsafe.Pointer(&instance_id))))
	return
}

func __gowrap__SDL_GetJoystickProductForID(instance_id JoystickID) (__result uint16) {
	__result = (uint16)(__SDL_GetJoystickProductForID(*(*int32)(unsafe.Pointer(&instance_id))))
	return
}

func __gowrap__SDL_GetJoystickProductVersionForID(instance_id JoystickID) (__result uint16) {
	__result = (uint16)(__SDL_GetJoystickProductVersionForID(*(*int32)(unsafe.Pointer(&instance_id))))
	return
}

func __gowrap__SDL_GetJoystickTypeForID(instance_id JoystickID) (__result JoystickType) {
	__result = (JoystickType)(__SDL_GetJoystickTypeForID(*(*int32)(unsafe.Pointer(&instance_id))))
	return
}

func __gowrap__SDL_OpenJoystick(instance_id JoystickID) (__result Joystick) {
	__result = (Joystick)(unsafe.Pointer(__SDL_OpenJoystick(*(*int32)(unsafe.Pointer(&instance_id)))))
	return
}

func __gowrap__SDL_GetJoystickFromID(instance_id JoystickID) (__result Joystick) {
	__result = (Joystick)(unsafe.Pointer(__SDL_GetJoystickFromID(*(*int32)(unsafe.Pointer(&instance_id)))))
	return
}

func __gowrap__SDL_GetJoystickFromPlayerIndex(player_index int32) (__result Joystick) {
	__result = (Joystick)(unsafe.Pointer(__SDL_GetJoystickFromPlayerIndex(*(*int32)(unsafe.Pointer(&player_index)))))
	return
}

func __gowrap__SDL_AttachVirtualJoystick(desc *VirtualJoystickDesc) (__result JoystickID) {
	__result = (JoystickID)(__SDL_AttachVirtualJoystick(uintptr(unsafe.Pointer(desc))))
	return
}

func __gowrap__SDL_DetachVirtualJoystick(instance_id JoystickID) (__result bool) {
	__result = (bool)(__SDL_DetachVirtualJoystick(*(*int32)(unsafe.Pointer(&instance_id))) != 0)
	return
}

func __gowrap__SDL_IsJoystickVirtual(instance_id JoystickID) (__result bool) {
	__result = (bool)(__SDL_IsJoystickVirtual(*(*int32)(unsafe.Pointer(&instance_id))) != 0)
	return
}

func __gowrap__SDL_SetJoystickVirtualAxis(joystick Joystick, axis int32, value int16) (__result bool) {
	__result = (bool)(__SDL_SetJoystickVirtualAxis(uintptr(unsafe.Pointer(joystick)), *(*int32)(unsafe.Pointer(&axis)), *(*int32)(unsafe.Pointer(&value))) != 0)
	return
}

func __gowrap__SDL_SetJoystickVirtualBall(joystick Joystick, ball int32, xrel int16, yrel int16) (__result bool) {
	__result = (bool)(__SDL_SetJoystickVirtualBall(uintptr(unsafe.Pointer(joystick)), *(*int32)(unsafe.Pointer(&ball)), *(*int32)(unsafe.Pointer(&xrel)), *(*int32)(unsafe.Pointer(&yrel))) != 0)
	return
}

func __gowrap__SDL_SetJoystickVirtualButton(joystick Joystick, button int32, down bool) (__result bool) {
	__result = (bool)(__SDL_SetJoystickVirtualButton(uintptr(unsafe.Pointer(joystick)), *(*int32)(unsafe.Pointer(&button)), *(*int32)(unsafe.Pointer(&down))) != 0)
	return
}

func __gowrap__SDL_SetJoystickVirtualHat(joystick Joystick, hat int32, value uint8) (__result bool) {
	__result = (bool)(__SDL_SetJoystickVirtualHat(uintptr(unsafe.Pointer(joystick)), *(*int32)(unsafe.Pointer(&hat)), *(*int32)(unsafe.Pointer(&value))) != 0)
	return
}

func __gowrap__SDL_SetJoystickVirtualTouchpad(joystick Joystick, touchpad int32, finger int32, down bool, x float32, y float32, pressure float32) (__result bool) {
	__result = (bool)(__SDL_SetJoystickVirtualTouchpad(uintptr(unsafe.Pointer(joystick)), *(*int32)(unsafe.Pointer(&touchpad)), *(*int32)(unsafe.Pointer(&finger)), *(*int32)(unsafe.Pointer(&down)), *(*float32)(unsafe.Pointer(&x)), *(*float32)(unsafe.Pointer(&y)), *(*float32)(unsafe.Pointer(&pressure))) != 0)
	return
}

func __gowrap__SDL_SendJoystickVirtualSensorData(joystick Joystick, type_ SensorType, sensor_timestamp uint64, data *float32, num_values int32) (__result bool) {
	__result = (bool)(__SDL_SendJoystickVirtualSensorData(uintptr(unsafe.Pointer(joystick)), *(*int32)(unsafe.Pointer(&type_)), *(*int64)(unsafe.Pointer(&sensor_timestamp)), uintptr(unsafe.Pointer(data)), *(*int32)(unsafe.Pointer(&num_values))) != 0)
	return
}

func __gowrap__SDL_GetJoystickProperties(joystick Joystick) (__result PropertiesID) {
	__result = (PropertiesID)(__SDL_GetJoystickProperties(uintptr(unsafe.Pointer(joystick))))
	return
}

func __gowrap__SDL_GetJoystickName(joystick Joystick) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetJoystickName(uintptr(unsafe.Pointer(joystick)))).Collect()))
	return
}

func __gowrap__SDL_GetJoystickPath(joystick Joystick) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetJoystickPath(uintptr(unsafe.Pointer(joystick)))).Collect()))
	return
}

func __gowrap__SDL_GetJoystickPlayerIndex(joystick Joystick) (__result int32) {
	__result = (int32)(__SDL_GetJoystickPlayerIndex(uintptr(unsafe.Pointer(joystick))))
	return
}

func __gowrap__SDL_SetJoystickPlayerIndex(joystick Joystick, player_index int32) (__result bool) {
	__result = (bool)(__SDL_SetJoystickPlayerIndex(uintptr(unsafe.Pointer(joystick)), *(*int32)(unsafe.Pointer(&player_index))) != 0)
	return
}

func __gowrap__SDL_GetJoystickGUID(joystick Joystick) (__result GUID) {
	__SDL_GetJoystickGUID(uintptr(unsafe.Pointer(&__result)), uintptr(unsafe.Pointer(joystick)))
	return
}

func __gowrap__SDL_GetJoystickVendor(joystick Joystick) (__result uint16) {
	__result = (uint16)(__SDL_GetJoystickVendor(uintptr(unsafe.Pointer(joystick))))
	return
}

func __gowrap__SDL_GetJoystickProduct(joystick Joystick) (__result uint16) {
	__result = (uint16)(__SDL_GetJoystickProduct(uintptr(unsafe.Pointer(joystick))))
	return
}

func __gowrap__SDL_GetJoystickProductVersion(joystick Joystick) (__result uint16) {
	__result = (uint16)(__SDL_GetJoystickProductVersion(uintptr(unsafe.Pointer(joystick))))
	return
}

func __gowrap__SDL_GetJoystickFirmwareVersion(joystick Joystick) (__result uint16) {
	__result = (uint16)(__SDL_GetJoystickFirmwareVersion(uintptr(unsafe.Pointer(joystick))))
	return
}

func __gowrap__SDL_GetJoystickSerial(joystick Joystick) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetJoystickSerial(uintptr(unsafe.Pointer(joystick)))).Collect()))
	return
}

func __gowrap__SDL_GetJoystickType(joystick Joystick) (__result JoystickType) {
	__result = (JoystickType)(__SDL_GetJoystickType(uintptr(unsafe.Pointer(joystick))))
	return
}

func __gowrap__SDL_GetJoystickGUIDInfo(guid *GUID, vendor *uint16, product *uint16, version *uint16, crc16 *uint16) {
	__SDL_GetJoystickGUIDInfo(uintptr(unsafe.Pointer(guid)), uintptr(unsafe.Pointer(vendor)), uintptr(unsafe.Pointer(product)), uintptr(unsafe.Pointer(version)), uintptr(unsafe.Pointer(crc16)))
}

func __gowrap__SDL_JoystickConnected(joystick Joystick) (__result bool) {
	__result = (bool)(__SDL_JoystickConnected(uintptr(unsafe.Pointer(joystick))) != 0)
	return
}

func __gowrap__SDL_GetJoystickID(joystick Joystick) (__result JoystickID) {
	__result = (JoystickID)(__SDL_GetJoystickID(uintptr(unsafe.Pointer(joystick))))
	return
}

func __gowrap__SDL_GetNumJoystickAxes(joystick Joystick) (__result int32) {
	__result = (int32)(__SDL_GetNumJoystickAxes(uintptr(unsafe.Pointer(joystick))))
	return
}

func __gowrap__SDL_GetNumJoystickBalls(joystick Joystick) (__result int32) {
	__result = (int32)(__SDL_GetNumJoystickBalls(uintptr(unsafe.Pointer(joystick))))
	return
}

func __gowrap__SDL_GetNumJoystickHats(joystick Joystick) (__result int32) {
	__result = (int32)(__SDL_GetNumJoystickHats(uintptr(unsafe.Pointer(joystick))))
	return
}

func __gowrap__SDL_GetNumJoystickButtons(joystick Joystick) (__result int32) {
	__result = (int32)(__SDL_GetNumJoystickButtons(uintptr(unsafe.Pointer(joystick))))
	return
}

func __gowrap__SDL_SetJoystickEventsEnabled(enabled bool) {
	__SDL_SetJoystickEventsEnabled(*(*int32)(unsafe.Pointer(&enabled)))
}

func __gowrap__SDL_JoystickEventsEnabled() (__result bool) {
	__result = (bool)(__SDL_JoystickEventsEnabled() != 0)
	return
}

func __gowrap__SDL_UpdateJoysticks() {
	__SDL_UpdateJoysticks()
}

func __gowrap__SDL_GetJoystickAxis(joystick Joystick, axis int32) (__result int16) {
	__result = (int16)(__SDL_GetJoystickAxis(uintptr(unsafe.Pointer(joystick)), *(*int32)(unsafe.Pointer(&axis))))
	return
}

func __gowrap__SDL_GetJoystickAxisInitialState(joystick Joystick, axis int32, state *int16) (__result bool) {
	__result = (bool)(__SDL_GetJoystickAxisInitialState(uintptr(unsafe.Pointer(joystick)), *(*int32)(unsafe.Pointer(&axis)), uintptr(unsafe.Pointer(state))) != 0)
	return
}

func __gowrap__SDL_GetJoystickBall(joystick Joystick, ball int32, dx *int32, dy *int32) (__result bool) {
	__result = (bool)(__SDL_GetJoystickBall(uintptr(unsafe.Pointer(joystick)), *(*int32)(unsafe.Pointer(&ball)), uintptr(unsafe.Pointer(dx)), uintptr(unsafe.Pointer(dy))) != 0)
	return
}

func __gowrap__SDL_GetJoystickHat(joystick Joystick, hat Hat) (__result uint8) {
	__result = (uint8)(__SDL_GetJoystickHat(uintptr(unsafe.Pointer(joystick)), *(*int32)(unsafe.Pointer(&hat))))
	return
}

func __gowrap__SDL_GetJoystickButton(joystick Joystick, button int32) (__result bool) {
	__result = (bool)(__SDL_GetJoystickButton(uintptr(unsafe.Pointer(joystick)), *(*int32)(unsafe.Pointer(&button))) != 0)
	return
}

func __gowrap__SDL_RumbleJoystick(joystick Joystick, low_frequency_rumble uint16, high_frequency_rumble uint16, duration_ms uint32) (__result bool) {
	__result = (bool)(__SDL_RumbleJoystick(uintptr(unsafe.Pointer(joystick)), *(*int32)(unsafe.Pointer(&low_frequency_rumble)), *(*int32)(unsafe.Pointer(&high_frequency_rumble)), *(*int32)(unsafe.Pointer(&duration_ms))) != 0)
	return
}

func __gowrap__SDL_RumbleJoystickTriggers(joystick Joystick, left_rumble uint16, right_rumble uint16, duration_ms uint32) (__result bool) {
	__result = (bool)(__SDL_RumbleJoystickTriggers(uintptr(unsafe.Pointer(joystick)), *(*int32)(unsafe.Pointer(&left_rumble)), *(*int32)(unsafe.Pointer(&right_rumble)), *(*int32)(unsafe.Pointer(&duration_ms))) != 0)
	return
}

func __gowrap__SDL_SetJoystickLED(joystick Joystick, red uint8, green uint8, blue uint8) (__result bool) {
	__result = (bool)(__SDL_SetJoystickLED(uintptr(unsafe.Pointer(joystick)), *(*int32)(unsafe.Pointer(&red)), *(*int32)(unsafe.Pointer(&green)), *(*int32)(unsafe.Pointer(&blue))) != 0)
	return
}

func __gowrap__SDL_SendJoystickEffect(joystick Joystick, data []byte, size int32) (__result bool) {
	__result = (bool)(__SDL_SendJoystickEffect(uintptr(unsafe.Pointer(joystick)), uintptr(unsafe.Pointer(&data[0])), *(*int32)(unsafe.Pointer(&size))) != 0)
	runtime.KeepAlive(data)
	return
}

func __gowrap__SDL_CloseJoystick(joystick Joystick) {
	__SDL_CloseJoystick(uintptr(unsafe.Pointer(joystick)))
}

func __gowrap__SDL_GetJoystickConnectionState(joystick Joystick) (__result JoystickConnectionState) {
	__result = (JoystickConnectionState)(__SDL_GetJoystickConnectionState(uintptr(unsafe.Pointer(joystick))))
	return
}

func __gowrap__SDL_GetJoystickPowerInfo(joystick Joystick, percent *int32) (__result PowerState) {
	__result = (PowerState)(__SDL_GetJoystickPowerInfo(uintptr(unsafe.Pointer(joystick)), uintptr(unsafe.Pointer(percent))))
	return
}

 
func init() {
	LockJoysticks = __gowrap__SDL_LockJoysticks
	UnlockJoysticks = __gowrap__SDL_UnlockJoysticks
	HasJoystick = __gowrap__SDL_HasJoystick
	GetJoysticks = __gowrap__SDL_GetJoysticks
	GetJoystickNameForID = __gowrap__SDL_GetJoystickNameForID
	GetJoystickPathForID = __gowrap__SDL_GetJoystickPathForID
	GetJoystickPlayerIndexForID = __gowrap__SDL_GetJoystickPlayerIndexForID
	GetJoystickGUIDForID = __gowrap__SDL_GetJoystickGUIDForID
	GetJoystickVendorForID = __gowrap__SDL_GetJoystickVendorForID
	GetJoystickProductForID = __gowrap__SDL_GetJoystickProductForID
	GetJoystickProductVersionForID = __gowrap__SDL_GetJoystickProductVersionForID
	GetJoystickTypeForID = __gowrap__SDL_GetJoystickTypeForID
	OpenJoystick = __gowrap__SDL_OpenJoystick
	GetJoystickFromID = __gowrap__SDL_GetJoystickFromID
	GetJoystickFromPlayerIndex = __gowrap__SDL_GetJoystickFromPlayerIndex
	AttachVirtualJoystick = __gowrap__SDL_AttachVirtualJoystick
	DetachVirtualJoystick = __gowrap__SDL_DetachVirtualJoystick
	IsJoystickVirtual = __gowrap__SDL_IsJoystickVirtual
	SetJoystickVirtualAxis = __gowrap__SDL_SetJoystickVirtualAxis
	SetJoystickVirtualBall = __gowrap__SDL_SetJoystickVirtualBall
	SetJoystickVirtualButton = __gowrap__SDL_SetJoystickVirtualButton
	SetJoystickVirtualHat = __gowrap__SDL_SetJoystickVirtualHat
	SetJoystickVirtualTouchpad = __gowrap__SDL_SetJoystickVirtualTouchpad
	SendJoystickVirtualSensorData = __gowrap__SDL_SendJoystickVirtualSensorData
	GetJoystickProperties = __gowrap__SDL_GetJoystickProperties
	GetJoystickName = __gowrap__SDL_GetJoystickName
	GetJoystickPath = __gowrap__SDL_GetJoystickPath
	GetJoystickPlayerIndex = __gowrap__SDL_GetJoystickPlayerIndex
	SetJoystickPlayerIndex = __gowrap__SDL_SetJoystickPlayerIndex
	GetJoystickGUID = __gowrap__SDL_GetJoystickGUID
	GetJoystickVendor = __gowrap__SDL_GetJoystickVendor
	GetJoystickProduct = __gowrap__SDL_GetJoystickProduct
	GetJoystickProductVersion = __gowrap__SDL_GetJoystickProductVersion
	GetJoystickFirmwareVersion = __gowrap__SDL_GetJoystickFirmwareVersion
	GetJoystickSerial = __gowrap__SDL_GetJoystickSerial
	GetJoystickType = __gowrap__SDL_GetJoystickType
	GetJoystickGUIDInfo = __gowrap__SDL_GetJoystickGUIDInfo
	JoystickConnected = __gowrap__SDL_JoystickConnected
	GetJoystickID = __gowrap__SDL_GetJoystickID
	GetNumJoystickAxes = __gowrap__SDL_GetNumJoystickAxes
	GetNumJoystickBalls = __gowrap__SDL_GetNumJoystickBalls
	GetNumJoystickHats = __gowrap__SDL_GetNumJoystickHats
	GetNumJoystickButtons = __gowrap__SDL_GetNumJoystickButtons
	SetJoystickEventsEnabled = __gowrap__SDL_SetJoystickEventsEnabled
	JoystickEventsEnabled = __gowrap__SDL_JoystickEventsEnabled
	UpdateJoysticks = __gowrap__SDL_UpdateJoysticks
	GetJoystickAxis = __gowrap__SDL_GetJoystickAxis
	GetJoystickAxisInitialState = __gowrap__SDL_GetJoystickAxisInitialState
	GetJoystickBall = __gowrap__SDL_GetJoystickBall
	GetJoystickHat = __gowrap__SDL_GetJoystickHat
	GetJoystickButton = __gowrap__SDL_GetJoystickButton
	RumbleJoystick = __gowrap__SDL_RumbleJoystick
	RumbleJoystickTriggers = __gowrap__SDL_RumbleJoystickTriggers
	SetJoystickLED = __gowrap__SDL_SetJoystickLED
	SendJoystickEffect = __gowrap__SDL_SendJoystickEffect
	CloseJoystick = __gowrap__SDL_CloseJoystick
	GetJoystickConnectionState = __gowrap__SDL_GetJoystickConnectionState
	GetJoystickPowerInfo = __gowrap__SDL_GetJoystickPowerInfo
}
