package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &SetClipboardText, "SDL_SetClipboardText" },
            { &GetClipboardText, "SDL_GetClipboardText" },
            { &HasClipboardText, "SDL_HasClipboardText" },
            { &SetPrimarySelectionText, "SDL_SetPrimarySelectionText" },
            { &GetPrimarySelectionText, "SDL_GetPrimarySelectionText" },
            { &HasPrimarySelectionText, "SDL_HasPrimarySelectionText" },
            { &SetClipboardData, "SDL_SetClipboardData" },
            { &ClearClipboardData, "SDL_ClearClipboardData" },
            { &GetClipboardData, "SDL_GetClipboardData" },
            { &HasClipboardData, "SDL_HasClipboardData" },
            { &GetClipboardMimeTypes, "SDL_GetClipboardMimeTypes" },
        }
    })
}
