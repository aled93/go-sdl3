package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &GetNumLogicalCPUCores, "SDL_GetNumLogicalCPUCores" },
            { &GetCPUCacheLineSize, "SDL_GetCPUCacheLineSize" },
            { &HasAltiVec, "SDL_HasAltiVec" },
            { &HasMMX, "SDL_HasMMX" },
            { &HasSSE, "SDL_HasSSE" },
            { &HasSSE2, "SDL_HasSSE2" },
            { &HasSSE3, "SDL_HasSSE3" },
            { &HasSSE41, "SDL_HasSSE41" },
            { &HasSSE42, "SDL_HasSSE42" },
            { &HasAVX, "SDL_HasAVX" },
            { &HasAVX2, "SDL_HasAVX2" },
            { &HasAVX512F, "SDL_HasAVX512F" },
            { &HasARMSIMD, "SDL_HasARMSIMD" },
            { &HasNEON, "SDL_HasNEON" },
            { &HasLSX, "SDL_HasLSX" },
            { &HasLASX, "SDL_HasLASX" },
            { &GetSystemRAM, "SDL_GetSystemRAM" },
            { &GetSIMDAlignment, "SDL_GetSIMDAlignment" },
        }
    })
}
