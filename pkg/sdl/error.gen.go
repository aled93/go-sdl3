package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &OutOfMemory, "SDL_OutOfMemory" },
            { &GetError, "SDL_GetError" },
            { &ClearError, "SDL_ClearError" },
        }
    })
}
