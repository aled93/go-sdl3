package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&SetHintWithPriority, lib, "SDL_SetHintWithPriority")
        purego.RegisterLibFunc(&SetHint, lib, "SDL_SetHint")
        purego.RegisterLibFunc(&ResetHint, lib, "SDL_ResetHint")
        purego.RegisterLibFunc(&ResetHints, lib, "SDL_ResetHints")
        purego.RegisterLibFunc(&GetHint, lib, "SDL_GetHint")
        purego.RegisterLibFunc(&GetHintBoolean, lib, "SDL_GetHintBoolean")
        purego.RegisterLibFunc(&AddHintCallback, lib, "SDL_AddHintCallback")
        purego.RegisterLibFunc(&RemoveHintCallback, lib, "SDL_RemoveHintCallback")
        
    })
}
