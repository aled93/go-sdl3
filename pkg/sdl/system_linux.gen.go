package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&SetX11EventHook, lib, "SDL_SetX11EventHook")
        purego.RegisterLibFunc(&SetLinuxThreadPriority, lib, "SDL_SetLinuxThreadPriority")
        purego.RegisterLibFunc(&SetLinuxThreadPriorityAndPolicy, lib, "SDL_SetLinuxThreadPriorityAndPolicy")
        
    })
}
