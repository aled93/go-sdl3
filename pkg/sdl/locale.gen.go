package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &GetPreferredLocales, "SDL_GetPreferredLocales" },
        }
    })
}
