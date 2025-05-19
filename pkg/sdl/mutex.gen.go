package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&ShouldInit, lib, "SDL_ShouldInit")
        purego.RegisterLibFunc(&ShouldQuit, lib, "SDL_ShouldQuit")
        purego.RegisterLibFunc(&SetInitialized, lib, "SDL_SetInitialized")
        
    })
}
