package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &Metal_CreateView, "SDL_Metal_CreateView" },
            { &Metal_DestroyView, "SDL_Metal_DestroyView" },
            { &Metal_GetLayer, "SDL_Metal_GetLayer" },
        }
    })
}
