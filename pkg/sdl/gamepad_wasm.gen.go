//go:build wasm

package sdl

import (
  "runtime"
  "unsafe"
)



//go:wasmimport sdl3 SDL_AddGamepadMapping
func __SDL_AddGamepadMapping(uintptr) int32

//go:wasmimport sdl3 SDL_AddGamepadMappingsFromIO
func __SDL_AddGamepadMappingsFromIO(uintptr, int32) int32

//go:wasmimport sdl3 SDL_AddGamepadMappingsFromFile
func __SDL_AddGamepadMappingsFromFile(uintptr) int32

//go:wasmimport sdl3 SDL_ReloadGamepadMappings
func __SDL_ReloadGamepadMappings() int32

//go:wasmimport sdl3 SDL_GetGamepadMappings
func __SDL_GetGamepadMappings(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetGamepadMappingForGUID
func __SDL_GetGamepadMappingForGUID(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetGamepadMapping
func __SDL_GetGamepadMapping(uintptr) uintptr

//go:wasmimport sdl3 SDL_SetGamepadMapping
func __SDL_SetGamepadMapping(int32, uintptr) int32

//go:wasmimport sdl3 SDL_HasGamepad
func __SDL_HasGamepad() int32

//go:wasmimport sdl3 SDL_GetGamepads
func __SDL_GetGamepads(uintptr) uintptr

//go:wasmimport sdl3 SDL_IsGamepad
func __SDL_IsGamepad(int32) int32

//go:wasmimport sdl3 SDL_GetGamepadNameForID
func __SDL_GetGamepadNameForID(int32) uintptr

//go:wasmimport sdl3 SDL_GetGamepadPathForID
func __SDL_GetGamepadPathForID(int32) uintptr

//go:wasmimport sdl3 SDL_GetGamepadPlayerIndexForID
func __SDL_GetGamepadPlayerIndexForID(int32) int32

//go:wasmimport sdl3 SDL_GetGamepadGUIDForID
func __SDL_GetGamepadGUIDForID(int32) uintptr

//go:wasmimport sdl3 SDL_GetGamepadVendorForID
func __SDL_GetGamepadVendorForID(int32) int32

//go:wasmimport sdl3 SDL_GetGamepadProductForID
func __SDL_GetGamepadProductForID(int32) int32

//go:wasmimport sdl3 SDL_GetGamepadProductVersionForID
func __SDL_GetGamepadProductVersionForID(int32) int32

//go:wasmimport sdl3 SDL_GetGamepadTypeForID
func __SDL_GetGamepadTypeForID(int32) int32

//go:wasmimport sdl3 SDL_GetRealGamepadTypeForID
func __SDL_GetRealGamepadTypeForID(int32) int32

//go:wasmimport sdl3 SDL_GetGamepadMappingForID
func __SDL_GetGamepadMappingForID(int32) uintptr

//go:wasmimport sdl3 SDL_OpenGamepad
func __SDL_OpenGamepad(int32) uintptr

//go:wasmimport sdl3 SDL_GetGamepadFromID
func __SDL_GetGamepadFromID(int32) uintptr

//go:wasmimport sdl3 SDL_GetGamepadFromPlayerIndex
func __SDL_GetGamepadFromPlayerIndex(int32) uintptr

//go:wasmimport sdl3 SDL_GetGamepadProperties
func __SDL_GetGamepadProperties(uintptr) int32

//go:wasmimport sdl3 SDL_GetGamepadID
func __SDL_GetGamepadID(uintptr) int32

//go:wasmimport sdl3 SDL_GetGamepadName
func __SDL_GetGamepadName(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetGamepadPath
func __SDL_GetGamepadPath(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetGamepadType
func __SDL_GetGamepadType(uintptr) int32

//go:wasmimport sdl3 SDL_GetRealGamepadType
func __SDL_GetRealGamepadType(uintptr) int32

//go:wasmimport sdl3 SDL_GetGamepadPlayerIndex
func __SDL_GetGamepadPlayerIndex(uintptr) int32

//go:wasmimport sdl3 SDL_SetGamepadPlayerIndex
func __SDL_SetGamepadPlayerIndex(uintptr, int32) int32

//go:wasmimport sdl3 SDL_GetGamepadVendor
func __SDL_GetGamepadVendor(uintptr) int32

//go:wasmimport sdl3 SDL_GetGamepadProduct
func __SDL_GetGamepadProduct(uintptr) int32

//go:wasmimport sdl3 SDL_GetGamepadProductVersion
func __SDL_GetGamepadProductVersion(uintptr) int32

//go:wasmimport sdl3 SDL_GetGamepadFirmwareVersion
func __SDL_GetGamepadFirmwareVersion(uintptr) int32

//go:wasmimport sdl3 SDL_GetGamepadSerial
func __SDL_GetGamepadSerial(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetGamepadSteamHandle
func __SDL_GetGamepadSteamHandle(uintptr) int64

//go:wasmimport sdl3 SDL_GetGamepadConnectionState
func __SDL_GetGamepadConnectionState(uintptr) int32

//go:wasmimport sdl3 SDL_GetGamepadPowerInfo
func __SDL_GetGamepadPowerInfo(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GamepadConnected
func __SDL_GamepadConnected(uintptr) int32

//go:wasmimport sdl3 SDL_GetGamepadJoystick
func __SDL_GetGamepadJoystick(uintptr) uintptr

//go:wasmimport sdl3 SDL_SetGamepadEventsEnabled
func __SDL_SetGamepadEventsEnabled(int32)

//go:wasmimport sdl3 SDL_GamepadEventsEnabled
func __SDL_GamepadEventsEnabled() int32

//go:wasmimport sdl3 SDL_GetGamepadBindings
func __SDL_GetGamepadBindings(uintptr, uintptr) uintptr

//go:wasmimport sdl3 SDL_UpdateGamepads
func __SDL_UpdateGamepads()

//go:wasmimport sdl3 SDL_GetGamepadTypeFromString
func __SDL_GetGamepadTypeFromString(uintptr) int32

//go:wasmimport sdl3 SDL_GetGamepadStringForType
func __SDL_GetGamepadStringForType(int32) uintptr

//go:wasmimport sdl3 SDL_GetGamepadAxisFromString
func __SDL_GetGamepadAxisFromString(uintptr) int32

//go:wasmimport sdl3 SDL_GetGamepadStringForAxis
func __SDL_GetGamepadStringForAxis(int32) uintptr

//go:wasmimport sdl3 SDL_GamepadHasAxis
func __SDL_GamepadHasAxis(uintptr, int32) int32

//go:wasmimport sdl3 SDL_GetGamepadAxis
func __SDL_GetGamepadAxis(uintptr, int32) int32

//go:wasmimport sdl3 SDL_GetGamepadButtonFromString
func __SDL_GetGamepadButtonFromString(uintptr) int32

//go:wasmimport sdl3 SDL_GetGamepadStringForButton
func __SDL_GetGamepadStringForButton(int32) uintptr

//go:wasmimport sdl3 SDL_GamepadHasButton
func __SDL_GamepadHasButton(uintptr, int32) int32

//go:wasmimport sdl3 SDL_GetGamepadButton
func __SDL_GetGamepadButton(uintptr, int32) int32

//go:wasmimport sdl3 SDL_GetGamepadButtonLabelForType
func __SDL_GetGamepadButtonLabelForType(int32, int32) int32

//go:wasmimport sdl3 SDL_GetGamepadButtonLabel
func __SDL_GetGamepadButtonLabel(uintptr, int32) int32

//go:wasmimport sdl3 SDL_GetNumGamepadTouchpads
func __SDL_GetNumGamepadTouchpads(uintptr) int32

//go:wasmimport sdl3 SDL_GetNumGamepadTouchpadFingers
func __SDL_GetNumGamepadTouchpadFingers(uintptr, int32) int32

//go:wasmimport sdl3 SDL_GetGamepadTouchpadFinger
func __SDL_GetGamepadTouchpadFinger(uintptr, int32, int32, uintptr, uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GamepadHasSensor
func __SDL_GamepadHasSensor(uintptr, int32) int32

//go:wasmimport sdl3 SDL_SetGamepadSensorEnabled
func __SDL_SetGamepadSensorEnabled(uintptr, int32, int32) int32

//go:wasmimport sdl3 SDL_GamepadSensorEnabled
func __SDL_GamepadSensorEnabled(uintptr, int32) int32

//go:wasmimport sdl3 SDL_GetGamepadSensorDataRate
func __SDL_GetGamepadSensorDataRate(uintptr, int32) float32

//go:wasmimport sdl3 SDL_GetGamepadSensorData
func __SDL_GetGamepadSensorData(uintptr, int32, uintptr, int32) int32

//go:wasmimport sdl3 SDL_RumbleGamepad
func __SDL_RumbleGamepad(uintptr, int32, int32, int32) int32

//go:wasmimport sdl3 SDL_RumbleGamepadTriggers
func __SDL_RumbleGamepadTriggers(uintptr, int32, int32, int32) int32

//go:wasmimport sdl3 SDL_SetGamepadLED
func __SDL_SetGamepadLED(uintptr, int32, int32, int32) int32

//go:wasmimport sdl3 SDL_SendGamepadEffect
func __SDL_SendGamepadEffect(uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_CloseGamepad
func __SDL_CloseGamepad(uintptr)

//go:wasmimport sdl3 SDL_GetGamepadAppleSFSymbolsNameForButton
func __SDL_GetGamepadAppleSFSymbolsNameForButton(uintptr, int32) uintptr

//go:wasmimport sdl3 SDL_GetGamepadAppleSFSymbolsNameForAxis
func __SDL_GetGamepadAppleSFSymbolsNameForAxis(uintptr, int32) uintptr



func __gowrap__SDL_AddGamepadMapping(mapping string) (__result int32) {
	__result = (int32)(__SDL_AddGamepadMapping(((*[2]uintptr)(unsafe.Pointer(&mapping)))[0]))
	runtime.KeepAlive(mapping)
	return
}

func __gowrap__SDL_AddGamepadMappingsFromIO(src IOStream, closeio bool) (__result int32) {
	__result = (int32)(__SDL_AddGamepadMappingsFromIO(uintptr(unsafe.Pointer(src)), *(*int32)(unsafe.Pointer(&closeio))))
	return
}

func __gowrap__SDL_AddGamepadMappingsFromFile(file string) (__result int32) {
	__result = (int32)(__SDL_AddGamepadMappingsFromFile(((*[2]uintptr)(unsafe.Pointer(&file)))[0]))
	runtime.KeepAlive(file)
	return
}

func __gowrap__SDL_ReloadGamepadMappings() (__result bool) {
	__result = (bool)(__SDL_ReloadGamepadMappings() != 0)
	return
}

func __gowrap__SDL_GetGamepadMappings(count *int32) (__result []string) {
	__result = ([]string)(newCSliceRaw[string](__SDL_GetGamepadMappings(uintptr(unsafe.Pointer(count)))).Collect())
	return
}

func __gowrap__SDL_GetGamepadMappingForGUID(guid *GUID) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetGamepadMappingForGUID(uintptr(unsafe.Pointer(guid)))).Collect()))
	return
}

func __gowrap__SDL_GetGamepadMapping(gamepad Gamepad) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetGamepadMapping(uintptr(unsafe.Pointer(gamepad)))).Collect()))
	return
}

func __gowrap__SDL_SetGamepadMapping(instance_id JoystickID, mapping string) (__result bool) {
	__result = (bool)(__SDL_SetGamepadMapping(*(*int32)(unsafe.Pointer(&instance_id)), ((*[2]uintptr)(unsafe.Pointer(&mapping)))[0]) != 0)
	runtime.KeepAlive(mapping)
	return
}

func __gowrap__SDL_HasGamepad() (__result bool) {
	__result = (bool)(__SDL_HasGamepad() != 0)
	return
}

func __gowrap__SDL_GetGamepads(count *int32) (__result *JoystickID) {
	__result = (*JoystickID)(unsafe.Pointer(__SDL_GetGamepads(uintptr(unsafe.Pointer(count)))))
	return
}

func __gowrap__SDL_IsGamepad(instance_id JoystickID) (__result bool) {
	__result = (bool)(__SDL_IsGamepad(*(*int32)(unsafe.Pointer(&instance_id))) != 0)
	return
}

func __gowrap__SDL_GetGamepadNameForID(instance_id JoystickID) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetGamepadNameForID(*(*int32)(unsafe.Pointer(&instance_id)))).Collect()))
	return
}

func __gowrap__SDL_GetGamepadPathForID(instance_id JoystickID) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetGamepadPathForID(*(*int32)(unsafe.Pointer(&instance_id)))).Collect()))
	return
}

func __gowrap__SDL_GetGamepadPlayerIndexForID(instance_id JoystickID) (__result int32) {
	__result = (int32)(__SDL_GetGamepadPlayerIndexForID(*(*int32)(unsafe.Pointer(&instance_id))))
	return
}

func __gowrap__SDL_GetGamepadGUIDForID(instance_id JoystickID) (__result *GUID) {
	__result = (*GUID)(unsafe.Pointer(__SDL_GetGamepadGUIDForID(*(*int32)(unsafe.Pointer(&instance_id)))))
	return
}

func __gowrap__SDL_GetGamepadVendorForID(instance_id JoystickID) (__result uint16) {
	__result = (uint16)(__SDL_GetGamepadVendorForID(*(*int32)(unsafe.Pointer(&instance_id))))
	return
}

func __gowrap__SDL_GetGamepadProductForID(instance_id JoystickID) (__result uint16) {
	__result = (uint16)(__SDL_GetGamepadProductForID(*(*int32)(unsafe.Pointer(&instance_id))))
	return
}

func __gowrap__SDL_GetGamepadProductVersionForID(instance_id JoystickID) (__result uint16) {
	__result = (uint16)(__SDL_GetGamepadProductVersionForID(*(*int32)(unsafe.Pointer(&instance_id))))
	return
}

func __gowrap__SDL_GetGamepadTypeForID(instance_id JoystickID) (__result GamepadType) {
	__result = (GamepadType)(__SDL_GetGamepadTypeForID(*(*int32)(unsafe.Pointer(&instance_id))))
	return
}

func __gowrap__SDL_GetRealGamepadTypeForID(instance_id JoystickID) (__result GamepadType) {
	__result = (GamepadType)(__SDL_GetRealGamepadTypeForID(*(*int32)(unsafe.Pointer(&instance_id))))
	return
}

func __gowrap__SDL_GetGamepadMappingForID(instance_id JoystickID) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetGamepadMappingForID(*(*int32)(unsafe.Pointer(&instance_id)))).Collect()))
	return
}

func __gowrap__SDL_OpenGamepad(instance_id JoystickID) (__result Gamepad) {
	__result = (Gamepad)(unsafe.Pointer(__SDL_OpenGamepad(*(*int32)(unsafe.Pointer(&instance_id)))))
	return
}

func __gowrap__SDL_GetGamepadFromID(instance_id JoystickID) (__result Gamepad) {
	__result = (Gamepad)(unsafe.Pointer(__SDL_GetGamepadFromID(*(*int32)(unsafe.Pointer(&instance_id)))))
	return
}

func __gowrap__SDL_GetGamepadFromPlayerIndex(player_index int32) (__result Gamepad) {
	__result = (Gamepad)(unsafe.Pointer(__SDL_GetGamepadFromPlayerIndex(*(*int32)(unsafe.Pointer(&player_index)))))
	return
}

func __gowrap__SDL_GetGamepadProperties(gamepad Gamepad) (__result PropertiesID) {
	__result = (PropertiesID)(__SDL_GetGamepadProperties(uintptr(unsafe.Pointer(gamepad))))
	return
}

func __gowrap__SDL_GetGamepadID(gamepad Gamepad) (__result JoystickID) {
	__result = (JoystickID)(__SDL_GetGamepadID(uintptr(unsafe.Pointer(gamepad))))
	return
}

func __gowrap__SDL_GetGamepadName(gamepad Gamepad) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetGamepadName(uintptr(unsafe.Pointer(gamepad)))).Collect()))
	return
}

func __gowrap__SDL_GetGamepadPath(gamepad Gamepad) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetGamepadPath(uintptr(unsafe.Pointer(gamepad)))).Collect()))
	return
}

func __gowrap__SDL_GetGamepadType(gamepad Gamepad) (__result GamepadType) {
	__result = (GamepadType)(__SDL_GetGamepadType(uintptr(unsafe.Pointer(gamepad))))
	return
}

func __gowrap__SDL_GetRealGamepadType(gamepad Gamepad) (__result GamepadType) {
	__result = (GamepadType)(__SDL_GetRealGamepadType(uintptr(unsafe.Pointer(gamepad))))
	return
}

func __gowrap__SDL_GetGamepadPlayerIndex(gamepad Gamepad) (__result int32) {
	__result = (int32)(__SDL_GetGamepadPlayerIndex(uintptr(unsafe.Pointer(gamepad))))
	return
}

func __gowrap__SDL_SetGamepadPlayerIndex(gamepad Gamepad, player_index int32) (__result bool) {
	__result = (bool)(__SDL_SetGamepadPlayerIndex(uintptr(unsafe.Pointer(gamepad)), *(*int32)(unsafe.Pointer(&player_index))) != 0)
	return
}

func __gowrap__SDL_GetGamepadVendor(gamepad Gamepad) (__result uint16) {
	__result = (uint16)(__SDL_GetGamepadVendor(uintptr(unsafe.Pointer(gamepad))))
	return
}

func __gowrap__SDL_GetGamepadProduct(gamepad Gamepad) (__result uint16) {
	__result = (uint16)(__SDL_GetGamepadProduct(uintptr(unsafe.Pointer(gamepad))))
	return
}

func __gowrap__SDL_GetGamepadProductVersion(gamepad Gamepad) (__result uint16) {
	__result = (uint16)(__SDL_GetGamepadProductVersion(uintptr(unsafe.Pointer(gamepad))))
	return
}

func __gowrap__SDL_GetGamepadFirmwareVersion(gamepad Gamepad) (__result uint16) {
	__result = (uint16)(__SDL_GetGamepadFirmwareVersion(uintptr(unsafe.Pointer(gamepad))))
	return
}

func __gowrap__SDL_GetGamepadSerial(gamepad Gamepad) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetGamepadSerial(uintptr(unsafe.Pointer(gamepad)))).Collect()))
	return
}

func __gowrap__SDL_GetGamepadSteamHandle(gamepad Gamepad) (__result uint64) {
	__result = (uint64)(__SDL_GetGamepadSteamHandle(uintptr(unsafe.Pointer(gamepad))))
	return
}

func __gowrap__SDL_GetGamepadConnectionState(gamepad Gamepad) (__result JoystickConnectionState) {
	__result = (JoystickConnectionState)(__SDL_GetGamepadConnectionState(uintptr(unsafe.Pointer(gamepad))))
	return
}

func __gowrap__SDL_GetGamepadPowerInfo(gamepad Gamepad, percent *int32) (__result PowerState) {
	__result = (PowerState)(__SDL_GetGamepadPowerInfo(uintptr(unsafe.Pointer(gamepad)), uintptr(unsafe.Pointer(percent))))
	return
}

func __gowrap__SDL_GamepadConnected(gamepad Gamepad) (__result bool) {
	__result = (bool)(__SDL_GamepadConnected(uintptr(unsafe.Pointer(gamepad))) != 0)
	return
}

func __gowrap__SDL_GetGamepadJoystick(gamepad Gamepad) (__result Joystick) {
	__result = (Joystick)(unsafe.Pointer(__SDL_GetGamepadJoystick(uintptr(unsafe.Pointer(gamepad)))))
	return
}

func __gowrap__SDL_SetGamepadEventsEnabled(enabled bool) {
	__SDL_SetGamepadEventsEnabled(*(*int32)(unsafe.Pointer(&enabled)))
}

func __gowrap__SDL_GamepadEventsEnabled() (__result bool) {
	__result = (bool)(__SDL_GamepadEventsEnabled() != 0)
	return
}

func __gowrap__SDL_GetGamepadBindings(gamepad Gamepad, count *int32) (__result []*GamepadBinding) {
	__result = ([]*GamepadBinding)(newCSliceRaw[*GamepadBinding](__SDL_GetGamepadBindings(uintptr(unsafe.Pointer(gamepad)), uintptr(unsafe.Pointer(count)))).Collect())
	return
}

func __gowrap__SDL_UpdateGamepads() {
	__SDL_UpdateGamepads()
}

func __gowrap__SDL_GetGamepadTypeFromString(str string) (__result GamepadType) {
	__result = (GamepadType)(__SDL_GetGamepadTypeFromString(((*[2]uintptr)(unsafe.Pointer(&str)))[0]))
	runtime.KeepAlive(str)
	return
}

func __gowrap__SDL_GetGamepadStringForType(type_ GamepadType) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetGamepadStringForType(*(*int32)(unsafe.Pointer(&type_)))).Collect()))
	return
}

func __gowrap__SDL_GetGamepadAxisFromString(str string) (__result GamepadAxis) {
	__result = (GamepadAxis)(__SDL_GetGamepadAxisFromString(((*[2]uintptr)(unsafe.Pointer(&str)))[0]))
	runtime.KeepAlive(str)
	return
}

func __gowrap__SDL_GetGamepadStringForAxis(axis GamepadAxis) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetGamepadStringForAxis(*(*int32)(unsafe.Pointer(&axis)))).Collect()))
	return
}

func __gowrap__SDL_GamepadHasAxis(gamepad Gamepad, axis GamepadAxis) (__result bool) {
	__result = (bool)(__SDL_GamepadHasAxis(uintptr(unsafe.Pointer(gamepad)), *(*int32)(unsafe.Pointer(&axis))) != 0)
	return
}

func __gowrap__SDL_GetGamepadAxis(gamepad Gamepad, axis GamepadAxis) (__result int16) {
	__result = (int16)(__SDL_GetGamepadAxis(uintptr(unsafe.Pointer(gamepad)), *(*int32)(unsafe.Pointer(&axis))))
	return
}

func __gowrap__SDL_GetGamepadButtonFromString(str string) (__result GamepadButton) {
	__result = (GamepadButton)(__SDL_GetGamepadButtonFromString(((*[2]uintptr)(unsafe.Pointer(&str)))[0]))
	runtime.KeepAlive(str)
	return
}

func __gowrap__SDL_GetGamepadStringForButton(button GamepadButton) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetGamepadStringForButton(*(*int32)(unsafe.Pointer(&button)))).Collect()))
	return
}

func __gowrap__SDL_GamepadHasButton(gamepad Gamepad, button GamepadButton) (__result bool) {
	__result = (bool)(__SDL_GamepadHasButton(uintptr(unsafe.Pointer(gamepad)), *(*int32)(unsafe.Pointer(&button))) != 0)
	return
}

func __gowrap__SDL_GetGamepadButton(gamepad Gamepad, button GamepadButton) (__result bool) {
	__result = (bool)(__SDL_GetGamepadButton(uintptr(unsafe.Pointer(gamepad)), *(*int32)(unsafe.Pointer(&button))) != 0)
	return
}

func __gowrap__SDL_GetGamepadButtonLabelForType(type_ GamepadType, button GamepadButton) (__result GamepadButtonLabel) {
	__result = (GamepadButtonLabel)(__SDL_GetGamepadButtonLabelForType(*(*int32)(unsafe.Pointer(&type_)), *(*int32)(unsafe.Pointer(&button))))
	return
}

func __gowrap__SDL_GetGamepadButtonLabel(gamepad Gamepad, button GamepadButton) (__result GamepadButtonLabel) {
	__result = (GamepadButtonLabel)(__SDL_GetGamepadButtonLabel(uintptr(unsafe.Pointer(gamepad)), *(*int32)(unsafe.Pointer(&button))))
	return
}

func __gowrap__SDL_GetNumGamepadTouchpads(gamepad Gamepad) (__result int32) {
	__result = (int32)(__SDL_GetNumGamepadTouchpads(uintptr(unsafe.Pointer(gamepad))))
	return
}

func __gowrap__SDL_GetNumGamepadTouchpadFingers(gamepad Gamepad, touchpad int32) (__result int32) {
	__result = (int32)(__SDL_GetNumGamepadTouchpadFingers(uintptr(unsafe.Pointer(gamepad)), *(*int32)(unsafe.Pointer(&touchpad))))
	return
}

func __gowrap__SDL_GetGamepadTouchpadFinger(gamepad Gamepad, touchpad int32, finger int32, down *bool, x *float32, y *float32, pressure *float32) (__result bool) {
	__result = (bool)(__SDL_GetGamepadTouchpadFinger(uintptr(unsafe.Pointer(gamepad)), *(*int32)(unsafe.Pointer(&touchpad)), *(*int32)(unsafe.Pointer(&finger)), uintptr(unsafe.Pointer(down)), uintptr(unsafe.Pointer(x)), uintptr(unsafe.Pointer(y)), uintptr(unsafe.Pointer(pressure))) != 0)
	return
}

func __gowrap__SDL_GamepadHasSensor(gamepad Gamepad, type_ SensorType) (__result bool) {
	__result = (bool)(__SDL_GamepadHasSensor(uintptr(unsafe.Pointer(gamepad)), *(*int32)(unsafe.Pointer(&type_))) != 0)
	return
}

func __gowrap__SDL_SetGamepadSensorEnabled(gamepad Gamepad, type_ SensorType, enabled bool) (__result bool) {
	__result = (bool)(__SDL_SetGamepadSensorEnabled(uintptr(unsafe.Pointer(gamepad)), *(*int32)(unsafe.Pointer(&type_)), *(*int32)(unsafe.Pointer(&enabled))) != 0)
	return
}

func __gowrap__SDL_GamepadSensorEnabled(gamepad Gamepad, type_ SensorType) (__result bool) {
	__result = (bool)(__SDL_GamepadSensorEnabled(uintptr(unsafe.Pointer(gamepad)), *(*int32)(unsafe.Pointer(&type_))) != 0)
	return
}

func __gowrap__SDL_GetGamepadSensorDataRate(gamepad Gamepad, type_ SensorType) (__result float32) {
	__result = (float32)(__SDL_GetGamepadSensorDataRate(uintptr(unsafe.Pointer(gamepad)), *(*int32)(unsafe.Pointer(&type_))))
	return
}

func __gowrap__SDL_GetGamepadSensorData(gamepad Gamepad, type_ SensorType, data *float32, num_values int32) (__result bool) {
	__result = (bool)(__SDL_GetGamepadSensorData(uintptr(unsafe.Pointer(gamepad)), *(*int32)(unsafe.Pointer(&type_)), uintptr(unsafe.Pointer(data)), *(*int32)(unsafe.Pointer(&num_values))) != 0)
	return
}

func __gowrap__SDL_RumbleGamepad(gamepad Gamepad, low_frequency_rumble uint16, high_frequency_rumble uint16, duration_ms uint32) (__result bool) {
	__result = (bool)(__SDL_RumbleGamepad(uintptr(unsafe.Pointer(gamepad)), *(*int32)(unsafe.Pointer(&low_frequency_rumble)), *(*int32)(unsafe.Pointer(&high_frequency_rumble)), *(*int32)(unsafe.Pointer(&duration_ms))) != 0)
	return
}

func __gowrap__SDL_RumbleGamepadTriggers(gamepad Gamepad, left_rumble uint16, right_rumble uint16, duration_ms uint32) (__result bool) {
	__result = (bool)(__SDL_RumbleGamepadTriggers(uintptr(unsafe.Pointer(gamepad)), *(*int32)(unsafe.Pointer(&left_rumble)), *(*int32)(unsafe.Pointer(&right_rumble)), *(*int32)(unsafe.Pointer(&duration_ms))) != 0)
	return
}

func __gowrap__SDL_SetGamepadLED(gamepad Gamepad, red uint8, green uint8, blue uint8) (__result bool) {
	__result = (bool)(__SDL_SetGamepadLED(uintptr(unsafe.Pointer(gamepad)), *(*int32)(unsafe.Pointer(&red)), *(*int32)(unsafe.Pointer(&green)), *(*int32)(unsafe.Pointer(&blue))) != 0)
	return
}

func __gowrap__SDL_SendGamepadEffect(gamepad Gamepad, data []byte, size int32) (__result bool) {
	__result = (bool)(__SDL_SendGamepadEffect(uintptr(unsafe.Pointer(gamepad)), uintptr(unsafe.Pointer(&data[0])), *(*int32)(unsafe.Pointer(&size))) != 0)
	runtime.KeepAlive(data)
	return
}

func __gowrap__SDL_CloseGamepad(gamepad Gamepad) {
	__SDL_CloseGamepad(uintptr(unsafe.Pointer(gamepad)))
}

func __gowrap__SDL_GetGamepadAppleSFSymbolsNameForButton(gamepad Gamepad, button GamepadButton) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetGamepadAppleSFSymbolsNameForButton(uintptr(unsafe.Pointer(gamepad)), *(*int32)(unsafe.Pointer(&button)))).Collect()))
	return
}

func __gowrap__SDL_GetGamepadAppleSFSymbolsNameForAxis(gamepad Gamepad, axis GamepadAxis) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetGamepadAppleSFSymbolsNameForAxis(uintptr(unsafe.Pointer(gamepad)), *(*int32)(unsafe.Pointer(&axis)))).Collect()))
	return
}

 
func init() {
	AddGamepadMapping = __gowrap__SDL_AddGamepadMapping
	AddGamepadMappingsFromIO = __gowrap__SDL_AddGamepadMappingsFromIO
	AddGamepadMappingsFromFile = __gowrap__SDL_AddGamepadMappingsFromFile
	ReloadGamepadMappings = __gowrap__SDL_ReloadGamepadMappings
	GetGamepadMappings = __gowrap__SDL_GetGamepadMappings
	GetGamepadMappingForGUID = __gowrap__SDL_GetGamepadMappingForGUID
	GetGamepadMapping = __gowrap__SDL_GetGamepadMapping
	SetGamepadMapping = __gowrap__SDL_SetGamepadMapping
	HasGamepad = __gowrap__SDL_HasGamepad
	GetGamepads = __gowrap__SDL_GetGamepads
	IsGamepad = __gowrap__SDL_IsGamepad
	GetGamepadNameForID = __gowrap__SDL_GetGamepadNameForID
	GetGamepadPathForID = __gowrap__SDL_GetGamepadPathForID
	GetGamepadPlayerIndexForID = __gowrap__SDL_GetGamepadPlayerIndexForID
	GetGamepadGUIDForID = __gowrap__SDL_GetGamepadGUIDForID
	GetGamepadVendorForID = __gowrap__SDL_GetGamepadVendorForID
	GetGamepadProductForID = __gowrap__SDL_GetGamepadProductForID
	GetGamepadProductVersionForID = __gowrap__SDL_GetGamepadProductVersionForID
	GetGamepadTypeForID = __gowrap__SDL_GetGamepadTypeForID
	GetRealGamepadTypeForID = __gowrap__SDL_GetRealGamepadTypeForID
	GetGamepadMappingForID = __gowrap__SDL_GetGamepadMappingForID
	OpenGamepad = __gowrap__SDL_OpenGamepad
	GetGamepadFromID = __gowrap__SDL_GetGamepadFromID
	GetGamepadFromPlayerIndex = __gowrap__SDL_GetGamepadFromPlayerIndex
	GetGamepadProperties = __gowrap__SDL_GetGamepadProperties
	GetGamepadID = __gowrap__SDL_GetGamepadID
	GetGamepadName = __gowrap__SDL_GetGamepadName
	GetGamepadPath = __gowrap__SDL_GetGamepadPath
	GetGamepadType = __gowrap__SDL_GetGamepadType
	GetRealGamepadType = __gowrap__SDL_GetRealGamepadType
	GetGamepadPlayerIndex = __gowrap__SDL_GetGamepadPlayerIndex
	SetGamepadPlayerIndex = __gowrap__SDL_SetGamepadPlayerIndex
	GetGamepadVendor = __gowrap__SDL_GetGamepadVendor
	GetGamepadProduct = __gowrap__SDL_GetGamepadProduct
	GetGamepadProductVersion = __gowrap__SDL_GetGamepadProductVersion
	GetGamepadFirmwareVersion = __gowrap__SDL_GetGamepadFirmwareVersion
	GetGamepadSerial = __gowrap__SDL_GetGamepadSerial
	GetGamepadSteamHandle = __gowrap__SDL_GetGamepadSteamHandle
	GetGamepadConnectionState = __gowrap__SDL_GetGamepadConnectionState
	GetGamepadPowerInfo = __gowrap__SDL_GetGamepadPowerInfo
	GamepadConnected = __gowrap__SDL_GamepadConnected
	GetGamepadJoystick = __gowrap__SDL_GetGamepadJoystick
	SetGamepadEventsEnabled = __gowrap__SDL_SetGamepadEventsEnabled
	GamepadEventsEnabled = __gowrap__SDL_GamepadEventsEnabled
	GetGamepadBindings = __gowrap__SDL_GetGamepadBindings
	UpdateGamepads = __gowrap__SDL_UpdateGamepads
	GetGamepadTypeFromString = __gowrap__SDL_GetGamepadTypeFromString
	GetGamepadStringForType = __gowrap__SDL_GetGamepadStringForType
	GetGamepadAxisFromString = __gowrap__SDL_GetGamepadAxisFromString
	GetGamepadStringForAxis = __gowrap__SDL_GetGamepadStringForAxis
	GamepadHasAxis = __gowrap__SDL_GamepadHasAxis
	GetGamepadAxis = __gowrap__SDL_GetGamepadAxis
	GetGamepadButtonFromString = __gowrap__SDL_GetGamepadButtonFromString
	GetGamepadStringForButton = __gowrap__SDL_GetGamepadStringForButton
	GamepadHasButton = __gowrap__SDL_GamepadHasButton
	GetGamepadButton = __gowrap__SDL_GetGamepadButton
	GetGamepadButtonLabelForType = __gowrap__SDL_GetGamepadButtonLabelForType
	GetGamepadButtonLabel = __gowrap__SDL_GetGamepadButtonLabel
	GetNumGamepadTouchpads = __gowrap__SDL_GetNumGamepadTouchpads
	GetNumGamepadTouchpadFingers = __gowrap__SDL_GetNumGamepadTouchpadFingers
	GetGamepadTouchpadFinger = __gowrap__SDL_GetGamepadTouchpadFinger
	GamepadHasSensor = __gowrap__SDL_GamepadHasSensor
	SetGamepadSensorEnabled = __gowrap__SDL_SetGamepadSensorEnabled
	GamepadSensorEnabled = __gowrap__SDL_GamepadSensorEnabled
	GetGamepadSensorDataRate = __gowrap__SDL_GetGamepadSensorDataRate
	GetGamepadSensorData = __gowrap__SDL_GetGamepadSensorData
	RumbleGamepad = __gowrap__SDL_RumbleGamepad
	RumbleGamepadTriggers = __gowrap__SDL_RumbleGamepadTriggers
	SetGamepadLED = __gowrap__SDL_SetGamepadLED
	SendGamepadEffect = __gowrap__SDL_SendGamepadEffect
	CloseGamepad = __gowrap__SDL_CloseGamepad
	GetGamepadAppleSFSymbolsNameForButton = __gowrap__SDL_GetGamepadAppleSFSymbolsNameForButton
	GetGamepadAppleSFSymbolsNameForAxis = __gowrap__SDL_GetGamepadAppleSFSymbolsNameForAxis
}
