//go:build wasm

package sdl

import (
  "runtime"
  "unsafe"
)



//go:wasmimport sdl3 SDL_Vulkan_LoadLibrary
func __SDL_Vulkan_LoadLibrary(uintptr) int32

//go:wasmimport sdl3 SDL_Vulkan_GetVkGetInstanceProcAddr
func __SDL_Vulkan_GetVkGetInstanceProcAddr() uintptr

//go:wasmimport sdl3 SDL_Vulkan_UnloadLibrary
func __SDL_Vulkan_UnloadLibrary()

//go:wasmimport sdl3 SDL_Vulkan_GetInstanceExtensions
func __SDL_Vulkan_GetInstanceExtensions(uintptr) uintptr

//go:wasmimport sdl3 SDL_Vulkan_CreateSurface
func __SDL_Vulkan_CreateSurface(uintptr, uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_Vulkan_DestroySurface
func __SDL_Vulkan_DestroySurface(uintptr, int64, uintptr)

//go:wasmimport sdl3 SDL_Vulkan_GetPresentationSupport
func __SDL_Vulkan_GetPresentationSupport(uintptr, uintptr, int32) int32



func __gowrap__SDL_Vulkan_LoadLibrary(path string) (__result bool) {
	__result = (bool)(__SDL_Vulkan_LoadLibrary(((*[2]uintptr)(unsafe.Pointer(&path)))[0]) != 0)
	runtime.KeepAlive(path)
	return
}

func __gowrap__SDL_Vulkan_GetVkGetInstanceProcAddr() (__result uintptr) {
	__result = (uintptr)(unsafe.Pointer(__SDL_Vulkan_GetVkGetInstanceProcAddr()))
	return
}

func __gowrap__SDL_Vulkan_UnloadLibrary() {
	__SDL_Vulkan_UnloadLibrary()
}

func __gowrap__SDL_Vulkan_GetInstanceExtensions(count *uint32) (__result *string) {
	__result = (*string)(unsafe.Pointer(__SDL_Vulkan_GetInstanceExtensions(uintptr(unsafe.Pointer(count)))))
	return
}

func __gowrap__SDL_Vulkan_CreateSurface(window Window, instance VkInstance, allocator VkAllocationCallbacks, surface *VkSurfaceKHR) (__result bool) {
	__result = (bool)(__SDL_Vulkan_CreateSurface(uintptr(unsafe.Pointer(window)), uintptr(unsafe.Pointer(instance)), 0 /* TODO: callbacks */, uintptr(unsafe.Pointer(surface))) != 0)
	return
}

func __gowrap__SDL_Vulkan_DestroySurface(instance VkInstance, surface VkSurfaceKHR, allocator VkAllocationCallbacks) {
	__SDL_Vulkan_DestroySurface(uintptr(unsafe.Pointer(instance)), *(*int64)(unsafe.Pointer(&surface)), 0 /* TODO: callbacks */)
}

func __gowrap__SDL_Vulkan_GetPresentationSupport(instance VkInstance, physicalDevice VkPhysicalDevice, queueFamilyIndex uint32) (__result bool) {
	__result = (bool)(__SDL_Vulkan_GetPresentationSupport(uintptr(unsafe.Pointer(instance)), uintptr(unsafe.Pointer(physicalDevice)), *(*int32)(unsafe.Pointer(&queueFamilyIndex))) != 0)
	return
}

 
func init() {
	Vulkan_LoadLibrary = __gowrap__SDL_Vulkan_LoadLibrary
	Vulkan_GetVkGetInstanceProcAddr = __gowrap__SDL_Vulkan_GetVkGetInstanceProcAddr
	Vulkan_UnloadLibrary = __gowrap__SDL_Vulkan_UnloadLibrary
	Vulkan_GetInstanceExtensions = __gowrap__SDL_Vulkan_GetInstanceExtensions
	Vulkan_CreateSurface = __gowrap__SDL_Vulkan_CreateSurface
	Vulkan_DestroySurface = __gowrap__SDL_Vulkan_DestroySurface
	Vulkan_GetPresentationSupport = __gowrap__SDL_Vulkan_GetPresentationSupport
}
