package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &ShouldInit, "SDL_ShouldInit" },
            { &ShouldQuit, "SDL_ShouldQuit" },
            { &SetInitialized, "SDL_SetInitialized" },
        }
    })
}
