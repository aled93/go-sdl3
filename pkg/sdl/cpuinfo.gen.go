package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&GetNumLogicalCPUCores, lib, "SDL_GetNumLogicalCPUCores")
        purego.RegisterLibFunc(&GetCPUCacheLineSize, lib, "SDL_GetCPUCacheLineSize")
        purego.RegisterLibFunc(&HasAltiVec, lib, "SDL_HasAltiVec")
        purego.RegisterLibFunc(&HasMMX, lib, "SDL_HasMMX")
        purego.RegisterLibFunc(&HasSSE, lib, "SDL_HasSSE")
        purego.RegisterLibFunc(&HasSSE2, lib, "SDL_HasSSE2")
        purego.RegisterLibFunc(&HasSSE3, lib, "SDL_HasSSE3")
        purego.RegisterLibFunc(&HasSSE41, lib, "SDL_HasSSE41")
        purego.RegisterLibFunc(&HasSSE42, lib, "SDL_HasSSE42")
        purego.RegisterLibFunc(&HasAVX, lib, "SDL_HasAVX")
        purego.RegisterLibFunc(&HasAVX2, lib, "SDL_HasAVX2")
        purego.RegisterLibFunc(&HasAVX512F, lib, "SDL_HasAVX512F")
        purego.RegisterLibFunc(&HasARMSIMD, lib, "SDL_HasARMSIMD")
        purego.RegisterLibFunc(&HasNEON, lib, "SDL_HasNEON")
        purego.RegisterLibFunc(&HasLSX, lib, "SDL_HasLSX")
        purego.RegisterLibFunc(&HasLASX, lib, "SDL_HasLASX")
        purego.RegisterLibFunc(&GetSystemRAM, lib, "SDL_GetSystemRAM")
        purego.RegisterLibFunc(&GetSIMDAlignment, lib, "SDL_GetSIMDAlignment")
        
    })
}
