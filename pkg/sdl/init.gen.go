package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &Init, "SDL_Init" },
            { &InitSubSystem, "SDL_InitSubSystem" },
            { &QuitSubSystem, "SDL_QuitSubSystem" },
            { &WasInit, "SDL_WasInit" },
            { &Quit, "SDL_Quit" },
            { &IsMainThread, "SDL_IsMainThread" },
            { &RunOnMainThread, "SDL_RunOnMainThread" },
            { &SetAppMetadata, "SDL_SetAppMetadata" },
            { &SetAppMetadataProperty, "SDL_SetAppMetadataProperty" },
            { &GetAppMetadataProperty, "SDL_GetAppMetadataProperty" },
        }
    })
}
