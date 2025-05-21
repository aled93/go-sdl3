package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &OpenURL, "SDL_OpenURL" },
        }
    })
}
