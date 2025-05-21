package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &SetX11EventHook, "SDL_SetX11EventHook" },
            { &SetLinuxThreadPriority, "SDL_SetLinuxThreadPriority" },
            { &SetLinuxThreadPriorityAndPolicy, "SDL_SetLinuxThreadPriorityAndPolicy" },
        }
    })
}
