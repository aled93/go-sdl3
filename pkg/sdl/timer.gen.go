package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &GetTicks, "SDL_GetTicks" },
            { &GetTicksNS, "SDL_GetTicksNS" },
            { &GetPerformanceCounter, "SDL_GetPerformanceCounter" },
            { &GetPerformanceFrequency, "SDL_GetPerformanceFrequency" },
            { &Delay, "SDL_Delay" },
            { &DelayNS, "SDL_DelayNS" },
            { &DelayPrecise, "SDL_DelayPrecise" },
            { &AddTimer, "SDL_AddTimer" },
            { &AddTimerNS, "SDL_AddTimerNS" },
            { &RemoveTimer, "SDL_RemoveTimer" },
        }
    })
}
