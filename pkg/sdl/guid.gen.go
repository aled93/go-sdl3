package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&GUIDToString, lib, "SDL_GUIDToString")
        purego.RegisterLibFunc(&StringToGUID, lib, "SDL_StringToGUID")
        
    })
}
