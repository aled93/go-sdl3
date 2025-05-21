package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &SetHintWithPriority, "SDL_SetHintWithPriority" },
            { &SetHint, "SDL_SetHint" },
            { &ResetHint, "SDL_ResetHint" },
            { &ResetHints, "SDL_ResetHints" },
            { &GetHint, "SDL_GetHint" },
            { &GetHintBoolean, "SDL_GetHintBoolean" },
            { &AddHintCallback, "SDL_AddHintCallback" },
            { &RemoveHintCallback, "SDL_RemoveHintCallback" },
        }
    })
}
