//go:build wasm

package sdl

import (
  "runtime"
  "unsafe"
)



//go:wasmimport sdl3 SDL_GetSensors
func __SDL_GetSensors(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetSensorNameForID
func __SDL_GetSensorNameForID(int32) uintptr

//go:wasmimport sdl3 SDL_GetSensorTypeForID
func __SDL_GetSensorTypeForID(int32) int32

//go:wasmimport sdl3 SDL_GetSensorNonPortableTypeForID
func __SDL_GetSensorNonPortableTypeForID(int32) int32

//go:wasmimport sdl3 SDL_OpenSensor
func __SDL_OpenSensor(int32) uintptr

//go:wasmimport sdl3 SDL_GetSensorFromID
func __SDL_GetSensorFromID(int32) uintptr

//go:wasmimport sdl3 SDL_GetSensorProperties
func __SDL_GetSensorProperties(uintptr) int32

//go:wasmimport sdl3 SDL_GetSensorName
func __SDL_GetSensorName(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetSensorType
func __SDL_GetSensorType(uintptr) int32

//go:wasmimport sdl3 SDL_GetSensorNonPortableType
func __SDL_GetSensorNonPortableType(uintptr) int32

//go:wasmimport sdl3 SDL_GetSensorID
func __SDL_GetSensorID(uintptr) int32

//go:wasmimport sdl3 SDL_GetSensorData
func __SDL_GetSensorData(uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_CloseSensor
func __SDL_CloseSensor(uintptr)

//go:wasmimport sdl3 SDL_UpdateSensors
func __SDL_UpdateSensors()



func __gowrap__SDL_GetSensors(count *int32) (__result *SensorID) {
	__result = (*SensorID)(unsafe.Pointer(__SDL_GetSensors(uintptr(unsafe.Pointer(count)))))
	return
}

func __gowrap__SDL_GetSensorNameForID(instance_id SensorID) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetSensorNameForID(*(*int32)(unsafe.Pointer(&instance_id)))).Collect()))
	return
}

func __gowrap__SDL_GetSensorTypeForID(instance_id SensorID) (__result SensorType) {
	__result = (SensorType)(__SDL_GetSensorTypeForID(*(*int32)(unsafe.Pointer(&instance_id))))
	return
}

func __gowrap__SDL_GetSensorNonPortableTypeForID(instance_id SensorID) (__result int32) {
	__result = (int32)(__SDL_GetSensorNonPortableTypeForID(*(*int32)(unsafe.Pointer(&instance_id))))
	return
}

func __gowrap__SDL_OpenSensor(instance_id SensorID) (__result *Sensor) {
	__result = (*Sensor)(unsafe.Pointer(__SDL_OpenSensor(*(*int32)(unsafe.Pointer(&instance_id)))))
	return
}

func __gowrap__SDL_GetSensorFromID(instance_id SensorID) (__result *Sensor) {
	__result = (*Sensor)(unsafe.Pointer(__SDL_GetSensorFromID(*(*int32)(unsafe.Pointer(&instance_id)))))
	return
}

func __gowrap__SDL_GetSensorProperties(sensor Sensor) (__result PropertiesID) {
	__result = (PropertiesID)(__SDL_GetSensorProperties(uintptr(unsafe.Pointer(sensor))))
	return
}

func __gowrap__SDL_GetSensorName(sensor Sensor) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetSensorName(uintptr(unsafe.Pointer(sensor)))).Collect()))
	return
}

func __gowrap__SDL_GetSensorType(sensor Sensor) (__result SensorType) {
	__result = (SensorType)(__SDL_GetSensorType(uintptr(unsafe.Pointer(sensor))))
	return
}

func __gowrap__SDL_GetSensorNonPortableType(sensor Sensor) (__result int32) {
	__result = (int32)(__SDL_GetSensorNonPortableType(uintptr(unsafe.Pointer(sensor))))
	return
}

func __gowrap__SDL_GetSensorID(sensor Sensor) (__result SensorID) {
	__result = (SensorID)(__SDL_GetSensorID(uintptr(unsafe.Pointer(sensor))))
	return
}

func __gowrap__SDL_GetSensorData(sensor Sensor, data []float32, num_values int32) (__result bool) {
	__result = (bool)(__SDL_GetSensorData(uintptr(unsafe.Pointer(sensor)), uintptr(unsafe.Pointer(&data[0])), *(*int32)(unsafe.Pointer(&num_values))) != 0)
	runtime.KeepAlive(data)
	return
}

func __gowrap__SDL_CloseSensor(sensor Sensor) {
	__SDL_CloseSensor(uintptr(unsafe.Pointer(sensor)))
}

func __gowrap__SDL_UpdateSensors() {
	__SDL_UpdateSensors()
}

 
func init() {
	GetSensors = __gowrap__SDL_GetSensors
	GetSensorNameForID = __gowrap__SDL_GetSensorNameForID
	GetSensorTypeForID = __gowrap__SDL_GetSensorTypeForID
	GetSensorNonPortableTypeForID = __gowrap__SDL_GetSensorNonPortableTypeForID
	OpenSensor = __gowrap__SDL_OpenSensor
	GetSensorFromID = __gowrap__SDL_GetSensorFromID
	GetSensorProperties = __gowrap__SDL_GetSensorProperties
	GetSensorName = __gowrap__SDL_GetSensorName
	GetSensorType = __gowrap__SDL_GetSensorType
	GetSensorNonPortableType = __gowrap__SDL_GetSensorNonPortableType
	GetSensorID = __gowrap__SDL_GetSensorID
	GetSensorData = __gowrap__SDL_GetSensorData
	CloseSensor = __gowrap__SDL_CloseSensor
	UpdateSensors = __gowrap__SDL_UpdateSensors
}
