package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&GetPixelFormatName, lib, "SDL_GetPixelFormatName")
        purego.RegisterLibFunc(&GetMasksForPixelFormat, lib, "SDL_GetMasksForPixelFormat")
        purego.RegisterLibFunc(&GetPixelFormatForMasks, lib, "SDL_GetPixelFormatForMasks")
        purego.RegisterLibFunc(&GetPixelFormatDetails, lib, "SDL_GetPixelFormatDetails")
        purego.RegisterLibFunc(&CreatePalette, lib, "SDL_CreatePalette")
        purego.RegisterLibFunc(&SetPaletteColors, lib, "SDL_SetPaletteColors")
        purego.RegisterLibFunc(&DestroyPalette, lib, "SDL_DestroyPalette")
        purego.RegisterLibFunc(&MapRGB, lib, "SDL_MapRGB")
        purego.RegisterLibFunc(&MapRGBA, lib, "SDL_MapRGBA")
        purego.RegisterLibFunc(&GetRGB, lib, "SDL_GetRGB")
        purego.RegisterLibFunc(&GetRGBA, lib, "SDL_GetRGBA")
        
    })
}
