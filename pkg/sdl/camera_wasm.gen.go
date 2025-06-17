//go:build wasm

package sdl

import (
  "unsafe"
)



//go:wasmimport sdl3 SDL_GetNumCameraDrivers
func __SDL_GetNumCameraDrivers() int32

//go:wasmimport sdl3 SDL_GetCameraDriver
func __SDL_GetCameraDriver(int32) uintptr

//go:wasmimport sdl3 SDL_GetCurrentCameraDriver
func __SDL_GetCurrentCameraDriver() uintptr

//go:wasmimport sdl3 SDL_GetCameras
func __SDL_GetCameras(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetCameraSupportedFormats
func __SDL_GetCameraSupportedFormats(int32, uintptr) uintptr

//go:wasmimport sdl3 SDL_GetCameraName
func __SDL_GetCameraName(int32) uintptr

//go:wasmimport sdl3 SDL_GetCameraPosition
func __SDL_GetCameraPosition(int32) int32

//go:wasmimport sdl3 SDL_OpenCamera
func __SDL_OpenCamera(int32, uintptr) uintptr

//go:wasmimport sdl3 SDL_GetCameraPermissionState
func __SDL_GetCameraPermissionState(uintptr) int32

//go:wasmimport sdl3 SDL_GetCameraID
func __SDL_GetCameraID(uintptr) int32

//go:wasmimport sdl3 SDL_GetCameraProperties
func __SDL_GetCameraProperties(uintptr) int32

//go:wasmimport sdl3 SDL_GetCameraFormat
func __SDL_GetCameraFormat(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_AcquireCameraFrame
func __SDL_AcquireCameraFrame(uintptr, uintptr) uintptr

//go:wasmimport sdl3 SDL_ReleaseCameraFrame
func __SDL_ReleaseCameraFrame(uintptr, uintptr)

//go:wasmimport sdl3 SDL_CloseCamera
func __SDL_CloseCamera(uintptr)



func __gowrap__SDL_GetNumCameraDrivers() (__result int32) {
	__result = (int32)(__SDL_GetNumCameraDrivers())
	return
}

func __gowrap__SDL_GetCameraDriver(index int32) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetCameraDriver(*(*int32)(unsafe.Pointer(&index)))).Collect()))
	return
}

func __gowrap__SDL_GetCurrentCameraDriver() (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetCurrentCameraDriver()).Collect()))
	return
}

func __gowrap__SDL_GetCameras(count *int32) (__result *CameraID) {
	__result = (*CameraID)(unsafe.Pointer(__SDL_GetCameras(uintptr(unsafe.Pointer(count)))))
	return
}

func __gowrap__SDL_GetCameraSupportedFormats(instance_id CameraID, count *int32) (__result **CameraSpec) {
	__result = (**CameraSpec)(unsafe.Pointer(__SDL_GetCameraSupportedFormats(*(*int32)(unsafe.Pointer(&instance_id)), uintptr(unsafe.Pointer(count)))))
	return
}

func __gowrap__SDL_GetCameraName(instance_id CameraID) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetCameraName(*(*int32)(unsafe.Pointer(&instance_id)))).Collect()))
	return
}

func __gowrap__SDL_GetCameraPosition(instance_id CameraID) (__result CameraPosition) {
	__result = (CameraPosition)(__SDL_GetCameraPosition(*(*int32)(unsafe.Pointer(&instance_id))))
	return
}

func __gowrap__SDL_OpenCamera(instance_id CameraID, spec *CameraSpec) (__result Camera) {
	__result = (Camera)(unsafe.Pointer(__SDL_OpenCamera(*(*int32)(unsafe.Pointer(&instance_id)), uintptr(unsafe.Pointer(spec)))))
	return
}

func __gowrap__SDL_GetCameraPermissionState(camera Camera) (__result int32) {
	__result = (int32)(__SDL_GetCameraPermissionState(uintptr(unsafe.Pointer(camera))))
	return
}

func __gowrap__SDL_GetCameraID(camera Camera) (__result CameraID) {
	__result = (CameraID)(__SDL_GetCameraID(uintptr(unsafe.Pointer(camera))))
	return
}

func __gowrap__SDL_GetCameraProperties(camera Camera) (__result PropertiesID) {
	__result = (PropertiesID)(__SDL_GetCameraProperties(uintptr(unsafe.Pointer(camera))))
	return
}

func __gowrap__SDL_GetCameraFormat(camera Camera, spec *CameraSpec) (__result bool) {
	__result = (bool)(__SDL_GetCameraFormat(uintptr(unsafe.Pointer(camera)), uintptr(unsafe.Pointer(spec))) != 0)
	return
}

func __gowrap__SDL_AcquireCameraFrame(camera Camera, timestampNS *uint64) (__result *Surface) {
	__result = (*Surface)(unsafe.Pointer(__SDL_AcquireCameraFrame(uintptr(unsafe.Pointer(camera)), uintptr(unsafe.Pointer(timestampNS)))))
	return
}

func __gowrap__SDL_ReleaseCameraFrame(camera Camera, frame *Surface) {
	__SDL_ReleaseCameraFrame(uintptr(unsafe.Pointer(camera)), uintptr(unsafe.Pointer(frame)))
}

func __gowrap__SDL_CloseCamera(camera Camera) {
	__SDL_CloseCamera(uintptr(unsafe.Pointer(camera)))
}

 
func init() {
	GetNumCameraDrivers = __gowrap__SDL_GetNumCameraDrivers
	GetCameraDriver = __gowrap__SDL_GetCameraDriver
	GetCurrentCameraDriver = __gowrap__SDL_GetCurrentCameraDriver
	GetCameras = __gowrap__SDL_GetCameras
	GetCameraSupportedFormats = __gowrap__SDL_GetCameraSupportedFormats
	GetCameraName = __gowrap__SDL_GetCameraName
	GetCameraPosition = __gowrap__SDL_GetCameraPosition
	OpenCamera = __gowrap__SDL_OpenCamera
	GetCameraPermissionState = __gowrap__SDL_GetCameraPermissionState
	GetCameraID = __gowrap__SDL_GetCameraID
	GetCameraProperties = __gowrap__SDL_GetCameraProperties
	GetCameraFormat = __gowrap__SDL_GetCameraFormat
	AcquireCameraFrame = __gowrap__SDL_AcquireCameraFrame
	ReleaseCameraFrame = __gowrap__SDL_ReleaseCameraFrame
	CloseCamera = __gowrap__SDL_CloseCamera
}
