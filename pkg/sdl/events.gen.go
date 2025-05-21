package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &PumpEvents, "SDL_PumpEvents" },
            { &PeepEvents, "SDL_PeepEvents" },
            { &HasEvent, "SDL_HasEvent" },
            { &HasEvents, "SDL_HasEvents" },
            { &FlushEvent, "SDL_FlushEvent" },
            { &FlushEvents, "SDL_FlushEvents" },
            { &PollEvent, "SDL_PollEvent" },
            { &WaitEvent, "SDL_WaitEvent" },
            { &WaitEventTimeout, "SDL_WaitEventTimeout" },
            { &PushEvent, "SDL_PushEvent" },
            { &SetEventFilter, "SDL_SetEventFilter" },
            { &GetEventFilter, "SDL_GetEventFilter" },
            { &AddEventWatch, "SDL_AddEventWatch" },
            { &RemoveEventWatch, "SDL_RemoveEventWatch" },
            { &FilterEvents, "SDL_FilterEvents" },
            { &SetEventEnabled, "SDL_SetEventEnabled" },
            { &EventEnabled, "SDL_EventEnabled" },
            { &RegisterEvents, "SDL_RegisterEvents" },
            { &GetWindowFromEvent, "SDL_GetWindowFromEvent" },
            { &GetEventDescription, "SDL_GetEventDescription" },
        }
    })
}
