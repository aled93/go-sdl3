package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &SetWindowsMessageHook, "SDL_SetWindowsMessageHook" },
        }
    })
}
