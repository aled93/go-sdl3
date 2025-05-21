package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &GetPixelFormatName, "SDL_GetPixelFormatName" },
            { &GetMasksForPixelFormat, "SDL_GetMasksForPixelFormat" },
            { &GetPixelFormatForMasks, "SDL_GetPixelFormatForMasks" },
            { &GetPixelFormatDetails, "SDL_GetPixelFormatDetails" },
            { &CreatePalette, "SDL_CreatePalette" },
            { &SetPaletteColors, "SDL_SetPaletteColors" },
            { &DestroyPalette, "SDL_DestroyPalette" },
            { &MapRGB, "SDL_MapRGB" },
            { &MapRGBA, "SDL_MapRGBA" },
            { &GetRGB, "SDL_GetRGB" },
            { &GetRGBA, "SDL_GetRGBA" },
        }
    })
}
