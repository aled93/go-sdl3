package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &Vulkan_LoadLibrary, "SDL_Vulkan_LoadLibrary" },
            { &Vulkan_GetVkGetInstanceProcAddr, "SDL_Vulkan_GetVkGetInstanceProcAddr" },
            { &Vulkan_UnloadLibrary, "SDL_Vulkan_UnloadLibrary" },
            { &Vulkan_GetInstanceExtensions, "SDL_Vulkan_GetInstanceExtensions" },
            { &Vulkan_CreateSurface, "SDL_Vulkan_CreateSurface" },
            { &Vulkan_DestroySurface, "SDL_Vulkan_DestroySurface" },
            { &Vulkan_GetPresentationSupport, "SDL_Vulkan_GetPresentationSupport" },
        }
    })
}
