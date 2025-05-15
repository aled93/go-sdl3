package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&GetTicks, lib, "SDL_GetTicks")
        purego.RegisterLibFunc(&GetTicksNS, lib, "SDL_GetTicksNS")
        purego.RegisterLibFunc(&GetPerformanceCounter, lib, "SDL_GetPerformanceCounter")
        purego.RegisterLibFunc(&GetPerformanceFrequency, lib, "SDL_GetPerformanceFrequency")
        purego.RegisterLibFunc(&Delay, lib, "SDL_Delay")
        purego.RegisterLibFunc(&DelayNS, lib, "SDL_DelayNS")
        purego.RegisterLibFunc(&DelayPrecise, lib, "SDL_DelayPrecise")
        purego.RegisterLibFunc(&AddTimer, lib, "SDL_AddTimer")
        purego.RegisterLibFunc(&AddTimerNS, lib, "SDL_AddTimerNS")
        purego.RegisterLibFunc(&RemoveTimer, lib, "SDL_RemoveTimer")
        
    })
}
