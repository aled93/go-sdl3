package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&GetPreferredLocales, lib, "SDL_GetPreferredLocales")
        
    })
}
