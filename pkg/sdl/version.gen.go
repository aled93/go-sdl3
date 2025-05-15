package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&GetVersion, lib, "SDL_GetVersion")
        purego.RegisterLibFunc(&GetRevision, lib, "SDL_GetRevision")
        
    })
}
