package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &SetiOSAnimationCallback, "SDL_SetiOSAnimationCallback" },
            { &SetiOSEventPump, "SDL_SetiOSEventPump" },
            { &OnApplicationDidChangeStatusBarOrientation, "SDL_OnApplicationDidChangeStatusBarOrientation" },
        }
    })
}
