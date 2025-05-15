package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&OutOfMemory, lib, "SDL_OutOfMemory")
        purego.RegisterLibFunc(&GetError, lib, "SDL_GetError")
        purego.RegisterLibFunc(&ClearError, lib, "SDL_ClearError")
        
    })
}
