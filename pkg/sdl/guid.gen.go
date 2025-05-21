package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &GUIDToString, "SDL_GUIDToString" },
            { &StringToGUID, "SDL_StringToGUID" },
        }
    })
}
