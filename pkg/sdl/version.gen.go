package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &GetVersion, "SDL_GetVersion" },
            { &GetRevision, "SDL_GetRevision" },
        }
    })
}
