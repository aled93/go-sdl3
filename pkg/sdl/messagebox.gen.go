package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &ShowMessageBox, "SDL_ShowMessageBox" },
            { &ShowSimpleMessageBox, "SDL_ShowSimpleMessageBox" },
        }
    })
}
