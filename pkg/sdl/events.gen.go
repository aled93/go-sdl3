package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&PumpEvents, lib, "SDL_PumpEvents")
        purego.RegisterLibFunc(&PeepEvents, lib, "SDL_PeepEvents")
        purego.RegisterLibFunc(&HasEvent, lib, "SDL_HasEvent")
        purego.RegisterLibFunc(&HasEvents, lib, "SDL_HasEvents")
        purego.RegisterLibFunc(&FlushEvent, lib, "SDL_FlushEvent")
        purego.RegisterLibFunc(&FlushEvents, lib, "SDL_FlushEvents")
        purego.RegisterLibFunc(&PollEvent, lib, "SDL_PollEvent")
        purego.RegisterLibFunc(&WaitEvent, lib, "SDL_WaitEvent")
        purego.RegisterLibFunc(&WaitEventTimeout, lib, "SDL_WaitEventTimeout")
        purego.RegisterLibFunc(&PushEvent, lib, "SDL_PushEvent")
        purego.RegisterLibFunc(&SetEventFilter, lib, "SDL_SetEventFilter")
        purego.RegisterLibFunc(&GetEventFilter, lib, "SDL_GetEventFilter")
        purego.RegisterLibFunc(&AddEventWatch, lib, "SDL_AddEventWatch")
        purego.RegisterLibFunc(&RemoveEventWatch, lib, "SDL_RemoveEventWatch")
        purego.RegisterLibFunc(&FilterEvents, lib, "SDL_FilterEvents")
        purego.RegisterLibFunc(&SetEventEnabled, lib, "SDL_SetEventEnabled")
        purego.RegisterLibFunc(&EventEnabled, lib, "SDL_EventEnabled")
        purego.RegisterLibFunc(&RegisterEvents, lib, "SDL_RegisterEvents")
        purego.RegisterLibFunc(&GetWindowFromEvent, lib, "SDL_GetWindowFromEvent")
        purego.RegisterLibFunc(&GetEventDescription, lib, "SDL_GetEventDescription")
        
    })
}
