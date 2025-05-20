package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&Vulkan_LoadLibrary, lib, "SDL_Vulkan_LoadLibrary")
        purego.RegisterLibFunc(&Vulkan_GetVkGetInstanceProcAddr, lib, "SDL_Vulkan_GetVkGetInstanceProcAddr")
        purego.RegisterLibFunc(&Vulkan_UnloadLibrary, lib, "SDL_Vulkan_UnloadLibrary")
        purego.RegisterLibFunc(&Vulkan_GetInstanceExtensions, lib, "SDL_Vulkan_GetInstanceExtensions")
        purego.RegisterLibFunc(&Vulkan_CreateSurface, lib, "SDL_Vulkan_CreateSurface")
        purego.RegisterLibFunc(&Vulkan_DestroySurface, lib, "SDL_Vulkan_DestroySurface")
        purego.RegisterLibFunc(&Vulkan_GetPresentationSupport, lib, "SDL_Vulkan_GetPresentationSupport")
        
    })
}
